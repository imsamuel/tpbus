package transform

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"testing"
	"tpbus/api"
	"tpbus/models"
	"tpbus/transform"
)

// Tests the ResponseToServices function.
func TestResponseToServices(t *testing.T) {
	var response api.Response
	raw, _ := ioutil.ReadFile("../test_data/response.json")
	json.Unmarshal(raw, &response)

	var transformed models.Services
	raw, _ = ioutil.ReadFile("../test_data/transformed.json")
	json.Unmarshal(raw, &transformed)

	want := transformed
	got := transform.ResponseToServices(response)

	if !cmp.Equal(got, want) {
		t.Errorf("response not converted to services: \ngot %+v \nwant %+v", got, want)
	}
}
