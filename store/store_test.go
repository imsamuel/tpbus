package store

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"tpbus-backend/constants"
	"tpbus-backend/store2/services"
)

// Cutting down the name space.
var busStops = constants.BusStopsAtGates

// Tests the getServices method on type Store.
func TestGetServices(t *testing.T) {
	t.Run("passing invalid bus stop location", func(t *testing.T) {
		// Creating an instance of store to test its getServices method.
		store := New()
		// Passing a bus stop location that does not exist, hoping for error.
		_, err := store.GetServices("south-east")

 		if err != ErrInvalidBusStopLocation {
 			t.Errorf("expected ErrInvalidBusStopLocation but got none")
		}
	})

	t.Run("expecting no error for valid bus stop locations", func(t *testing.T) {
		// Creating an instance of store to test its getServices method.
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

		// Ranging over the valid inputs and hoping for no error.
		for _, location := range validBusStopLocations {
			_, err := store.GetServices(location)
			if err != nil {
				t.Errorf("expected no error from input %q but got one", location)
			}
		}
	})

	t.Run("testing if bus services retrieved", func(t *testing.T) {
		// Creating an instance of store to test its getServices method.
		store := New()

		// Creating an instance of services.Services to use as data.
		fakeBusServices := &services.Services{
			services.Service{
				ServiceNumber: "15",
				NextBus:       "2017-06-05T14:46:27+08:00",
				SubsequentBus: "2017-06-05T14:57:09+08:00",
				Destination:   "Pasir Ris Int",
			},
		}

		// Assigning the test data to the MainGate field of store.
		store.MainGate = fakeBusServices

		// Trying to get bus services of the bus stop at TP's main gate.
		retrieved, _ := store.GetServices("main")

		// Hoping for retrieved value to be what was inserted into the Main field.
		if !cmp.Equal(retrieved, *fakeBusServices) {
			t.Errorf("getServices did not retrieve the correct data")
		}
	})
}

// Tests the setServices method on type Store.
func TestSetServices(t *testing.T) {
	t.Run("expecting error for invalid bus stop location", func(t *testing.T) {
		// Creating an instance of store to test its setServices method.
		store := New()
		// Passing a bus stop location that does not exist, hoping for error.
		err := store.SetServices("south-east", services.Services{})

		if err != ErrInvalidBusStopLocation {
			t.Errorf("expected ErrInvalidBusStopLocation but got none")
		}
	})

	t.Run("expecting no error for valid bus stop locations", func(t *testing.T) {
		// Creating an instance of store to test its setServices method.
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

		// Ranging over the valid inputs and hoping for no error.
		for _, location := range validBusStopLocations {
			err := store.SetServices(location, *services.New())
			if err != nil {
				t.Errorf("expected no error from input %q but got one", location)
			}
		}
	})

	t.Run("setting all store fields", func(t *testing.T) {
		// Creating an instance of store to test its setServices method.
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
				NextBus:       "2017-06-05T14:46:27+08:00",
				SubsequentBus: "2017-06-05T14:57:09+08:00",
				Destination:   "Pasir Ris Int",
			},
		}

		// Setting every field of Store to fakeBusServices.
		for _, location := range validBusStopLocations {
			_ = store.SetServices(location, *fakeBusServices)
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