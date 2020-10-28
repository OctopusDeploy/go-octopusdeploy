package octopusdeploy

type ConfigurationSections struct {
	Items []*ConfigurationSection `json:"Items"`
	PagedResults
}

type ConfigurationSection struct {
	Description string `json:"Description,omitempty"`
	Name        string `json:"Name,omitempty"`

	resource
}

func NewConfigurationSection() *ConfigurationSection {
	return &ConfigurationSection{
		resource: *newResource(),
	}
}
