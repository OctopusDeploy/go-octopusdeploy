package configuration

type FeatureToggleConfigurationQuery struct {
	Name string `uri:"Name,omitempty" url:"Name,omitempty"`
}

type FeatureToggleConfigurationResponse struct {
	FeatureToggles []ConfiguredFeatureToggle `json:"FeatureToggles"`
}

type ConfiguredFeatureToggle struct {
	Name      string `json:"Name"`
	IsEnabled bool   `json:"IsEnabled"`
}
