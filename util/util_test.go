package util

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"tpbus-backend/api"
	"tpbus-backend/models"
)

func TestGetBusStopLocationFromBusStopCode(t *testing.T) {
	testCases := []struct{
		busStopCode string // expected result
		busStopLocation string  // input
	}{
		{"75249", "west"},
		{"75241", "opp-west"},
		{"75239", "main"},
		{"75231", "opp-main"},
		{"75229", "east"},
		{"75221", "opp-east"},
	}

	for _, tt := range testCases {
		got := BusStopCodeToBusStopName[tt.busStopCode]
		want := tt.busStopLocation
		if got != want {
			t.Errorf("got %q want %q given %s\n", got, want, tt.busStopCode)
		}
	}
}

func TestGetDestinationFromDestinationCode(t *testing.T) {
	testCases := []struct{
		destinationCode string // input
		destination string // expected result
	}{
		{"75009", "Tamp Int"},
		{"52009", "Toa Payoh Int"},
		{"77009", "Pasir Ris Int"},
		{"65009", "Punggol Temp Int"},
		{"97009", "Changi Biz Pk Ter"},
		{"75019", "Tamp Conc Int"},
		{"52499", "St. Michael's Ter"},
	}

	for _, tt := range testCases {
		got := DestinationCodeToDestination[tt.destinationCode]
		want := tt.destination
		if got != want {
			t.Errorf("got %q want %q given %s\n", got, want, tt.destinationCode)
		}
	}
}

func TestTransformAPIResponse(t *testing.T) {
	mockAPIResponse := &api.Response{
		Services: []api.Service{
			api.Service{
				ServiceNumber: "15",
				NextBus: api.IncomingBusDetails{
					DestinationCode:  "77009",
					EstimatedArrival:  "2017-06-05T14:46:27+08:00",
				},
				SubsequentBus: api.IncomingBusDetails{
					DestinationCode:  "77009",
					EstimatedArrival:  "2017-06-05T14:57:09+08:00",
				},
			},
		},
	}

	want := models.BusServices{
		models.BusService{
			ServiceNumber: "15",
			NextBus:       "2017-06-05T14:46:27+08:00",
			SubsequentBus: "2017-06-05T14:57:09+08:00",
			Destination:   "Pasir Ris Int",
		},
	}

	got := TransformAPIResponse(mockAPIResponse)

	if !cmp.Equal(got, want) {
		t.Error("TransformAPIResponse did not convert response struct correctly\n")
	}
}
