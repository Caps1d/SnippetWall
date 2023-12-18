package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
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
