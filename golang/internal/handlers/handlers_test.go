package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStatus(t *testing.T) {
	// Create a new instance of the gin.Engine
	router := gin.New()

	// Set up the handler for the /health route
	Repo = &Repository{}
	router.GET("/health", Repo.Status)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder
	res := httptest.NewRecorder()

	// Call the router.ServeHTTP function
	router.ServeHTTP(res, req)

	// Assert that the response recorder's status code and body match the expected values
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	expectedBody := `{"status":200}`
	if res.Body.String() != expectedBody {
		t.Errorf("Expected body %s but got %s", expectedBody, res.Body.String())
	}
}
