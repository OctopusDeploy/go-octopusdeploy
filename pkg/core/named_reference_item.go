package octopusdeploy

type NamedReferenceItem struct {
	DisplayIDAndName bool   `json:"DisplayIdAndName,omitempty"`
	DisplayName      string `json:"DisplayName,omitempty"`
	ID               string `json:"Id,omitempty"`
}

func NewNamedReferenceItem() *NamedReferenceItem {
	return &NamedReferenceItem{}
}
