package model

type CommunityActionTemplates struct {
	Items []CommunityActionTemplate `json:"Items"`
	PagedResults
}

type CommunityActionTemplate struct {
	Author      string                             `json:"Author,omitempty"`
	Description string                             `json:"Description,omitempty"`
	HistoryURL  string                             `json:"HistoryUrl,omitempty"`
	ID          string                             `json:"Id,omitempty"`
	Links       map[string]string                  `json:"Links,omitempty"`
	Name        string                             `json:"Name,omitempty"`
	Parameters  []*ActionTemplateParameterResource `json:"Parameters"`
	Properties  map[string]PropertyValueResource   `json:"Properties,omitempty"`
	Type        string                             `json:"Type,omitempty"`
	Version     int32                              `json:"Version,omitempty"`
	Website     string                             `json:"Website,omitempty"`
}
