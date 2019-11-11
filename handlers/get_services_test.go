package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Testing the GetServices handler
func TestGetServices(t *testing.T) {
	t.Run("status code", func(t *testing.T) {
		router := httprouter.New()
		router.GET("/services/:busStopLocation", GetServices)

		req, _ := http.NewRequest("GET", "/services/west", nil)
		rr := httptest.NewRecorder()

		// Mocking a GET request to /services with a path parameter of "west"
		router.ServeHTTP(rr, req)

		// Check if correct status code received
		received := rr.Code
		expected := http.StatusOK

		if received != expected {
			t.Errorf("received a %d status code but expected %d", received, expected)
		}
	})

	t.Run("content type", func(t *testing.T) {
		router := httprouter.New()
		router.GET("/services/:busStopLocation", GetServices)

		req, _ := http.NewRequest("GET", "/services/west", nil)
		rr := httptest.NewRecorder()

		// Mocking a GET request to /services with a path parameter of "west"
		router.ServeHTTP(rr, req)

		// Check if expected content type received
		received := rr.Header().Get("Content-Type")
		expected := "application/json"

		if received != expected {
			t.Errorf("type of content received is %q but expected %q", received, expected)
		}
	})
}

