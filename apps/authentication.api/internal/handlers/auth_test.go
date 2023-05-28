package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSignup(t *testing.T) { 
	api := MustNewAPI(nil)
	api.SetupV1()
	
    tt := []struct{
        name string
        status int
        want any
        body *SignupRequest
    }{
        {
            name: "should return status 400 if required fields are missing",
            status: http.StatusBadRequest,
            want: &SignupResponse{},
            body: &SignupRequest{},
        },
    }

    for _, tc := range tt {
        t.Run(tc.name, func(t *testing.T) {
            byts, err := json.Marshal(tc.body)
            require.NoError(t, err)
            buff := bytes.NewBuffer(byts)
            req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/signup", buff)
            response := executeRequest(req, api)
            checkResponseCode(t, tc.status, response.Code)
        })
    }
} 
