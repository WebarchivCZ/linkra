package handlers

import (
	"linkra/server/components"

	"linkra/utils"
	"net/http"

	"github.com/labstack/echo/v5"
)

// Main handler for routes "/" and "/index.html"
type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (handler *IndexHandler) ServeHTTP(c *echo.Context) error {
	r := c.Request()
	w := c.Response()
	return handler.View(w, r)
}

func (handler *IndexHandler) View(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set(utils.ContentType, utils.TextHTML)
	return components.IndexView().Render(r.Context(), w)
}
