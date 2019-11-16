package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Handler for route /services
func GetServicesFromAllBusStops(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	services := appStore.GetServicesAtAllBusStops()
	marshalledServices, err := getPrettyPrint(services)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalledServices)
}