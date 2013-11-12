package hljson

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMissions2(t *testing.T) {
	mis, err := GetMissions()

	if err != nil {
		t.Fatalf("Error retreiving missions: %s", err)
	}

	for _, m := range mis {
		if len(m.Title) <= 0 {
			t.Errorf("mid %d has empty title", m.Id)
		}

		if m.BudgetEst <= 0 {
			t.Errorf("mid %d budget_est <= 0", m.Id)
		}

		t.Logf("Title: %s\n", m.Title)
		t.Logf("Description: %s\n", m.Description)
		t.Logf("Budget Est: %f\n", m.BudgetEst)
	}
}
