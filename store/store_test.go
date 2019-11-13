package store

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"tpbus/constants"
	"tpbus/store/services"
)

// Cutting down the name space.
var busStops = constants.BusStopsAtGates

// Tests the GetAService method on type Store.
func TestGetAService(t *testing.T) {
	t.Run("testing if bus service retrieved", func(t *testing.T) {
		// Creating an instance of store to test its GetAService method.
		store := New()

		fakeBusService := services.Service{
			ServiceNumber: "15",
			DestinationCode: "97009",
			NextBus:       "2017-06-05T14:46:27+08:00",
			NextBus2: "2017-06-05T14:57:09+08:00",
			NextBus3: "2017-06-05T14:57:09+08:00",
			Type: "single deck",
			IsWheelChairAccessible: true,
		}

		// Assigning test data to the MainGate field of store.
		store.MainGate = &services.Services{
			fakeBusService,
		}

		// Trying to get details of bus 15 at bus stop at TP's main gate.
		retrieved := store.GetAService("main-gate", "15")

		// Hoping for retrieved value to be what was inserted into the Main field.
		if !cmp.Equal(retrieved, fakeBusService) {
			t.Errorf("getServices did not retrieve the correct data")
		}
	})
}

// Tests the GetAllServices method on type Store.
func TestGetAllServices(t *testing.T) {
	t.Run("testing if bus services retrieved", func(t *testing.T) {
		// Creating an instance of store to test its GetAllServices method.
		store := New()

		// Creating an instance of services.Services to use as data.
		fakeBusServices := &services.Services{
			services.Service{
				ServiceNumber: "15",
				DestinationCode: "97009",
				NextBus:       "2017-06-05T14:46:27+08:00",
				NextBus2: "2017-06-05T14:57:09+08:00",
				NextBus3: "2017-06-05T14:57:09+08:00",
				Type: "single deck",
				IsWheelChairAccessible: true,
			},
		}

		// Assigning the test data to the MainGate field of store.
		store.MainGate = fakeBusServices

		// Trying to get bus services of the bus stop at TP's main gate.
		retrieved := store.GetAllServices("main-gate")

		// Hoping for retrieved value to be what was inserted into the Main field.
		if !cmp.Equal(retrieved, *fakeBusServices) {
			t.Errorf("getServices did not retrieve the correct data")
		}
	})
}

// Tests the SetServices method on type Store.
func TestSetServices(t *testing.T) {
	t.Run("setting all store fields", func(t *testing.T) {
		// Creating an instance of store to test its SetServices method.
		store := New()

		// Creating a slice to hold all valid bus stop locations.
		validBusStopLocations := []string{
			busStops.West.Name,
			busStops.OppWest.Name,
			busStops.Main.Name,
			busStops.OppMain.Name,
			busStops.East.Name,
			busStops.OppEast.Name,
		}

		// Creating an instance of services.Services to use as data.
		fakeBusServices := &services.Services{
			services.Service{
				ServiceNumber: "15",
				DestinationCode: "97009",
				NextBus:       "2017-06-05T14:46:27+08:00",
				NextBus2: "2017-06-05T14:57:09+08:00",
				NextBus3: "2017-06-05T14:57:09+08:00",
				Type: "single deck",
				IsWheelChairAccessible: true,
			},
		}

		// Setting every field of Store to fakeBusServices.
		for _, location := range validBusStopLocations {
			store.SetServices(location, *fakeBusServices)
		}

		// Creating a slice of store fields with their fields names to range over.
		storeFields := []struct{
			name string
			value services.Services
		}{
			{"WestGate", *store.WestGate},
			{"OppWestGate", *store.OppWestGate},
			{"MainGate", *store.MainGate},
			{"OppMainGate", *store.OppMainGate},
			{"EastGate", *store.EastGate},
			{"OppEastGate", *store.OppEastGate},
		}

		// Hoping that all store fields have been set from the previous for loop.
		for _, storeField := range storeFields {
			if !cmp.Equal(storeField.value, *fakeBusServices) {
				t.Errorf("field %q not set", storeField.name)
			}
		}
	})
}