package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ErrNotEnoughParametrs указывает на необходимость других параметров.
var ErrNotEnoughParametrs = errors.New("не достаточно параметров для запроса")

type filterConfig struct {
	songs       []string
	groups      []string
	limits      int
	page        int
	Link        string
	ReleaseDate time.Time
	SongID      int    `json:"songId"`
	Song        string `json:"song"`
	Group       string `json:"group"`
}

func (f filterConfig) GetSongFilter() []string {
	return f.songs
}

func (f filterConfig) GetGroupFilter() []string {
	return f.groups
}

func (f filterConfig) GetLimits() int {
	return f.limits
}

func (f filterConfig) GetPage() int {
	return f.page
}

func (f filterConfig) GetSongID() int {
	return f.SongID
}

func (f filterConfig) GetSong() string {
	return f.Song
}

func (f filterConfig) GetGroup() string {
	return f.Group
}

func (f filterConfig) GetReleaseDate() time.Time {
	return f.ReleaseDate
}

func (f filterConfig) GetLink() string {
	return f.Link
}

func (f *filterConfig) SetSongID(songID int) {
	f.SongID = songID
}

func createFilter() *filterConfig {
	return &filterConfig{
		limits: 3,
		page:   1,
	}
}

func setupGroupsFilter(r *http.Request, filter *filterConfig) error {
	if groups, ok := r.URL.Query()["group"]; ok {
		for i := 0; i < len(groups); i++ {
			groups[i] = strings.Trim(groups[i], "\"")
		}

		filter.groups = groups
	}

	return nil
}

func setupFilterJSON(r *http.Request, filter *filterConfig) error {
	return getJSONData(r, filter)
}

func setupSongsFilter(r *http.Request, filter *filterConfig) error {
	if songs, ok := r.URL.Query()["song"]; ok {
		for i := 0; i < len(songs); i++ {
			songs[i] = strings.Trim(songs[i], "\"")
		}
		filter.songs = songs
	}

	return nil
}

func setupLimitsFilter(r *http.Request, filter *filterConfig) error {
	if limits, ok := r.URL.Query()["limits"]; ok {
		if lim, err := strconv.Atoi(limits[0]); err == nil && lim < 5 && lim > 0 {
			filter.limits = lim
		}
	}

	return nil
}

func setupPageFilter(r *http.Request, filter *filterConfig) error {
	if p, ok := r.URL.Query()["page"]; ok {
		if page, err := strconv.Atoi(p[0]); err == nil {
			filter.page = page
		}
	}

	return nil
}

func setupSongIDFilter(r *http.Request, filter *filterConfig) error {
	if id, ok := r.URL.Query()["song_id"]; ok {
		if page, err := strconv.Atoi(id[0]); err == nil {
			filter.SongID = page
		}
	}

	return nil
}

func getJSONData(r *http.Request, reqStruct *filterConfig) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("ошибка чтения body %w", err)
	}

	err = json.Unmarshal(body, reqStruct)
	if err != nil {
		return fmt.Errorf("неправильный JSON формат %w", err)
	}

	return nil
}

func setupFilter(r *http.Request, filterFuncs []func(*http.Request, *filterConfig) error,
	requiredConf filterConfig) (*filterConfig, error) {
	filter := createFilter()
	for _, v := range filterFuncs {
		if err := v(r, filter); err != nil {
			return filter, err
		}
	}

	if err := isSetupCorrect(filter, requiredConf); err != nil {
		return filter, err
	}

	return filter, nil
}

func isSetupCorrect(sourceConf *filterConfig, destConf filterConfig) error {
	if destConf.Link == "yotube.com/placeholder" {
		sourceConf.Link = "yotube.com/placeholder"
	}
	if destConf.groups != nil && sourceConf.groups == nil {
		return fmt.Errorf("нехватает groups в запросе: %w", ErrNotEnoughParametrs)
	}

	if destConf.songs != nil && sourceConf.songs == nil {
		return fmt.Errorf("нехватает songs в запросе: %w", ErrNotEnoughParametrs)
	}

	if destConf.SongID != 0 && sourceConf.SongID == 0 {
		return fmt.Errorf("нехватает songID в запросе: %w", ErrNotEnoughParametrs)
	}

	if destConf.Group != "" && sourceConf.Group == "" {
		return fmt.Errorf("нехватает Group в запросе: %w", ErrNotEnoughParametrs)
	}

	if destConf.Song != "" && sourceConf.Song == "" {
		return fmt.Errorf("нехватает Song в запросе: %w", ErrNotEnoughParametrs)
	}

	return nil
}
