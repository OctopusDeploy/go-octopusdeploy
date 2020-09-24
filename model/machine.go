package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Machines defines a collection of machines with built-in support for paged results from the API.
type Machines struct {
	Items []Machine `json:"Items"`
	PagedResults
}

// Machine represents deployment targets (or machine).
type Machine struct {
	DeploymentMode    string           `json:"TenantedDeploymentParticipation,omitempty" validate:"required,oneof=Untenanted TenantedOrUntenanted Tenanted"`
	Endpoint          *MachineEndpoint `json:"Endpoint" validate:"required"`
	EnvironmentIDs    []string         `json:"EnvironmentIds"`
	HasLatestCalamari bool             `json:"HasLatestCalamari"`
	HealthStatus      string           `json:"HealthStatus,omitempty"`
	IsDisabled        bool             `json:"IsDisabled"`
	IsInProcess       bool             `json:"IsInProcess"`
	MachinePolicyID   string           `json:"MachinePolicyId,omitempty"`
	Name              string           `json:"Name,omitempty"`
	OperatingSystem   string           `json:"OperatingSystem,omitempty"`
	Roles             []string         `json:"Roles"`
	ShellName         string           `json:"ShellName,omitempty"`
	ShellVersion      string           `json:"ShellVersion,omitempty"`
	Status            string           `json:"Status,omitempty"`
	StatusSummary     string           `json:"StatusSummary,omitempty"`
	TenantIDs         []string         `json:"TenantIds"`
	TenantTags        []string         `json:"TenantTags"`
	Thumbprint        string           `json:"Thumbprint,omitempty"`
	URI               string           `json:"Uri,omitempty" validate:"omitempty,uri"`

	Resource
}

func NewMachine(name string, isDisabled bool, environments []string, roles []string, machinePolicy string, deploymentMode string, tenantIDs []string, tenantTags []string) (*Machine, error) {
	return &Machine{
		DeploymentMode:  deploymentMode,
		EnvironmentIDs:  environments,
		IsDisabled:      isDisabled,
		MachinePolicyID: machinePolicy,
		Name:            name,
		Roles:           roles,
		TenantIDs:       tenantIDs,
		TenantTags:      tenantTags,
	}, nil
}

// GetID returns the ID value of the Machine.
func (resource Machine) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Machine.
func (resource Machine) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Machine was changed.
func (resource Machine) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Machine.
func (resource Machine) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the Machine and returns an error if invalid.
func (resource Machine) Validate() error {
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

var _ ResourceInterface = &Machine{}
