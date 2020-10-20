package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
)

type Projects struct {
	Items []*Project `json:"Items"`
	PagedResults
}

type Project struct {
	AutoCreateRelease               bool                         `json:"AutoCreateRelease"`
	AutoDeployReleaseOverrides      []*AutoDeployReleaseOverride `json:"AutoDeployReleaseOverrides,omitempty"`
	ClonedFromProjectID             string                       `json:"ClonedFromProjectId,omitempty"`
	DefaultGuidedFailureMode        string                       `json:"DefaultGuidedFailureMode,omitempty"`
	DefaultToSkipIfAlreadyInstalled bool                         `json:"DefaultToSkipIfAlreadyInstalled"`
	DeploymentChangesTemplate       string                       `json:"DeploymentChangesTemplate,omitempty"`
	DeploymentProcessID             string                       `json:"DeploymentProcessId,omitempty"`
	Description                     string                       `json:"Description,omitempty"`
	DiscreteChannelRelease          bool                         `json:"DiscreteChannelRelease"`
	ExtensionSettings               []*ExtensionSettingsValues   `json:"ExtensionSettings,omitempty"`
	IncludedLibraryVariableSetIDs   []string                     `json:"IncludedLibraryVariableSetIds,omitempty"`
	IsDisabled                      bool                         `json:"IsDisabled"`
	IsVersionControlled             bool                         `json:"IsVersionControlled"`
	LifecycleID                     string                       `json:"LifecycleId" validate:"required"`
	Name                            string                       `json:"Name" validate:"required"`
	ProjectConnectivityPolicy       *ProjectConnectivityPolicy   `json:"ProjectConnectivityPolicy,omitempty"`
	ProjectGroupID                  string                       `json:"ProjectGroupId" validate:"required"`
	ReleaseCreationStrategy         *ReleaseCreationStrategy     `json:"ReleaseCreationStrategy,omitempty"`
	ReleaseNotesTemplate            string                       `json:"ReleaseNotesTemplate,omitempty"`
	Slug                            string                       `json:"Slug,omitempty"`
	SpaceID                         string                       `json:"SpaceId,omitempty"`
	Templates                       []*ActionTemplateParameter   `json:"Templates,omitempty"`
	TenantedDeploymentMode          string                       `json:"TenantedDeploymentMode" validate:"required,oneof=Untenanted TenantedOrUntenanted Tenanted"`
	VariableSetID                   string                       `json:"VariableSetId,omitempty"`
	VersionControlSettings          *VersionControlSettings      `json:"VersionControlSettings,omitempty"`
	VersioningStrategy              VersioningStrategy           `json:"VersioningStrategy"`

	Resource
}

func NewProject(name, lifeCycleID, projectGroupID string) *Project {
	return &Project{
		DefaultGuidedFailureMode: "EnvironmentDefault",
		LifecycleID:              lifeCycleID,
		Name:                     name,
		ProjectConnectivityPolicy: &ProjectConnectivityPolicy{
			AllowDeploymentsToNoTargets: false,
			SkipMachineBehavior:         "None",
		},
		ProjectGroupID:         projectGroupID,
		TenantedDeploymentMode: "Untenanted",
		VersioningStrategy: VersioningStrategy{
			Template: "#{Octopus.Version.LastMajor}.#{Octopus.Version.LastMinor}.#{Octopus.Version.NextPatch}",
		},
		Resource: *newResource(),
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
		ValidatePropertyValues("SkipMachineBehavior", resource.ProjectConnectivityPolicy.SkipMachineBehavior, ValidProjectConnectivityPolicySkipMachineBehaviors),
		ValidatePropertyValues("DefaultGuidedFailureMode", resource.DefaultGuidedFailureMode, ValidProjectDefaultGuidedFailureModes),
		ValidateRequiredPropertyValue("LifecycleID", resource.LifecycleID),
		ValidateRequiredPropertyValue("Name", resource.Name),
		ValidateRequiredPropertyValue("ProjectGroupID", resource.ProjectGroupID),
	})
}
