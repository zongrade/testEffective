package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	db "github.com/zongrade/testeffective/internal/database"
)

var (

	// ErrNoDeleted ничего не получилось удалить.
	ErrNoDeleted = errors.New("ничего не удалено")

	// ErrNothingToChange нет параметров для изменения.
	ErrNothingToChange = errors.New("нечего менять")
)

// Song provides struct for postgre field types.
type Song struct {
	ID          int       `json:"id"`
	Group       string    `json:"group"`
	Song        string    `json:"song"`
	Link        string    `json:"link"`
	ReleaseDate time.Time `json:"releaseDate"`
}

// GetSongs returns filtered songs.
func GetSongs(filter filtered) ([]Song, error) {
	songFilter := filter.GetSongFilter()
	groupFilter := filter.GetGroupFilter()

	var query strings.Builder
	query.WriteString("SELECT id, group_name, song_name,release_date,link FROM songs WHERE 1=1")

	args := make([]any, 0, len(groupFilter)+len(songFilter))
	var count int

	filterArgs := addFilterString(filterConf{
		stringWriter:     &query,
		fieldRestriction: songFilter,
		filedName:        "song_name",
		count:            count,
	})
	count += len(filterArgs)
	args = append(args, filterArgs...)

	filterArgs = addFilterString(filterConf{
		stringWriter:     &query,
		fieldRestriction: groupFilter,
		filedName:        "group_name",
		count:            count,
	})
	args = append(args, filterArgs...)

	query.WriteString(" ORDER BY id")
	query.WriteString(fmt.Sprintf(" LIMIT %d", filter.GetLimits()))
	query.WriteString(fmt.Sprintf(" OFFSET %d;", filter.GetPage()*filter.GetLimits()-filter.GetLimits()))

	return querySong(query, args)
}

func querySong(query strings.Builder, args []any) ([]Song, error) {
	var errSong error
	var rows *sql.Rows
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err = db.DBcfg.Connection.QueryContext(ctx, query.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	var songs []Song

	defer func() {
		if err := rows.Close(); err != nil {
			if errSong != nil {
				errSong = fmt.Errorf("%w: ошибка закрытия строк результата бд: %w", errSong, err)
			}

			errSong = fmt.Errorf("ошибка закрытия строк результата бд: %w", err)
		}
	}()

	for rows.Next() {
		var song Song
		if err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Link); err != nil {
			errSong = fmt.Errorf("ошибка сканирования строки: %w", err)

			return nil, errSong
		}

		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при чтении строк: %w", err)
	}

	return songs, errSong
}

type filterConf struct {
	stringWriter     io.StringWriter
	fieldRestriction []string
	filedName        string
	count            int
}

func addFilterString(conf filterConf) []any {
	args := make([]any, 0, len(conf.fieldRestriction))

	if len(conf.fieldRestriction) > 0 {
		// ошибки по WriteString будут игнорироваться
		// Наиболее вероятная ошибка это недостаток памяти, вероятность чего не высока
		// Предполагается разумная длина строки
		_, _ = conf.stringWriter.WriteString(fmt.Sprintf(" AND %s IN (", conf.filedName))

		for i, song := range conf.fieldRestriction {
			if i > 0 {
				_, _ = conf.stringWriter.WriteString(",")
			}

			_, _ = conf.stringWriter.WriteString("$")

			conf.count++
			_, _ = conf.stringWriter.WriteString(strconv.Itoa(conf.count))

			args = append(args, song)
		}

		_, _ = conf.stringWriter.WriteString(")")
	}

	return args
}

// DeleteSong удаляет песню по её ID и возвращает ID удалённой песни или ошибку.
func DeleteSong(filter filtered) (int, error) {
	arg := filter.GetSongID()
	query := "DELETE FROM songs WHERE songs.id = $1;"

	var err error
	var result sql.Result
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	result, err = db.DBcfg.Connection.ExecContext(ctx, query, arg)
	if err != nil {
		return 0, fmt.Errorf("не получилось удалить песню с id = %v: %w", arg, err)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return 0, fmt.Errorf("ошибка получения количества затронутых строк: %w", err)
	} else if rowsAffected < 1 {
		return 0, ErrNoDeleted
	}

	return arg, nil
}

// PatchSong изменяет параметры песни.
func PatchSong(filter filtered) (int, error) {
	var query strings.Builder
	var err error
	var result sql.Result
	var args []any

	if query, args, err = patchQuery(filter); err != nil {
		return 0, fmt.Errorf("не удалось правильно построить запрос: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	result, err = db.DBcfg.Connection.ExecContext(ctx, query.String(), args...)
	if err != nil {
		return 0, fmt.Errorf("не изменить песню с id = %v: %w", filter.GetSongID(), err)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return 0, fmt.Errorf("ошибка получения количества затронутых строк: %w", err)
	} else if rowsAffected < 1 {
		return 0, fmt.Errorf("не было затронута не 1 строка: %w", ErrNothingToChange)
	}

	return filter.GetSongID(), nil
}

func patchQuery(filter filtered) (strings.Builder, []any, error) {
	var query strings.Builder
	query.WriteString("UPDATE songs SET")
	var count int
	var args []any

	group := filter.GetGroup()
	if group != "" {
		count++

		args = append(args, group)

		query.WriteString(fmt.Sprintf(" group_name = $%d", count))
	}

	song := filter.GetSong()
	if song != "" {
		if count != 0 {
			query.WriteString(", ")
		}

		count++

		args = append(args, song)

		query.WriteString(fmt.Sprintf(" song_name = $%d", count))
	}

	if count < 1 {
		return query, nil, fmt.Errorf("нет параметров для изменения: %w", ErrNothingToChange)
	}

	songID := filter.GetSongID()
	if songID < 1 {
		return query, nil, fmt.Errorf("id меньше 1 не сушествует: %w", ErrNothingToChange)
	}

	args = append(args, songID)

	count++
	query.WriteString(fmt.Sprintf(" WHERE songs.id = $%d;", count))

	return query, args, nil
}

// CreateSong создаёт песню.
func CreateSong(filter filtered) (int, error) {
	var query strings.Builder
	var songID int
	var err error
	var args []any

	if query, args, err = postQuery(filter); err != nil {
		return 0, fmt.Errorf("не удалось правильно построить запрос: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err = db.DBcfg.Connection.QueryRowContext(ctx, query.String(), args...).Scan(&songID); err != nil {
		return 0, fmt.Errorf("не удалось создать песню: %w", err)
	}

	return songID, nil
}

func postQuery(filter filtered) (strings.Builder, []any, error) {
	var query strings.Builder
	query.WriteString("INSERT INTO songs (group_name, song_name,release_date,link) VALUES ($1,$2,$3,$4) RETURNING id;")
	var args []any

	group := filter.GetGroup()
	if group == "" {
		return query, nil, fmt.Errorf("отсутсвует необходимые параметры group для создания: %w", ErrNothingToChange)
	}

	args = append(args, group)

	song := filter.GetSong()
	if song == "" {
		return query, nil, fmt.Errorf("отсутсвует необходимые параметры song для создания: %w", ErrNothingToChange)
	}

	args = append(args, song)

	date := filter.GetReleaseDate().Format("2000-01-29")
	if date == "" {
		return query, nil, fmt.Errorf("отсутсвует необходимые параметры date для создания: %w", ErrNothingToChange)
	}

	args = append(args, date)

	link := filter.GetLink()
	if link == "" {
		return query, nil, fmt.Errorf("отсутсвует необходимые параметры link для создания: %w", ErrNothingToChange)
	}

	args = append(args, link)

	return query, args, nil
}
