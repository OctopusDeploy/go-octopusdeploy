package releases

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type DeploymentPromotionTarget struct {
	ID    string            `json:"Id"`
	Name  string            `json:"Name"`
	Links map[string]string `json:"Links"`
}

type ReleaseDeploymentTemplate struct {
	IsDeploymentProcessModified  bool                        `json:"IsDeploymentProcessModified"`
	DeploymentNotes              string                      `json:"DeploymentNotes"`
	IsVariableSetModified        bool                        `json:"IsVariableSetModified"`
	IsLibraryVariableSetModified bool                        `json:"IsLibraryVariableSetModified"`
	PromoteTo                    []DeploymentPromotionTarget `json:"PromoteTo"`
	TenantPromotions             []DeploymentPromotionTarget `json:"TenantPromotions"`

	resources.Resource
}
