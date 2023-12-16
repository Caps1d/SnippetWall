package main

import (
	"log"
	"net/http"
	"path"
)

func main() {
	// servemux is a router
	mux := http.NewServeMux()

	// added url path sanitation just in case
	fileServer := http.FileServer(http.Dir(path.Clean("./ui/static/")))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting servern on: 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
