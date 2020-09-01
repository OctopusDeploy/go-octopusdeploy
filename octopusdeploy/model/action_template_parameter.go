package model

type ActionTemplateParameter struct {
	DefaultValue    *PropertyValueResource `json:"DefaultValue,omitempty"`
	DisplaySettings map[string]string      `json:"DisplaySettings,omitempty"`
	HelpText        string                 `json:"HelpText,omitempty"`
	Label           string                 `json:"Label,omitempty"`
	Name            string                 `json:"Name,omitempty"`
	Resource
}
