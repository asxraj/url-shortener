package main

import (
	"net/http"
)

func (app *application) errorResponseJSON(w http.ResponseWriter, r *http.Request, status int, errorJson any) {

	err := app.writeJSON(w, status, errorJson, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}
