package main

import (
	"authenticationapi/internal/accounts"
	"authenticationapi/internal/config"
	"authenticationapi/internal/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewMux()
	c, err := config.New()
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
	if err := run(r, c); err != nil {
		log.Print(err.Error())
	}
}

func run(r *chi.Mux, c *config.Config) error {
	accs := accounts.NewAccountService()
	api := handlers.MustNewAPI(accs)
	v1, _ := api.SetupV1()
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", c.Port),
		Handler: v1,
	}

	return srv.ListenAndServe()
}
