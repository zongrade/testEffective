// Package models provides models working with db
package models

import "time"

type filtered interface {
	GetSongFilter() []string
	GetGroupFilter() []string
	GetLimits() int
	GetPage() int
	GetSongID() int
	GetSong() string
	GetGroup() string
	GetReleaseDate() time.Time
	GetLink() string
	SetSongID(songID int)
}
