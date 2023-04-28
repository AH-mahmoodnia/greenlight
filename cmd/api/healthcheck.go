package main

import (
	"fmt"
	"net/http"
)

// Declare a handler method over application struct which return the status, environment
// and version of api and for now it's plain-text.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a fixed-format JSON response from a string.
	js := `{"status": "available", "environment": %q, "version": %q}`
	fmt.Sprintf(js, app.config.env, version)

	// Set the Content-Type header to application/json so the
	// browser will know the type of response is json.
	// the default value is "Content-Type: text/plain; charset=utf-8"
	w.Header().Set("Content-Type", "application/json")

	// write the JSON as the HTTP response body.
	w.Write([]byte(js))
}
