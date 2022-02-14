package resources

type TagSet struct {
	Description string `json:"Description,omitempty"`
	Name        string `json:"Name"`
	SortOrder   int32  `json:"SortOrder,omitempty"`
	SpaceID     string `json:"SpaceId,omitempty"`
	Tags        []Tag  `json:"Tags,omitempty"`

	Resource
}

// NewTagSet initializes a TagSet with a name.
func NewTagSet(name string) *TagSet {
	return &TagSet{
		Name:     name,
		Resource: *NewResource(),
	}
}
