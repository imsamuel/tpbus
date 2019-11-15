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


// Tests the GetAService method on type Store.
func TestGetAService(t *testing.T) {
	appStore := store.New()
	appStore.WestGate = fakeServices

	targetServiceNumber := "15"

	got := appStore.GetAService("west-gate", targetServiceNumber)
	want := fakeService

	if !cmp.Equal(got, want) {
		t.Errorf("service not retrieved: got %+v want %+v", got, want)
	}
}

// Tests the GetAllServices method on type Store.
func GetAllServices(t *testing.T) {
	appStore := store.New()
	appStore.WestGate = fakeServices
	want := fakeServices
	got := appStore.GetAllServices("west-gate")
	if !cmp.Equal(got, want) {
		t.Errorf("services not retrieved: got %+v want %+v", got, want)
	}
}

// Tests the SetServices method on type Store.
func SetServices(t *testing.T) {
	appStore := store.New()
	appStore.SetServices("west-gate", fakeServices)
	got := appStore.GetAllServices("west-gate")
	want := fakeServices
	if !cmp.Equal(got, want) {
		t.Errorf("services not set: got %+v want %+v", got, want)
	}
}


