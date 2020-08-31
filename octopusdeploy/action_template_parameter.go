package octopusdeploy

type ActionTemplateParameter struct {
	DefaultValue    *PropertyValueResource `json:"DefaultValue,omitempty"`
	DisplaySettings map[string]string      `json:"DisplaySettings,omitempty"`
	HelpText        string                 `json:"HelpText,omitempty"`
	ID              string                 `json:"Id,omitempty"`
	Label           string                 `json:"Label,omitempty"`
	Name            string                 `json:"Name,omitempty"`
	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn string `json:"LastModifiedOn,omitempty"` // datetime

	// links
	Links Links `json:"Links,omitempty"`
}
