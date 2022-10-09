package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheck)

	// URL
	router.HandlerFunc(http.MethodGet, "/v1/resolve/:url", app.resolveURL)
	router.HandlerFunc(http.MethodPost, "/v1/shorten", app.shortenURL)

	// USERS
	router.HandlerFunc(http.MethodPost, "/v1/users/register", app.registerUser)
	router.HandlerFunc(http.MethodPost, "/v1/users/login", app.loginUser)

	return app.recoverPanic(app.logRequest(app.enableCORS(app.authenticate(router))))
}
