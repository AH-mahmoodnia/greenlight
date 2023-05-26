package main

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

// a struct type which represent the resource for routing
// contains the method, regex which match the desired path and corresponding handlers.
type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

// Create route struct based on the input regex pattern.
func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

// Empty struct used as the key type in request contexts.
type ctxkey struct{}

// Return the ith elemnt of parameters returned from the regex matching.
func getParam(r *http.Request, index int) string {
	fields := r.Context().Value(ctxkey{}).([]string)
	return fields[index]
}

// Create a handler which use the above functions for creating the
// regex table and serve them in a closure.
func (app *application) routes() http.Handler {
	var routes []route = []route{
		newRoute(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler),
		newRoute(http.MethodPost, "/v1/movies", app.createMovieHandler),
		newRoute(http.MethodGet, "/v1/movies/(-?[0-9]+)", app.showMovieHandler),
		newRoute(http.MethodPatch, "/v1/movies/(-?[0-9]+)", app.partialyUpdateMovieHandler),
		newRoute(http.MethodDelete, "/v1/movies/(-?[0-9]+)", app.deleteMovieHandler),
		newRoute(http.MethodGet, "/v1/movies", app.listMoviesHandler),
	}
	router := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var allow []string
		for _, route := range routes {
			matches := route.regex.FindStringSubmatch(r.URL.Path)
			if len(matches) > 0 {
				if r.Method != route.method {
					allow = append(allow, route.method)
					continue
				}
				ctx := context.WithValue(r.Context(), ctxkey{}, matches[1:])
				route.handler(w, r.WithContext(ctx))
				return
			}
		}
		if len(allow) > 0 {
			w.Header().Set("Allow", strings.Join(allow, ", "))
			app.methodNotAllowedResponse(w, r)
			return
		}
		app.notFoundResponse(w, r)
	})
	return app.recoverPanic(router)
}
