package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"tpbus/handlers"
	"tpbus/constants"
	"tpbus/scheduler"
)

var gates = constants.BusStopsAtGates
var callAndStoreEveryMin = scheduler.CallAndStoreEveryMin

func main() {
	// exit early if account key has not been set.
	if os.Getenv("ACCOUNT_KEY") == "" {
		panic("account key has to be set to access the bus arrival API")
	}

	go callAndStoreEveryMin(gates.West.Code)
	go callAndStoreEveryMin(gates.OppWest.Code)
	go callAndStoreEveryMin(gates.Main.Code)
	go callAndStoreEveryMin(gates.OppMain.Code)
	go callAndStoreEveryMin(gates.East.Code)
	go callAndStoreEveryMin(gates.OppEast.Code)

	router := httprouter.New()
	router.GET("/services", handlers.GetServicesFromAllBusStops)
	router.GET("/services/:busStopLocation", handlers.GetServicesFromBusStop)
	router.GET("/services/:busStopLocation/:serviceNumber", handlers.GetServiceFromBusStop)
	router.NotFound = http.HandlerFunc(handlers.HandleNotFound)

	http.ListenAndServe(":8080", router)
}
