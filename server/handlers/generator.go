package handlers

import (
	"fmt"
	"linkra/assert"
	"linkra/entities"
	"linkra/server/components"
	"linkra/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

// The handler for citation generator page
type GeneratorHandler struct {
	SeedService  *services.SeedService
	ErrorHandler *ErrorHandler
}

func NewGeneratorHandler(seedService *services.SeedService, errorHandler *ErrorHandler) *GeneratorHandler {
	assert.Must(seedService != nil, "NewGeneratorHandler: seedService can't be nil")
	assert.Must(errorHandler != nil, "NewGeneratorHandler: errorHandler can't be nil")
	return &GeneratorHandler{
		SeedService:  seedService,
		ErrorHandler: errorHandler,
	}
}

func (handler *GeneratorHandler) ServeHTTP(c *echo.Context) error {
	r := c.Request()
	w := c.Response()
	var group *entities.SeedsGroup
	var err error
	id := c.Param("id")
	// If path value id is undefined then group will remain nil. That is expected.
	// An unpopulated generator will be rendered in that case.
	if id != "" {
		group, err = handler.SeedService.GetGroup(id)
		if err != nil {
			handler.ErrorHandler.InternalServerError(w, r)
			return fmt.Errorf("GeneratorHandler.ServeHTTP failed to fetch SeedsGroup data; %w", err)
		}
	}
	err = handler.View(w, r, components.NewGeneratorViewData(group))
	if err != nil {
		return fmt.Errorf("GeneratorHandler.ServeHTTP failed to render view; %w", err)
	}
	return nil
}

func (handler *GeneratorHandler) View(w http.ResponseWriter, r *http.Request, data *components.GeneratorViewData) error {
	return components.GeneratorView(data).Render(r.Context(), w)
}
