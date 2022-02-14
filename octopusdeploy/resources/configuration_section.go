package resources

type ConfigurationSections struct {
	Items []*ConfigurationSection `json:"Items"`
	PagedResults
}

type ConfigurationSection struct {
	Description string `json:"Description,omitempty"`
	Name        string `json:"Name,omitempty"`

	Resource
}

func NewConfigurationSection() *ConfigurationSection {
	return &ConfigurationSection{
		Resource: *NewResource(),
	}
}
