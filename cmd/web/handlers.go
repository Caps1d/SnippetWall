package main

import (
	"encoding/json"
	"errors"

	// "html/template"
	"net/http"
	"strconv"

	"github.com/Caps1d/Lets-Go/internal/models"
	"github.com/julienschmidt/httprouter"
)

func (app *applicaiton) home(w http.ResponseWriter, r *http.Request) {
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
	// getting Clean URL parameters from request context
	// we also want to make sure that the id is an int
	// we parse the str and convert it to an int
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
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
	w.Write([]byte("Display the form for creating a new snippet..."))
}
func (app *applicaiton) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
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
