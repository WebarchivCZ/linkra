package handlers

import (
	"bytes"
	"fmt"
	"linkra/assert"
	"linkra/entities"
	"linkra/server/components"
	"linkra/services"
	"linkra/utils"
	"net/http"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type ExportGroupHandler struct {
	SeedService     *services.SeedService
	ExporterService *services.ExporterService
	ErrorHandler    *ErrorHandler
}

func NewExportGroupHandler(
	seedService *services.SeedService,
	exporterService *services.ExporterService,
	errorHandler *ErrorHandler,
) *ExportGroupHandler {
	assert.Must(seedService != nil, "NewExportGroupHandler: seedService can't be nil")
	assert.Must(exporterService != nil, "NewExportGroupHandler: exporterService can't be nil")
	assert.Must(errorHandler != nil, "NewExportGroupHandler: errorHandler can't be nil")
	return &ExportGroupHandler{
		SeedService:     seedService,
		ExporterService: exporterService,
		ErrorHandler:    errorHandler,
	}
}

func (handler *ExportGroupHandler) ServeHTTP(c *echo.Context) error {
	r := c.Request()
	w := c.Response()

	groupId := c.Param("id")

	group, err := handler.SeedService.GetGroup(groupId)
	// TODO: Create common error for services to communicate that record does not exist so we don't have to break the layer model all the time
	if err == gorm.ErrRecordNotFound {
		handler.ErrorHandler.PageNotFound(w, r)
		return fmt.Errorf("ExportGroupHandler.ServeHTTP group not found; %w", err)
	}
	if err != nil {
		handler.ErrorHandler.InternalServerError(w, r)
		return fmt.Errorf("ExportGroupHandler.ServeHTTP failed to fetch SeedsGroup data; %w", err)
	}

	format := c.Param("format")
	switch format {
	case "excel":
		handler.RespondExcel(c, group)
	case "csv":
		handler.RespondCsv(c, group)
	default:
		// TODO: This should maybe return pure 400. Investigate.
		handler.ErrorHandler.PageNotFound(w, r)
		return fmt.Errorf("ExportGroupHandler.ServeHTTP user requested unknown format")
	}
	return nil
}

func (handler *ExportGroupHandler) RespondExcel(c *echo.Context, group *entities.SeedsGroup) error {
	r := c.Request()
	w := c.Response()
	buffer := new(bytes.Buffer)
	err := handler.ExporterService.GroupToExcel(group, buffer)
	if err != nil {
		handler.ErrorHandler.InternalServerError(w, r)
		return fmt.Errorf("ExportGroupHandler.ServeHTTP got error from exporter service; %w", err)
	}

	header := w.Header()
	const XlsxMimetype = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	header.Set(utils.ContentType, XlsxMimetype)
	filename := components.ExcelFilename(group)
	header.Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	w.WriteHeader(http.StatusOK)
	_, _ = buffer.WriteTo(w)
	return nil
}

func (handler *ExportGroupHandler) RespondCsv(c *echo.Context, group *entities.SeedsGroup) error {
	r := c.Request()
	w := c.Response()
	buffer := new(bytes.Buffer)
	err := handler.ExporterService.GroupToCsv(group, buffer)
	if err != nil {
		handler.ErrorHandler.InternalServerError(w, r)
		return fmt.Errorf("ExportGroupHandler.ServeHTTP got error from exporter service; %w", err)
	}

	header := w.Header()
	const CsvMimetype = "text/csv"
	header.Set(utils.ContentType, CsvMimetype)
	filename := components.CsvFilename(group)
	header.Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	w.WriteHeader(http.StatusOK)
	_, _ = buffer.WriteTo(w)
	return nil
}
