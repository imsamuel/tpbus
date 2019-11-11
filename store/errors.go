package store

import "errors"

var (
	// ErrInvalidBusStopLocation is returned when bus stop location does not exist.
	ErrInvalidBusStopLocation = errors.New("bus stop location is invalid")
)
