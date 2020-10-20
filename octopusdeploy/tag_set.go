package octopusdeploy

type TagSets struct {
	Items []*TagSet `json:"Items"`
	PagedResults
}

type TagSet struct {
	Name string `json:"Name"`
	Tags []Tag  `json:"Tags,omitempty"`

	Resource
}

// NewTagSet initializes a TagSet with a name.
func NewTagSet(name string) *TagSet {
	return &TagSet{
		Name:     name,
		Resource: *newResource(),
	}
}
