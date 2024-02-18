package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func (app *applicaiton) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// shows where the error originated from by setting the frame depth to 2
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *applicaiton) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *applicaiton) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *applicaiton) render(w http.ResponseWriter, status int, page string, data *templateData) {
	// Retrieve the appropriate template set from the cache based on the page
	// name (like 'home.tmpl'). If no entry exists in the cache with the
	// provided name, then create a new error and call the serverError() helper
	// method that we made earlier and return.
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	// Write out the provided HTTP status code ('200 OK', '400 Bad Request', etc)
	// w.WriteHeader(status)
	// Execute the template set and write the response body. Again, if there
	// is any error we call the the serverError() helper
	// err := ts.ExecuteTemplate(w, "base", data) if err != nil {
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (app *applicaiton) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
	}
}
