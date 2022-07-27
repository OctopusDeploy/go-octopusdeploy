package configuration

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

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
