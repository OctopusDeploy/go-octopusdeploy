package resources

type ProjectedTeamReferenceDataItem struct {
	ExternalSecurityGroups []NamedReferenceItem `json:"ExternalSecurityGroups"`
	ID                     string               `json:"Id,omitempty"`
	IsDirectlyAssigned     bool                 `json:"IsDirectlyAssigned,omitempty"`
	Name                   string               `json:"Name,omitempty"`
	SpaceID                string               `json:"SpaceId,omitempty"`
}
