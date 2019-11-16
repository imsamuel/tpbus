package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
	"tpbus/handlers"
)

// Tests the GetServicesFromAllBusStops handler
func TestGetServicesFromAllBusStops(t *testing.T) {
	router := httprouter.New()
	router.GET(SERVICES_PATH, handlers.GetServicesFromAllBusStops)
	req, _ := http.NewRequest("GET", "/services", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assertResponseCode(t, rr.Code, http.StatusOK)
	assertResponseType(t, rr.Header().Get(handlers.CONTENT_TYPE), handlers.APPLICATION_JSON)
}
