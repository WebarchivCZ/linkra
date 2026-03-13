package handlers

import (
	"errors"
	"fmt"
	"linkra/assert"
	"linkra/server/components"
	"linkra/services"
	"net/http"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type SeedHandler struct {
	SeedService  *services.SeedService
	ErrorHandler *ErrorHandler
}

func NewSeedHandler(seedService *services.SeedService, errorHandler *ErrorHandler) *SeedHandler {
	assert.Must(seedService != nil, "NewSeedHandler: seedService can't be nil")
	assert.Must(errorHandler != nil, "NewSeedHandler: errorHandler can't be nil")
	return &SeedHandler{
		SeedService:  seedService,
		ErrorHandler: errorHandler,
	}
}

func (handler *SeedHandler) ServeHTTP(c *echo.Context) error {
	r := c.Request()
	w := c.Response()
	requestedID := c.Param("id")
	seed, err := handler.SeedService.GetSeed(requestedID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		handler.ErrorHandler.PageNotFound(w, r) // Less scary and more informative than 500
		return fmt.Errorf("SeedHandler.ServeHTTP seed not found; %w", err)
	}
	if err != nil {
		handler.ErrorHandler.InternalServerError(w, r)
		return fmt.Errorf("SeedHandler.ServeHTTP failed to get Seed data from SeedService; %w", err)
	}
	data := components.NewSeedViewData(seed, &components.Translations{
		Czech:   "Linkra - Detail " + seed.URL,
		English: "Linkra - Detail " + seed.URL,
	})
	err = handler.View(w, r, data)
	if err != nil {
		return fmt.Errorf("SeedHandler.ServeHTTP failed to render view; %w", err)
	}
	return nil
}

func (handler *SeedHandler) View(w http.ResponseWriter, r *http.Request, data *components.SeedViewData) error {
	return components.SeedView(data).Render(r.Context(), w)
}
