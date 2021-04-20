package octopusdeploy

type VariableScope struct {
	Environments []string `json:"Environment,omitempty"`
	Machines     []string `json:"Machine,omitempty"`
	Actions      []string `json:"Action,omitempty"`
	Roles        []string `json:"Role,omitempty"`
	Channels     []string `json:"Channel,omitempty"`
	TenantTags   []string `json:"TenantTag,omitempty"`
}

func (scope VariableScope) IsEmpty() bool {
	return len(scope.Actions) == 0 &&
		len(scope.Channels) == 0 &&
		len(scope.Environments) == 0 &&
		len(scope.Machines) == 0 &&
		len(scope.Roles) == 0 &&
		len(scope.TenantTags) == 0
}
