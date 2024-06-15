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

	fs := flag.NewFlagSet("config", flag.ExitOnError)

	// app
	fs.StringVar(&cfg.Addr, "addr", ":8080", "HTTP network address")
	fs.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	fs.StringVar(&cfg.DBUrl, "db", "postgres://postgres:test@localhost:5432/snippetwall", "DataBase URL")

	flag.Parse()

	// cfg.TLS = &tls.Config{
	// 	CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	// }
	//
	return cfg
}
