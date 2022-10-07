package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/:url", app.resolveURL)
	router.HandlerFunc(http.MethodPost, "/v1/shorten", app.shortenURL)

	return app.enableCORS(router)
}
