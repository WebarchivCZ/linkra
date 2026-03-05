package handlers

import (
	"errors"
	"linkra/assert"
	"linkra/server/components"
	"linkra/services"
	"linkra/utils"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type SeedHandler struct {
	Log          *slog.Logger
	SeedService  *services.SeedService
	ErrorHandler *ErrorHandler
}

func NewSeedHandler(log *slog.Logger, seedService *services.SeedService, errorHandler *ErrorHandler) *SeedHandler {
	assert.Must(log != nil, "NewSeedHandler: log can't be nil")
	assert.Must(seedService != nil, "NewSeedHandler: seedService can't be nil")
	assert.Must(errorHandler != nil, "NewSeedHandler: errorHandler can't be nil")
	return &SeedHandler{
		Log:          log,
		SeedService:  seedService,
		ErrorHandler: errorHandler,
	}
}

func (handler *SeedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, c *echo.Context) {
	requestedID := c.Param("id")
	seed, err := handler.SeedService.GetSeed(requestedID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		handler.Log.Warn("SeedHandler.ServeHTTP seed not found", "error", err.Error(), utils.LogRequestInfo(r))
		handler.ErrorHandler.PageNotFound(w, r) // Less scary and more informative than 500
		return
	}
	if err != nil {
		handler.Log.Error("SeedHandler.ServeHTTP failed to get Seed data from SeedService", "error", err.Error(), utils.LogRequestInfo(r))
		handler.ErrorHandler.InternalServerError(w, r)
		return
	}
	data := components.NewSeedViewData(seed, "Linkra - Detail "+seed.URL)
	err = handler.View(w, r, data)
	if err != nil {
		handler.Log.Error("SeedHandler.ServeHTTP failed to render view", "error", err.Error(), utils.LogRequestInfo(r))
		return
	}
	handler.Log.Info("SeedHandler.ServeHTTP sucessfully responded", utils.LogRequestInfo(r))
}

func (handler *SeedHandler) View(w http.ResponseWriter, r *http.Request, data *components.SeedViewData) error {
	return components.SeedView(data).Render(r.Context(), w)
}
