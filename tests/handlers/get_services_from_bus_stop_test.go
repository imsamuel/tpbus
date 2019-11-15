package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
	"tpbus/constants"
	"tpbus/handlers"
)

// Path to retrieve the bus service timings from a specified bus stop.
var busStopLocationPath = "/services/:busStopLocation"

var westGate = constants.BusStopsAtGates.West.Name

func assertResponseCode(t *testing.T, receivedResponseCode, expectedResponseCode int) {
	if receivedResponseCode != expectedResponseCode {
		t.Errorf(
			"expected a %d response code but received a %d",
			expectedResponseCode,
			receivedResponseCode,
		)
	}
}

func assertResponseType(t *testing.T, receivedResponseType, expectedResponseType string) {
	if receivedResponseType != expectedResponseType {
		t.Errorf(
			"expected the type of data in response to be %q but it is of type %q",
			receivedResponseType,
			expectedResponseType)
	}
}

// Tests the GetServices handler.
func TestGetServicesFromBusStop(t *testing.T) {
	t.Run("invalid <busStopCode> parameter given", func(t*testing.T) {
		router := httprouter.New()
		router.GET(busStopLocationPath, handlers.GetServicesFromBusStop)
		req, _ := http.NewRequest("GET", "/services/some-invalid-value", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assertResponseCode(t, rr.Code, http.StatusNotFound)
		assertResponseType(t, rr.Header().Get("Content-Type"), "application/json")
	})

	t.Run("ServiceNumber query param not given", func(t *testing.T) {
		router := httprouter.New()
		router.GET("/services/:busStopLocation", handlers.GetServicesFromBusStop)
		req, _ := http.NewRequest("GET", fmt.Sprintf("/services/%s", westGate), nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assertResponseCode(t, rr.Code, http.StatusOK)
		assertResponseType(t, rr.Header().Get("Content-Type"), "application/json")
	})
}
