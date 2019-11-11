/*
Package services defines the structure of how the
data of bus services at a single bus stop is stored.
*/
package services

// Represents a single bus service.
type Service struct {
	ServiceNumber string
	NextBus string
	SubsequentBus string
	Destination string
}

// Represents all bus bus_services at a single bus stop.
type Services []Service

// Creates an instance of BusServices and returns a pointer to it.
func New() *Services { return &Services{} }
