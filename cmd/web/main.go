package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Caps1d/Lets-Go/internal/config"
)

// application struct for dependency injection
type applicaiton struct {
	cfg      *config.Config
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	// add command line flags
	// addr := flag.String("addr", ":4000", "HTTP network address")
	// Below I am storing the flag values in a config struct for convenience
	cfg := config.NewConfig()

	// levelled logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// app struct
	app := &applicaiton{
		cfg:      &cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	// initialize a new Server struct which containing our config
	// this is how we hadle errors with errorLog
	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting servern on %v", cfg.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
