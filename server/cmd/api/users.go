package main

import (
	"net/http"

	"github.com/asxraj/url-shortener/internal/models"
	"github.com/asxraj/url-shortener/internal/validator"
)

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := app.readJSON(w, r, &user)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	if models.ValidateUser(v, &user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = user.Set()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.models.Users.Insert(&user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, map[string]any{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
