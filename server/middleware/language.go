package middleware

import (
	"context"

	"github.com/labstack/echo/v5"
	"golang.org/x/text/language"
)

type LanguageKeyType int

// The key used to store and obtain the language value from context.Context.
const LanguageKey = 0

var DefaultLanguageTag = language.AmericanEnglish

// Middleware that extracts user language preferences from requests and stores
// it in requests context.Context (Instead of echos context. This allows us to pass the context to templ components).
func Language() echo.MiddlewareFunc {
	matcher := language.NewMatcher([]language.Tag{
		DefaultLanguageTag,
		language.Czech,
	})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			r := c.Request()
			acceptLang := r.Header.Get("Accept-Language")
			matchedLang, _ := language.MatchStrings(matcher, acceptLang)
			ctx := r.Context()
			newCtx := context.WithValue(ctx, LanguageKey, matchedLang)
			newRequest := r.WithContext(newCtx)
			c.SetRequest(newRequest)
			return next(c)
		}
	}
}
