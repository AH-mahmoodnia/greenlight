package main

import (
	"fmt"
	"net/http"
)

// Declare a handler method over application struct which return the status, environment
// and version of api and for now it's plain-text.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
