package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Handler for route /services
func GetServicesFromAllBusStops(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enableCors(&w)

	services := appStore.GetServicesAtAllBusStops()
	marshalledServices, _ := getPrettyPrint(services)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalledServices)
}