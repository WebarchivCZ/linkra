package handlers

import (
	"errors"
	"fmt"
	"linkra/assert"
	"linkra/server/components"
	"linkra/utils"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

// This handler is used to handle errors propagated through echo middleware chain.
// It should be registered directly into Echo.HTTPErrorHandler.
//
// For now, most handlers call ErrorHandler directly before propagating errors,
// so only some errors (mostly 404s) will be caught by this.
// TODO: Use error handling mechanisms provided by echo.
func NewEchoErrorHandler(errorHandler *ErrorHandler) echo.HTTPErrorHandler {
	return func(c *echo.Context, err error) {
		// Don't do anything if some handler already responded.
		if resp, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
			if resp.Committed {
				return
			}
		}

		// Find error that implements HTTPStatusCoder and use it's status code or use 500 Internal Server Error.
		code := http.StatusInternalServerError
		var sc echo.HTTPStatusCoder
		if errors.As(err, &sc) {
			if tmp := sc.StatusCode(); tmp != 0 {
				code = tmp
			}
		}

		// Var for error returned by attempt to serve error page.
		var cErr error

		if c.Request().Method == http.MethodHead {
			cErr = c.NoContent(code)
		} else if code == 404 {
			cErr = errorHandler.PageNotFound(c.Response(), c.Request())
		} else if code == 405 {
			cErr = errorHandler.MethodNotAllowed(c.Response(), c.Request())
		} else if code == 500 {
			cErr = errorHandler.InternalServerError(c.Response(), c.Request())
		} else {
			// TODO: Add special handling for more client errors (400, 401, 413, 414, etc.). Other server error are usually produced by proxy in situations where responding is not possible anyway.
			desc := http.StatusText(code)
			cErr = errorHandler.ServeError(
				c.Response(),
				c.Request(),
				desc,
				code,
				desc,
				"Try again later.",
			)
		}

		var uce *UnusualStatusCodeError
		if errors.As(err, &uce) {
			c.Logger().Warn("POTENTIAL BUG", "error", uce)
		} else if cErr != nil {
			c.Logger().Error("serving error page failed", "error", errors.Join(cErr, err))
		}
	}
}

// This handler is little bit different. It has no ServeHTTP or Routes method.
// It should be used from other handlers to render error pages.
type ErrorHandler struct {
	Log *slog.Logger
}

func NewErrorHandler(log *slog.Logger) *ErrorHandler {
	assert.Must(log != nil, "NewErrorHandler: log can't be nil")
	return &ErrorHandler{
		Log: log,
	}
}

// Serve 404 page.
func (handler *ErrorHandler) PageNotFound(w http.ResponseWriter, r *http.Request) error {
	title := "404 - Stránka nenalezena"
	code := http.StatusNotFound
	description := "Stránka nenalezena"
	message := "Vámi hledaná adresa: '" + r.Host + r.URL.Path + "' neexistuje. " +
		"Zkontrolujte zda je adresa správně zadaná a zkuste ji vyhledat znovu. " +
		"Případně se vraťte na předchozí stranu."
	return handler.ServeError(w, r, title, code, description, message)
}

// Serve 405 page.
func (handler *ErrorHandler) MethodNotAllowed(w http.ResponseWriter, r *http.Request) error {
	title := "405 - Metoda nepodporována"
	code := http.StatusMethodNotAllowed
	description := "Metoda nepodporována"
	message := "HTTP metoda: '" + r.Method + "' není podporována pro tento endpoint." +
		"Pokud vidíte tuto zprávu po odeslání formuláře, tak se může jednat o chybu aplikace." +
		"Prosím kontaktujte provozovatele stránek a popište co se stalo."
	return handler.ServeError(w, r, title, code, description, message)
}

// Serve 500 page.
func (handler *ErrorHandler) InternalServerError(w http.ResponseWriter, r *http.Request) error {
	// This will likely have special handling in the future, so don't use ServeError and handle the request directly
	title := "500 - Chyba"
	code := strconv.Itoa(http.StatusInternalServerError)
	description := "Chyba na straně serveru"
	message := "Omlouváme se, došlo k chybě a nebyli jsme schopni splnit váš požadavek. Zkuste to prosím později."
	data := components.NewErrorViewData(title, code, description, message)
	err := handler.View(w, r, data)
	if err != nil {
		return fmt.Errorf("NewErrorHandler.InternalServerError failed to render view; %w", err)
	}
	return nil
}

// Generic error handler.
// 'title' is optional and can be left empty.
// 'code' must be the http response code of the error.
// 'description' should contain human readable description of the error (Like: Page Not Found).
// 'message' short explanation of what happened and instructions on how to proceed (if possible) for user.
func (handler *ErrorHandler) ServeError(w http.ResponseWriter, r *http.Request, title string, code int, description, message string) error {
	var warnStatusCode *UnusualStatusCodeError
	if code < 200 || code > 599 {
		warnStatusCode = &UnusualStatusCodeError{
			Message: "ErrorHandler.ServeError received unusual error code, this may be bug pls fix!",
			Code:    code,
		}
	}
	strcode := strconv.Itoa(code)
	if title == "" {
		title = strcode + " - Chyba"
	}
	if description == "" {
		description = "Něco se pokazilo :("
	}
	data := components.NewErrorViewData(title, strcode, description, message)
	err := handler.View(w, r, data)
	if err != nil {
		return fmt.Errorf("NewErrorHandler.ServeError failed to render view; %w", err)
	}
	return warnStatusCode
}

func (handler *ErrorHandler) View(w http.ResponseWriter, r *http.Request, data *components.ErrorViewData) error {
	w.Header().Set(utils.ContentType, utils.TextHTML)
	return components.ErrorView(data).Render(r.Context(), w)
}

type UnusualStatusCodeError struct {
	Message string
	Code    int
}

func (err *UnusualStatusCodeError) Error() string {
	return err.Message + fmt.Sprintf("; status code: %d", err.Code)
}
