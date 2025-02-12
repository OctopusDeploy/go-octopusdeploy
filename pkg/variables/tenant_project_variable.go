package variables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actiontemplates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type TenantProjectVariablesResponse struct {
	TenantID         string                  `json:"TenantId,omitempty"`
	ProjectVariables []TenantProjectVariable `json:"ProjectVariables,omitempty"`

	resources.Resource
}

type TenantProjectVariable struct {
	ProjectID   string                                  `json:"ProjectId"`
	ProjectName string                                  `json:"ProjectName,omitempty"`
	TemplateID  string                                  `json:"TemplateId"`
	Template    actiontemplates.ActionTemplateParameter `json:"Template"`
	Value       core.PropertyValue                      `json:"Value"`
	Scope       TenantVariableScope                     `json:"Scope"`
	Links       map[string]string                       `json:"Links,omitempty"`

	resources.Resource
}

type ModifyTenantProjectVariablesCommand struct {
	Variables []TenantProjectVariableCommand `json:"Variables"`
}

type TenantProjectVariableCommand struct {
	ID         string              `json:"Id,omitempty"`
	ProjectID  string              `json:"ProjectId"`
	TemplateID string              `json:"TemplateId"`
	Value      core.PropertyValue  `json:"Value"`
	Scope      TenantVariableScope `json:"Scope"`
}
