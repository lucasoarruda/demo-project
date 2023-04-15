package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	app.Port = ":8080" // set port for testing

	// create a new instance of the server
	server := httptest.NewServer(routes(&app))
	defer server.Close()

	// make a GET request to the server
	resp, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatalf("Error making GET request: %s", err)
	}
	defer resp.Body.Close()

	// assert that the response has a 200 status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// assert that the response body is not empty
	buf := make([]byte, 1024)
	n, _ := resp.Body.Read(buf)
	if n == 0 {
		t.Errorf("Expected non-empty response body")
	}
}
