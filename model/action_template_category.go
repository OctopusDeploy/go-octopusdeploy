package model

// ActionTemplateCategory represents an action template category.
type ActionTemplateCategory struct {
	DisplayOrder int32             `json:"DisplayOrder,omitempty"`
	ID           string            `json:"Id,omitempty"`
	Links        map[string]string `json:"Links,omitempty"`
	Name         string            `json:"Name,omitempty"`
}
