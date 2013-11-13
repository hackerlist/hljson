package hljson

const (
	UserEndpoint = "/users"
)

// XXX: add users endpoint struct

type HlUser struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Website   string    `json:"website"`
	Created   string    `json:"created"` // XXX: change to time.Time
	Type      string    `json:"type"`
	Username  string    `json:"username"`
	Bio       string    `json:"bio"`
	AvatarUrl string    `json:"avatar_url"`
	Account   HlAccount `json:"account"`
}
