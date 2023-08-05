package main

import (
	"github.com/lucasoarruda/demo-project/golang/internal/config"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestMainFunction(t *testing.T) {
	// Save the current value of GO_PORT
	oldPort := os.Getenv("GO_PORT")

	// Set GO_PORT to a test value
	os.Setenv("GO_PORT", ":8080")

	// Call the main function
	go main()

	// Restore the original value of GO_PORT
	os.Setenv("GO_PORT", oldPort)
}

func TestMain(t *testing.T) {
	// Create a new instance of AppConfig
	app := &config.AppConfig{}

	// Call the main function with the app instance
	go main()

	// Wait for the server to start up
	time.Sleep(1 * time.Second)

	// Make a GET request to the server
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
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

	// Stop the server
	resp, err := http.DefaultClient.Get("http://localhost:8000")
	if err != nil {
		t.Fatalf("Error stopping server: %s", err)
	}
	defer resp.Body.Close()
	// Assert that the response has a 200 status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}
}
