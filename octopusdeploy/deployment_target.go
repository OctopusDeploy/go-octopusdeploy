package octopusdeploy

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// TODO: research the JSON marshalling to include SpaceID

// DeploymentTargets defines a collection of deployment targets with built-in
// support for paged results from the API.
type DeploymentTargets struct {
	Items []*DeploymentTarget `json:"Items"`
	PagedResults
}

type DeploymentTarget struct {
	EnvironmentIDs         []string               `json:"EnvironmentIds,omitempty"`
	Roles                  []string               `json:"Roles,omitempty"`
	SpaceID                string                 `json:"SpaceId,omitempty"`
	TenantedDeploymentMode TenantedDeploymentMode `json:"TenantedDeploymentParticipation,omitempty"`
	TenantIDs              []string               `json:"TenantIds,omitempty"`
	TenantTags             []string               `json:"TenantTags,omitempty"`

	machine
}

func NewDeploymentTarget(name string, endpoint IEndpoint, environmentIDs []string, roles []string) *DeploymentTarget {
	return &DeploymentTarget{
		EnvironmentIDs:         environmentIDs,
		Roles:                  roles,
		TenantIDs:              []string{},
		TenantTags:             []string{},
		TenantedDeploymentMode: TenantedDeploymentMode("Untenanted"),

		machine: *newMachine(name, endpoint),
	}
}

// MarshalJSON returns a deployment target as its JSON encoding.
func (d *DeploymentTarget) MarshalJSON() ([]byte, error) {
	deploymentTarget := struct {
		EnvironmentIDs         []string               `json:"EnvironmentIds,omitempty"`
		Roles                  []string               `json:"Roles,omitempty"`
		SpaceID                string                 `json:"SpaceId,omitempty"`
		TenantedDeploymentMode TenantedDeploymentMode `json:"TenantedDeploymentParticipation,omitempty"`
		TenantIDs              []string               `json:"TenantIds,omitempty"`
		TenantTags             []string               `json:"TenantTags,omitempty"`

		machine
	}{
		EnvironmentIDs:         d.EnvironmentIDs,
		Roles:                  d.Roles,
		SpaceID:                d.SpaceID,
		TenantedDeploymentMode: d.TenantedDeploymentMode,
		TenantIDs:              d.TenantIDs,
		TenantTags:             d.TenantTags,

		machine: d.machine,
	}

	return json.Marshal(deploymentTarget)
}

// UnmarshalJSON sets this deployment target to its representation in JSON.
func (resource *DeploymentTarget) UnmarshalJSON(b []byte) error {
	var deploymentTarget map[string]*json.RawMessage
	err := json.Unmarshal(b, &deploymentTarget)
	if err != nil {
		return err
	}

	var m machine
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	resource.machine = m

	for deploymentTargetKey, deploymentTargetValue := range deploymentTarget {
		switch deploymentTargetKey {
		case "EnvironmentIds":
			err = json.Unmarshal(*deploymentTargetValue, &resource.EnvironmentIDs)
			if err != nil {
				return err
			}
		case "Roles":
			err = json.Unmarshal(*deploymentTargetValue, &resource.Roles)
			if err != nil {
				return err
			}
		case "SpaceId":
			err = json.Unmarshal(*deploymentTargetValue, &resource.SpaceID)
			if err != nil {
				return err
			}
		case "TenantedDeploymentParticipation":
			err = json.Unmarshal(*deploymentTargetValue, &resource.TenantedDeploymentMode)
			if err != nil {
				return err
			}
		case "TenantIds":
			err = json.Unmarshal(*deploymentTargetValue, &resource.TenantIDs)
			if err != nil {
				return err
			}
		case "TenantTags":
			err = json.Unmarshal(*deploymentTargetValue, &resource.TenantTags)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Validate checks the state of the deployment target and returns an error if
// invalid.
func (d *DeploymentTarget) Validate() error {
	return validator.New().Struct(d)
}
