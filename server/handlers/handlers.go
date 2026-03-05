package handlers

import (
	"io/fs"
	"linkra/services"
	"log/slog"
)

type Handlers struct {
	ErrorHandler       *ErrorHandler
	IndexHandler       *IndexHandler
	SeedHandler        *SeedHandler
	RedirectHandler    *RedirectHandler
	GeneratorHandler   *GeneratorHandler
	StaticHandler      *StaticHandler
	GroupHandler       *GroupHandler
	ExportGroupHandler *ExportGroupHandler
	SaveGroupHandler   *SaveGroupHandler
}

func NewHandlers(log *slog.Logger, services *services.Services, staticFs fs.FS) *Handlers {
	errorHandler := NewErrorHandler(log)
	indexHandler := NewIndexHandler()
	seedHandler := NewSeedHandler(services.SeedService, errorHandler)
	redirectHandler := NewRedirectHandler(log, services.SeedService, errorHandler)
	generatorHandler := NewGeneratorHandler(log, services.SeedService, errorHandler)
	staticHandler := NewStaticHandler(log, staticFs)
	groupHandler := NewGroupHandler(log, services.SeedService, errorHandler)
	exporterGroupHandler := NewExportGroupHandler(log, services.SeedService, services.ExporterService, errorHandler)
	saveGroupHandler := NewSaveGroupHandler(log, services.SeedService, services.CaptureService, errorHandler)
	return &Handlers{
		ErrorHandler:       errorHandler,
		IndexHandler:       indexHandler,
		SeedHandler:        seedHandler,
		RedirectHandler:    redirectHandler,
		GeneratorHandler:   generatorHandler,
		StaticHandler:      staticHandler,
		GroupHandler:       groupHandler,
		ExportGroupHandler: exporterGroupHandler,
		SaveGroupHandler:   saveGroupHandler,
	}
}
