package route

import (
	"mini-project-golang/constants"
	"mini-project-golang/controller"
	m "mini-project-golang/middleware"

	mid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo{
	// intance echo
	e := echo.New()

	m.LogMiddleware(e)

	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_KEY_JWT)))

	// user
	eJwt.GET("/users", controller.GetUsersController)

	return e
}