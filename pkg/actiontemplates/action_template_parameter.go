package actiontemplates

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"
)

// ActionTemplateParameter represents an action template parameter.
type ActionTemplateParameter struct {
	DefaultValue    *core.PropertyValue `json:"DefaultValue,omitempty"`
	DisplaySettings map[string]string   `json:"DisplaySettings,omitempty"`
	HelpText        string              `json:"HelpText,omitempty"`
	Label           string              `json:"Label,omitempty"`
	Name            string              `json:"Name,omitempty"`

	resources.Resource
}

func NewActionTemplateParameter() *ActionTemplateParameter {
	return &ActionTemplateParameter{
		Resource: *resources.NewResource(),
	}
}
