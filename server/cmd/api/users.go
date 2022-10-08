package main

import (
	"net/http"
)

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Firstname string `json:"first_name"`
		Lastname  string `json:"last_name"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, map[string]any{"user": input}, nil)
}
