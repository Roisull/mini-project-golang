package controller

import (
	"mini-project-golang/config"
	"mini-project-golang/middleware"
	"mini-project-golang/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "rois@gmail.com" && password == "AdminPlaylist" {
		token, err := middleware.CreateToken(int(admin))
	}
}