package resources

type OptionsMetadata struct {
	SelectMode string            `json:"SelectMode,omitempty"`
	Values     map[string]string `json:"Values,omitempty"`
}

func NewOptionsMetadata() *OptionsMetadata {
	return &OptionsMetadata{}
}
