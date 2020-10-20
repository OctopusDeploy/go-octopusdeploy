package octopusdeploy

type DeploymentProcesses struct {
	Items []*DeploymentProcess `json:"Items"`
	PagedResults
}

type DeploymentProcess struct {
	LastSnapshotID string           `json:"LastSnapshotId,omitempty"`
	ProjectID      string           `json:"ProjectId,omitempty"`
	Steps          []DeploymentStep `json:"Steps,omitempty"`
	Version        int32            `json:"Version"`

	Resource
}

// NewDeploymentProcess initializes a deployment process.
func NewDeploymentProcess(projectID string) *DeploymentProcess {
	return &DeploymentProcess{
		ProjectID: projectID,
		Resource:  *newResource(),
	}
}

type DeploymentStepPackageRequirement string

const (
	DeploymentStepPackageRequirementLetOctopusDecide         = DeploymentStepPackageRequirement("LetOctopusDecide")
	DeploymentStepPackageRequirementBeforePackageAcquisition = DeploymentStepPackageRequirement("BeforePackageAcquisition")
	DeploymentStepPackageRequirementAfterPackageAcquisition  = DeploymentStepPackageRequirement("AfterPackageAcquisition")
)

type DeploymentStepCondition string

const (
	DeploymentStepConditionSuccess  = DeploymentStepCondition("Success")
	DeploymentStepConditionFailure  = DeploymentStepCondition("Failure")
	DeploymentStepConditionAlways   = DeploymentStepCondition("Always")
	DeploymentStepConditionVariable = DeploymentStepCondition("Variable")
)

type DeploymentStepStartTrigger string

const (
	DeploymentStepStartTriggerStartAfterPrevious = DeploymentStepStartTrigger("StartAfterPrevious")
	DeploymentStepStartTriggerStartWithPrevious  = DeploymentStepStartTrigger("StartWithPrevious")
)

const (
	PackageAcquisitionLocationServer          = "Server"
	PackageAcquisitionLocationExecutionTarget = "ExecutionTarget"
	PackageAcquisitionLocationNotAcquired     = "NotAcquired"
)
