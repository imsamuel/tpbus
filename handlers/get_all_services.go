package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetAllServices(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get a JSON-encoded value of all bus services.
	JSON, _ := json.MarshalIndent(appStore, "", "    ") // no error handling needed; valid arg

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(JSON)
}