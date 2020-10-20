package octopusdeploy

type TypeMetadata struct {
	Name       string              `json:"Name,omitempty"`
	Properties []*PropertyMetadata `json:"Properties"`
}
