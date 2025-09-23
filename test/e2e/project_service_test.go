package e2e

import (
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/lifecycles"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projectgroups"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/spaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualProjects(t *testing.T, expected *projects.Project, actual *projects.Project) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	// IResource
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.True(t, internal.IsLinksEqual(expected.GetLinks(), actual.GetLinks()))

	// Project
	assert.Equal(t, expected.AllowIgnoreChannelRules, actual.AllowIgnoreChannelRules)
	assert.Equal(t, expected.AutoCreateRelease, actual.AutoCreateRelease)
	assert.Equal(t, expected.AutoDeployReleaseOverrides, actual.AutoDeployReleaseOverrides)
	assert.Equal(t, expected.ClonedFromProjectID, actual.ClonedFromProjectID)
	assert.Equal(t, expected.ConnectivityPolicy, actual.ConnectivityPolicy)
	assert.Equal(t, expected.DefaultGuidedFailureMode, actual.DefaultGuidedFailureMode)
	assert.Equal(t, expected.DefaultToSkipIfAlreadyInstalled, actual.DefaultToSkipIfAlreadyInstalled)
	assert.Equal(t, expected.DeploymentChangesTemplate, actual.DeploymentChangesTemplate)
	assert.Equal(t, expected.DeploymentProcessID, actual.DeploymentProcessID)
	assert.Equal(t, expected.DeprovisioningRunbookID, actual.DeprovisioningRunbookID)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.ExtensionSettings, actual.ExtensionSettings)
	assert.Equal(t, expected.IncludedLibraryVariableSets, actual.IncludedLibraryVariableSets)
	assert.Equal(t, expected.IsDisabled, actual.IsDisabled)
	assert.Equal(t, expected.IsDiscreteChannelRelease, actual.IsDiscreteChannelRelease)
	assert.Equal(t, expected.IsVersionControlled, actual.IsVersionControlled)
	assert.Equal(t, expected.LifecycleID, actual.LifecycleID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.ProjectGroupID, actual.ProjectGroupID)
	assert.Equal(t, expected.ProvisioningRunbookID, actual.ProvisioningRunbookID)
	assert.Equal(t, expected.ReleaseCreationStrategy, actual.ReleaseCreationStrategy)
	assert.Equal(t, expected.ReleaseNotesTemplate, actual.ReleaseNotesTemplate)
	assert.Equal(t, expected.Slug, actual.Slug)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
	assert.Equal(t, expected.Templates, actual.Templates)
	assert.Equal(t, expected.TenantedDeploymentMode, actual.TenantedDeploymentMode)
	assert.Equal(t, expected.VariableSetID, actual.VariableSetID)
}

func CreateTestProject(t *testing.T, client *client.Client, space *spaces.Space, lifecycle *lifecycles.Lifecycle, projectGroup *projectgroups.ProjectGroup) *projects.Project {
	require.NotNil(t, space)
	require.NotNil(t, lifecycle)
	require.NotNil(t, projectGroup)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	project := projects.NewProject(name, lifecycle.GetID(), projectGroup.GetID())
	require.NotNil(t, project)

	createdProject, err := client.Projects.Add(project)
	require.NoError(t, err)
	require.NotNil(t, createdProject)
	require.NotEmpty(t, createdProject.GetID())

	// verify the add operation was successful
	projectToCompare, err := client.Projects.GetByID(createdProject.GetID())
	require.NoError(t, err)
	require.NotNil(t, projectToCompare)
	AssertEqualProjects(t, createdProject, projectToCompare)

	return createdProject
}

func DeleteTestProject(t *testing.T, client *client.Client, project *projects.Project) {
	require.NotNil(t, project)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Projects.DeleteByID(project.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedProject, err := client.Projects.GetByID(project.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedProject)
}

func TestProjectAddWithPersistenceSettings(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	name := internal.GetRandomName()

	project := projects.NewProject(name, lifecycle.GetID(), projectGroup.GetID())
	require.NotNil(t, project)

	basePath := internal.GetRandomName()
	credentials := credentials.NewAnonymous()
	defaultBranch := "main"
	protectedBranchNamePatterns := []string{}
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)

	project.PersistenceSettings = projects.NewGitPersistenceSettings(basePath, credentials, defaultBranch, protectedBranchNamePatterns, url)

	createdProject, err := client.Projects.Add(project)
	require.NoError(t, err)
	require.NotNil(t, createdProject)
	require.NotEmpty(t, createdProject.GetID())

	defer DeleteTestProject(t, client, createdProject)

	// verify the add operation was successful
	projectToCompare, err := client.Projects.GetByID(createdProject.GetID())
	require.NoError(t, err)
	require.NotNil(t, projectToCompare)
	AssertEqualProjects(t, createdProject, projectToCompare)
}

func TestProjectAddGetDelete(t *testing.T) {
	octopus := getOctopusClient()
	require.NotNil(t, octopus)

	space := GetDefaultSpace(t, octopus)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, octopus)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, octopus, lifecycle)

	projectGroup := CreateTestProjectGroup(t, octopus)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, octopus, projectGroup)

	project := CreateTestProject(t, octopus, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, octopus, project)
}

func TestProjectServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	projects, err := client.Projects.GetAll()
	require.NoError(t, err)
	require.NotNil(t, projects)

	for _, project := range projects {
		if project.Name != "Test Project" {
			defer DeleteTestProject(t, client, project)
		}
	}
}

func TestProjectGetThatDoesNotExist(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := internal.GetRandomName()
	resource, err := client.Projects.GetByID(id)
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestProjectGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	// create many projects to test pagination
	projectsToCreate := 32
	sum := 0
	for i := 0; i < projectsToCreate; i++ {
		project := CreateTestProject(t, client, space, lifecycle, projectGroup)
		require.NotNil(t, project)
		defer DeleteTestProject(t, client, project)

		sum += i
	}

	allProjects, err := client.Projects.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	numberOfProjects := len(allProjects)

	// check there are greater than or equal to the amount of projects requested to be created, otherwise pagination isn't working
	if numberOfProjects < projectsToCreate {
		t.Fatalf("There should be at least %d projects created but there was only %d. Pagination is likely not working.", projectsToCreate, numberOfProjects)
	}

	project := CreateTestProject(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	allProjectsAfterCreatingAdditional, err := client.Projects.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	assert.NoError(t, err, "error when looking for project when not expected")
	assert.Equal(t, len(allProjectsAfterCreatingAdditional), numberOfProjects+1, "created an additional project and expected number of projects to increase by 1")
}

func TestProjectUpdate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	newProjectName := internal.GetRandomName()
	newDescription := internal.GetRandomName()
	newSkipMachineBehavior := core.SkipMachineBehaviorNone

	project.Name = newProjectName
	project.ConnectivityPolicy.SkipMachineBehavior = newSkipMachineBehavior
	project.Description = newDescription

	updatedProject, err := client.Projects.Update(project)
	require.NoError(t, err)
	require.Equal(t, newProjectName, updatedProject.Name, "project name was not updated")
	require.Equal(t, newDescription, updatedProject.Description, "project description was not updated")
	require.Equal(t, newSkipMachineBehavior, project.ConnectivityPolicy.SkipMachineBehavior, "project connectivity policy name was not updated")
}

func TestProjectGetByName(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	query := projects.ProjectsQuery{
		Name: project.Name,
		Take: 1,
	}

	projects, err := client.Projects.Get(query)
	require.NoError(t, err)
	require.NotNil(t, projects)
}

// ----- new -----
// todo: There are a few client calls that use the old client as they have not been migrated over yet;
// 		 revisit and update these to use the new client calls once migrated.

func CreateTestProject_NewClient(t *testing.T, client *client.Client, space *spaces.Space, lifecycle *lifecycles.Lifecycle, projectGroup *projectgroups.ProjectGroup) *projects.Project {
	require.NotNil(t, space)
	require.NotNil(t, lifecycle)
	require.NotNil(t, projectGroup)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	project := projects.NewProject(name, lifecycle.GetID(), projectGroup.GetID())
	require.NotNil(t, project)

	createdProject, err := projects.Add(client, project)
	require.NoError(t, err)
	require.NotNil(t, createdProject)
	require.NotEmpty(t, createdProject.GetID())

	// verify the add operation was successful
	projectToCompare, err := projects.GetByID(client, project.SpaceID, createdProject.GetID())
	require.NoError(t, err)
	require.NotNil(t, projectToCompare)
	AssertEqualProjects(t, createdProject, projectToCompare)

	return createdProject
}

func DeleteTestProject_NewClient(t *testing.T, client *client.Client, project *projects.Project) {
	require.NotNil(t, project)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := projects.DeleteByID(client, project.SpaceID, project.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedProject, err := projects.GetByID(client, project.SpaceID, project.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedProject)
}

func TestProjectAddGetDelete_NewClient(t *testing.T) {
	octopus := getOctopusClient()
	require.NotNil(t, octopus)

	space := GetDefaultSpace(t, octopus)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, octopus)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, octopus, lifecycle)

	projectGroup := CreateTestProjectGroup(t, octopus)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, octopus, projectGroup)

	project := CreateTestProject_NewClient(t, octopus, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject_NewClient(t, octopus, project)
}

func TestProjectUpdate_NewClient(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject_NewClient(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	newProjectName := internal.GetRandomName()
	newDescription := internal.GetRandomName()
	newSkipMachineBehavior := core.SkipMachineBehaviorNone

	project.Name = newProjectName
	project.ConnectivityPolicy.SkipMachineBehavior = newSkipMachineBehavior
	project.Description = newDescription

	updatedProject, err := projects.Update(client, project)
	require.NoError(t, err)
	require.Equal(t, newProjectName, updatedProject.Name, "project name was not updated")
	require.Equal(t, newDescription, updatedProject.Description, "project description was not updated")
	require.Equal(t, newSkipMachineBehavior, project.ConnectivityPolicy.SkipMachineBehavior, "project connectivity policy name was not updated")
}
