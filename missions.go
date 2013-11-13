package hljson

import (
	"encoding/json"
	"fmt"
)

const (
	MissionEndpoint = "/missions"
)

// HlMissionResponse is a response from the 'missions' api call.
type HlMissionResponse struct {
	Limit    int         `json:"limit"`
	Offset   int         `json:"offset"`
	Missions []HlMission `json:"missions"`
}

// HlMission contains data about each individual mission.
type HlMission struct {
	Id          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Slug        string         `json:"slug"`
	Openings    int            `json:"openings"`
	Org         HlOrganization `json:"org"`
	SuperId     int            `json:"superid"`
	OrgId       int            `json:"orgid"`
	Status      string         `json:"status"`
	Urgency     string         `json:"urgency"`
	Visibility  string         `json:"visibility"`
	BudgetEst   float64        `json:"est_budget"`
	Enabled     bool           `json:"enabled"`
	// XXX These should be *time.Time when the server emits
	// the proper format.
	Created   string `json:"created"`
	Modified  string `json:"modified"`
	StartTime string `json:"starttime"`
	Deadline  string `json:"deadline"`
	EndTime   string `json:"endtime"`
}

func GetMissions() ([]HlMission, error) {
	// XXX: make a way to switch between real api and test
	resp, err := doReq(ApiTest, MissionEndpoint)

	if err != nil {
		return nil, err
	}

	var hlmissions HlMissionResponse

	err = json.Unmarshal([]byte(*resp.Response), &hlmissions)

	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling mission json: %s", err)
	}

	return hlmissions.Missions, nil
}
