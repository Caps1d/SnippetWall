package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// since this handles a subtree path, it will match any requests
	// that start with "/", hence we need to add a check to handle unwanted behaviour
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	log.Print("Home endpoint reached")

	files := []string{
		"./ui/html/pages/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	// go can use ParseFiles to read the template file into a template set
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	// we use ExecuteTemplate to write the content of the "base" template
	// from the template set into the response body. We have 4 templates in the template set:
	// "base", "title", "main", "nav" where "base" invokes the other 3
	// The last parameter represents any dynamic content
	// that we would like to pass to the template - will use it later
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// getting url query string parameters
	// we also want to make sure that the id is an int
	// we parse the str and convert it to an int
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		w.Header().Set("allow-id", ">=1")
		http.NotFound(w, r) // will send a 404 page not found response
		return
	}
	// write a formatted string response
	fmt.Fprintf(w, "Displaying snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// we can write "POST" or use constants from net/http
	if r.Method != http.MethodPost {
		// adds a custom header Allow: POST
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Request Method Not Allowed"))
		// A shortcut to this is the http.Error helper func
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
