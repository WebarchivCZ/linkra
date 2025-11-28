package server

import (
	"context"
	"linkra/server/handlers"
	"linkra/server/handlers/generator"
	"linkra/server/handlers/group"
	"linkra/server/handlers/httperror"
	"linkra/server/handlers/index"
	"linkra/server/handlers/redirect"
	"linkra/server/handlers/seed"
	"linkra/server/handlers/static"
	"linkra/services"
	"log/slog"
	"net"
	"net/http"
	"time"
)

func NewServer(ctx context.Context, log *slog.Logger, addr string, services *services.Services) *http.Server {
	// Create router
	mux := http.NewServeMux()
	router := handlers.NewRouterHandler(mux)

	// Create the error handler
	errorHandler := httperror.NewErrorHandler(log)

	// Add all handlers to the router
	router.AddHandlers(
		index.NewIndexHandler(log, errorHandler),
		static.NewStaticHandler(log, staticFiles /* from embed.go */),
		group.NewGroupHandler(log, services.SeedService, services.ExporterService, services.CaptureService, errorHandler),
		seed.NewSeedHandler(log, services.SeedService, errorHandler),
		generator.NewGeneratorHandler(log, services.SeedService, errorHandler),
		redirect.NewRedirectHandler(log, services.SeedService, errorHandler),
	)

	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		BaseContext:  func(l net.Listener) context.Context { return ctx },
	}

	return server
}
