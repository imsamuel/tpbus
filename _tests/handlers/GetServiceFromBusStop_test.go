package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
	"tpbus/handlers"
)

// Tests the GetServicesFromBusStop handler
func TestGetServiceFromBusStop(t *testing.T) {
	t.Run("invalid <busStopCode> parameter given", func(t *testing.T) {
		router := httprouter.New()
		router.GET(SERVICE_PATH, handlers.GetServiceFromBusStop)
		req, _ := http.NewRequest("GET", "/services/invalid-path/15", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assertResponseCode(t, rr.Code, http.StatusNotFound)
		assertResponseType(t, rr.Header().Get(handlers.CONTENT_TYPE), handlers.APPLICATION_JSON)
	})

	t.Run("invalid <serviceNumber> parameter given", func(t *testing.T) {
		router := httprouter.New()
		router.GET(SERVICE_PATH, handlers.GetServiceFromBusStop)
		req, _ := http.NewRequest("GET", "/services/east-gate/123123", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assertResponseCode(t, rr.Code, http.StatusNotFound)
		assertResponseType(t, rr.Header().Get(handlers.CONTENT_TYPE), handlers.APPLICATION_JSON)
	})

	t.Run("valid <busStopLocation> and <serviceNumber> parameters given", func(t *testing.T) {
		router := httprouter.New()
		router.GET(SERVICE_PATH, handlers.GetServiceFromBusStop)
		req, _ := http.NewRequest("GET", "/services/east-gate/15", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assertResponseCode(t, rr.Code, http.StatusOK)
		assertResponseType(t, rr.Header().Get(handlers.CONTENT_TYPE), handlers.APPLICATION_JSON)
	})
}
