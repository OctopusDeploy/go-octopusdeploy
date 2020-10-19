package model

type ProjectGroups struct {
	Items []*ProjectGroup `json:"Items"`
	PagedResults
}

type ProjectGroup struct {
	Description       string   `json:"Description,omitempty"`
	EnvironmentIDs    []string `json:"EnvironmentIds"`
	Name              string   `json:"Name,omitempty" validate:"required"`
	RetentionPolicyID string   `json:"RetentionPolicyId,omitempty"`

	Resource
}

func NewProjectGroup(name string) *ProjectGroup {
	return &ProjectGroup{
		Name:     name,
		Resource: *newResource(),
	}
}
