package tagsets

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type TagSet struct {
	Description string   `json:"Description"`
	Name        string   `json:"Name"`
	Scopes      []string `json:"Scopes,omitempty"`
	SortOrder   int32    `json:"SortOrder,omitempty"`
	SpaceID     string   `json:"SpaceId,omitempty"`
	Tags        []*Tag   `json:"Tags,omitempty"`
	Type        string   `json:"Type,omitempty"`

	resources.Resource
}

// NewTagSet initializes a TagSet with a name.
func NewTagSet(name string) *TagSet {
	return &TagSet{
		Name:     name,
		Scopes:   []string{"Tenant"},
		Type:     "MultiSelect",
		Resource: *resources.NewResource(),
	}
}
