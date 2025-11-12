package generator

import (
	"linkra/assert"
	"linkra/entities"
	"linkra/server/components"
	"linkra/server/handlers/httperror"
	"linkra/services"
	"linkra/utils"
	"log/slog"
	"net/http"
)

// The handler for citation generator page
type GeneratorHandler struct {
	Log          *slog.Logger
	SeedService  *services.SeedService
	ErrorHandler *httperror.ErrorHandler
}

func NewGeneratorHandler(log *slog.Logger, seedService *services.SeedService, errorHandler *httperror.ErrorHandler) *GeneratorHandler {
	assert.Must(log != nil, "NewGeneratorHandler: log can't be nil")
	assert.Must(seedService != nil, "NewGeneratorHandler: seedService can't be nil")
	assert.Must(errorHandler != nil, "NewGeneratorHandler: errorHandler can't be nil")
	return &GeneratorHandler{
		Log:          log,
		SeedService:  seedService,
		ErrorHandler: errorHandler,
	}
}

func (handler *GeneratorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var group *entities.SeedsGroup
	var err error
	id := r.PathValue("id")
	// If path value id is undefined then group will remain nil. That is expected.
	if id != "" {
		group, err = handler.SeedService.GetGroup(id)
		if err != nil {
			handler.Log.Error("GeneratorHandler.ServeHTTP failed to fetch SeedsGroup data", "error", err.Error(), utils.LogRequestInfo(r))
			handler.ErrorHandler.InternalServerError(w, r)
			return
		}
	}
	err = handler.View(w, r, components.NewGeneratorViewData(group))
	if err != nil {
		handler.Log.Error("GeneratorHandler.ServeHTTP failed to render view", "error", err.Error(), utils.LogRequestInfo(r))
		return
	}
	handler.Log.Info("GeneratorHandler.ServeHTTP sucessfully responded", utils.LogRequestInfo(r))
}

func (handler *GeneratorHandler) View(w http.ResponseWriter, r *http.Request, data *components.GeneratorViewData) error {
	return components.GeneratorView(data).Render(r.Context(), w)
}

func (handler *GeneratorHandler) Routes(mux *http.ServeMux) {
	mux.Handle("GET /citace/", handler)
	mux.Handle("GET /citace/{id}", handler)
}
