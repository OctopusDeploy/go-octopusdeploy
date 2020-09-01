package model

type VersioningStrategy struct {
	DonorPackage *DeploymentActionPackage `json:"DonorPackage,omitempty"`
	Template     string                   `json:"Template,omitempty"`
}
