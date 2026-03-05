package server

import (
	"linkra/server/handlers"

	"github.com/labstack/echo/v5"
)

func RegisterRoutes(router *echo.Echo, handlers *handlers.Handlers) {
	router.GET("/", IntoFunc(handlers.IndexHandler))
	router.GET("/seeds/export/:format/:id", IntoFunc(handlers.ExportGroupHandler))
	router.GET("/citace", IntoFunc(handlers.GeneratorHandler))
	router.GET("/citace/:id", IntoFunc(handlers.GeneratorHandler))
	router.GET("/seeds/:id", IntoFunc(handlers.GroupHandler))
	router.GET("/wa/:id", IntoFunc(handlers.RedirectHandler))
	router.POST("/seeds/save", IntoFunc(handlers.SaveGroupHandler))
	router.GET("/seed/:id", IntoFunc(handlers.SeedHandler))
	router.GET("/static/*", IntoFunc(handlers.StaticHandler))
}

type EchoHandler interface {
	ServeHTTP(*echo.Context) error
}

func IntoFunc(h EchoHandler) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return h.ServeHTTP(c)
	}
}
