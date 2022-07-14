package configuration

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"

type ConfigurationSections struct {
	Items []*ConfigurationSection `json:"Items"`
	resources.PagedResults
}

type ConfigurationSection struct {
	Description string `json:"Description,omitempty"`
	Name        string `json:"Name,omitempty"`

	resources.Resource
}

func NewConfigurationSection() *ConfigurationSection {
	return &ConfigurationSection{
		Resource: *resources.NewResource(),
	}
}
