package projectgroups

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type ProjectGroup struct {
	Description       string   `json:"Description,omitempty"`
	EnvironmentIDs    []string `json:"EnvironmentIds,omitempty"`
	Name              string   `json:"Name,omitempty" validate:"required"`
	RetentionPolicyID string   `json:"RetentionPolicyId,omitempty"`
	SpaceID           string   `json:"SpaceId,omitempty"`

	resources.Resource
}

func NewProjectGroup(name string) *ProjectGroup {
	return &ProjectGroup{
		Name:     name,
		Resource: *resources.NewResource(),
	}
}

func (s *ProjectGroup) GetName() string {
	return s.Name
}
