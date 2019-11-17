package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"tpbus/store"
)

var appStore = store.Value

// Handler for route /services/<bus-stop-location>
func GetServicesFromBusStop(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enableCors(&w)

	// Extract the bus stop location provided from path parameter
	busStopLocation := ps.ByName("busStopLocation")

	/*
	Check if a valid bus stop location has been provided.
	If not, return an error message as part of the HTTP response.
	*/
	if !isBusStopLocationValid(busStopLocation) {
		WriteNotFoundMessage(w, r.URL.Path)
		return
	}

	// Retrieve the timings of all bus services of specified bus stop.
	services := appStore.GetServicesFromOneBusStop(busStopLocation)
	marshalledServices, _ := getPrettyPrint(services)

	w.WriteHeader(http.StatusOK)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Write(marshalledServices)
}

/*
Note:
- getNotFoundMessage called as early as possible to panic if error produced.
  Panicking because if an error is produced it's a development issue.
*/