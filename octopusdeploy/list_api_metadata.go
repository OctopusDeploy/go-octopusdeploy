package octopusdeploy

type ListAPIMetadata struct {
	APIEndpoint string `json:"ApiEndpoint,omitempty"`
	SelectMode  string `json:"SelectMode,omitempty"`
}

func NewListAPIMetadata() *ListAPIMetadata {
	return &ListAPIMetadata{}
}
