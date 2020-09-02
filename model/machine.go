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
	Name              string           `json:"Name"`
	Thumbprint        string           `json:"Thumbprint"`
	URI               string           `json:"Uri" validate:"omitempty,uri"`
	IsDisabled        bool             `json:"IsDisabled"`
	EnvironmentIDs    []string         `json:"EnvironmentIds"`
	Roles             []string         `json:"Roles"`
	MachinePolicyID   string           `json:"MachinePolicyId"`
	DeploymentMode    string           `json:"TenantedDeploymentParticipation" validate:"required,oneof=Untenanted TenantedOrUntenanted Tenanted"`
	TenantIDs         []string         `json:"TenantIDs"`
	TenantTags        []string         `json:"TenantTags"`
	Status            string           `json:"Status"`
	HealthStatus      string           `json:HealthStatus`
	HasLatestCalamari bool             `json:"HasLatestCalamari"`
	StatusSummary     string           `json:"StatusSummary"`
	IsInProcess       bool             `json:"IsInProcess"`
	Endpoint          *MachineEndpoint `json:"Endpoint" validate:"required"`
	OperatingSystem   string           `json:OperatingSystem`
	ShellName         string           `json:ShellName`
	ShellVersion      string           `json:ShellVersion`
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
