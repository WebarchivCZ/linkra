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

// This handler redirects shortened archival URLs to wayback
type RedirectHandler struct {
	SeedService  *services.SeedService
	ErrorHandler *ErrorHandler
}

func NewRedirectHandler(seedService *services.SeedService, errorHandler *ErrorHandler) *RedirectHandler {
	assert.Must(seedService != nil, "NewRedirectHandler: seedService can't be nil")
	assert.Must(errorHandler != nil, "NewRedirectHandler: errorHandler can't be nil")
	return &RedirectHandler{
		SeedService:  seedService,
		ErrorHandler: errorHandler,
	}
}

func (handler *RedirectHandler) ServeHTTP(c *echo.Context) error {
	r := c.Request()
	w := c.Response()
	seedId := c.Param("id")
	seed, err := handler.SeedService.GetSeed(seedId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		handler.ErrorHandler.PageNotFound(w, r) // Less scary and more informative than 500
		return fmt.Errorf("RedirectHandler.ServeHTTP seed not found; %w", err)
	}
	if err != nil {
		handler.ErrorHandler.InternalServerError(w, r)
		return fmt.Errorf("RedirectHandler.ServeHTTP failed to get Seed data from SeedService; %w", err)
	}

	if seed.ArchivalURL == "" {
		c.Logger().Warn("RedirectHandler.ServeHTTP seed is not harvested or archival URL is missing", "seed", seed.ShadowID)

		data := components.NewRedirectErrorViewData(seed)
		err = handler.ViewError(w, r, data)
		if err != nil {
			return fmt.Errorf("RedirectHandler.ServeHTTP failed to render error page; %w", err)
		}

		return nil // We did early serve, don't continue.
	}

	http.Redirect(w, r, seed.ArchivalURL, http.StatusMovedPermanently)
	return nil
}

func (handler *RedirectHandler) ViewError(w http.ResponseWriter, r *http.Request, data *components.RedirectErrorViewData) error {
	return components.RedirectErrorView(data).Render(r.Context(), w)
}
