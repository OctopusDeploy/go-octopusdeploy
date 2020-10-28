package octopusdeploy

type PropertyMetadata struct {
	DisplayInfo *DisplayInfo `json:"DisplayInfo,omitempty"`
	Name        string       `json:"Name,omitempty"`
	Type        string       `json:"Type,omitempty"`
}

func NewPropertyMetadata() *PropertyMetadata {
	return &PropertyMetadata{}
}
