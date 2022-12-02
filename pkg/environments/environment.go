package environments

import (
	"encoding/json"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type Environment struct {
	AllowDynamicInfrastructure bool                           `json:"AllowDynamicInfrastructure"`
	Description                string                         `json:"Description,omitempty"`
	ExtensionSettings          []extensions.ExtensionSettings `json:"ExtensionSettings,omitempty"`
	Name                       string                         `json:"Name" validate:"required"`
	Slug                       string                         `json:"Slug"` // deliberately send empty string
	SortOrder                  int                            `json:"SortOrder"`
	SpaceID                    string                         `json:"SpaceId"`
	UseGuidedFailure           bool                           `json:"UseGuidedFailure"`

	resources.Resource
}

func NewEnvironment(name string) *Environment {
	return &Environment{
		AllowDynamicInfrastructure: false,
		ExtensionSettings:          []extensions.ExtensionSettings{},
		Name:                       name,
		SortOrder:                  0,
		UseGuidedFailure:           false,
		Resource:                   *resources.NewResource(),
	}
}

// UnmarshalJSON sets an environment to its representation in JSON.
func (e *Environment) UnmarshalJSON(data []byte) error {
	var fields struct {
		AllowDynamicInfrastructure bool   `json:"AllowDynamicInfrastructure"`
		Description                string `json:"Description,omitempty"`
		Name                       string `json:"Name" validate:"required"`
		Slug                       string `json:"Slug"`
		SortOrder                  int    `json:"SortOrder"`
		SpaceID                    string `json:"SpaceId"`
		UseGuidedFailure           bool   `json:"UseGuidedFailure"`
		resources.Resource
	}

	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	// validate JSON representation
	validate := validator.New()
	if err := validate.Struct(fields); err != nil {
		return err
	}

	e.AllowDynamicInfrastructure = fields.AllowDynamicInfrastructure
	e.Description = fields.Description
	e.Name = fields.Name
	e.Slug = fields.Slug
	e.SortOrder = fields.SortOrder
	e.SpaceID = fields.SpaceID
	e.UseGuidedFailure = fields.UseGuidedFailure
	e.Resource = fields.Resource

	var environment map[string]*json.RawMessage
	if err := json.Unmarshal(data, &environment); err != nil {
		return err
	}

	var extensionSettings *json.RawMessage
	var extensionSettingsCollection []*json.RawMessage

	if environment["ExtensionSettings"] != nil {
		extensionSettingsValue := environment["ExtensionSettings"]

		if err := json.Unmarshal(*extensionSettingsValue, &extensionSettings); err != nil {
			return err
		}

		if err := json.Unmarshal(*extensionSettings, &extensionSettingsCollection); err != nil {
			return err
		}

		for _, v := range extensionSettingsCollection {
			var extensionSettingsItem map[string]*json.RawMessage
			if err := json.Unmarshal(*v, &extensionSettingsItem); err != nil {
				return err
			}

			if extensionSettingsItem["ExtensionId"] != nil {
				var extensionID extensions.ExtensionID
				json.Unmarshal(*extensionSettingsItem["ExtensionId"], &extensionID)

				switch extensionID {
				case extensions.ExtensionIDJira:
					var jiraExtensionSettings *JiraExtensionSettings
					if err := json.Unmarshal(*v, &jiraExtensionSettings); err != nil {
						return err
					}
					e.ExtensionSettings = append(e.ExtensionSettings, jiraExtensionSettings)
				case extensions.ExtensionIDJiraServiceManagement:
					var jiraServiceManagementExtensionSettings *JiraServiceManagementExtensionSettings
					if err := json.Unmarshal(*v, &jiraServiceManagementExtensionSettings); err != nil {
						return err
					}
					e.ExtensionSettings = append(e.ExtensionSettings, jiraServiceManagementExtensionSettings)
				case extensions.ExtensionIDServiceNow:
					var serviceNowExtensionSettings *ServiceNowExtensionSettings
					if err := json.Unmarshal(*v, &serviceNowExtensionSettings); err != nil {
						return err
					}
					e.ExtensionSettings = append(e.ExtensionSettings, serviceNowExtensionSettings)
				}
			}
		}
	}

	return nil
}

// Validate checks the state of the environment and returns an error if
// invalid.
func (e *Environment) Validate() error {
	return validator.New().Struct(e)
}

// GetName returns the name of the environment.
func (e *Environment) GetName() string {
	return e.Name
}

// SetName sets the name of the environment.
func (e *Environment) SetName(name string) {
	e.Name = name
}

var _ resources.IHasName = &Environment{}
