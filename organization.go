package hljson

import ()

const (
	OrganizationEndpoint = "/orgs"
)

// XXX: define a struct for orgs api response

type HlOrganization struct {
	Id      int    `json:"id"`
	Created string `json:"created"`
	User    HlUser `json:"user"`
}
