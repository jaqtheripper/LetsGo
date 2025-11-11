package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	files := []string{
		"ui/html/base.tmpl.html",
		"ui/html/pages/home.tmpl.html",
		"ui/html/partials/nav.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID: %d", id)
	// msg := fmt.Sprintf("Display a specific snippet with ID: %d", id)
	// w.Write([]byte(msg))
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(201)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Create a new snippet..."))
}

// func productHandler(w http.ResponseWriter, r *http.Request) {
// 	category := r.PathValue("category")
// 	itemID := r.PathValue("itemID")
// }
