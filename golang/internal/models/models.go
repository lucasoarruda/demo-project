package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// User is the user model
type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	AccessLevel int       `json:"access_level"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// AccessLevel is the AccessLevel model
type AccessLevel struct {
	ID              uuid.UUID `json:"id"`
	AccessLevelName string    `json:"access_level_name"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ListOfAnimes struct {
	AnimesList []Animes
}

// Animes is the animes model
type Animes struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name" binding:"required"`
	Enabled     bool         `json:"enabled" default:"true"`
	Destination string       `json:"destination" binding:"required"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	SearchUrls  []SearchUrls `json:"search_urls"`
	Torrents    []Torrents   `json:"torrents"`
}

// SearchUrls is the SearchUrls model
type SearchUrls struct {
	ID        uuid.UUID `json:"id"`
	AnimeID   uuid.UUID `json:"anime_id" binding:"required"`
	Url       string    `json:"url" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Torrents is the room Torrents model
type Torrents struct {
	ID              uuid.UUID `json:"id"`
	AnimeID         uuid.UUID `json:"anime_id" binding:"required"`
	SearchUrlID     uuid.UUID `json:"search_url_id" binding:"required"`
	Magnet          string    `json:"magnet" binding:"required"`
	TorrentDelugeID string    `json:"torrent_deluge_id" default:""`
	Progress        int       `json:"progress" default:""`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Tlist struct {
	TorrentList struct {
		Torrents []struct {
			Name   string `json:"name"`
			Magnet string `json:"magnet"`
		} `json:"Torrents"`
	} `json:"TorrentList"`
}

type SUrl struct {
	URL string
}

type AddTorrent struct {
	Magnet          string         `json:"magnet" binding:"required"`
	TorrentDelugeID string         `json:"torrent_deluge_id"`
	Progress        int            `json:"progress"`
	Pathdestination string         `json:"path"`
	Newanime        bool           `json:"newanime"`
	Options         TorrentOptions `json:"options"`
}

type TorrentOptions struct {
	MoveCompletedPath string `json:"move_completed_path" default:"/downloads/media/videos/animes/"`
	AddPaused         bool   `json:"add_paused" default:"false"`
	MoveCompleted     bool   `json:"move_completed" default:"true"`
}
