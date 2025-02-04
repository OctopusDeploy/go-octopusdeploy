package variables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type GetTenantProjectVariablesQuery struct {
	TenantID string `json:"TenantId"`
	SpaceID  string `json:"SpaceId"`
}

type TenantProjectVariablesResource struct {
	TenantID         string                  `json:"TenantId,omitempty"`
	ProjectVariables []TenantProjectVariable `json:"ProjectVariables,omitempty"`

	resources.Resource
}

type TenantProjectVariable struct {
	Id         string              `json:"Links,omitempty"`
	ProjectID  string              `json:"ProjectId"`
	TemplateID string              `json:"TemplateId"`
	Value      core.PropertyValue  `json:"Value"`
	Scope      TenantVariableScope `json:"Scope"`
	Links      map[string]string   `json:"Links,omitempty"`
}
