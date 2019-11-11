/*
Package util contains utility helper maps/functions that are used across other
sub-packages.
*/
package util

import (
	"tpbus/api"
	"tpbus/constants"
	"tpbus/store/services"
)

// Details of the bus stops at Temasek Poly. (cutting down namespace)
var gates = constants.BusStopsAtGates

// Details of the bus stops the buses start from/emd at. (cutting down namespace)
var serviceConstants = constants.BusStopsOfServices

// Maps the bus stop code of bus stops to the bus stop's name.
var BusStopCodeToBusStopName = map[string]string{
	gates.West.Code: gates.West.Name,
	gates.OppWest.Code: gates.OppWest.Name,
	gates.Main.Code: gates.Main.Name,
	gates.OppMain.Code: gates.OppMain.Name,
	gates.East.Code: gates.East.Name,
	gates.OppEast.Code: gates.OppEast.Name,
}

// Maps the destination code of a bus stop to the location.
var DestinationCodeToDestination = map[string]string{
	serviceConstants.TampinesInt.Code: serviceConstants.TampinesInt.Name,
	serviceConstants.ToaPayohInt.Code: serviceConstants.ToaPayohInt.Name,
	serviceConstants.PasirRisInt.Code: serviceConstants.PasirRisInt.Name,
	serviceConstants.PunggolTempInt.Code: serviceConstants.PunggolTempInt.Name,
	serviceConstants.ChangiBusinessPkTer.Code: serviceConstants.ChangiBusinessPkTer.Name,
	serviceConstants.TampinesConcourseInt.Code: serviceConstants.TampinesConcourseInt.Name,
	serviceConstants.StMichaelTer.Code: serviceConstants.StMichaelTer.Name,
	serviceConstants.OppITECol.Code: serviceConstants.OppITECol.Name,
	serviceConstants.BefPunggolRd.Code: serviceConstants.BefPunggolRd.Name,
}

/*
Returns an instance of models.BusServices populated with the values from the
passed in struct (in which the body of the API response got marshalled into).
*/
func TransformAPIResponse(responseStruct api.Response) services.Services {
	/*
	Initializing an empty instance of models.BusServices which will be
	appended with the newly populated instances of model.BusService.
	(see for loop)
	*/
	var busServices services.Services

	for _, service := range responseStruct.Services {
		/*
		Initializing an empty instance of models.BusService is going to be
		populated with the data from the struct in which the API response
		got decoded into.
		*/
		var busService services.Service

		busService.ServiceNumber = service.ServiceNumber
		busService.NextBus = service.NextBus.EstimatedArrival
		busService.SubsequentBus = service.SubsequentBus.EstimatedArrival

		destinationCode := service.NextBus.DestinationCode
		busService.Destination = DestinationCodeToDestination[destinationCode]

		// Appending the populated models.BusService instance to busServices.
		busServices = append(busServices, busService)
	}

	return busServices
}
