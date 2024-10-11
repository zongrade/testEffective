package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zongrade/testeffective/internal/models"
)

// GetCouplet processes requests for verses.
// @Summary Get Couplets
// @Description получить куплеты
// @Tags Couplets
// @Param song_id query int true "ID песни"
// @Param page query int false "номер куплета"
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /couplets [get].
func GetCouplet(w http.ResponseWriter, r *http.Request) {
	var text string
	var err error
	var filter *filterConfig

	if filter, err = setupFilter(r, []func(*http.Request, *filterConfig) error{
		setupSongIDFilter,
		setupPageFilter,
	}, filterConfig{
		SongID: 1,
	}); err != nil {
		fmt.Printf("Ошибка в фильтрах %v", err)
		http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)

		return
	}

	if text, err = models.GetCouplet(filter); err != nil {
		http.Error(w, "Ошибка получения песен", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(text); err != nil {
		log.Printf("Ошибка кодирования JSON: %v", err)
		http.Error(w, "Ошибка обработки запроса", http.StatusInternalServerError)
	}
}
