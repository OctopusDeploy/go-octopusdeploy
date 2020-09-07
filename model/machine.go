package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Machines defines a collection of machines with built-in support for paged
// results from the API.
type Machines struct {
	Items []Machine `json:"Items"`
	PagedResults
}

// Machine represents deployment targets (or machine).
type Machine struct {
	Name              string           `json:"Name,omitempty"`
	Thumbprint        string           `json:"Thumbprint,omitempty"`
	URI               string           `json:"Uri,omitempty" validate:"omitempty,uri"`
	IsDisabled        bool             `json:"IsDisabled"`
	EnvironmentIDs    []string         `json:"EnvironmentIds"`
	Roles             []string         `json:"Roles"`
	MachinePolicyID   string           `json:"MachinePolicyId,omitempty"`
	DeploymentMode    string           `json:"TenantedDeploymentParticipation,omitempty" validate:"required,oneof=Untenanted TenantedOrUntenanted Tenanted"`
	TenantIDs         []string         `json:"TenantIds"`
	TenantTags        []string         `json:"TenantTags"`
	Status            string           `json:"Status,omitempty"`
	HealthStatus      string           `json:"HealthStatus,omitempty"`
	HasLatestCalamari bool             `json:"HasLatestCalamari"`
	StatusSummary     string           `json:"StatusSummary,omitempty"`
	IsInProcess       bool             `json:"IsInProcess"`
	Endpoint          *MachineEndpoint `json:"Endpoint" validate:"required"`
	OperatingSystem   string           `json:"OperatingSystem,omitempty"`
	ShellName         string           `json:"ShellName,omitempty"`
	ShellVersion      string           `json:"ShellVersion,omitempty"`
	Resource
}

func NewMachine(name string, isDisabled bool, environments []string, roles []string, machinePolicy string, deploymentMode string, tenantIDs []string, tenantTags []string) (*Machine, error) {
	return &Machine{
		Name:            name,
		IsDisabled:      isDisabled,
		EnvironmentIDs:  environments,
		Roles:           roles,
		MachinePolicyID: machinePolicy,
		DeploymentMode:  deploymentMode,
		TenantIDs:       tenantIDs,
		TenantTags:      tenantTags,
	}, nil
}

func (m *Machine) GetID() string {
	return m.ID
}

// Validate returns a collection of validation errors against the machine's
// internal values.
func (m *Machine) Validate() error {
	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}

		return err
	}

	return nil
}

var _ ResourceInterface = &Machine{}
