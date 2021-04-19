package octopusdeploy

type VariableScope struct {
	Actions           []string `json:"Action,omitempty"`
	Channels          []string `json:"Channel,omitempty"`
	Environments      []string `json:"Environment,omitempty"`
	Machines          []string `json:"Machine,omitempty"`
	ParentDeployments []string `json:"ParentDeployment,omitempty"`
	Private           []string `json:"Private,omitempty"`
	ProcessOwners     []string `json:"ProcessOwner,omitempty"`
	Projects          []string `json:"Project,omitempty"`
	Roles             []string `json:"Role,omitempty"`
	TargetRoles       []string `json:"TargetRole,omitempty"`
	Tenants           []string `json:"Tenant,omitempty"`
	TenantTags        []string `json:"TenantTag,omitempty"`
	Triggers          []string `json:"Trigger,omitempty"`
	Users             []string `json:"User,omitempty"`
}

func (scope VariableScope) IsEmpty() bool {
	return len(scope.Actions) == 0 &&
		len(scope.Channels) == 0 &&
		len(scope.Environments) == 0 &&
		len(scope.Machines) == 0 &&
		len(scope.ParentDeployments) == 0 &&
		len(scope.Private) == 0 &&
		len(scope.ProcessOwners) == 0 &&
		len(scope.Projects) == 0 &&
		len(scope.Roles) == 0 &&
		len(scope.TargetRoles) == 0 &&
		len(scope.Tenants) == 0 &&
		len(scope.TenantTags) == 0 &&
		len(scope.Triggers) == 0 &&
		len(scope.Users) == 0
}
