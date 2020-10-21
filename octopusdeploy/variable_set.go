package octopusdeploy

type VariableSet struct {
	OwnerID     string               `json:"OwnerId,omitempty"`
	ScopeValues *VariableScopeValues `json:"ScopeValues,omitempty"`
	SpaceID     string               `json:"SpaceId,omitempty"`
	Variables   []*Variable          `json:"Variables"`
	Version     int32                `json:"Version,omitempty"`

	resource
}

func NewVariableSet() *VariableSet {
	return &VariableSet{
		resource: *newResource(),
	}
}
