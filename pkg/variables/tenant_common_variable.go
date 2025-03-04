package variables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actiontemplates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type GetTenantCommonVariablesQuery struct {
	TenantID                string `uri:"id" url:"id"`
	SpaceID                 string `uri:"spaceId" url:"spaceId"`
	IncludeMissingVariables bool   `uri:"includeMissingVariables" url:"includeMissingVariables"`
}

type GetTenantCommonVariablesResponse struct {
	TenantID         string                 `json:"TenantId,omitempty"`
	Variables        []TenantCommonVariable `json:"Variables,omitempty"`
	MissingVariables []TenantCommonVariable `json:"MissingCommonVariables,omitempty"`

	resources.Resource
}

type ModifyTenantCommonVariablesResponse struct {
	TenantID  string                 `json:"TenantId,omitempty"`
	Variables []TenantCommonVariable `json:"CommonVariables,omitempty"`

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
	Variables []TenantCommonVariablePayload `json:"Variables"`
}

type TenantCommonVariablePayload struct {
	ID                   string              `json:"Id,omitempty"`
	LibraryVariableSetId string              `json:"LibraryVariableSetId"`
	TemplateID           string              `json:"TemplateId"`
	Value                core.PropertyValue  `json:"Value"`
	Scope                TenantVariableScope `json:"Scope"`
}
