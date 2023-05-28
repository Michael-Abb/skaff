package handlers

import (
	"authenticationapi/internal/accounts"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type API struct {
	r   *chi.Mux
	acc accounts.AccountsService
}

func (a *API) SetupV1() (*chi.Mux, error) {
	a.r.Use(middleware.Logger)
	a.r.Use(middleware.RequestID)
	a.r.Use(middleware.RealIP)
	a.r.Use(middleware.Logger)
	a.r.Use(middleware.Recoverer)
    a.r.Use(render.SetContentType(render.ContentTypeJSON))

	ah := NewAuth(a.acc)

	a.r.Route("/api/v1", func(r chi.Router) {
		r.Get("/heartbeat", a.heartbeat)
        r.Route("/auth", func(r chi.Router) {
		    r.Post("/login", ah.Login)
            r.Post("/signup", ah.Signup)
        })
	})

	return a.r, nil
}

func (a *API) heartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func MustNewAPI(accs accounts.AccountsService) *API {
	return &API{
		r:   chi.NewMux(),
		acc: accs,
	}
}
