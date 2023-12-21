package config

import "flag"

type Config struct {
	Addr      string
	DBUrl     string
	StaticDir string
}

func NewConfig() Config {
	var cfg Config

	// app
	flag.StringVar(&cfg.Addr, "addr", ":3000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.DBUrl, "db", "postgres://web:test@localhost:5432/snippetbox", "DataBase URL")

	flag.Parse()

	return cfg
}
