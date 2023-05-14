package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// Function will get the specified parameter and then convert it into integer and returns that.
func (app *application) ReadNthIDParam(r *http.Request, index int) (int64, error) {
	param := getParam(r, index)
	id, err := strconv.Atoi(param)
	if err != nil {
		return 0, errors.New("invalid id parameter")
	}
	return int64(id), nil
}

type envelope map[string]interface{}

// Define a helper function for sending responses.
// the parameters are the destination w, the HTTP status code, the data which is gonna be
// encoded, and a header map containing any additional HTTP headers we want to include in the response.
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// Encode the data to JSON, return error if there was one.
	// Use the MarshalIndent instead of Marshal to print the output json
	// in terminal pretier.
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Append a newline to make it easier to read in terminal.
	js = append(js, '\n')

	// add the headers input into response header map and if the headers input
	// is nil it's ok.
	for key, val := range headers {
		w.Header()[key] = val
	}

	// add the application/json type to Content-Type and set the status code.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
