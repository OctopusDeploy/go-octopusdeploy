package octopusdeploy

type ConnectivityPolicy struct {
	AllowDeploymentsToNoTargets bool                `json:"AllowDeploymentsToNoTargets,omitempty"`
	ExcludeUnhealthyTargets     bool                `json:"ExcludeUnhealthyTargets,omitempty"`
	SkipMachineBehavior         SkipMachineBehavior `json:"SkipMachineBehavior,omitempty"`
	TargetRoles                 []string            `json:"TargetRoles,omitempty"`
}
