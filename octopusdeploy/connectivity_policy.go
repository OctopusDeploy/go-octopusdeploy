package octopusdeploy

type ConnectivityPolicy struct {
	AllowDeploymentsToNoTargets bool                `json:"AllowDeploymentsToNoTargets,omitempty"`
	ExcludeUnhealthyTargets     bool                `json:"ExcludeUnhealthyTargets,omitempty"`
	SkipMachineBehavior         SkipMachineBehavior `json:"SkipMachineBehavior,omitempty"`
	TargetRoles                 []string            `json:"TargetRoles,omitempty"`
}

func NewConnectivityPolicy() *ConnectivityPolicy {
	return &ConnectivityPolicy{
		AllowDeploymentsToNoTargets: false,
		ExcludeUnhealthyTargets:     false,
		SkipMachineBehavior:         "None",
		TargetRoles:                 []string{},
	}
}
