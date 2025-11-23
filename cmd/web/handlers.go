package main

import (
	"errors"
	"fmt"
	"jaqtheripper/LetsGo/internal/models"
	"jaqtheripper/LetsGo/internal/validator"
	"net/http"
	"strconv"
)

type snippetCreateForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Server", "Go")
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	// for _, snippet := range snippets {
	// 	fmt.Fprintf(w, "%+v\n", snippet)
	// }
	// files := []string{
	// 	"ui/html/base.tmpl.html",
	// 	"ui/html/pages/home.tmpl.html",
	// 	"ui/html/partials/nav.tmpl.html",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }
	// data := templateData{
	// 	Snippets: snippets,
	// }
	// err = ts.ExecuteTemplate(w, "base", data)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// }

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl.html", data)

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// }
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// fmt.Fprintf(w, "Display a specific snippet with ID: %d...\n", id)
	// msg := fmt.Sprintf("Display a specific snippet with ID: %d", id)
	// w.Write([]byte(msg))
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}
	// files := []string{
	// 	"./ui/html/base.tmpl.html",
	// 	"./ui/html/partials/nav.tmpl.html",
	// 	"./ui/html/pages/view.tmpl.html",
	// }
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }
	// data := templateData{
	// 	Snippet: snippet,
	// }
	// err = ts.ExecuteTemplate(w, "base", data)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// }

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.tmpl.html", data)

	// fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, r, http.StatusOK, "create.tmpl.html", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	var form snippetCreateForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// err = app.formDecoder.Decode(&form, r.PostForm)
	// if err != nil {
	// 	app.clientError(w, http.StatusBadRequest)
	// }

	// expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	// if err != nil {
	// 	app.clientError(w, http.StatusBadRequest)
	// }

	// form := snippetCreateForm{
	// 	Title:   r.PostForm.Get("title"),
	// 	Content: r.PostForm.Get("content"),
	// 	Expires: expires,
	// }

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7, or 365")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl.html", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	// w.WriteHeader(http.StatusCreated)
	// w.Write([]byte("Create a new snippet..."))
	app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

// func productHandler(w http.ResponseWriter, r *http.Request) {
// 	category := r.PathValue("category")
// 	itemID := r.PathValue("itemID")
// }
