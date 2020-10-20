package octopusdeploy

type VersioningStrategy struct {
	DonorPackage       *DeploymentActionPackage `json:"DonorPackage,omitempty"`
	DonorPackageStepID *string                  `json:"DonorPackageStepId,omitempty"`
	Template           string                   `json:"Template,omitempty"`
}
