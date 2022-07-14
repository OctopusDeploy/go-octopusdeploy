package tagsets

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"

type TagSets struct {
	Items []*TagSet `json:"Items"`
	resources.PagedResults
}

type TagSet struct {
	Description string `json:"Description,omitempty"`
	Name        string `json:"Name"`
	SortOrder   int32  `json:"SortOrder,omitempty"`
	SpaceID     string `json:"SpaceId,omitempty"`
	Tags        []Tag  `json:"Tags,omitempty"`

	resources.Resource
}

// NewTagSet initializes a TagSet with a name.
func NewTagSet(name string) *TagSet {
	return &TagSet{
		Name:     name,
		Resource: *resources.NewResource(),
	}
}
