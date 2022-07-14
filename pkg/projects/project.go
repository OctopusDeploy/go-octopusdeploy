package projects

import (
	"encoding/json"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actiontemplates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type Project struct {
	AutoCreateRelease               bool                                      `json:"AutoCreateRelease,omitempty"`
	AutoDeployReleaseOverrides      []AutoDeployReleaseOverride               `json:"AutoDeployReleaseOverrides,omitempty"`
	ClonedFromProjectID             string                                    `json:"ClonedFromProjectId,omitempty"`
	ConnectivityPolicy              *core.ConnectivityPolicy                  `json:"ProjectConnectivityPolicy,omitempty"`
	DefaultGuidedFailureMode        string                                    `json:"DefaultGuidedFailureMode,omitempty"`
	DefaultToSkipIfAlreadyInstalled bool                                      `json:"DefaultToSkipIfAlreadyInstalled,omitempty"`
	DeploymentChangesTemplate       string                                    `json:"DeploymentChangesTemplate,omitempty"`
	DeploymentProcessID             string                                    `json:"DeploymentProcessId,omitempty"`
	Description                     string                                    `json:"Description,omitempty"`
	ExtensionSettings               []ExtensionSettingsValues                 `json:"ExtensionSettings,omitempty"`
	IncludedLibraryVariableSets     []string                                  `json:"IncludedLibraryVariableSetIds,omitempty"`
	IsDisabled                      bool                                      `json:"IsDisabled,omitempty"`
	IsDiscreteChannelRelease        bool                                      `json:"DiscreteChannelRelease,omitempty"`
	IsVersionControlled             bool                                      `json:"IsVersionControlled,omitempty"`
	LifecycleID                     string                                    `json:"LifecycleId" validate:"required"`
	Name                            string                                    `json:"Name" validate:"required"`
	PersistenceSettings             IPersistenceSettings                      `json:"PersistenceSettings,omitempty"`
	ProjectGroupID                  string                                    `json:"ProjectGroupId" validate:"required"`
	ReleaseCreationStrategy         *ReleaseCreationStrategy                  `json:"ReleaseCreationStrategy,omitempty"`
	ReleaseNotesTemplate            string                                    `json:"ReleaseNotesTemplate,omitempty"`
	Slug                            string                                    `json:"Slug,omitempty"`
	SpaceID                         string                                    `json:"SpaceId,omitempty"`
	Templates                       []actiontemplates.ActionTemplateParameter `json:"Templates,omitempty"`
	TenantedDeploymentMode          core.TenantedDeploymentMode               `json:"TenantedDeploymentMode,omitempty"`
	VariableSetID                   string                                    `json:"VariableSetId,omitempty"`
	VersioningStrategy              *VersioningStrategy                       `json:"VersioningStrategy,omitempty"`

	resources.Resource
}

type Projects struct {
	Items []*Project `json:"Items"`
	resources.PagedResults
}

func NewProjects() *Projects {
	return &Projects{
		Items: []*Project{},
		PagedResults: resources.PagedResults{
			ItemType:       "Project",
			LastPageNumber: 0,
			NumberOfPages:  1,
			ItemsPerPage:   30,
			TotalResults:   0,
		},
	}
}

func NewProject(name string, lifecycleID string, projectGroupID string) *Project {
	return &Project{
		LifecycleID:    lifecycleID,
		Name:           name,
		ProjectGroupID: projectGroupID,
		Resource:       *resources.NewResource(),
	}
}

// UnmarshalJSON sets this project to its representation in JSON.
func (p *Project) UnmarshalJSON(data []byte) error {
	var fields struct {
		AutoCreateRelease               bool                                      `json:"AutoCreateRelease,omitempty"`
		AutoDeployReleaseOverrides      []AutoDeployReleaseOverride               `json:"AutoDeployReleaseOverrides,omitempty"`
		ClonedFromProjectID             string                                    `json:"ClonedFromProjectId,omitempty"`
		ConnectivityPolicy              *core.ConnectivityPolicy                  `json:"ProjectConnectivityPolicy,omitempty"`
		DefaultGuidedFailureMode        string                                    `json:"DefaultGuidedFailureMode,omitempty"`
		DefaultToSkipIfAlreadyInstalled bool                                      `json:"DefaultToSkipIfAlreadyInstalled,omitempty"`
		DeploymentChangesTemplate       string                                    `json:"DeploymentChangesTemplate,omitempty"`
		DeploymentProcessID             string                                    `json:"DeploymentProcessId,omitempty"`
		Description                     string                                    `json:"Description,omitempty"`
		ExtensionSettings               []ExtensionSettingsValues                 `json:"ExtensionSettings,omitempty"`
		IncludedLibraryVariableSets     []string                                  `json:"IncludedLibraryVariableSetIds,omitempty"`
		IsDisabled                      bool                                      `json:"IsDisabled,omitempty"`
		IsDiscreteChannelRelease        bool                                      `json:"DiscreteChannelRelease,omitempty"`
		IsVersionControlled             bool                                      `json:"IsVersionControlled,omitempty"`
		LifecycleID                     string                                    `json:"LifecycleId" validate:"required"`
		Name                            string                                    `json:"Name" validate:"required"`
		ProjectGroupID                  string                                    `json:"ProjectGroupId" validate:"required"`
		ReleaseCreationStrategy         *ReleaseCreationStrategy                  `json:"ReleaseCreationStrategy,omitempty"`
		ReleaseNotesTemplate            string                                    `json:"ReleaseNotesTemplate,omitempty"`
		Slug                            string                                    `json:"Slug,omitempty"`
		SpaceID                         string                                    `json:"SpaceId,omitempty"`
		Templates                       []actiontemplates.ActionTemplateParameter `json:"Templates,omitempty"`
		TenantedDeploymentMode          core.TenantedDeploymentMode               `json:"TenantedDeploymentMode,omitempty"`
		VariableSetID                   string                                    `json:"VariableSetId,omitempty"`
		VersioningStrategy              *VersioningStrategy                       `json:"VersioningStrategy,omitempty"`
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

	p.AutoCreateRelease = fields.AutoCreateRelease
	p.AutoDeployReleaseOverrides = fields.AutoDeployReleaseOverrides
	p.ClonedFromProjectID = fields.ClonedFromProjectID
	p.ConnectivityPolicy = fields.ConnectivityPolicy
	p.DefaultGuidedFailureMode = fields.DefaultGuidedFailureMode
	p.DefaultToSkipIfAlreadyInstalled = fields.DefaultToSkipIfAlreadyInstalled
	p.DeploymentChangesTemplate = fields.DeploymentChangesTemplate
	p.DeploymentProcessID = fields.DeploymentProcessID
	p.Description = fields.Description
	p.ExtensionSettings = fields.ExtensionSettings
	p.IncludedLibraryVariableSets = fields.IncludedLibraryVariableSets
	p.IsDisabled = fields.IsDisabled
	p.IsDiscreteChannelRelease = fields.IsDiscreteChannelRelease
	p.IsVersionControlled = fields.IsVersionControlled
	p.LifecycleID = fields.LifecycleID
	p.Name = fields.Name
	p.ProjectGroupID = fields.ProjectGroupID
	p.ReleaseCreationStrategy = fields.ReleaseCreationStrategy
	p.ReleaseNotesTemplate = fields.ReleaseNotesTemplate
	p.Slug = fields.Slug
	p.SpaceID = fields.SpaceID
	p.Templates = fields.Templates
	p.TenantedDeploymentMode = fields.TenantedDeploymentMode
	p.VariableSetID = fields.VariableSetID
	p.VersioningStrategy = fields.VersioningStrategy
	p.Resource = fields.Resource

	var project map[string]*json.RawMessage
	if err := json.Unmarshal(data, &project); err != nil {
		return err
	}

	var persistenceSettings *json.RawMessage
	var persistenceSettingsProperties map[string]*json.RawMessage
	var persistenceSettingsType string

	if project["PersistenceSettings"] != nil {
		persistenceSettingsValue := project["PersistenceSettings"]

		if err := json.Unmarshal(*persistenceSettingsValue, &persistenceSettings); err != nil {
			return err
		}

		if err := json.Unmarshal(*persistenceSettings, &persistenceSettingsProperties); err != nil {
			return err
		}

		if persistenceSettingsProperties["Type"] != nil {
			pst := persistenceSettingsProperties["Type"]
			json.Unmarshal(*pst, &persistenceSettingsType)
		}
	}

	switch persistenceSettingsType {
	case "Database":
		var databasePersistenceSettings *DatabasePersistenceSettings
		if err := json.Unmarshal(*persistenceSettings, &databasePersistenceSettings); err != nil {
			return err
		}
		p.PersistenceSettings = databasePersistenceSettings
	case "VersionControlled":
		var gitPersistenceSettings *GitPersistenceSettings
		if err := json.Unmarshal(*persistenceSettings, &gitPersistenceSettings); err != nil {
			return err
		}
		p.PersistenceSettings = gitPersistenceSettings
	}

	return nil
}

// Validate checks the state of the project and returns an error if invalid.
func (resource Project) Validate() error {
	err := validator.New().Struct(resource)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}

		return err
	}

	return internal.ValidateMultipleProperties([]error{
		internal.ValidateRequiredPropertyValue("LifecycleID", resource.LifecycleID),
		internal.ValidateRequiredPropertyValue("Name", resource.Name),
		internal.ValidateRequiredPropertyValue("ProjectGroupID", resource.ProjectGroupID),
	})
}
