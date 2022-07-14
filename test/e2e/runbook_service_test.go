package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/lifecycles"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projectgroups"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/runbooks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualRunbooks(t *testing.T, expected *runbooks.Runbook, actual *runbooks.Runbook) {
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
	assert.Equal(t, expected.ConnectivityPolicy, actual.ConnectivityPolicy)
	assert.Equal(t, expected.DefaultGuidedFailureMode, actual.DefaultGuidedFailureMode)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.EnvironmentScope, actual.EnvironmentScope)
	assert.Equal(t, expected.Environments, actual.Environments)
	assert.Equal(t, expected.MultiTenancyMode, actual.MultiTenancyMode)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.ProjectID, actual.ProjectID)
	assert.Equal(t, expected.PublishedRunbookSnapshotID, actual.PublishedRunbookSnapshotID)
	assert.Equal(t, expected.RunRetentionPolicy, actual.RunRetentionPolicy)
	assert.Equal(t, expected.RunbookProcessID, actual.RunbookProcessID)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
}

func CreateTestRunbook(t *testing.T, client *client.Client, lifecycle *lifecycles.Lifecycle, projectGroup *projectgroups.ProjectGroup, project *projects.Project) *runbooks.Runbook {
	require.NotNil(t, lifecycle)
	require.NotNil(t, projectGroup)
	require.NotNil(t, project)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	runbook := runbooks.NewRunbook(name, project.GetID())
	require.NotNil(t, runbook)
	require.NoError(t, runbook.Validate())

	createdRunbook, err := client.Runbooks.Add(runbook)
	require.NoError(t, err)
	require.NotNil(t, createdRunbook)
	require.NotEmpty(t, createdRunbook.GetID())

	// verify the add operation was successful
	runbookToCompare, err := client.Runbooks.GetByID(createdRunbook.GetID())
	require.NoError(t, err)
	require.NotNil(t, runbookToCompare)
	AssertEqualRunbooks(t, createdRunbook, runbookToCompare)

	return createdRunbook
}

func DeleteTestRunbook(t *testing.T, client *client.Client, runbook *runbooks.Runbook) {
	require.NotNil(t, runbook)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Runbooks.DeleteByID(runbook.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedRunbook, err := client.Projects.GetByID(runbook.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedRunbook)
}

func TestRunbookServiceDeleteAll(t *testing.T) {
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

	// create 30 test runbooks (to be deleted)
	for i := 0; i < 30; i++ {
		runbook := CreateTestRunbook(t, client, lifecycle, projectGroup, project)
		require.NotNil(t, runbook)
		defer DeleteTestRunbook(t, client, runbook)
	}

	allRunbooks, err := client.Runbooks.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allRunbooks)
	require.True(t, len(allRunbooks) >= 30)
}

func TestRunbookServiceAddGetDelete(t *testing.T) {
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

	runbook := CreateTestRunbook(t, client, lifecycle, projectGroup, project)
	require.NotNil(t, runbook)
	defer DeleteTestRunbook(t, client, runbook)
}

func TestRunbookServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	runbooks, err := client.Runbooks.GetAll()
	require.NoError(t, err)
	require.NotNil(t, runbooks)

	for _, runbook := range runbooks {
		runbookToCompare, err := client.Runbooks.GetByID(runbook.GetID())
		require.NoError(t, err)
		require.NotNil(t, runbookToCompare)
		AssertEqualRunbooks(t, runbook, runbookToCompare)
	}
}
