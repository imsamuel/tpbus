package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"tpbus/handlers"
	"tpbus/constants"
	"tpbus/scheduler"
)

var gates = constants.BusStopsAtGates
var callAndStoreEveryMin = scheduler.CallAndStoreEveryMin
var port string

func main() {
	// exit early if account key has not been set.
	if os.Getenv("ACCOUNT_KEY") == "" {
		panic("account key has to be set to access the bus arrival API")
	}

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}

	go callAndStoreEveryMin(gates.West.Code)
	go callAndStoreEveryMin(gates.OppWest.Code)
	go callAndStoreEveryMin(gates.Main.Code)
	go callAndStoreEveryMin(gates.OppMain.Code)
	go callAndStoreEveryMin(gates.East.Code)
	go callAndStoreEveryMin(gates.OppEast.Code)

	router := httprouter.New()
	router.GET("/", handlers.Index)
	router.GET("/services", handlers.GetServicesFromAllBusStops)
	router.GET("/services/:busStopLocation", handlers.GetServicesFromBusStop)
	router.GET("/services/:busStopLocation/:serviceNumber", handlers.GetServiceFromBusStop)

	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
