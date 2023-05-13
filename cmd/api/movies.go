package main

import (
	"fmt"
	"net/http"
)

// a movie handler for POST /v1/movies endpoint.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new movies")
}

// Shows the details of the specified movie in the GET /v1/movies/:id endpoint.
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// use the ReadNthIDParam func to read the 0th integer id in the path."
	id, err := app.ReadNthIDParam(r, 0)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// interpolate the movie ID in a placeholder response.
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
