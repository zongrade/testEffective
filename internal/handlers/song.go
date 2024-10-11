package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/zongrade/testeffective/internal/models"
)

// GetSongs получает песню.
// @Summary Get Songs
// @Description Получить список песен с фильтрацией
// @Tags songs
// @Param group query string false "Фильтр по названию группы"
// @Param song query string false "Фильтр по названию песни"
// @Param limit query int false "лимит"
// @Param page query int false "страница"
// @Produce json
// @Success 200 {object} SongApi
// @Failure 500 {array} SongApi
// @Router /songs [get].
func GetSongs(w http.ResponseWriter, r *http.Request, reqParam [2][]string) {
	var songs []models.Song
	var err error
	var filter *filterConfig
	var reqFilte = filterConfig{
		songs:  reqParam[0],
		groups: reqParam[1],
	}

	if filter, err = setupFilter(r, []func(*http.Request, *filterConfig) error{
		setupGroupsFilter,
		setupSongsFilter,
		setupLimitsFilter,
		setupPageFilter,
	}, reqFilte); err != nil {
		fmt.Printf("Ошибка в фильтрах %v", err)
		http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)

		return
	}

	if songs, err = models.GetSongs(filter); err != nil {
		log.Println(err)
		http.Error(w, "Ошибка получения песен", http.StatusInternalServerError)

		return
	}

	if songs == nil {
		songs = []models.Song{}
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(songs); err != nil {
		log.Printf("Ошибка кодирования JSON: %v", err)
		http.Error(w, "Ошибка обработки запроса", http.StatusInternalServerError)
	}
}

type changeResponse struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

// DeleteSongs обрабатывает попытку удаления песни с указанным id
// @Summary Delete Songs
// @Description Удалить песню
// @Tags songs
// @Param song_id body int true "ID песни для удаления"
// @Produce json
// @Success 200 {object} changeResponse
// @Failure 500 {object} string
// @Router /songs [delete].
func DeleteSongs(w http.ResponseWriter, r *http.Request) {
	var err error
	var filter *filterConfig
	var delResponse changeResponse

	if filter, err = setupFilter(r, []func(*http.Request, *filterConfig) error{
		setupFilterJSON,
	}, filterConfig{
		SongID: 1,
	}); err != nil {
		fmt.Printf("Ошибка в фильтрах %v", err)
		http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)

		return
	}

	if delResponse.ID, err = models.DeleteSong(filter); err != nil {
		if errors.Is(err, models.ErrNoDeleted) {
			log.Println(err)
			http.Error(w, models.ErrNoDeleted.Error(), http.StatusInternalServerError)

			return
		}

		log.Println(err)
		http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	delResponse.Message = "удалена песня с ID "

	if err := json.NewEncoder(w).Encode(delResponse); err != nil {
		log.Printf("Ошибка кодирования JSON: %v", err)
		http.Error(w, "Ошибка обработки запроса", http.StatusInternalServerError)
	}
}

// PatchSongs обрабатывает попытку изменения песни с указанным id
// @Summary Patch Songs
// @Description обновить песню
// @Tags songs
// @Param song_id body int true "ID песни для обновления"
// @Param group body int false "название группы для обновления"
// @Param song body int false "название песни для обновления"
// @Produce json
// @Success 200 {object} changeResponse
// @Failure 500 {object} string
// @Router /songs [patch].
func PatchSongs(w http.ResponseWriter, r *http.Request) {
	var err error
	var filter *filterConfig
	var delResponse changeResponse

	if filter, err = setupFilter(r, []func(*http.Request, *filterConfig) error{
		setupSongIDFilter,
		setupFilterJSON,
	}, filterConfig{
		SongID: 1,
	}); err != nil {
		fmt.Printf("Ошибка в фильтрах %v", err)
		http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)

		return
	}

	if delResponse.ID, err = models.PatchSong(filter); err != nil {
		if errors.Is(err, models.ErrNoDeleted) {
			log.Println(err)
			http.Error(w, models.ErrNoDeleted.Error(), http.StatusInternalServerError)

			return
		}

		log.Println(err)
		http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	delResponse.Message = "изменение успешно песни с ID"

	if err := json.NewEncoder(w).Encode(delResponse); err != nil {
		log.Printf("Ошибка кодирования JSON: %v", err)
		http.Error(w, "Ошибка обработки запроса", http.StatusInternalServerError)
	}
}

type postResponse struct {
	ReleaseDate time.Time `json:"releaseDate"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}

// PostSongs обрабатывает попытку добавления песни
// @Summary Post Songs
// @Description добавить песню
// @Tags songs
// @Param group body int true "название группы"
// @Param song body int true "название песни"
// @Produce json
// @Success 200 {object} postResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /songs [post].
func PostSongs(w http.ResponseWriter, r *http.Request) {
	var err error
	var filter *filterConfig
	var pResponse postResponse

	config := filterConfig{
		Song:  "non zero value",
		Group: "non zero value",
	}

	config.ReleaseDate = time.Now()
	config.Link = "yotube.com/placeholder"

	if filter, err = setupFilter(r, []func(*http.Request, *filterConfig) error{
		setupFilterJSON,
	}, config); err != nil {
		fmt.Printf("Ошибка в фильтрах %v", err)
		http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)

		return
	}

	var group = filter.Group
	var song = filter.Song

	if group == "" || song == "" {
		fmt.Println("недостаток параметров для создания песни")
		http.Error(w, "Ошибка параметров", http.StatusBadRequest)
	}

	if _, err = models.CreateSong(filter); err != nil {
		if errors.Is(err, models.ErrNoDeleted) {
			log.Println(err)
			http.Error(w, models.ErrNoDeleted.Error(), http.StatusInternalServerError)

			return
		}

		log.Println(err)
		http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)

		return
	}

	filter.groups = []string{group}
	filter.songs = []string{song}
	var data []models.Song

	if data, err = models.GetSongs(filter); err != nil {
		log.Printf("Второго запроса после добавления: %v", err)
		http.Error(w, "Ошибка обработки запроса", http.StatusInternalServerError)
	}

	pResponse.ReleaseDate = data[0].ReleaseDate
	pResponse.Link = data[0].Link

	if pResponse.Text, err = models.GetAllCouplets(filter); err != nil {
		log.Printf("ошибка получения всех куплетов: %v", err)
		http.Error(w, "Ошибка обработки запроса", http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(pResponse); err != nil {
		log.Printf("Ошибка кодирования JSON: %v", err)
		http.Error(w, "Ошибка обработки запроса", http.StatusInternalServerError)
	}
}
