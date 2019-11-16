package store

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"testing"
	"tpbus/models"
	"tpbus/store"
)

// Reads a JSON file and marshals its contents into a models.Service struct.
func getService() models.Service {
	newService := models.Service{} // will get marshalled into
	raw, _ := ioutil.ReadFile("../test_data/service.json")
	json.Unmarshal(raw, &newService)
	return newService
}

// Reads a JSON file and marshals its contents into a models.Services struct.
func getServices() models.Services {
	newServices := models.Services{} // will get marshalled into
	raw, _ := ioutil.ReadFile("../test_data/services.json")
	json.Unmarshal(raw, &newServices)
	return newServices
}

// An instance of models.Service filled with dummy test data.
var fakeService = getService()

// An instance of models.Services filled with dummy test_data.
var fakeServices = getServices()


// Tests the GetServiceFromOneBusStop method on type Store.
func TestGetServiceFromOneBusStop(t *testing.T) {
	appStore := store.New()
	appStore.WestGate = fakeServices

	targetServiceNumber := "15"

	got := appStore.GetServiceFromOneBusStop("west-gate", targetServiceNumber)
	want := fakeService

	if !cmp.Equal(got, want) {
		t.Errorf("service not retrieved: got %+v want %+v", got, want)
	}
}

// Tests the GetServicesFromOneBusStop method on type Store.
func TestGetServicesFromOneBusStop(t *testing.T) {
	appStore := store.New()
	appStore.WestGate = fakeServices
	want := fakeServices
	got := appStore.GetServicesFromOneBusStop("west-gate")
	if !cmp.Equal(got, want) {
		t.Errorf("services from %s not retrieved: got %+v want %+v","west-gate", got, want)
	}
}

// Tests the GetServicesFromAllBusStops method on type Store.
func TestGetServicesFromAllBusStops(t *testing.T) {
	appStore := store.New()
	appStore.WestGate = fakeServices
	appStore.OppWestGate = fakeServices
	appStore.MainGate = fakeServices
	appStore.OppMainGate = fakeServices
	appStore.EastGate = fakeServices
	appStore.OppEastGate = fakeServices
	got := *appStore
	want := appStore.GetServicesAtAllBusStops()
	if !cmp.Equal(got, want) {
		t.Errorf("services from all bus stops not retrieved: got %+v want %+v", got, want)
	}
}

// Tests the SetServices method on type Store.
func TestSetServices(t *testing.T) {
	appStore := store.New()
	appStore.SetServices("west-gate", fakeServices)
	got := appStore.GetServicesFromOneBusStop("west-gate")
	want := fakeServices
	if !cmp.Equal(got, want) {
		t.Errorf("services not set: got %+v want %+v", got, want)
	}
}


