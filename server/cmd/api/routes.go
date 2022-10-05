package main

import (
	"net/http"

	"github.com/asxraj/url-shortener/routes"
	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/:url", routes.ResolveURL)
	router.HandlerFunc(http.MethodPost, "/v1/shorten", routes.ShortenURL)

	return router
}
