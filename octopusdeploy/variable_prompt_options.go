package octopusdeploy

type VariablePromptOptions struct {
	Description     string           `json:"Description"`
	DisplaySettings *DisplaySettings `json:"DisplaySettings,omitempty"`
	IsRequired      bool             `json:"Required"`
	Label           string           `json:"Label"`
}
