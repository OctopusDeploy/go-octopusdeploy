package configuration

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type FeatureToggleConfigurationQuery struct {
	Name string `uri:"Name,omitempty" url:"Name,omitempty"`
}

type ConfiguredFeatureToggle struct {
	Name      string `json:"Name"`
	IsEnabled bool   `json:"IsEnabled"`

	resources.Resource
}
