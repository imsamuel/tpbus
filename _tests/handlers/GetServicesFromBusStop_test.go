package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
	"tpbus/handlers"
)

// Tests the GetServicesFromBusStop handler.
func TestGetServicesFromBusStop(t *testing.T) {
	t.Run("invalid <busStopCode> parameter given", func(t*testing.T) {
		router := httprouter.New()
		router.GET(BUS_STOP_LOCATION_PATH, handlers.GetServicesFromBusStop)
		req, _ := http.NewRequest("GET", "/services/some-invalid-value", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assertResponseCode(t, rr.Code, http.StatusNotFound)
		assertResponseType(t, rr.Header().Get(handlers.CONTENT_TYPE), handlers.APPLICATION_JSON)
	})

	t.Run("valid <busStopCode> parameter given", func(t *testing.T) {
		router := httprouter.New()
		router.GET(BUS_STOP_LOCATION_PATH, handlers.GetServicesFromBusStop)
		req, _ := http.NewRequest("GET", fmt.Sprintf("/services/%s", WEST_GATE), nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assertResponseCode(t, rr.Code, http.StatusOK)
		assertResponseType(t, rr.Header().Get(handlers.CONTENT_TYPE), handlers.APPLICATION_JSON)
	})
}
