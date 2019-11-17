/*
Package scheduler contains a function which makes a call to the BusArrival
API every minute and stores the result in the app's in memory store.
*/
package scheduler

import (
	"log"
	"time"
	"tpbus/api"
	"tpbus/constants"
	"tpbus/store"
	"tpbus/transform"
)

// Cutting down name space.
var appStore = store.Value

// Details of the bus stops at Temasek Poly. (cutting down namespace)
var gates = constants.BusStopsAtGates

// Maps the bus stop code of bus stops to the bus stop's name.
var busStopCodeToBusStopName = map[string]string{
	gates.West.Code: gates.West.Name,
	gates.OppWest.Code: gates.OppWest.Name,
	gates.Main.Code: gates.Main.Name,
	gates.OppMain.Code: gates.OppMain.Name,
	gates.East.Code: gates.East.Name,
	gates.OppEast.Code: gates.OppEast.Name,
}

/*
Calls the BusArrival API every minute and stores the result in the app's
in-memory store.
*/
func CallAndStoreEveryMin(busStopCode string) {
	ticker := time.NewTicker(60 * time.Second)
	for {
		select {
		case <- ticker.C:
			respStruct, err := api.GetBusServicesFromAPI(busStopCode) // GetBusServicesFromAPI SHOULD NOT BE RETURNING A NIL POINTER
			if err != nil {
				log.Printf("error occured in GetBusServicesFromAPI: %v", err.Error())
			}

			services := transform.ResponseToServices(respStruct)

			busStopLocation := busStopCodeToBusStopName[busStopCode]
			appStore.SetServices(busStopLocation, services)
		}
	}
}