package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Handler for route /services/<busStopLocation>/<serviceNumber>.
// Returns the timings of the specified bus service at specified bus stop.
func GetServiceFromBusStop(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enableCors(&w)

	// Check if <busStopLocation> argument is valid
	busStopLocation := ps.ByName("busStopLocation")
	if !isBusStopLocationValid(busStopLocation) {
		WriteNotFoundMessage(w, r.URL.Path)
		return
	}

	// Check if <serviceNumber> argument is valid
	serviceNumber := ps.ByName("serviceNumber")
	if !isServiceNumberValid(serviceNumber) {
		WriteNotFoundMessage(w, r.URL.Path)
		return
	}

	service := appStore.GetServiceFromOneBusStop(busStopLocation, serviceNumber)
	marshalledService, _ := getPrettyPrint(service)

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Write(marshalledService)
}
