package octopusdeploy

type TagSets struct {
	Items []*TagSet `json:"Items"`
	PagedResults
}

type TagSet struct {
	Description string `json:"Description,omitempty"`
	Name        string `json:"Name"`
	SortOrder   int32  `json:"SortOrder,omitempty"`
	SpaceID     string `json:"SpaceId,omitempty"`
	Tags        []Tag  `json:"Tags,omitempty"`

	resource
}

// NewTagSet initializes a TagSet with a name.
func NewTagSet(name string) *TagSet {
	return &TagSet{
		Name:     name,
		resource: *newResource(),
	}
}
