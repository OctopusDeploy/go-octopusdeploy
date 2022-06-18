package metadata

type TypeMetadata struct {
	Name       string              `json:"Name,omitempty"`
	Properties []*PropertyMetadata `json:"Properties"`
}

func NewTypeMetadata() *TypeMetadata {
	return &TypeMetadata{}
}
