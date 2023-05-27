package config

import "os"

type Config struct {
	Port string
}

const (
	DefaultPort = "3001"
)

func New() (*Config, error) {
	port := DefaultPort

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	return &Config{
		Port: port,
	}, nil
}
