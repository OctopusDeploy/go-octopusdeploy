package resources

type EventGroup struct {
	EventCategories []string          `json:"EventCategories"`
	ID              string            `json:"Id,omitempty"`
	Links           map[string]string `json:"Links,omitempty"`
	Name            string            `json:"Name,omitempty"`
}
