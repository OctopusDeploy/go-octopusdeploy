package v2

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"

type EnvironmentV2 struct {
	ID               string   `json:"Id"`
	Name             string   `json:"Name"`
	SpaceID          string   `json:"SpaceId"`
	Slug             string   `json:"Slug"`
	Description      string   `json:"Description,omitempty"`
	Type             string   `json:"Type"`
	SortOrder        int      `json:"SortOrder"`
	UseGuidedFailure bool     `json:"UseGuidedFailure"`
	EnvironmentTags  []string `json:"EnvironmentTags,omitempty"`

	// Fields for Static environments
	AllowDynamicInfrastructure *bool                          `json:"AllowDynamicInfrastructure,omitempty"`
	ExtensionSettings          []extensions.ExtensionSettings `json:"ExtensionSettings,omitempty"`

	// Fields for Ephemeral environments
	ParentEnvironmentId string `json:"ParentEnvironmentId,omitempty"`
}
