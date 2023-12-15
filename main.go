package main

import (
	// "fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	log.Print("Home endpoint reached")
	// since this handles a subtree path, it will match any requests
	// that start with "/", hence we need to add a check to handle unwanted behaviour
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Get request to home endpoint"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// servemux is a router
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting servern on: 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
