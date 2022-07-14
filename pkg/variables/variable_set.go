package variables

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type VariableSet struct {
	OwnerID     string               `json:"OwnerId,omitempty"`
	ScopeValues *VariableScopeValues `json:"ScopeValues,omitempty"`
	SpaceID     string               `json:"SpaceId,omitempty"`
	Variables   []*Variable          `json:"Variables"`
	Version     int32                `json:"Version"`

	resources.Resource
}

func NewVariableSet() *VariableSet {
	return &VariableSet{
		Resource: *resources.NewResource(),
	}
}
