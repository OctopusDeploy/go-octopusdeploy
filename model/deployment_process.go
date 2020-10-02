package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type DeploymentProcesses struct {
	Items []DeploymentProcess `json:"Items"`
	PagedResults
}

type DeploymentProcess struct {
	LastSnapshotID string           `json:"LastSnapshotId,omitempty"`
	ProjectID      string           `json:"ProjectId,omitempty"`
	Steps          []DeploymentStep `json:"Steps,omitempty"`
	Version        int32            `json:"Version"`

	Resource
}

// NewDeploymentProcess initializes a deployment process If any of the input
// parameters are invalid, it will return nil and an error.
func NewDeploymentProcess(projectID string) (*DeploymentProcess, error) {
	if isEmpty(projectID) {
		return nil, createInvalidParameterError("NewDeploymentProcess", "projectID")
	}

	return &DeploymentProcess{
		ProjectID: projectID,
	}, nil
}

// GetID returns the ID value of the DeploymentProcess.
func (resource DeploymentProcess) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this DeploymentProcess.
func (resource DeploymentProcess) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this DeploymentProcess was changed.
func (resource DeploymentProcess) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this DeploymentProcess.
func (resource DeploymentProcess) GetLinks() map[string]string {
	return resource.Links
}

func (resource DeploymentProcess) SetID(id string) {
	resource.ID = id
}

func (resource DeploymentProcess) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

func (resource DeploymentProcess) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the DeploymentProcess and returns an error if invalid.
func (resource DeploymentProcess) Validate() error {
	validate := validator.New()

	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
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
