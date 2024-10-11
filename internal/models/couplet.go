package models

import (
	"context"
	"fmt"
	"strings"
	"time"

	db "github.com/zongrade/testeffective/internal/database"
)

// GetCouplet returns couplet by page.
func GetCouplet(filter filtered) (string, error) {
	var text string
	songID := filter.GetSongID()
	page := filter.GetPage()

	var query strings.Builder
	query.WriteString("SELECT couplet_text FROM couplets WHERE couplets.song_id = $1 OFFSET $2 LIMIT 1")

	return queryCouplets(query, []any{songID, page - 1}, text)
}

// GetAllCouplets placholder.
func GetAllCouplets(filter filtered) (string, error) {
	var text struct {
		coupletText string
	}

	var fullText strings.Builder

	songID := filter.GetSongID()

	var query strings.Builder
	query.WriteString("SELECT couplet_text FROM couplets WHERE couplets.song_id = $1")

	if _, err := queryCouplets(query, []any{songID}, text, func(t struct {
		coupletText string
	}) {
		fullText.WriteString(t.coupletText)
	}); err != nil {
		return "", err
	}

	return fullText.String(), nil
}

func queryCouplets[T any](query strings.Builder, args []any, target T, extraFunc ...any) (T, error) {
	var errSong error

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err := db.DBcfg.Connection.QueryContext(ctx, query.String(), args...)
	if err != nil {
		return target, fmt.Errorf("ошибка выполнения запроса queryCouplets: %w", err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			if errSong != nil {
				errSong = fmt.Errorf("%w: ошибка закрытия строк результата бд: %w", errSong, err)
			}

			errSong = fmt.Errorf("ошибка закрытия строк результата бд: %w", err)
		}
	}()

	for rows.Next() {
		if err := rows.Scan(&target); err != nil {
			errSong = fmt.Errorf("ошибка сканирования строки: %w", err)

			return target, errSong
		}

		for _, fn := range extraFunc {
			if f, ok := fn.(func(T)); ok {
				f(target)
			}
		}
	}

	if err := rows.Err(); err != nil {
		return target, fmt.Errorf("ошибка при чтении строк: %w", err)
	}

	return target, errSong
}
