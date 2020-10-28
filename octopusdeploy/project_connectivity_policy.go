package octopusdeploy

type ProjectConnectivityPolicy struct {
	AllowDeploymentsToNoTargets bool                `json:"AllowDeploymentsToNoTargets,omitempty"`
	TargetRoles                 []string            `json:"TargetRoles,omitempty"`
	SkipMachineBehavior         SkipMachineBehavior `json:"SkipMachineBehavior,omitempty"`
}
