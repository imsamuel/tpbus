/*
Package transform is responsible for the transformation of api.Response
instances into new model.Services instances.
*/
package transform

import (
	"tpbus/api"
	"tpbus/models"
)

// Utility map for getting new bus type from original response value.
var getBusType = map[string]string{
	"SD": "single deck",
	"DD": "double deck",
	"BD": "bendy",
}

// Utility map for getting a bool from whether a bus is wheelchair accessible.
var isWheelChairAccessible = map[string]bool{
	"WAB": false,
	"": true,
}

/*
responseToServices takes a api.Response instance and uses its values to
construct an instance of model.Services, and returns it.
*/
func ResponseToServices(response api.Response) models.Services {
	var newServices models.Services

	for _, respService := range response.Services {
		var newService models.Service

		newService.ServiceNumber = respService.ServiceNumber
		newService.DestinationCode = respService.NextBus.DestinationCode
		newService.NextBus = models.IncomingBus{
			EstimatedArrival:       respService.NextBus.EstimatedArrival,
			Type:                   getBusType[respService.NextBus.Type],
			IsWheelChairAccessible: isWheelChairAccessible[respService.NextBus.Feature],
		}
		newService.NextBus2 = models.IncomingBus{
			EstimatedArrival:       respService.NextBus2.EstimatedArrival,
			Type:                   getBusType[respService.NextBus2.Type],
			IsWheelChairAccessible: isWheelChairAccessible[respService.NextBus2.Feature],
		}
		newService.NextBus3 = models.IncomingBus{
			EstimatedArrival:       respService.NextBus3.EstimatedArrival,
			Type:                   getBusType[respService.NextBus3.Type],
			IsWheelChairAccessible: isWheelChairAccessible[respService.NextBus3.Feature],
		}

		newServices = append(newServices, newService)
	}

	return newServices
}
