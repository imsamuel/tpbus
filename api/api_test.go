package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
)

/*
The bus stop code at the bus stop located at TP's west gate.
Made package-level for reuse across across test functions.
*/
var busStopCodeAtWestGate = "75249"

// Tests the newRequest function
func TestNewRequest(t *testing.T) {
	t.Run("expecting no error for valid parameters", func(t *testing.T) {
		_, err := newRequest(busStopCodeAtWestGate)
		if err != nil {
			t.Errorf("expected no error but got one: %s\n", err.Error())
		}
	})

	t.Run("checking if account key attached to Request", func(t *testing.T) {
		req, _ := newRequest(busStopCodeAtWestGate)

		/*
			Note: req.Header.Get(key) returns the value associated
			      with the specified key, else returns "".
		*/
		if req.Header.Get("AccountKey") == "" {
			t.Errorf("expected account key to be set on Request but it was not")
		}
	})
}

// Tests the decodeToStruct function
func TestDecodeToStruct(t *testing.T) {
	t.Run("should be able to be decoded into Response", func(t *testing.T) {
		dummyResponse := Response{
			Services: []Service{
				{
					ServiceNumber: "15",
					NextBus:       IncomingBusDetails{
						DestinationCode:  "77009",
						EstimatedArrival: "2017-06-05T14:46:27+08:00",
					},
					SubsequentBus: IncomingBusDetails{
						DestinationCode:  "77009",
						EstimatedArrival: "2017-06-05T14:57:09+08:00",
					},
				},
			},
		}

		// Converting the Response struct to an io.ReadCloser
		toBytes, _ := json.Marshal(dummyResponse) // convert to bytes
		byteReader := bytes.NewReader(toBytes) // convert bytes to io.Reader
		readCloser := ioutil.NopCloser(byteReader) // wrapped with ioutil.NopCloser

		// get the struct representation of readCloser
		_, err := decodeToStruct(readCloser)

		if err != nil {
			t.Errorf("not expecting error but got one: %q", err.Error())
		}
	})

	t.Run("should not be able to be decoded into Response", func(t *testing.T) {
		stringReader := strings.NewReader("hello world")
		invalidReadCloser := ioutil.NopCloser(stringReader)

		_, err := decodeToStruct(invalidReadCloser)

		if err == nil {
			t.Errorf("expected an error but did not get one")
		}
	})
}