package octopusdeploy

type ExtensionSettingsValues struct {
	ExtensionID string      `json:"ExtensionId,omitempty"`
	Values      interface{} `json:"Values,omitempty"`
}
