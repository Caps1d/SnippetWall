package main

import (
	"flag"
	"log"
	"net/http"
	"path"

	"github.com/Caps1d/Lets-Go/internal/config"
)

func main() {
	// add command line flags
	// addr := flag.String("addr", ":4000", "HTTP network address")
	// Below I am storing the flag values in a config struct for convenience
	var cfg config.Config
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	// servemux is a router
	mux := http.NewServeMux()

	// added url path sanitation just in case
	fileServer := http.FileServer(http.Dir(path.Clean("./ui/static/")))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting servern on: %v", cfg.Addr)
	err := http.ListenAndServe(cfg.Addr, mux)
	log.Fatal(err)
}
