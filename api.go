package hljson

import (
	"encoding/json"
)

var (
	// Base api url
	ApiBase = "http://hackerlist.net"

	// Test server url
	ApiTest = "http://test.hackerlist.net"
)

const (
	ApiEndpoint = "/api/v1"
)

// Hlresponse is a type encapsulating all responses from the
// hackerlist API.
type HlResponse struct {
	Status string `json:"status"` // error or success

	// This is valid if status == "success"
	Response *json.RawMessage `json:"response"` // api response data

	// These are valid if status == "error"
	Message string `json:"msg"`   // a full error message
	Cause   string `json:"cause"` // single word describing the cause
}
