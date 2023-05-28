package handlers

import (
	"authenticationapi/internal/accounts"
	"encoding/json"
	"net/http"
)

type Auth struct {
	accs accounts.AccountsService
}

func (ah Auth) Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (ah Auth) Signup(w http.ResponseWriter, r *http.Request) {
    var (
        req *SignupRequest
        res *Response[*SignupResponse]
    )
    defer r.Body.Close()

    if err := json.NewDecoder(r.Body).Decode(req); err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }
    
}

func NewAuth(accs accounts.AccountsService) *Auth {
	return &Auth{
		accs: accs,
	}
}
