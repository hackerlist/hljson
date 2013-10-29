package hljson

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
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Slug        string  `json:"slug"`
	OrgId       string  `json:"orgid"`
	BudgetMin   float64 `json:"budget_min"`
	BudgetMax   float64 `json:"budget_max"`
	Created     string  `json:"created"`
	Modified    string  `json:"modified"`
	StartTime   string  `json:"starttime"`
	EndTime     string  `json:"endtime"`
}
