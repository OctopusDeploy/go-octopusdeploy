package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Machines struct {
	Items []Machine `json:"Items"`
	PagedResults
}

type Machine struct {
	Name              string           `json:"Name,omitempty"`
	Thumbprint        string           `json:"Thumbprint,omitempty"`
	URI               string           `json:"Uri,omitempty" validate:"omitempty,uri"`
	IsDisabled        bool             `json:"IsDisabled"`
	EnvironmentIDs    []string         `json:"EnvironmentIds"`
	Roles             []string         `json:"Roles"`
	MachinePolicyID   string           `json:"MachinePolicyId,omitempty"`
	DeploymentMode    string           `json:"TenantedDeploymentParticipation,omitempty" validate:"required,oneof=Untenanted TenantedOrUntenanted Tenanted"`
	TenantIDs         []string         `json:"TenantIDs"`
	TenantTags        []string         `json:"TenantTags"`
	Status            string           `json:"Status,omitempty"`
	HealthStatus      string           `json:"HealthStatus,omitempty"`
	HasLatestCalamari bool             `json:"HasLatestCalamari"`
	StatusSummary     string           `json:"StatusSummary,omitempty"`
	IsInProcess       bool             `json:"IsInProcess"`
	Endpoint          *MachineEndpoint `json:"Endpoint" validate:"required"`
	OperatingSystem   string           `json:OperatingSystem,omitempty`
	ShellName         string           `json:ShellName,omitempty`
	ShellVersion      string           `json:ShellVersion,omitempty`
	Resource
}

func (machine *Machine) Validate() error {
	validate := validator.New()
	err := validate.Struct(machine)

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
