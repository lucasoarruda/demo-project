package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lucasoarruda/demo-project/golang/internal/config"
)

func TestRoutes(t *testing.T) {
	// Create a new instance of AppConfig
	app := &config.AppConfig{}

	// Call the routes function with the app instance
	router := routes(app)

	// Create a new HTTP request and record the response
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/health", nil)

	// Call the router with the request and recorder
	router.ServeHTTP(w, req)

	// Assert that the response code is 200
	assert.Equal(t, http.StatusOK, w.Code)
}
