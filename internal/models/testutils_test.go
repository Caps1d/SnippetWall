package models

import (
	"context"
	"os"
	"testing"

	"github.com/Caps1d/Lets-Go/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func newTestDB(t *testing.T) *pgxpool.Pool {
	cfg := config.NewConfig()
	cfg.DBUrl = "postgres://test_web:test@localhost:5432/test_snippetbox"
	db, err := pgxpool.New(context.Background(), cfg.DBUrl)
	if err != nil {
		t.Fatal(err)
	}

	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(context.Background(), string(script))
	if err != nil {
		t.Fatal(err)
	}
	// Use the t.Cleanup() to register a function *which will automatically be
	// called by Go when the current test (or sub-test) which calls newTestDB()
	// has finished*.
	t.Cleanup(func() {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}

		_, err = db.Exec(context.Background(), string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
	})
	return db
}
