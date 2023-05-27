package handlers

import (
	"authenticationapi/internal/accounts"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func executeRequest(req *http.Request, api *API) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	require.Equal(t, expected, actual, fmt.Sprintf("expected response code %d, got %d", expected, actual))
}

func TestHeartbeat(t *testing.T) {
	accs := accounts.NewAccountService()

	api := MustNewAPI(accs)
	api.SetupV1()

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/heartbeat", nil)

	response := executeRequest(req, api)

	checkResponseCode(t, http.StatusOK, response.Code)
}
