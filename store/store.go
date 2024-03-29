// Package store exposes a Store struct which has accessor and mutator methods.
package store

import (
	"tpbus/constants"
	"tpbus/models"
)

type Store models.BusStops

// Assigning the struct of gates to a variable(cutting down name space).
var gates = constants.BusStopsAtGates

// Initializes a pointer to a Store instance.
func New() *Store { return &Store{} }

// Gets the test_data of a single bus service at specified bus stop.
func (s *Store) GetServiceFromOneBusStop(busStopLocation string, serviceNumber string) models.Service {
	var services models.Services

	switch busStopLocation {
	case gates.West.Name:
		services = s.WestGate
	case gates.OppWest.Name:
		services = s.OppWestGate
	case gates.Main.Name:
		services = s.MainGate
	case gates.OppMain.Name:
		services = s.OppMainGate
	case gates.East.Name:
		services = s.EastGate
	case gates.OppEast.Name:
		services = s.OppEastGate
	}

	for _, service := range services {
		if service.ServiceNumber == serviceNumber {
			return service
		}
	}

	return models.Service{} // to satisfy the compulsory return at top scope
}

// Gets the timings of the bus services at the specified bus stop.
func (s *Store) GetServicesFromOneBusStop(busStopLocation string) models.Services {
	var services models.Services

	switch busStopLocation {
	case gates.West.Name:
		services = s.WestGate
	case gates.OppWest.Name:
		services = s.OppWestGate
	case gates.Main.Name:
		services = s.MainGate
	case gates.OppMain.Name:
		services = s.OppMainGate
	case gates.East.Name:
		services = s.EastGate
	case gates.OppEast.Name:
		services = s.OppEastGate
	}

	return services
}

// Gets the timings of the bus services at all bus stops.
func (s *Store) GetServicesAtAllBusStops() Store { return *s }

// Sets the specified Store field with the given test_data of bus services.
func (s *Store) SetServices(busStopLocation string, services models.Services) {
	switch busStopLocation {
	case gates.West.Name:
		s.WestGate = services
	case gates.OppWest.Name:
		s.OppWestGate = services
	case gates.Main.Name:
		s.MainGate = services
	case gates.OppMain.Name:
		s.OppMainGate = services
	case gates.East.Name:
		s.EastGate = services
	case gates.OppEast.Name:
		s.OppEastGate = services
	}
}

// Exported store
var Value = New()
