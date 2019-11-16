package handlers

import (
	"testing"
)

/*
Fails the test if response status code is not equal to the expected status
code.
*/
func assertResponseCode(t *testing.T, receivedResponseCode, expectedResponseCode int) {
	if receivedResponseCode != expectedResponseCode {
		t.Errorf(
			"expected a %d response code but received a %d",
			expectedResponseCode,
			receivedResponseCode,
		)
	}
}

/*
Fails the test if the response Content-Type header's value is not equal to the
expected content type.
*/
func assertResponseType(t *testing.T, receivedResponseType, expectedResponseType string) {
	if receivedResponseType != expectedResponseType {
		t.Errorf(
			"expected the type of data in response to be %q but it is of type %q",
			receivedResponseType,
			expectedResponseType)
	}
}