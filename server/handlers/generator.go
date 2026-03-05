package handlers

import (
	"linkra/assert"
	"linkra/entities"
	"linkra/server/components"
	"linkra/services"
	"linkra/utils"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
)

// The handler for citation generator page
type GeneratorHandler struct {
	Log          *slog.Logger
	SeedService  *services.SeedService
	ErrorHandler *ErrorHandler
}

func NewGeneratorHandler(log *slog.Logger, seedService *services.SeedService, errorHandler *ErrorHandler) *GeneratorHandler {
	assert.Must(log != nil, "NewGeneratorHandler: log can't be nil")
	assert.Must(seedService != nil, "NewGeneratorHandler: seedService can't be nil")
	assert.Must(errorHandler != nil, "NewGeneratorHandler: errorHandler can't be nil")
	return &GeneratorHandler{
		Log:          log,
		SeedService:  seedService,
		ErrorHandler: errorHandler,
	}
}

func (handler *GeneratorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, c *echo.Context) {
	var group *entities.SeedsGroup
	var err error
	id := c.Param("id")
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
