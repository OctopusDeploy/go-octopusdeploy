package model

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
)

type Projects struct {
	Items []Project `json:"Items"`
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
	LifecycleID                     string                       `json:"LifecycleId"`
	Name                            string                       `json:"Name,omitempty"`
	ProjectConnectivityPolicy       *ProjectConnectivityPolicy   `json:"ProjectConnectivityPolicy,omitempty"`
	ProjectGroupID                  string                       `json:"ProjectGroupId,omitempty"`
	ReleaseCreationStrategy         *ReleaseCreationStrategy     `json:"ReleaseCreationStrategy,omitempty"`
	ReleaseNotesTemplate            string                       `json:"ReleaseNotesTemplate,omitempty"`
	Slug                            string                       `json:"Slug,omitempty"`
	SpaceID                         string                       `json:"SpaceId,omitempty"`
	Templates                       []*ActionTemplateParameter   `json:"Templates,omitempty"`
	TenantedDeploymentMode          enum.TenantedDeploymentMode  `json:"TenantedDeploymentMode"`
	VariableSetID                   string                       `json:"VariableSetId,omitempty"`
	VersionControlSettings          *VersionControlSettings      `json:"VersionControlSettings,omitempty"`
	VersioningStrategy              VersioningStrategy           `json:"VersioningStrategy"`

	Resource
}

func NewProject(name, lifeCycleID, projectGroupID string) *Project {
	return &Project{
		Name:                     name,
		DefaultGuidedFailureMode: "EnvironmentDefault",
		LifecycleID:              lifeCycleID,
		ProjectGroupID:           projectGroupID,
		VersioningStrategy: VersioningStrategy{
			Template: "#{Octopus.Version.LastMajor}.#{Octopus.Version.LastMinor}.#{Octopus.Version.NextPatch}",
		},
		ProjectConnectivityPolicy: &ProjectConnectivityPolicy{
			AllowDeploymentsToNoTargets: false,
			SkipMachineBehavior:         "None",
		},
	}
}

// GetID returns the ID value of the Project.
func (resource Project) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Project.
func (resource Project) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Project was changed.
func (resource Project) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Project.
func (resource Project) GetLinks() map[string]string {
	return resource.Links
}

func (resource Project) SetID(id string) {
	resource.ID = id
}

func (resource Project) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

func (resource Project) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the Project and returns an error if invalid.
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

var _ ResourceInterface = &Project{}
