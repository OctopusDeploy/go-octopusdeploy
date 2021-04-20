package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
)

type Project struct {
	AutoCreateRelease               bool                        `json:"AutoCreateRelease"`
	AutoDeployReleaseOverrides      []AutoDeployReleaseOverride `json:"AutoDeployReleaseOverrides,omitempty"`
	ClonedFromProjectID             string                      `json:"ClonedFromProjectId,omitempty"`
	DefaultGuidedFailureMode        string                      `json:"DefaultGuidedFailureMode,omitempty"`
	DefaultToSkipIfAlreadyInstalled bool                        `json:"DefaultToSkipIfAlreadyInstalled"`
	DeploymentChangesTemplate       string                      `json:"DeploymentChangesTemplate,omitempty"`
	DeploymentProcessID             string                      `json:"DeploymentProcessId,omitempty"`
	Description                     string                      `json:"Description,omitempty"`
	ExtensionSettings               []ExtensionSettingsValues   `json:"ExtensionSettings,omitempty"`
	IncludedLibraryVariableSets     []string                    `json:"IncludedLibraryVariableSetIds,omitempty"`
	IsDisabled                      bool                        `json:"IsDisabled"`
	IsDiscreteChannelRelease        bool                        `json:"DiscreteChannelRelease"`
	IsVersionControlled             bool                        `json:"IsVersionControlled"`
	LifecycleID                     string                      `json:"LifecycleId" validate:"required"`
	Name                            string                      `json:"Name" validate:"required"`
	ConnectivityPolicy              *ConnectivityPolicy         `json:"ProjectConnectivityPolicy,omitempty"`
	ProjectGroupID                  string                      `json:"ProjectGroupId" validate:"required"`
	ReleaseCreationStrategy         *ReleaseCreationStrategy    `json:"ReleaseCreationStrategy,omitempty"`
	ReleaseNotesTemplate            string                      `json:"ReleaseNotesTemplate,omitempty"`
	Slug                            string                      `json:"Slug,omitempty"`
	SpaceID                         string                      `json:"SpaceId,omitempty"`
	Templates                       []ActionTemplateParameter   `json:"Templates,omitempty"`
	TenantedDeploymentMode          TenantedDeploymentMode      `json:"TenantedDeploymentMode"`
	VariableSetID                   string                      `json:"VariableSetId,omitempty"`
	VersionControlSettings          *VersionControlSettings     `json:"VersionControlSettings,omitempty"`
	VersioningStrategy              VersioningStrategy          `json:"VersioningStrategy"`

	resource
}

type Projects struct {
	Items []*Project `json:"Items"`
	PagedResults
}

func NewProject(name string, lifeCycleID string, projectGroupID string) *Project {
	return &Project{
		AutoDeployReleaseOverrides: []AutoDeployReleaseOverride{},
		DefaultGuidedFailureMode:   "EnvironmentDefault",
		ExtensionSettings:          []ExtensionSettingsValues{},
		LifecycleID:                lifeCycleID,
		Name:                       name,
		ConnectivityPolicy: &ConnectivityPolicy{
			AllowDeploymentsToNoTargets: false,
			SkipMachineBehavior:         "None",
		},
		ProjectGroupID:         projectGroupID,
		Templates:              []ActionTemplateParameter{},
		TenantedDeploymentMode: TenantedDeploymentMode("Untenanted"),
		VersioningStrategy: VersioningStrategy{
			Template: "#{Octopus.Version.LastMajor}.#{Octopus.Version.LastMinor}.#{Octopus.Version.NextPatch}",
		},
		resource: *newResource(),
	}
}

// Validate checks the state of the project and returns an error if invalid.
func (resource Project) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}

		return err
	}

	return ValidateMultipleProperties([]error{
		ValidatePropertyValues("DefaultGuidedFailureMode", resource.DefaultGuidedFailureMode, ValidProjectDefaultGuidedFailureModes),
		ValidateRequiredPropertyValue("LifecycleID", resource.LifecycleID),
		ValidateRequiredPropertyValue("Name", resource.Name),
		ValidateRequiredPropertyValue("ProjectGroupID", resource.ProjectGroupID),
	})
}
