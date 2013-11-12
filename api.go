package hljson

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func doReq(base, endpoint string) (*HlResponse, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	url := base + ApiEndpoint + endpoint

	resp, err := client.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Bad response: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Error reading body: %s", err)
	}

	var hlresp HlResponse

	err = json.Unmarshal(body, &hlresp)

	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling: %s", err)
	}

	if hlresp.Status != "success" {
		return nil, fmt.Errorf("API Error: (%s) %s", hlresp.Message, hlresp.Cause)
	}

	return &hlresp, nil
}
