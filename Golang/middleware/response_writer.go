package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/api-sekejap/internal/entity/base"
)

func GenericMiddleware(next func(http.ResponseWriter, *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the handler and capture its response and error.
		response, err := next(w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the response data to JSON.
		mapData := base.Response[interface{}]{
			Data:    response,
			Success: err == nil,
		}
		jsonResponse, err := json.Marshal(mapData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to indicate JSON response.
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response.
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse) //nolint:all
	}
}
