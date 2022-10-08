package main

import "net/http"

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {

	data := map[string]any{"status": "available", "system-info": map[string]any{
		"version":     version,
		"environment": app.config.env,
	}}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
