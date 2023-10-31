package repositories

import (
	"errors"
	"mini-project-golang/config"
	"mini-project-golang/model"
)

func SelectPlaylist() ([]model.Playlist, error){
	// menggunakan DB
	var dataPlaylist []model.Playlist
	// SELECT * FROM USER
	tx := config.DB.Preload("User").Find(&dataPlaylist)
	if tx.Error != nil{
		return nil, tx.Error
	}
	return dataPlaylist, nil
}

func InsertPlaylist(data model.Playlist) error {
	tx := config.DB.Create(&data)
	if tx.Error != nil {
		return errors.New("failed insert data")
	}

	return nil
}