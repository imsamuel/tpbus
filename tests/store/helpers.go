package store

import (
	"encoding/json"
	"io/ioutil"
	"tpbus/models"
	"tpbus/store"
)

// Initializes an instance of store.Store and returns a pointer to it. (helper)
func newStore() *store.Store { return &store.Store{} }

// Reads a JSON file and marshals its contents into a models.Services struct.
func getServices() models.Services {
	newServices := models.Services{} // will get marshalled into
	raw, _ := ioutil.ReadFile("./data/example_services.json")
	json.Unmarshal(raw, &newServices)
	return newServices
}