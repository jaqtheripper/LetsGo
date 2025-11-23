package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	// It is possible to include domain names in the path
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	// Paths that end with / are treated like wildcards, to prevent that, you would need to use the {$} pattern
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetCreatePost))
	// Wildcard statements would look like
	// mux.HandleFunc("products/{category}/item/{itemID}", productHandler)
	// There can be only one wildcard statement in a path segment (between slashes)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
