package resources

type DeploymentStepPackageRequirement string

const (
	DeploymentStepPackageRequirementLetOctopusDecide         DeploymentStepPackageRequirement = "LetOctopusDecide"
	DeploymentStepPackageRequirementBeforePackageAcquisition DeploymentStepPackageRequirement = "BeforePackageAcquisition"
	DeploymentStepPackageRequirementAfterPackageAcquisition  DeploymentStepPackageRequirement = "AfterPackageAcquisition"
)
