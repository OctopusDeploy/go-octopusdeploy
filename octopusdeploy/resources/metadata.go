package resources

type Metadata struct {
	Description string          `json:"Description,omitempty"`
	Types       []*TypeMetadata `json:"Types"`
}

func NewMetadata() *Metadata {
	return &Metadata{}
}
