package model

type Metadata struct {
	Description string          `json:"Description,omitempty"`
	Types       []*TypeMetadata `json:"Types"`
}
