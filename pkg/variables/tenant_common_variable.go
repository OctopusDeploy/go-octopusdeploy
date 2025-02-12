package variables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actiontemplates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type TenantCommonVariablesResponse struct {
	TenantID        string                 `json:"TenantId,omitempty"`
	CommonVariables []TenantCommonVariable `json:"CommonVariables,omitempty"`

	resources.Resource
}

type TenantCommonVariable struct {
	LibraryVariableSetId   string                                  `json:"LibraryVariableSetId"`
	LibraryVariableSetName string                                  `json:"LibraryVariableSetName,omitempty"`
	TemplateID             string                                  `json:"TemplateId"`
	Template               actiontemplates.ActionTemplateParameter `json:"Template"`
	Value                  core.PropertyValue                      `json:"Value"`
	Scope                  TenantVariableScope                     `json:"Scope"`

	resources.Resource
}

type ModifyTenantCommonVariablesCommand struct {
	Variables []TenantCommonVariableCommand `json:"Variables"`
}

type TenantCommonVariableCommand struct {
	ID                   string              `json:"Id,omitempty"`
	LibraryVariableSetId string              `json:"LibraryVariableSetId"`
	TemplateID           string              `json:"TemplateId"`
	Value                core.PropertyValue  `json:"Value"`
	Scope                TenantVariableScope `json:"Scope"`
}
