package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	// It is possible to include domain names in the path
	mux.HandleFunc("GET /{$}", home)
	// Paths that end with / are treated like wildcards, to prevent that, you would need to use the {$} pattern
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	// Wildcard statements would look like
	// mux.HandleFunc("products/{category}/item/{itemID}", productHandler)
	// There can be only one wildcard statement in a path segment (between slashes)

	log.Printf("Starting server on 127.0.0.1:4000")

	err := http.ListenAndServe("127.0.0.1:4000", mux)
	log.Fatal(err)
}
