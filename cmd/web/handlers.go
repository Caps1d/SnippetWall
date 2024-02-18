package main

import (
	"encoding/json"
	"errors"

	// "html/template"
	"net/http"
	"strconv"

	"github.com/Caps1d/Lets-Go/internal/models"
)

func (app *applicaiton) home(w http.ResponseWriter, r *http.Request) {
	// since this handles a subtree path, it will match any requests
	// that start with "/", hence we need to add a check to handle unwanted behaviour
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	app.infoLog.Println("Home endpoint reached")

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, http.StatusOK, "home.tmpl.html", data)
}

func (app *applicaiton) snippetView(w http.ResponseWriter, r *http.Request) {

	// getting url query string parameters
	// we also want to make sure that the id is an int
	// we parse the str and convert it to an int
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		w.Header().Set("allow-id", ">=1")
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = s

	app.render(w, http.StatusOK, "view.tmpl.html", data)

	app.infoLog.Printf("Displaying snippet with ID %d...", id)
}

func (app *applicaiton) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// we can write "POST" or use constants from net/http
	if r.Method != http.MethodPost {
		// adds a custom header Allow: POST
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Request Method Not Allowed"))
		// A shortcut to this is the http.Error helper func -> using a custom helper now, defined in helpers.go
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// snippet struct
	var s models.Snippet

	// decode the requests body into our post struct declared as p
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Printf("Received post: %v", s)

	lastId, err := app.snippets.Insert(s.Title, s.Content, s.ExpiresInt)

	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New snippet created with id %d", lastId)

	w.Write([]byte("Create a new snippet..."))
}
