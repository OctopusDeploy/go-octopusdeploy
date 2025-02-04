package configuration

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type FeatureToggleConfigurationQuery struct {
	Name string `json:"Name,omitempty"`

	resources.Resource
}

type FeatureToggleConfigurationResponse struct {
	FeatureToggles []ConfiguredFeatureToggle `json:"FeatureToggles"`

	resources.Resource
}

type ConfiguredFeatureToggle struct {
	Name      string `json:"Name"`
	IsEnabled bool   `json:"IsEnabled"`
}
