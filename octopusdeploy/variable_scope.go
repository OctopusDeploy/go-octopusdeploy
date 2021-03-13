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

// NewVariableScope initializes a variable scope.
func NewVariableScope() *VariableScope {
	return &VariableScope{}
}
