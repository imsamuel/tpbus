// Package store exposes a Store struct which has accessor and mutator methods.
package store

import (
	"tpbus/constants"
	"tpbus/store/services"
)

// Assigning the struct of gates to a variable(cutting down name space).
var gates = constants.BusStopsAtGates

type store struct {
	WestGate *services.Services
	OppWestGate *services.Services
	MainGate *services.Services
	OppMainGate *services.Services
	EastGate *services.Services
	OppEastGate *services.Services
}

func New() *store {
	return &store{
		WestGate:    services.New(),
		OppWestGate: services.New(),
		MainGate:    services.New(),
		OppMainGate: services.New(),
		EastGate:    services.New(),
		OppEastGate: services.New(),
	}
}

// Gets the data of the bus services at the specified bus stop.
func (s *store) GetServices(busStopLocation string) (services.Services, error) {
	switch busStopLocation {
	case gates.West.Name:
		return *s.WestGate, nil
	case gates.OppWest.Name:
		return *s.OppWestGate, nil
	case gates.Main.Name:
		return *s.MainGate, nil
	case gates.OppMain.Name:
		return *s.OppMainGate, nil
	case gates.East.Name:
		return *s.EastGate, nil
	case gates.OppEast.Name:
		return *s.OppEastGate, nil
	default:
		return nil, ErrInvalidBusStopLocation
	}
}

// Sets the specified Store field with the given data of bus services.
func (s *store) SetServices(busStopLocation string, services services.Services) error {
	switch busStopLocation {
	case gates.West.Name:
		*s.WestGate = services
		return nil
	case gates.OppWest.Name:
		*s.OppWestGate = services
		return nil
	case gates.Main.Name:
		*s.MainGate = services
		return nil
	case gates.OppMain.Name:
		*s.OppMainGate = services
		return nil
	case gates.East.Name:
		*s.EastGate = services
		return nil
	case gates.OppEast.Name:
		*s.OppEastGate = services
		return nil
	default:
		return ErrInvalidBusStopLocation
	}
}

// Exported store
var Store = New()
