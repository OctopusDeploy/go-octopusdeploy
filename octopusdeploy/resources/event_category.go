package resources

type EventCategory struct {
	ID    string            `json:"Id,omitempty"`
	Links map[string]string `json:"Links,omitempty"`
	Name  string            `json:"Name,omitempty"`
}
