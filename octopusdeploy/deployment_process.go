package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type DeploymentProcessService struct {
	sling *sling.Sling
}

func NewDeploymentProcessService(sling *sling.Sling) *DeploymentProcessService {
	return &DeploymentProcessService{
		sling: sling,
	}
}

type DeploymentProcesses struct {
	Items []DeploymentProcess `json:"Items"`
	PagedResults
}

type DeploymentProcess struct {
	ID             string           `json:"Id,omitempty"`
	LastModifiedBy string           `json:"LastModifiedBy,omitempty"`
	LastModifiedOn string           `json:"LastModifiedOn,omitempty"`
	LastSnapshotID string           `json:"LastSnapshotId,omitempty"`
	Links          Links            `json:"Links,omitempty"`
	ProjectID      string           `json:"ProjectId,omitempty"`
	Steps          []DeploymentStep `json:"Steps,omitempty"`
	Version        int32            `json:"Version"`
}

type DeploymentStep struct {
	ID                 string                           `json:"Id,omitempty"`
	Name               string                           `json:"Name"`
	PackageRequirement DeploymentStepPackageRequirement `json:"PackageRequirement,omitempty"`                                         // may need its own model / enum
	Properties         map[string]string                `json:"Properties"`                                                           // TODO: refactor to use the PropertyValueResource for handling sensitive values - https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
	Condition          DeploymentStepCondition          `json:"Condition,omitempty" validate:"oneof=Success Failure Always Variable"` // variable option adds a Property "Octopus.Action.ConditionVariableExpression"
	StartTrigger       DeploymentStepStartTrigger       `json:"StartTrigger,omitempty" validate:"oneof=StartAfterPrevious StartWithPrevious"`
	Actions            []DeploymentAction               `json:"Actions,omitempty"`
}

type DeploymentAction struct {
	ID                            string             `json:"Id,omitempty"`
	Name                          string             `json:"Name"`
	ActionType                    string             `json:"ActionType"`
	IsDisabled                    bool               `json:"IsDisabled"`
	IsRequired                    bool               `json:"IsRequired"`
	WorkerPoolID                  string             `json:"WorkerPoolId,omitempty"`
	CanBeUsedForProjectVersioning bool               `json:"CanBeUsedForProjectVersioning"`
	Environments                  []string           `json:"Environments,omitempty"`
	ExcludedEnvironments          []string           `json:"ExcludedEnvironments,omitempty"`
	Channels                      []string           `json:"Channels,omitempty"`
	TenantTags                    []string           `json:"TenantTags,omitempty"`
	Properties                    map[string]string  `json:"Properties"` // TODO: refactor to use the PropertyValueResource for handling sensitive values - https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
	Packages                      []PackageReference `json:"Packages,omitempty"`
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

type PackageReference struct {
	ID                  string            `json:"Id,omitempty"`
	Name                string            `json:"Name,omitempty"`
	PackageID           string            `json:"PackageId,omitempty"`
	FeedID              string            `json:"FeedId"`
	AcquisitionLocation string            `json:"AcquisitionLocation"` // This can be an expression
	Properties          map[string]string `json:"Properties"`
}

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

func (s *DeploymentProcessService) Get(deploymentProcessID string) (*DeploymentProcess, error) {
	path := fmt.Sprintf("deploymentprocesses/%s", deploymentProcessID)
	resp, err := apiGet(s.sling, new(DeploymentProcess), path)

	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentProcess), nil
}

func (s *DeploymentProcessService) GetAll() (*[]DeploymentProcess, error) {
	var dp []DeploymentProcess

	path := "deploymentprocesses"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(DeploymentProcesses), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*DeploymentProcesses)

		dp = append(dp, r.Items...)

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &dp, nil
}

func (s *DeploymentProcessService) Update(deploymentProcess *DeploymentProcess) (*DeploymentProcess, error) {
	path := fmt.Sprintf("deploymentprocesses/%s", deploymentProcess.ID)
	resp, err := apiUpdate(s.sling, deploymentProcess, new(DeploymentProcess), path)

	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentProcess), nil
}
