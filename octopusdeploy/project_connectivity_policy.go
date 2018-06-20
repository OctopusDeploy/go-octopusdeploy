package octopusdeploy

type ProjectConnectivityPolicy struct {
	AllowDeploymentsToNoTargets bool     `json:"AllowDeploymentsToNoTargets"`
	TargetRoles                 []string `json:"TargetRoles"`
	SkipMachineBehavior         string   `json:"SkipMachineBehavior,omitempty"`
}
