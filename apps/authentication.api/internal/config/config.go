package config

import "os"

type Config struct {
	Port string
    TokenSecret string
}

const (
	DefaultPort = "3001"
    DefaultTokenSecret = "qwertyuiop"
)

func New() (*Config, error) {
	port := DefaultPort
    secret := DefaultTokenSecret

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	if os.Getenv("TOKEN_SECRET") != "" {
		secret = os.Getenv("TOKEN_SECRET")
	}

	return &Config{
		Port: port,
        TokenSecret: secret,
	}, nil
}
