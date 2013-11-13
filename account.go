package hljson

const ()

type HlAccount struct {
	Id        int    `json:"id"`
	Created   string `json:"created"`
	SiteAdmin bool   `json:"site_admin"`
}
