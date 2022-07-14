package variables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actiontemplates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
)

type Library struct {
	LibraryVariableSetID   string                                     `json:"LibraryVariableSetId,omitempty"`
	LibraryVariableSetName string                                     `json:"LibraryVariableSetName,omitempty"`
	Links                  map[string]string                          `json:"Links,omitempty"`
	Templates              []*actiontemplates.ActionTemplateParameter `json:"Templates"`
	Variables              map[string]core.PropertyValue              `json:"Variables,omitempty"`
}
