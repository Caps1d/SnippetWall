package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Caps1d/Lets-Go/internal/config"
	"github.com/Caps1d/Lets-Go/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

// application struct for dependency injection
type applicaiton struct {
	cfg      *config.Config
	snippets *models.SnippetModel
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	// add command line flags
	// addr := flag.String("addr", ":4000", "HTTP network address")
	// below I am storing the flag values in a config struct for convenience
	cfg := config.NewConfig()

	// levelled logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(cfg.DBUrl)
	if err != nil {
		errorLog.Fatal(err)
	}
	infoLog.Print("DB connection established...")
	defer db.Close()

	// app struct
	app := &applicaiton{
		cfg:      &cfg,
		snippets: &models.SnippetModel{DB: db},
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
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	// db connection
	conn, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return conn, nil
}
