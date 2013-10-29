package hljson

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMissions(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	url := ApiTest + ApiEndpoint + MissionEndpoint

	resp, err := client.Get(url)

	if err != nil {
		t.Fatal("http.Get error:", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var hlresp HlResponse
	err = json.Unmarshal(body, &hlresp)

	if err != nil {
		t.Fatal("unmarshal error:", err)
	}

	var hlmissions HlMissionResponse

	err = json.Unmarshal([]byte(*hlresp.Response), &hlmissions)

	if err != nil {
		t.Fatal("mission unmarshal error:", err)
	}

	for _, m := range hlmissions.Missions {
		t.Logf("Title: %s\n", m.Title)
		t.Logf("Description: %s\n", m.Description)
		t.Logf("Budget: %f-%f\n", m.BudgetMin, m.BudgetMax)
	}
}
