package main

import (
	"net/http"
)

// Declare a handler method over application struct which return the status, environment
// and version of api and for now it's plain-text.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a map which holds the information we want to send in the response.
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	if err := app.writeJSON(w, http.StatusOK, data, nil); err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request.", http.StatusInternalServerError)
	}
}
