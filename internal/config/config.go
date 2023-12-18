package config

import "flag"

type Config struct {
	Addr      string
	StaticDir string
}

func NewConfig() Config {
	var cfg Config

	// app
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	return cfg
}
