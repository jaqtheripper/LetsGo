package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	// It is possible to include domain names in the path
	mux.HandleFunc("GET /{$}", app.home)
	// Paths that end with / are treated like wildcards, to prevent that, you would need to use the {$} pattern
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	// Wildcard statements would look like
	// mux.HandleFunc("products/{category}/item/{itemID}", productHandler)
	// There can be only one wildcard statement in a path segment (between slashes)
	return mux
}
