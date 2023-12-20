package main

import (
	"net/http"
	"path"
)

func (app *applicaiton) routes() *http.ServeMux {
	// servemux is a router
	mux := http.NewServeMux()

	// app struct contains config with static dir path
	fileServer := http.FileServer(http.Dir(path.Clean(app.cfg.StaticDir)))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
