package model

import (
	"gorm.io/gorm"
)

type Playlist struct {
	gorm.Model
	Name      string     `json:"name" form:"name"`
	UserID    uint       `json:"user_id" form:"user_id"` // Ini adalah foreign key ke tabel users
	User      User       `gorm:"foreignKey:UserID"` // Relasi dengan User
	Playlist_tracks []Playlist_track `json:"playlist_track" gorm:"foreignKey:PlaylistID"`
}