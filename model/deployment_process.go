package model

import (
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

func (d *DeploymentProcess) Validate() error {
	validate := validator.New()

	err := validate.Struct(d)

	if err != nil {
		return err
	}

	return nil
}
