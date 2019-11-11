package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"tpbus/store"
)

var appStore = store.Store

// Handler for route /services/<bus-stop-location> | Accepts a BusStopCode param
func GetServices(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract the bus stop location provided from path parameter
	busStopLocation := ps.ByName("busStopLocation")

	// Retrieve the slice of bus bus_services from store.
	services, err := appStore.GetServices(busStopLocation)
	if err != nil {
		panic("passed invalid bus stop location")
	}

	// Get a JSON-encoded value of the slice of bus_services.
	JSON, _ := json.MarshalIndent(services, "", "    ") // no error handling needed; valid arg

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(JSON)
}