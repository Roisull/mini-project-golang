package controller

import (
	"mini-project-golang/config"
	"mini-project-golang/middleware"
	"mini-project-golang/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateTrackController(c echo.Context) error{
	// mengambil data dari request
	track := model.Track{}
	c.Bind(&track)

	// mendapatkan token dari header
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Kamu belum bisa membuat track jika belum login")
	}

	// mendapatkan userID dari token JWT
	userID, err := middleware.ExtractTokenUserId(c)
	if err != nil{
		return echo.NewHTTPError(http.StatusUnauthorized, "Token anda tidak vali")
	}

	// set userID ke track
	track.UserID = uint(userID)

	// simpan track ke database
	if err := config.DB.Create(&track).Error; err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Playlist created successfully",
		"playlist": track,
	})
}