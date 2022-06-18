package variables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/actiontemplates"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
)

type ProjectVariable struct {
	Links       map[string]string                                 `json:"Links,omitempty"`
	ProjectID   string                                            `json:"ProjectId,omitempty"`
	ProjectName string                                            `json:"ProjectName,omitempty"`
	Templates   []*actiontemplates.ActionTemplateParameter        `json:"Templates"`
	Variables   map[string]map[string]core.PropertyValue `json:"Variables,omitempty"`
}
