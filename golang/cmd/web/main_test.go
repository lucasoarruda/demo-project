package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/lucasoarruda/demo-project/golang/internal/config"
)

var startMain sync.Once

// TestMainFunction starts main() once (it blocks until SIGINT, so it cannot be
// started twice without racing on the global app config) and verifies that it
// honours GO_PORT and serves the homepage.
func TestMainFunction(t *testing.T) {
	t.Setenv("GO_PORT", ":18080")

	startMain.Do(func() { go main() })

	// Wait for the server to start up
	var resp *http.Response
	var err error
	deadline := time.Now().Add(10 * time.Second)
	for time.Now().Before(deadline) {
		resp, err = http.Get("http://localhost:18080/")
		if err == nil {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
	if err != nil {
		t.Fatalf("server did not start on GO_PORT: %s", err)
	}
	defer func() { _ = resp.Body.Close() }()

	// Assert that the response has a 200 status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestHomepage(t *testing.T) {
	// Create a new instance of AppConfig
	app := &config.AppConfig{}

	// Make a GET request to the router
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	router := routes(app)
	router.ServeHTTP(w, req)

	// Assert that the response has a 200 status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	// Assert that the response body is not empty
	if w.Body.Len() == 0 {
		t.Errorf("Expected non-empty response body")
	}
}
