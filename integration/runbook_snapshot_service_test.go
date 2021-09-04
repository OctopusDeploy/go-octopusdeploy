package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualRunbookSnapshots(
	t *testing.T,
	expected *octopusdeploy.RunbookSnapshot,
	actual *octopusdeploy.RunbookSnapshot) {
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
	assert.True(t, IsEqualLinks(expected.GetLinks(), actual.GetLinks()))

	// Project
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.ProjectID, actual.ProjectID)
	assert.Equal(t, expected.RunbookID, actual.RunbookID)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
}

func CreateTestRunbookSnapshot(
	t *testing.T,
	client *octopusdeploy.Client,
	lifecycle *octopusdeploy.Lifecycle,
	projectGroup *octopusdeploy.ProjectGroup,
	project *octopusdeploy.Project,
	runbook *octopusdeploy.Runbook) *octopusdeploy.RunbookSnapshot {

	require.NotNil(t, lifecycle)
	require.NotNil(t, projectGroup)
	require.NotNil(t, project)
	require.NotNil(t, runbook)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()

	runbookSnapshot := octopusdeploy.NewRunbookSnapshot(name, project.GetID(), runbook.GetID())
	require.NotNil(t, runbookSnapshot)
	require.NoError(t, runbookSnapshot.Validate())

	createdRunbookSnapshot, err := client.RunbookSnapshots.Add(runbookSnapshot)
	require.NoError(t, err)
	require.NotNil(t, createdRunbookSnapshot)
	require.NotEmpty(t, createdRunbookSnapshot.GetID())

	// verify the add operation was successful
	runbookSnapshotToCompare, err := client.RunbookSnapshots.GetByID(createdRunbookSnapshot.GetID())
	require.NoError(t, err)
	require.NotNil(t, runbookSnapshotToCompare)
	AssertEqualRunbookSnapshots(t, createdRunbookSnapshot, runbookSnapshotToCompare)

	return createdRunbookSnapshot
}

func DeleteTestRunbookSnapshot(t *testing.T, client *octopusdeploy.Client, runbookSnapshot *octopusdeploy.RunbookSnapshot) {
	require.NotNil(t, runbookSnapshot)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.RunbookSnapshots.DeleteByID(runbookSnapshot.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedRunbookSnapshot, err := client.Projects.GetByID(runbookSnapshot.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedRunbookSnapshot)
}

func TestRunbookSnapshotServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	runbook := CreateTestRunbook(t, client, lifecycle, projectGroup, project)
	require.NotNil(t, runbook)
	defer DeleteTestRunbook(t, client, runbook)

	runbookSnapshot := CreateTestRunbookSnapshot(t, client, lifecycle, projectGroup, project, runbook)
	require.NotNil(t, runbookSnapshot)
	defer DeleteTestRunbookSnapshot(t, client, runbookSnapshot)
}
