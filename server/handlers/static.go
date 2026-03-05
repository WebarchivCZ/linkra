package handlers

import (
	"io/fs"
	"linkra/assert"
	"net/http"

	"github.com/labstack/echo/v5"
)

// Handler for the "/static" route. Servers files from the fs.FS injected into the constructor function.
type StaticHandler struct {
	FileServer http.Handler
}

func NewStaticHandler(fs fs.FS) *StaticHandler {
	assert.Must(fs != nil, "NewStaticHandler: fs can't be nil")
	fileServer := http.FileServerFS(fs)
	return &StaticHandler{
		FileServer: fileServer,
	}
}

func (handler *StaticHandler) ServeHTTP(c *echo.Context) error {
	handler.FileServer.ServeHTTP(c.Response(), c.Request())
	return nil
}
