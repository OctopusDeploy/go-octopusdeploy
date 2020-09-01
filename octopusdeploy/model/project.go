package model

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/enum"

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
	IncludedLibraryVariableSetIds   []string                     `json:"IncludedLibraryVariableSetIds,omitempty"`
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

// ValidateProjectValues checks the values of a Project object to see if they are suitable for
// sending to Octopus Deploy. Used when adding or updating projects.
func ValidateProjectValues(Project *Project) error {
	return ValidateMultipleProperties([]error{
		ValidatePropertyValues("SkipMachineBehavior", Project.ProjectConnectivityPolicy.SkipMachineBehavior, ValidProjectConnectivityPolicySkipMachineBehaviors),
		ValidatePropertyValues("DefaultGuidedFailureMode", Project.DefaultGuidedFailureMode, ValidProjectDefaultGuidedFailureModes),
		ValidateRequiredPropertyValue("LifecycleID", Project.LifecycleID),
		ValidateRequiredPropertyValue("Name", Project.Name),
		ValidateRequiredPropertyValue("ProjectGroupID", Project.ProjectGroupID),
	})
}
