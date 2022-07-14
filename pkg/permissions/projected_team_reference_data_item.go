package permissions

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/core"

type ProjectedTeamReferenceDataItem struct {
	ExternalSecurityGroups []core.NamedReferenceItem `json:"ExternalSecurityGroups"`
	ID                     string                    `json:"Id,omitempty"`
	IsDirectlyAssigned     bool                      `json:"IsDirectlyAssigned,omitempty"`
	Name                   string                    `json:"Name,omitempty"`
	SpaceID                string                    `json:"SpaceId,omitempty"`
}
