package handlers

import (
	"authenticationapi/internal/accounts"
	"net/http"
)

type Auth struct {
	accs accounts.AccountsService
}

func (ah Auth) Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func NewAuth(accs accounts.AccountsService) *Auth {
	return &Auth{
		accs: accs,
	}
}
