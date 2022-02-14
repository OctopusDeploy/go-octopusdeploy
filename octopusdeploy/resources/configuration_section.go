package resources

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
