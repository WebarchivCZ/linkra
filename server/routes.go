package server

import (
	"linkra/server/handlers"
	"net/http"

	"github.com/labstack/echo/v5"
)

func RegisterRoutes(router *echo.Echo, handlers *handlers.Handlers) {
	router.GET("/", IntoFunc(handlers.IndexHandler))
	router.GET("/seeds/export/:format/:id", partialWrapHandler(handlers.ExportGroupHandler))
	router.GET("/citace", partialWrapHandler(handlers.GeneratorHandler))
	router.GET("/citace/:id", partialWrapHandler(handlers.GeneratorHandler))
	router.GET("/seeds/:id", partialWrapHandler(handlers.GroupHandler))
	router.GET("/wa/:id", partialWrapHandler(handlers.RedirectHandler))
	router.POST("/seeds/save", echo.WrapHandler(handlers.SaveGroupHandler))
	router.GET("/seed/:id", IntoFunc(handlers.SeedHandler))
	router.GET("/static/*", echo.WrapHandler(handlers.StaticHandler))
}

// Mixed handler to speed up migration to echo
// TODO: Migrate all handlers to echo handlers
type mixedHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, c *echo.Context)
}

func partialWrapHandler(h mixedHandler) echo.HandlerFunc {
	return func(c *echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request(), c)
		return nil
	}
}

type EchoHandler interface {
	ServeHTTP(*echo.Context) error
}

func IntoFunc(h EchoHandler) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return h.ServeHTTP(c)
	}
}
