package main

import (
	"net/http"

	"github.com/asxraj/url-shortener/internal/database"
	"github.com/go-redis/redis/v8"
	"github.com/julienschmidt/httprouter"
)

func (app *application) resolveURL(w http.ResponseWriter, r *http.Request) {
	url := httprouter.ParamsFromContext(r.Context()).ByName("url")

	rdb := database.CreateClient()
	defer rdb.Close()

	val, err := rdb.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		app.errorResponseJSON(w, r, http.StatusNotFound, map[string]any{"error": "short url not found in the database"})
		return
	} else if err != nil {
		app.errorResponseJSON(w, r, http.StatusInternalServerError, "cannot connect to DB")
		return
	}

	app.writeJSON(w, http.StatusMovedPermanently, map[string]any{
		"status":   http.StatusMovedPermanently,
		"redirect": true,
		"to":       val,
	}, nil)
}
