package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Deployment struct {
	Changes                  []*ReleaseChanges `json:"Changes"`
	ChangesMarkdown          string            `json:"ChangesMarkdown,omitempty"`
	ChannelID                string            `json:"ChannelId,omitempty"`
	Comments                 string            `json:"Comments,omitempty"`
	Created                  *time.Time        `json:"Created,omitempty"`
	DeployedBy               string            `json:"DeployedBy,omitempty"`
	DeployedByID             string            `json:"DeployedById,omitempty"`
	DeployedToMachineIDs     []string          `json:"DeployedToMachineIds"`
	DeploymentProcessID      string            `json:"DeploymentProcessId,omitempty"`
	EnvironmentID            *string           `json:"EnvironmentId" validate:"required"`
	ExcludedMachineIDs       []string          `json:"ExcludedMachineIds"`
	FailureEncountered       bool              `json:"FailureEncountered,omitempty"`
	ForcePackageDownload     bool              `json:"ForcePackageDownload,omitempty"`
	ForcePackageRedeployment bool              `json:"ForcePackageRedeployment,omitempty"`
	FormValues               map[string]string `json:"FormValues,omitempty"`
	ManifestVariableSetID    string            `json:"ManifestVariableSetId,omitempty"`
	Name                     string            `json:"Name,omitempty"`
	ProjectID                string            `json:"ProjectId,omitempty"`
	QueueTime                *time.Time        `json:"QueueTime,omitempty"`
	QueueTimeExpiry          *time.Time        `json:"QueueTimeExpiry,omitempty"`
	ReleaseID                *string           `json:"ReleaseId" validate:"required"`
	SkipActions              []string          `json:"SkipActions"`
	SpaceID                  string            `json:"SpaceId,omitempty"`
	SpecificMachineIDs       []string          `json:"SpecificMachineIds"`
	TaskID                   string            `json:"TaskId,omitempty"`
	TenantID                 string            `json:"TenantId,omitempty"`
	TentacleRetentionPeriod  *RetentionPeriod  `json:"TentacleRetentionPeriod,omitempty"`
	UseGuidedFailure         bool              `json:"UseGuidedFailure,omitempty"`

	Resource
}

// Deployments defines a collection of Deployment instances with built-in support for paged results.
type Deployments struct {
	Items []Deployment `json:"Items"`
	PagedResults
}

// NewDeployment initializes a Deployment with a name, environment ID, and release ID.
func NewDeployment(name string, environmentID string, releaseID string) (*Deployment, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewDeployment", "name")
	}

	return &Deployment{
		EnvironmentID: &environmentID,
		Name:          name,
		ReleaseID:     &releaseID,
	}, nil
}

// GetID returns the ID value of the Deployment.
func (resource Deployment) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Deployment.
func (resource Deployment) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Deployment was changed.
func (resource Deployment) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Deployment.
func (resource Deployment) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the Deployment and returns an error if invalid.
func (resource Deployment) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		return err
	}

	return nil
}

var _ ResourceInterface = &Deployment{}
