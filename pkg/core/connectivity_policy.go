package core

type ConnectivityPolicy struct {
	AllowDeploymentsToNoTargets bool                `json:"AllowDeploymentsToNoTargets"`
	ExcludeUnhealthyTargets     bool                `json:"ExcludeUnhealthyTargets"`
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
