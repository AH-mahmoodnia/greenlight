package main

import (
	"encoding/json"
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

	// Using the json.Marshal to encode the data into json and return it as a []byte.
	// In case of error we log it and send the client a generic error message.
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	// Append a newline character to the encoded json []byte
	// just for better view in terminal
	js = append(js, '\n')

	// Set the Content-Type header to application/json so the
	// browser will know the type of response is json.
	// the default value is "Content-Type: text/plain; charset=utf-8"
	w.Header().Set("Content-Type", "application/json")

	// write the JSON as the HTTP response body.
	w.Write(js)
}
