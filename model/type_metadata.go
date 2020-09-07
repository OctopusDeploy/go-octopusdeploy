package model

type TypeMetadata struct {
	Name       string              `json:"Name,omitempty"`
	Properties []*PropertyMetadata `json:"Properties"`
}
