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

	// Register Route
	e.POST("/users", controller.CreateUserController)
	e.POST("/login", controller.LoginController)

	// Delete Without JWT
	// e.DELETE("/users/:id", controller.DeleteUserByIdController)

	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_KEY_JWT)))

	// user
	eJwt.GET("/users", controller.GetUsersController)
	eJwt.GET("/users/:id", controller.GetUserByIdController)
	eJwt.DELETE("/users/:id", controller.DeleteUserByIdController)
	eJwt.PUT("/users/:id", controller.UpdateUserByIdController)

	// playlist
	playlist := e.Group("/playlists", m.JwtMiddleware())
	playlist.POST("", controller.CreatePlaylistController)
	playlist.GET("", controller.GetAllPlaylistController)
	playlist.PUT("/:id", controller.UpdatePlaylistController)
	playlist.DELETE("/:id", controller.DeletePlaylistController)

	// track
	track := e.Group("/tracks", m.JwtMiddleware())
	track.POST("", controller.CreateTrackController)

	// middleware log
	m.LogMiddleware(e)

	return e
}