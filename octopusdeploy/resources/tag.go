package resources

type Tag struct {
	ID               string `json:"Id,omitempty"`
	Name             string `json:"Name,omitempty"`
	Color            string `json:"Color,omitempty"`
	CanonicalTagName string `json:"CanonicalTagName,omitempty"`
	Description      string `json:"Description,omitempty"`
	SortOrder        int    `json:"SortOrder"`
}
