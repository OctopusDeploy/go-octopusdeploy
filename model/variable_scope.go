package model

type VariableScope struct {
	Project     []string `json:"Project,omitempty"`
	Environment []string `json:"Environment,omitempty"`
	Machine     []string `json:"Machine,omitempty"`
	Role        []string `json:"Role,omitempty"`
	TargetRole  []string `json:"TargetRole,omitempty"`
	Action      []string `json:"Action,omitempty"`
	User        []string `json:"User,omitempty"`
	Private     []string `json:"Private,omitempty"`
	Channel     []string `json:"Channel,omitempty"`
	TenantTag   []string `json:"TenantTag,omitempty"`
	Tenant      []string `json:"Tenant,omitempty"`
}

// NewVariableScope initializes a variable scope.
func NewVariableScope() *VariableScope {
	return &VariableScope{}
}
