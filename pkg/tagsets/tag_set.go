package tagsets

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type TagSetScope string

const (
	TagSetScopeTenant      TagSetScope = "Tenant"
	TagSetScopeEnvironment TagSetScope = "Environment"
	TagSetScopeProject     TagSetScope = "Project"
	TagSetScopeRunbook     TagSetScope = "Runbook"
	TagSetScopeTarget      TagSetScope = "Target"
)

type TagSetType string

const (
	TagSetTypeSingleSelect TagSetType = "SingleSelect"
	TagSetTypeMultiSelect  TagSetType = "MultiSelect"
	TagSetTypeFreeText     TagSetType = "FreeText"
)

type TagSet struct {
	Description string        `json:"Description"`
	Name        string        `json:"Name"`
	Scopes      []TagSetScope `json:"Scopes,omitempty"`
	SortOrder   int32         `json:"SortOrder,omitempty"`
	SpaceID     string        `json:"SpaceId,omitempty"`
	Tags        []*Tag        `json:"Tags,omitempty"`
	Type        TagSetType    `json:"Type,omitempty"`

	resources.Resource
}

// NewTagSet initializes a TagSet with a name.
func NewTagSet(name string) *TagSet {
	return &TagSet{
		Name:     name,
		Scopes:   []TagSetScope{TagSetScopeTenant},
		Type:     TagSetTypeMultiSelect,
		Resource: *resources.NewResource(),
	}
}
