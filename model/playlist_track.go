package model

import "gorm.io/gorm"

type Playlist_track struct {
	gorm.Model
	PlaylistID uint `json:"playlist_id" form:"playlist_id"`
	Playlist Playlist `gorm:"foreignKey:PlaylistID"`
	TrackID uint `json:"track_id" form:"track_id"`
	Track Track `gorm:"foreignKey:TrackID"`
}