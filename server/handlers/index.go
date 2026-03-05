package handlers

import (
	"linkra/assert"
	"linkra/server/components"

	"linkra/utils"
	"log/slog"
	"net/http"
)

// Main handler for routes "/" and "/index.html"
type IndexHandler struct {
	Log          *slog.Logger
	ErrorHandler *ErrorHandler
}

func NewIndexHandler(log *slog.Logger, errorHandler *ErrorHandler) *IndexHandler {
	assert.Must(log != nil, "NewIndexHandler: log can't be nil")
	assert.Must(errorHandler != nil, "NewIndexHandler: errorHandler can't be nil")
	return &IndexHandler{
		Log:          log,
		ErrorHandler: errorHandler,
	}
}

func (handler *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := handler.View(w, r)
	if err != nil {
		handler.Log.Error("IndexHandler.ServeHTTP failed to render view", "error", err.Error(), utils.LogRequestInfo(r))
		return
	}
	handler.Log.Info("IndexHandler sucessfully responded", utils.LogRequestInfo(r))
}

func (handler *IndexHandler) View(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set(utils.ContentType, utils.TextHTML)
	return components.IndexView().Render(r.Context(), w)
}
