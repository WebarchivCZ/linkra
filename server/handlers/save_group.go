package handlers

import (
	"errors"
	"fmt"
	"linkra/assert"
	"linkra/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

func NewSaveGroupHandler(
	seedService *services.SeedService,
	captureService *services.CaptureService,
	errorHandler *ErrorHandler,
) *SaveGroupHandler {
	assert.Must(seedService != nil, "NewSaveGroupHandler: seedService can't be nil")
	assert.Must(captureService != nil, "NewSaveGroupHandler: captureService can't be nil")
	assert.Must(errorHandler != nil, "NewSaveGroupHandler: errorHandler can't be nil")
	return &SaveGroupHandler{
		SeedService:    seedService,
		CaptureService: captureService,
		ErrorHandler:   errorHandler,
	}
}

type SaveGroupHandler struct {
	SeedService    *services.SeedService
	CaptureService *services.CaptureService
	ErrorHandler   *ErrorHandler
}

func (handler *SaveGroupHandler) ServeHTTP(c *echo.Context) error {
	r := c.Request()
	w := c.Response()

	const urlKey = "url-list"
	// TODO: Check that server has correct setting for request size.
	seedURL := r.FormValue(urlKey)
	group, err := handler.SeedService.Save(seedURL, true)
	// TODO: Return different error pages/messages when different errors are received. This should help user understand what they did wrong.
	if errors.Is(err, services.ErrEmptyList) {
		handler.ErrorHandler.ServeError(w, r, "Prázdný požadavek", 400, "Prázdný požadavek", "Požadavek který jsme obdrželi obsahoval jen prázdné řádky. Prosím vraťte se na hlavní stránku a zadejte platnou URL adresu.")
		return errors.New("SaveGroupHandler.ServeHTTP received empty seed list")
	}
	if err != nil {
		handler.ErrorHandler.InternalServerError(w, r)
		return fmt.Errorf("SaveGroupHandler.ServeHTTP SeedService returned error when trying to save group; %w", err)
	}

	// Enqueue seeds for capture (and let the request succeed, it can be reenqueue later)
	err = handler.CaptureService.CaptureGroup(r.Context(), group)
	http.Redirect(w, r, "/seeds/"+group.ShadowID, http.StatusSeeOther)
	// The request completed, this will only log the potential error
	return fmt.Errorf("SaveGroupHandler.ServeHTTP CaptureService returned error when trying to enqueue group; %w", err)
}
