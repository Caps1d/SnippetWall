package main

import (
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *applicaiton) routes() http.Handler {
	// servemux is a router
	// mux := http.NewServeMux()

	router := httprouter.New()
	// enable redirecting with trailing slashes
	router.RedirectTrailingSlash = true

	// create a handler func which wraps around our notFound helper and assign
	// as the custom handler for 404 not found responses
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})
	// app struct contains config with static dir path
	fileServer := http.FileServer(http.Dir(path.Clean(app.cfg.StaticDir)))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
	router.Handler(http.MethodGet, "/snippet/create", dynamic.ThenFunc(app.snippetCreate))
	router.Handler(http.MethodPost, "/snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	// using justinas/alice package to manage middleware chains
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
