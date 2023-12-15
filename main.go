package main

import (
	// "fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	log.Print("Home endpoint reached")
	w.Write([]byte("Get request to home endpoint"))
}

func main() {
	// servemux is a router
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	log.Print("Starting servern on: 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
