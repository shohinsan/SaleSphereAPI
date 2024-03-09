package web

import (
	"context"
	"net/http"

	"github.com/go-json-experiment/json"
)

func Respond(ctx context.Context, w http.ResponseWriter, data any, statusCode int) error {

	SetStatusCode(ctx, statusCode)

	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	// Set the status code for the request logger middleware.
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
