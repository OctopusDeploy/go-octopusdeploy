package model

import (
	"fmt"

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

func (p *Project) GetID() string {
	return p.ID
}

// Validate returns a collection of validation errors against the project's
// internal values.
func (p *Project) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}

		return err
	}

	return ValidateMultipleProperties([]error{
		ValidatePropertyValues("SkipMachineBehavior", p.ProjectConnectivityPolicy.SkipMachineBehavior, ValidProjectConnectivityPolicySkipMachineBehaviors),
		ValidatePropertyValues("DefaultGuidedFailureMode", p.DefaultGuidedFailureMode, ValidProjectDefaultGuidedFailureModes),
		ValidateRequiredPropertyValue("LifecycleID", p.LifecycleID),
		ValidateRequiredPropertyValue("Name", p.Name),
		ValidateRequiredPropertyValue("ProjectGroupID", p.ProjectGroupID),
	})

}

var _ ResourceInterface = &Project{}
