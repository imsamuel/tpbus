/*
Package scheduler contains a function which makes a call to the BusArrival
API every minute and stores the result in the app's in memory store.
*/
package scheduler

import (
	"log"
	"time"
	"tpbus/api"
	"tpbus/store"
	"tpbus/util"
)

// Cutting down name space.
var appStore = store.Store

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
				*respStruct = api.Response{}
			}

			transformed := util.TransformAPIResponse(*respStruct)

			busStopLocation := util.BusStopCodeToBusStopName[busStopCode]
			appStore.SetServices(busStopLocation, transformed)
		}
	}
}