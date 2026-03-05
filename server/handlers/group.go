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

// Handler for seed groups. Used to create/show list of seeds to make tracking of progress of individual seeds easier.
type GroupHandler struct {
	SeedService    *services.SeedService
	CaptureService *services.CaptureService
	ErrorHandler   *ErrorHandler
}

func NewGroupHandler(
	seedService *services.SeedService,
	errorHandler *ErrorHandler,
) *GroupHandler {
	assert.Must(seedService != nil, "NewGroupHandler: seedService can't be nil")
	assert.Must(errorHandler != nil, "NewGroupHandler: errorHandler can't be nil")
	return &GroupHandler{
		SeedService:  seedService,
		ErrorHandler: errorHandler,
	}
}

func (handler *GroupHandler) ServeHTTP(c *echo.Context) error {
	r := c.Request()
	w := c.Response()
	requestedID := c.Param("id")
	group, err := handler.SeedService.GetGroup(requestedID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		handler.ErrorHandler.PageNotFound(w, r) // Less scary and more informative than 500
		return fmt.Errorf("GroupHandler.ServeHTTP group not found; %w", err)
	}
	if err != nil {
		handler.ErrorHandler.InternalServerError(w, r)
		return fmt.Errorf("GroupHandler.ServeHTTP failed to fetch SeedsGroup data; %w", err)
	}
	data := components.NewGroupViewData(group)
	err = handler.View(w, r, data)
	if err != nil {
		return fmt.Errorf("GroupHandler.ServeHTTP failed to render view; %w", err)
	}
	return nil
}

func (handler *GroupHandler) View(w http.ResponseWriter, r *http.Request, data *components.GroupViewData) error {
	return components.GroupView(data).Render(r.Context(), w)
}
