package model

type ActionTemplates struct {
	Items []ActionTemplate `json:"Items"`
	PagedResults
}

type ActionTemplate struct {
	ActionType                *string                    `json:"ActionType"`
	CommunityActionTemplateID string                     `json:"CommunityActionTemplateId,omitempty"`
	Description               string                     `json:"Description,omitempty"`
	Name                      *string                    `json:"Name"`
	Packages                  []*PackageReference        `json:"Packages"`
	Parameters                []*ActionTemplateParameter `json:"Parameters"`
	Properties                map[string]PropertyValue   `json:"Properties,omitempty"`
	SpaceID                   string                     `json:"SpaceId,omitempty"`
	Version                   int32                      `json:"Version,omitempty"`
	Resource
}
