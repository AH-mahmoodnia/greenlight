package main

import (
	"fmt"
	"net/http"
)

// Declare a handler method over application struct which return the status, environment
// and version of api and for now it's plain-text.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status": "available", "environment": %q, "version": %q}`
	fmt.Sprintf(js, app.config.env, version)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}
