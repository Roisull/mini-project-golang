package config

import (
	"fmt"
	"mini-project-golang/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

const DB_USER = "root"
const DB_PASS = "i2226915September20012020"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"
const DB_NAME = "song_playlist_db"

func InitDB(){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// memanggil func InitMigrate() yang berisi auto migrate
	InitMigrate()
}

func InitMigrate(){
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Track{})
	DB.AutoMigrate(&model.Playlist{})
	DB.AutoMigrate(&model.Playlist_track{})
	// DB.AutoMigrate(&model.Admin{})
}