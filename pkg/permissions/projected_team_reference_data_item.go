package permissions

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"

type ProjectedTeamReferenceDataItem struct {
	ExternalSecurityGroups []core.NamedReferenceItem `json:"ExternalSecurityGroups"`
	ID                     string                    `json:"Id,omitempty"`
	IsDirectlyAssigned     bool                      `json:"IsDirectlyAssigned"`
	Name                   string                    `json:"Name,omitempty"`
	SpaceID                string                    `json:"SpaceId,omitempty"`
}
