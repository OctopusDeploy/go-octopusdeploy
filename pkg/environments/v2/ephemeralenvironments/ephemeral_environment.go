package ephemeralenvironments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core/pagination"
)

type EphemeralEnvironment struct {
	ID                  string `json:"Id"`
	Name                string `json:"Name"`
	SpaceID             string `json:"SpaceId"`
	Slug                string `json:"Slug"`
	Description         string `json:"Description"`
	Type                string `json:"Type"`
	SortOrder           int    `json:"SortOrder"`
	UseGuidedFailure    bool   `json:"UseGuidedFailure"`
	ParentEnvironmentId string `json:"ParentEnvironmentId"`
}

type EphemeralEnvironmentResponse struct {
	Items []*EphemeralEnvironment `json:"Items"`

	pagination.PaginatedCollectionFields
}
