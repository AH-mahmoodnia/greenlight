package main

import (
	"errors"
	"net/http"
	"strconv"
)

func (app *application) ReadNthIDParam(r *http.Request, index int) (int, error) {
	param := getParam(r, index)
	id, err := strconv.Atoi(param)
	if err != nil {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}