package config

import "flag"

type Config struct {
	Port int
	Env  string
}

func InitConfig() *Config {
	var cfg Config

	flag.IntVar(&cfg.Port, "port", 4000, "Mock API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	return &cfg
}
