package octopusdeploy

type OptionsMetadata struct {
	SelectMode string            `json:"SelectMode,omitempty"`
	Values     map[string]string `json:"Values,omitempty"`
}
