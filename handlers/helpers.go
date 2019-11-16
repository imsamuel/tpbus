package handlers

import (
	"encoding/json"
	"fmt"
	"tpbus/constants"
)

// Defines the message sent as part of a HTTP response to a Not Found error.
type NotFoundMessage struct {
	ID string `json:"id"`
	Description string `json:"description"`
}

// Slice of valid bus stop locations.
var validBusStopLocations = []string{
	constants.BusStopsAtGates.West.Name,
	constants.BusStopsAtGates.OppWest.Name,
	constants.BusStopsAtGates.Main.Name,
	constants.BusStopsAtGates.OppMain.Name,
	constants.BusStopsAtGates.East.Name,
	constants.BusStopsAtGates.OppEast.Name,
}

// Slice of valid bus service numbers.
var validServiceNumbers = []string{
	"8", "15", "23", "69",
	"118", "118B", "129", "518",
}

/*
Generates a JSON-formatted value that is returned to the client in the case of
a Not Found error.
*/
func getNotFoundMessage(path string) ([]byte, error) {
	notFoundMessage := NotFoundMessage{
		ID:          "Resource not found",
		Description: fmt.Sprintf("API could not find a resource mapped to path %s", path),
	}
	marshalled, err := getPrettyPrint(notFoundMessage)
	return marshalled, err
}

/*
Takes any value(a struct will be passed to it for this app's scenario) and
marshals it to be JSON formatted which will have 4 spaces for indentation.
*/
func getPrettyPrint(v interface{}) ([]byte, error) {
	pretty, err :=  json.MarshalIndent(v, "", "    ")
	return pretty, err
}

// Returns a boolean depending on whether given bus stop location is valid.
func isBusStopLocationValid(busStopLocation string) bool {
	for _, v := range validBusStopLocations {
		if busStopLocation == v {
			return true
		}
	}
	return false
}

// Returns a boolean depending on whether given service number is valid.
func isServiceNumberValid(serviceNumber string) bool {
	for _, v := range validServiceNumbers {
		if serviceNumber == v {
			return true
		}
	}

	return false
}