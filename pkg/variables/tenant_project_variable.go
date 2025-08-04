package variables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actiontemplates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type GetTenantProjectVariablesQuery struct {
	TenantID                string `uri:"id" url:"id"`
	SpaceID                 string `uri:"spaceId" url:"spaceId"`
	IncludeMissingVariables bool   `uri:"includeMissingVariables" url:"includeMissingVariables"`
}

type GetTenantProjectVariablesResponse struct {
	TenantID         string                  `json:"TenantId,omitempty"`
	Variables        []TenantProjectVariable `json:"Variables,omitempty"`
	MissingVariables []TenantProjectVariable `json:"MissingVariables,omitempty"`

	resources.Resource
}

type ModifyTenantProjectVariablesResponse struct {
	TenantID  string                  `json:"TenantId,omitempty"`
	Variables []TenantProjectVariable `json:"Variables,omitempty"`

	resources.Resource
}

type TenantProjectVariable struct {
	ProjectID   string                                  `json:"ProjectId"`
	ProjectName string                                  `json:"ProjectName,omitempty"`
	TemplateID  string                                  `json:"TemplateId"`
	Template    actiontemplates.ActionTemplateParameter `json:"Template"`
	Value       core.PropertyValue                      `json:"Value"`
	Scope       TenantVariableScope                     `json:"Scope"`

	resources.Resource
}

type ModifyTenantProjectVariablesCommand struct {
	Variables []TenantProjectVariablePayload `json:"Variables"`
}

type TenantProjectVariablePayload struct {
	ID         string              `json:"Id,omitempty"`
	ProjectID  string              `json:"ProjectId"`
	TemplateID string              `json:"TemplateId"`
	Value      core.PropertyValue  `json:"Value"`
	Scope      TenantVariableScope `json:"Scope"`
}
