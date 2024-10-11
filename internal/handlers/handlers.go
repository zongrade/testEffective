// Package handlers provides handlers
package handlers

import (
	"net/http"
	"time"
)

type SongApi struct {
	ID          int       `json:"id"`
	Group       string    `json:"group"`
	Song        string    `json:"song"`
	Link        string    `json:"link"`
	ReleaseDate time.Time `json:"releaseDate"`
}

// InfoHandler orchestrates requests with song depending on their method.
// @Summary Получить информацию о песне
// @Description Возвращает информацию о конкретной песне по её ID.
// @Produce json
// @Param song_id query int true "ID песни"
// @Param group query string true "название группы"
// @Param song query string true "название песни"
// @Success 200 {object} SongApi
// @Failure 500 {array} SongApi
// @Router /info [get].
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetSongs(w, r, [2][]string{{"Non zero value"}, {"Non zero value"}})
	default:
		http.Error(w, "No registered method", http.StatusMethodNotAllowed)
	}
}

// SongsHandler orchestrates requests with song depending on their method.
func SongsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetSongs(w, r, [2][]string{})
	case http.MethodDelete:
		DeleteSongs(w, r)
	case http.MethodPatch:
		PatchSongs(w, r)
	case http.MethodPost:
		PostSongs(w, r)
	default:
		http.Error(w, "No registered method", http.StatusMethodNotAllowed)
	}
}

// CoupletsHandler orchestrates requests with couplet depending on their method.
func CoupletsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetCouplet(w, r)
	default:
		http.Error(w, "No registered method", http.StatusMethodNotAllowed)
	}
}
