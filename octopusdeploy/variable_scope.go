package octopusdeploy

type VariableScope struct {
	Environments []string `json:"Environments,omitempty"`
	Machines     []string `json:"Machines,omitempty"`
	Actions      []string `json:"Actions,omitempty"`
	Roles        []string `json:"Roles,omitempty"`
	Channels     []string `json:"Channels,omitempty"`
	TenantTags   []string `json:"TenantTags,omitempty"`
}

func (scope VariableScope) IsEmpty() bool {
	return len(scope.Actions) == 0 &&
		len(scope.Channels) == 0 &&
		len(scope.Environments) == 0 &&
		len(scope.Machines) == 0 &&
		len(scope.Roles) == 0 &&
		len(scope.TenantTags) == 0
}
