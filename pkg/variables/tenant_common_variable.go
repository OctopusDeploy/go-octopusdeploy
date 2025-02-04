package variables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type GetTenantCommonVariablesQuery struct {
	TenantID string `json:"TenantId"`
	SpaceID  string `json:"SpaceId"`
}

type TenantCommonVariablesResource struct {
	TenantID        string                 `json:"TenantId,omitempty"`
	CommonVariables []TenantCommonVariable `json:"CommonVariables,omitempty"`

	resources.Resource
}

type TenantCommonVariable struct {
	Id         string              `json:"Links,omitempty"`
	ProjectID  string              `json:"ProjectId"`
	TemplateID string              `json:"TemplateId"`
	Value      core.PropertyValue  `json:"Value"`
	Scope      TenantVariableScope `json:"Scope"`
	Links      map[string]string   `json:"Links,omitempty"`
}
