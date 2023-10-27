package model

import "gorm.io/gorm"

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

type User struct {
	gorm.Model
	Name     	string `json:"name" form:"name"`
	Email    	string `json:"email" form:"email"`
	Password 	string `json:"password" form:"password"`
	Playlists 	[]Playlist // relasi dengan playlist
	Tracks 		[]Track // realasi dengan track
}

type UserResponse struct {
	gorm.Model
	Name		string `json:"name" form:"name"`
	Email		string `json:"email" form:"email"`
	Password 	string `json:"password" form:"password"`
	Token 		string `json:"token" form:"token"`
}