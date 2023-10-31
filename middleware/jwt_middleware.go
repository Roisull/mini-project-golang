package middleware

import (
	"mini-project-golang/constants"
	"time"

	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware() echo.MiddlewareFunc{
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(constants.SECRET_KEY_JWT),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_KEY_JWT))
}

func ExtractTokenUserId(e echo.Context) (int, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid{
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
	return int(userId), nil
	}
	return 0, nil
}