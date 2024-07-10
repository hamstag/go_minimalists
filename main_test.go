package main

import (
	"context"
	"go-minimalists/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, s *app.App) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router().ServeHTTP(rr, req)

	return rr
}

// checkResponseCode is a simple utility to check the response code
// of the response
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestSabaideeHamstag(t *testing.T) {
	// Create a New Server Struct
	s := app.NewApp(context.Background())

	// Create a New Request
	req, _ := http.NewRequest("GET", "/sabaidee", nil)

	// Execute Request
	response := executeRequest(req, s)

	// Check the response code
	checkResponseCode(t, http.StatusOK, response.Code)

	// We can use testify/require to assert values, as it is more convenient
	require.Equal(t, "sabaidee hamstag", response.Body.String())
}
