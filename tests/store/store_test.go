package store

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"tpbus/models"
)

// An instance of models.Services filled with dummy data.
var fakeServices = getServices()

// Tests the GetAService method on type Store.
func TestGetAService(t *testing.T) {
	appStore := newStore()
	appStore.WestGate = fakeServices

	targetServiceNumber := "15"

	var want models.Service
	for _, service := range fakeServices {
		if service.ServiceNumber == targetServiceNumber {
			want = service
		}
	}

	got := appStore.GetAService("west-gate", targetServiceNumber)

	if !cmp.Equal(got, want) {
		t.Errorf("service not retrieved: got %+v want %+v", got, want)
	}
}

// Tests the GetAllServices method on type Store.
func GetAllServices(t *testing.T) {
	appStore := newStore()
	appStore.WestGate = fakeServices
	want := fakeServices
	got := appStore.GetAllServices("west-gate")
	if !cmp.Equal(got, want) {
		t.Errorf("services not retrieved: got %+v want %+v", got, want)
	}
}

// Tests the SetServices method on type Store.
func SetServices(t *testing.T) {
	appStore := newStore()
	appStore.SetServices("west-gate", fakeServices)
	got := appStore.GetAllServices("west-gate")
	want := fakeServices
	if !cmp.Equal(got, want) {
		t.Errorf("services not set: got %+v want %+v", got, want)
	}
}


