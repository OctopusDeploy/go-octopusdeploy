package actiontemplates

// ActionTemplateSearch represents an action template search.
type ActionTemplateSearch struct {
	Author                    string            `json:"Author,omitempty"`
	Categories                []string          `json:"Categories"`
	Category                  string            `json:"Category,omitempty"`
	CommunityActionTemplateID string            `json:"CommunityActionTemplateId,omitempty"`
	Description               string            `json:"Description,omitempty"`
	HasUpdate                 bool              `json:"HasUpdate"`
	ID                        string            `json:"Id,omitempty"`
	IsBuiltIn                 bool              `json:"IsBuiltIn"`
	IsInstalled               bool              `json:"IsInstalled"`
	Keywords                  string            `json:"Keywords,omitempty"`
	Links                     map[string]string `json:"Links,omitempty"`
	Name                      string            `json:"Name,omitempty"`
	Type                      string            `json:"Type,omitempty"`
	Website                   string            `json:"Website,omitempty"`
}
