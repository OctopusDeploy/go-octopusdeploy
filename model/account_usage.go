package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// NewAccountUsage initializes an AccountUsage.
func NewAccountUsage() (*AccountUsage, error) {
	return &AccountUsage{}, nil
}

// AccountUsage contains the projects and deployments which are using an
// account.
type AccountUsage struct {
	DeploymentProcesses []*StepUsage                    `json:"DeploymentProcesses,omitempty"`
	LibraryVariableSets []*LibraryVariableSetUsageEntry `json:"LibraryVariableSets,omitempty"`
	ProjectVariableSets []*ProjectVariableSetUsage      `json:"ProjectVariableSets,omitempty"`
	Releases            []*ReleaseUsage                 `json:"Releases,omitempty"`
	RunbookProcesses    []*RunbookStepUsage             `json:"RunbookProcesses,omitempty"`
	RunbookSnapshots    []*RunbookSnapshotUsage         `json:"RunbookSnapshots,omitempty"`
	Targets             []*TargetUsageEntry             `json:"Targets,omitempty"`

	Resource
}

// GetID returns the ID value of the AccountUsage struct instance.
func (resource AccountUsage) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this AccountUsage.
func (resource AccountUsage) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this AccountUsage was changed.
func (resource AccountUsage) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this AccountUsage.
func (resource AccountUsage) GetLinks() map[string]string {
	return resource.Links
}

// SetID
func (resource AccountUsage) SetID(id string) {
	resource.ID = id
}

// SetLastModifiedBy
func (resource AccountUsage) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

// SetLastModifiedOn
func (resource AccountUsage) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the AccountUsage and returns an error if invalid.
func (resource AccountUsage) Validate() error {
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

var _ ResourceInterface = &AccountUsage{}
