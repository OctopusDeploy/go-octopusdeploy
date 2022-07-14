package teams

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

// Teams defines a collection of teams with built-in support for paged results.
type Teams struct {
	Items []*Team `json:"Items"`
	resources.PagedResults
}

type Team struct {
	CanBeDeleted           bool                      `json:"CanBeDeleted,omitempty"`
	CanBeRenamed           bool                      `json:"CanBeRenamed,omitempty"`
	CanChangeMembers       bool                      `json:"CanChangeMembers,omitempty"`
	CanChangeRoles         bool                      `json:"CanChangeRoles,omitempty"`
	Description            string                    `json:"Description,omitempty"`
	ExternalSecurityGroups []core.NamedReferenceItem `json:"ExternalSecurityGroups,omitempty"`
	MemberUserIDs          []string                  `json:"MemberUserIds"`
	Name                   string                    `json:"Name" validate:"required"`
	SpaceID                string                    `json:"SpaceId,omitempty"`

	resources.Resource
}

func NewTeam(name string) *Team {
	return &Team{
		Name:     name,
		Resource: *resources.NewResource(),
	}
}

// Validate checks the state of the team and returns an error if invalid.
func (t *Team) Validate() error {
	return validator.New().Struct(t)
}
