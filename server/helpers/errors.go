package helpers

import (
	"net/http"
)

func ErrorResponseJSON(w http.ResponseWriter, r *http.Request, status int, message any) {

	wrap := wrapper{
		"error": message,
	}

	err := WriteJSON(w, status, wrap, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}
