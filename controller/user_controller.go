package controller

import (
	"mini-project-golang/config"
	"mini-project-golang/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GetUsersHandler(c echo.Context) error {
	return GetUsersController(c)
}

// get ID from JWT
func GetIdFromJwt(tokenString string) (string, error) {
	// Mendeklarasikan struktur Claims untuk menyimpan klaim dari token JWT
	claims := jwt.MapClaims{}

	// Mendekode token JWT
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Pastikan Anda menggunakan algoritma dan kunci yang sama yang digunakan untuk menghasilkan token
		return []byte("ROISJWT!!!"), nil
	})

	if err != nil {
		return "", err
	}

	// Periksa apakah token valid dan berisi klaim yang sesuai
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, ok := claims["id"].(string); ok {
			return id, nil
		}
	}

	return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
}

// get all users

func GetUsersController(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")

	if token == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is missing")
	}

	userId, err := GetIdFromJwt(token)

	if err != nil {
		return err
	}

	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "success get all users",
		"users":   users,
		"userId":  userId,
	})

}