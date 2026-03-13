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
			// TODO: Add special handling for more client errors (400, 401, 413, 414, etc.). Other server errors are usually produced by proxy in situations where responding is not possible anyway.
			cErr = errorHandler.ServeError(c.Response(), c.Request(), nil, code, nil, &components.Translations{
				Czech:   "Zkuste to prosím později.",
				English: "Please try again later.",
			})
		}

		var uce *UnusualStatusCodeError
		if errors.As(cErr, &uce) {
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
	title := &components.Translations{
		Czech:   "404 - Stránka nenalezena",
		English: "404 - Page Not Found",
	}
	code := http.StatusNotFound
	description := &components.Translations{
		Czech:   "Stránka nenalezena",
		English: "Page Not Found",
	}
	message := &components.Translations{
		Czech: "Vámi hledaná stránka: '" + r.Host + r.URL.Path + "' neexistuje. " +
			"Zkontrolujte zda je adresa správně zadaná a zkuste ji vyhledat znovu. " +
			"Případně se vraťte na předchozí stranu.",
		English: "The page you are looking for: '" + r.Host + r.URL.Path + "' does not exist. " +
			"Please check that you entered the correct address and try again. " +
			"Otherwise return to previous page.",
	}
	return handler.ServeError(w, r, title, code, description, message)
}

// Serve 405 page.
func (handler *ErrorHandler) MethodNotAllowed(w http.ResponseWriter, r *http.Request) error {
	title := &components.Translations{
		Czech:   "405 - Metoda nepodporována",
		English: "405 - Method Not Allowed",
	}
	code := http.StatusMethodNotAllowed
	description := &components.Translations{
		Czech:   "Metoda nepodporována",
		English: "Method Not Allowed",
	}
	message := &components.Translations{
		Czech: "HTTP metoda: '" + r.Method + "' není podporována pro tento endpoint. " +
			"Pokud vidíte tuto zprávu po odeslání formuláře, tak se může jednat o chybu aplikace. " +
			"Prosím kontaktujte provozovatele stránek a popište co se stalo.",
		English: "The HTTP method: '" + r.Method + "' is not supported for this endpoint. " +
			"If you see this message after submitting a form then it may be error in the application. " +
			"Please contact us and describe what happened.",
	}
	return handler.ServeError(w, r, title, code, description, message)
}

// Serve 500 page.
func (handler *ErrorHandler) InternalServerError(w http.ResponseWriter, r *http.Request) error {
	// This will likely have special handling in the future, so don't use ServeError and handle the request directly
	title := &components.Translations{
		Czech:   "500 - Chyba",
		English: "500 - Error",
	}
	code := strconv.Itoa(http.StatusInternalServerError)
	description := &components.Translations{
		Czech:   "Chyba na straně serveru",
		English: "Server error",
	}
	message := &components.Translations{
		Czech:   "Omlouváme se, došlo k chybě a nebyli jsme schopni splnit váš požadavek. Zkuste to prosím později.",
		English: "There was an error and we could't complete your request. Please try again later.",
	}
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
func (handler *ErrorHandler) ServeError(w http.ResponseWriter, r *http.Request, title *components.Translations, code int, description, message *components.Translations) error {
	// Don't dop this. This will mean that the return value of this function is always non nil, but also not really.
	// https://www.adityathebe.com/golang-nil-interface-check/
	// https://go.dev/doc/faq#nil_error
	//var warnStatusCode *UnusualStatusCodeError

	var warnStatusCode error
	if code < 200 || code > 599 {
		warnStatusCode = &UnusualStatusCodeError{
			Message: "ErrorHandler.ServeError received unusual error code, this may be bug pls fix!",
			Code:    code,
		}
	}
	strcode := strconv.Itoa(code)
	if title == nil {
		title = &components.Translations{
			Czech:   strcode + " - Chyba",
			English: strcode + " - Error",
		}
	}
	if description == nil {
		description = &components.Translations{
			Czech:   "Něco se pokazilo",
			English: "Something went wrong",
		}
	}
	w.Header().Set(utils.ContentType, utils.TextHTML)
	w.WriteHeader(code)
	data := components.NewErrorViewData(title, strcode, description, message)
	err := handler.View(w, r, data)
	if err != nil {
		return fmt.Errorf("NewErrorHandler.ServeError failed to render view; %w", err)
	}
	return warnStatusCode
}

func (handler *ErrorHandler) View(w http.ResponseWriter, r *http.Request, data *components.ErrorViewData) error {
	return components.ErrorView(data).Render(r.Context(), w)
}

type UnusualStatusCodeError struct {
	Message string
	Code    int
}

func (err *UnusualStatusCodeError) Error() string {
	return err.Message + fmt.Sprintf("; status code: %d", err.Code)
}
