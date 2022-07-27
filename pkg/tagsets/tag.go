package tagsets

type Tag struct {
	CanonicalTagName string `json:"CanonicalTagName,omitempty"`
	Color            string `json:"Color"`
	Description      string `json:"Description"`
	ID               string `json:"Id,omitempty"`
	Name             string `json:"Name"`
	SortOrder        int    `json:"SortOrder,omitempty"`
}

// NewTag initializes a tag with a name and a color.
func NewTag(name string, color string) *Tag {
	return &Tag{
		Color: color,
		Name:  name,
	}
}
