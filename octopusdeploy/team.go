package octopusdeploy

import "github.com/go-playground/validator/v10"

// Teams defines a collection of teams with built-in support for paged results.
type Teams struct {
	Items []*Team `json:"Items"`
	PagedResults
}

type Team struct {
	CanBeDeleted           bool                 `json:"CanBeDeleted,omitempty"`
	CanBeRenamed           bool                 `json:"CanBeRenamed,omitempty"`
	CanChangeMembers       bool                 `json:"CanChangeMembers,omitempty"`
	CanChangeRoles         bool                 `json:"CanChangeRoles,omitempty"`
	Description            string               `json:"Description,omitempty"`
	ExternalSecurityGroups []NamedReferenceItem `json:"ExternalSecurityGroups"`
	MemberUserIDs          []string             `json:"MemberUserIds"`
	Name                   string               `json:"Name" validate:"required"`
	SpaceID                string               `json:"SpaceId,omitempty"`

	resource
}

func NewTeam(name string) *Team {
	return &Team{
		Name:     name,
		resource: *newResource(),
	}
}

// Validate checks the state of the team and returns an error if invalid.
func (t *Team) Validate() error {
	return validator.New().Struct(t)
}
