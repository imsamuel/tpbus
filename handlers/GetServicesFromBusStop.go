package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"tpbus/store"
)

var appStore = store.Value

// Handler for route /services/<bus-stop-location> | accepts ServiceNumber query param
func GetServicesFromBusStop(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract the bus stop location provided from path parameter
	busStopLocation := ps.ByName("busStopLocation")

	/*
	Check if a valid bus stop location has been provided.
	If not, return an error message as part of the HTTP response.
	*/
	if !isBusStopLocationValid(busStopLocation) {
		errorMsg, err := getNotFoundMessage(r.URL.Path)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusNotFound)
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.Write(errorMsg)
		return
	}

	/*
	Check if valid ServiceNumber query parameter provided.
	If so, return to client only the timings of the specified bus service.
	*/
	serviceNumber := r.URL.Query().Get("ServiceNumber")
	if isServiceNumberValid(serviceNumber) {
		service := appStore.GetServiceFromOneBusStop(busStopLocation, serviceNumber)
		marshalledService, err := getPrettyPrint(service)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.Write(marshalledService)
		return
	}

	// Retrieve the timings of all bus services of specified bus stop.
	services := appStore.GetServicesFromOneBusStop(busStopLocation)
	marshalledServices, err := getPrettyPrint(services)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Write(marshalledServices)
}

/*
Note:
- getNotFoundMessage called as early as possible to panic if error produced.
  Panicking because if an error is produced it's a development issue.
*/