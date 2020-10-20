package octopusdeploy

// ActionTemplateParameter represents an action template parameter.
type ActionTemplateParameter struct {
	DefaultValue    *PropertyValueResource `json:"DefaultValue,omitempty"`
	DisplaySettings map[string]string      `json:"DisplaySettings,omitempty"`
	HelpText        string                 `json:"HelpText,omitempty"`
	Label           string                 `json:"Label,omitempty"`
	Name            string                 `json:"Name,omitempty"`

	Resource
}

func NewActionTemplateParameter() *ActionTemplateParameter {
	return &ActionTemplateParameter{
		Resource: *newResource(),
	}
}
