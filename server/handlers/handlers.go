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
	redirectHandler := NewRedirectHandler(services.SeedService, errorHandler)
	generatorHandler := NewGeneratorHandler(services.SeedService, errorHandler)
	staticHandler := NewStaticHandler(staticFs)
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
