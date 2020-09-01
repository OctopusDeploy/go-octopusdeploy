package model

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/enum"

type Machines struct {
	Items []Machine `json:"Items"`
	PagedResults
}

type Machine struct {
	Name                            string                      `json:"Name"`
	Thumbprint                      string                      `json:"Thumbprint"`
	URI                             string                      `json:"Uri"`
	IsDisabled                      bool                        `json:"IsDisabled"`
	EnvironmentIDs                  []string                    `json:"EnvironmentIds"`
	Roles                           []string                    `json:"Roles"`
	MachinePolicyID                 string                      `json:"MachinePolicyId"`
	TenantedDeploymentParticipation enum.TenantedDeploymentMode `json:"TenantedDeploymentParticipation"`
	TenantIDs                       []string                    `json:"TenantIDs"`
	TenantTags                      []string                    `json:"TenantTags"`
	Status                          string                      `json:"Status"`
	HasLatestCalamari               bool                        `json:"HasLatestCalamari"`
	StatusSummary                   string                      `json:"StatusSummary"`
	IsInProcess                     bool                        `json:"IsInProcess"`
	Endpoint                        *MachineEndpoint            `json:"Endpoint,omitempty"`
	Resource
}

func NewMachine(Name string, Disabled bool, EnvironmentIDs []string, Roles []string, MachinePolicyID string, TenantedDeploymentParticipation enum.TenantedDeploymentMode, TenantIDs, TenantTags []string) *Machine {
	return &Machine{
		Name:                            Name,
		IsDisabled:                      Disabled,
		EnvironmentIDs:                  EnvironmentIDs,
		Roles:                           Roles,
		MachinePolicyID:                 MachinePolicyID,
		TenantedDeploymentParticipation: TenantedDeploymentParticipation,
		TenantIDs:                       TenantIDs,
		TenantTags:                      TenantTags,
		Status:                          "Unknown",
		Thumbprint:                      "0123456789ABCDEF",
		URI:                             "https://localhost/",
	}
}

// ValidateMachineValues checks the values of a Machine object to see if they are suitable for
// sending to Octopus Deploy. Used when adding or updating machines.
func ValidateMachineValues(Machine *Machine) error {
	if Machine.Endpoint != nil {
		matchingPropertiesErr := ValidateMultipleProperties([]error{
			ValidatePropertiesMatch(Machine.Endpoint.Thumbprint, "Machine.Endpoint.Thumbprint", Machine.Thumbprint, "Machine.Thumbprint"),
			ValidatePropertiesMatch(Machine.Endpoint.URI, "Machine.Endpoint.URI", Machine.URI, "Machine.URI"),
		})

		if matchingPropertiesErr != nil {
			return matchingPropertiesErr
		}
	}

	return ValidateMultipleProperties([]error{
		ValidatePropertyValues("Machine.Status", Machine.Status, ValidMachineStatuses),
		ValidatePropertyValues("Machine.TenantedDeploymentParticipation", Machine.TenantedDeploymentParticipation.String(), enum.TenantedDeploymentModeNames()),
	})
}
