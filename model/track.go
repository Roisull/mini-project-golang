package model

import (
	"gorm.io/gorm"
)

type Track struct {
	gorm.Model
	Name      	string     	`json:"name" form:"name"`
	Artist		string 		`json:"artist" form:"artist"`
	UserID    	uint       	`json:"user_id" form:"user_id"` // Ini adalah foreign key ke tabel users
	User      	User       	`gorm:"foreignKey:UserID"` // Relasi dengan User
	Playlist_tracks []Playlist_track `json:"track_id" gorm:"foreignKey:TrackID"`
}