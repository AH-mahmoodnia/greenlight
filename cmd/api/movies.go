package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AH-mahmoodnia/greenlight/internal/data"
)

// a movie handler for POST /v1/movies endpoint.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new movies")
}

// Shows the details of the specified movie in the GET /v1/movies/:id endpoint.
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.ReadNthIDParam(r, 0)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romanc", "war"},
		Version:   1,
	}
	if err := app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
