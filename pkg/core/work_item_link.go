package core

type WorkItemLink struct {
	Description string `json:"Description,omitempty"`
	ID          string `json:"Id,omitempty"`
	LinkURL     string `json:"LinkUrl,omitempty"`
	Source      string `json:"Source,omitempty"`
}

func NewWorkItemLink() *WorkItemLink {
	return &WorkItemLink{}
}
