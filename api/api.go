package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// Represents the JSON response of a call to the BusArrival API as a Go ds.
type Response struct {
	Services []Service `json:"Services"`
}

type Service struct {
	ServiceNumber string `json:"ServiceNo"`
	NextBus IncomingBusDetails `json:"NextBus"`
	SubsequentBus IncomingBusDetails `json:"NextBus2"`
}

type IncomingBusDetails struct {
	DestinationCode string `json:"DestinationCode"`
	EstimatedArrival string `json"EstimatedArrival"`
}

/*
Creates and returns an instance of http.Request configured with:
- the URL of the BusArrival API
- the account key needed to access the API
*/
func newRequest(busStopCode string) (*http.Request, error) {
	baseURL := "http://datamall2.mytransport.sg/ltaodataservice/BusArrivalv2?BusStopCode="
	req, err := http.NewRequest("GET", baseURL + busStopCode, nil)
	req.Header.Set("AccountKey", os.Getenv("ACCOUNT_KEY"))
	return req, err
}

/*
Takes a pointer to http.Request and makes a HTTP request with it.
Returns the response's body or an error if one occurs.
*/
func makeRequest(r *http.Request) (io.ReadCloser, error) {
	client := &http.Client{}

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

/*
Decodes a response body into a Response instance and returns it.
*If an error is produced by (*Decoder).Decode, return a nil
 pointer to Response and the error, else return the set
 pointer to Response and a nil error.
*/
func decodeToStruct(rc io.ReadCloser) (*Response, error) {
	r := &Response{}

	err := json.NewDecoder(rc).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

/*
Makes a call to the BusArrival API, decodes the response body into an instance
of Response and returns a pointer to the instance.
*/
func GetBusServicesFromAPI(busStopCode string) (*Response, error) {
	// Error is ignored because err will never be produced, values are hardcoded
	request, _ := newRequest(busStopCode)

	// If issue occurs during API call, just return a nil pointer with error.
	respBody, err := makeRequest(request)
	if err != nil {
		return nil, err
	}

	/*
		Error is ignored because it will never be returned from decodeToStruct.
		For more info:
			- decodeToStruct calls (*Decoder).Decode which only returns an
		      error(UnsupportedTypeError) if passed a channel OR complex
		      OR function value(because they cannot be encoded in JSON).
		    - A io.ReadCloser value is neither of the three thus it is decode-able.
	*/
	decoded, _ := decodeToStruct(respBody)

	return decoded, nil
}

