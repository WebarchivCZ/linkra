package server

import (
	"context"
	"linkra/server/handlers"
	"linkra/services"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func NewServer(ctx context.Context, log *slog.Logger, addr string, services *services.Services, e *echo.Echo) *http.Server {
	// Create router
	// mux := http.NewServeMux()
	// router := handlers.NewRouterHandler(mux)

	// Create the error handler
	// errorHandler := httperror.NewErrorHandler(log)

	// Add all handlers to the router
	// router.AddHandlers(
	// 	index.NewIndexHandler(log, errorHandler),
	// 	static.NewStaticHandler(log, staticFiles /* from embed.go */),
	// 	group.NewGroupHandler(log, services.SeedService, services.ExporterService, services.CaptureService, errorHandler),
	// 	seed.NewSeedHandler(log, services.SeedService, errorHandler),
	// 	generator.NewGeneratorHandler(log, services.SeedService, errorHandler),
	// 	redirect.NewRedirectHandler(log, services.SeedService, errorHandler),
	// )

	handlers := handlers.NewHandlers(log, services, staticFiles /* from embed.go */)
	e.Pre(middleware.RemoveTrailingSlash())
	RegisterRoutes(e, handlers)

	server := &http.Server{
		Addr:         addr,
		Handler:      e,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		BaseContext:  func(l net.Listener) context.Context { return ctx },
	}

	return server
}
