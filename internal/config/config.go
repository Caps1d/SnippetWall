package config

import (
	"crypto/tls"
	"flag"
)

type Config struct {
	Addr      string
	DBUrl     string
	StaticDir string
	TLS       *tls.Config
}

func NewConfig() Config {
	var cfg Config

	// app
	flag.StringVar(&cfg.Addr, "addr", ":3000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.DBUrl, "db", "postgres://web:test@localhost:5432/snippetbox", "DataBase URL")

	flag.Parse()

	cfg.TLS = &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	return cfg
}
