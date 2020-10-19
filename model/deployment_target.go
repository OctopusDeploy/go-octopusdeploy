package model

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
	TenantedDeploymentMode string   `json:"TenantedDeploymentParticipation,omitempty" validate:"required,oneof=Untenanted TenantedOrUntenanted Tenanted"`
	EnvironmentIDs         []string `json:"EnvironmentIds"`
	Roles                  []string `json:"Roles"`
	SpaceID                string   `json:"SpaceId,omitempty"`
	TenantIDs              []string `json:"TenantIds"`
	TenantTags             []string `json:"TenantTags"`

	machine
}

func NewDeploymentTarget(name string, endpoint IEndpoint, environmentIDs []string, roles []string) *DeploymentTarget {
	return &DeploymentTarget{
		TenantedDeploymentMode: "Untenanted",
		EnvironmentIDs:         environmentIDs,
		Roles:                  roles,
		TenantIDs:              []string{},
		TenantTags:             []string{},
		machine:                *newMachine(name, endpoint),
	}
}

// MarshalJSON returns a deployment target as its JSON encoding.
func (d *DeploymentTarget) MarshalJSON() ([]byte, error) {
	deploymentTarget := struct {
		TenantedDeploymentMode string   `json:"TenantedDeploymentParticipation,omitempty" validate:"required,oneof=Untenanted TenantedOrUntenanted Tenanted"`
		EnvironmentIDs         []string `json:"EnvironmentIds"`
		Roles                  []string `json:"Roles"`
		SpaceID                string   `json:"SpaceId,omitempty"`
		TenantIDs              []string `json:"TenantIds"`
		TenantTags             []string `json:"TenantTags"`
		machine
	}{
		TenantedDeploymentMode: d.TenantedDeploymentMode,
		EnvironmentIDs:         d.EnvironmentIDs,
		Roles:                  d.Roles,
		SpaceID:                d.SpaceID,
		TenantIDs:              d.TenantIDs,
		TenantTags:             d.TenantTags,
		machine:                d.machine,
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
