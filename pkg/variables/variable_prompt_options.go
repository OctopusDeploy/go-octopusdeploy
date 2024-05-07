package variables

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type VariablePromptOptions struct {
	Description     string                     `json:"Description"`
	DisplaySettings *resources.DisplaySettings `json:"DisplaySettings,omitempty"`
	IsRequired      bool                       `json:"Required"`
	Label           string                     `json:"Label"`
}
