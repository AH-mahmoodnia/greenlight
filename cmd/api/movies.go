package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new movies")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(getParam(r, 0))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
