package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"golang.org/x/text/language"
)

type LanguageKeyType int

// The key used to store and obtain the language value from context.Context.
const LanguageKey = 0

var DefaultLanguageTag = language.AmericanEnglish

// Middleware that extracts user language preferences from requests and stores
// it in requests context.Context (Instead of echos context. This allows us to pass the context to templ components).
//
// By default, language preference is extracted from Accept-Language header,
// but it can be overridden with cookie for testing purposes.
//
// The cookie will be set, if there is a 'lang' path parameter and
func Language() echo.MiddlewareFunc {
	const aproxYear = int(365 * 24 * time.Hour)

	matcher := language.NewMatcher([]language.Tag{
		DefaultLanguageTag,
		language.Czech,
	})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			r := c.Request()

			// Get value from cookie if it exists
			cookieLang := ""
			cookie, err := r.Cookie("lang")
			if err == nil {
				cookieLang = cookie.Value
			}

			// Get value from query lang and set new cookie if it exists
			queryLang := c.QueryParam("lang")
			if queryLang == "reset" {
				// If lang is 'reset' than we delete the cookie and fall back to accept header
				cookieLang = ""
				c.SetCookie(&http.Cookie{
					Name:   "lang",
					Path:   "/",
					MaxAge: -1,
				})
			} else if queryLang != "" {
				cookieLang = queryLang
				c.SetCookie(&http.Cookie{
					Name:   "lang",
					Value:  queryLang,
					Path:   "/",
					MaxAge: aproxYear,
				})
				// TODO: If the cookie feature becomes available to normal users, then the cookie should be resend with every response to reset max-age.
			}

			acceptLang := r.Header.Get("Accept-Language")

			matchedLang, _ := language.MatchStrings(matcher, cookieLang, acceptLang)
			ctx := r.Context()
			newCtx := context.WithValue(ctx, LanguageKey, matchedLang)
			newRequest := r.WithContext(newCtx)
			c.SetRequest(newRequest)
			return next(c)
		}
	}
}
