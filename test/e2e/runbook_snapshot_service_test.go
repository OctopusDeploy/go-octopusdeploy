package e2e

import (
	"net/http"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/configuration"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/lifecycles"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projectgroups"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/runbooks"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualRunbookSnapshots(
	t *testing.T,
	expected *runbooks.RunbookSnapshot,
	actual *runbooks.RunbookSnapshot) {
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
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.ProjectID, actual.ProjectID)
	assert.Equal(t, expected.RunbookID, actual.RunbookID)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
}

func CreateTestRunbookSnapshot(
	t *testing.T,
	client *client.Client,
	lifecycle *lifecycles.Lifecycle,
	projectGroup *projectgroups.ProjectGroup,
	project *projects.Project,
	runbook *runbooks.Runbook) *runbooks.RunbookSnapshot {

	require.NotNil(t, lifecycle)
	require.NotNil(t, projectGroup)
	require.NotNil(t, project)
	require.NotNil(t, runbook)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	runbookSnapshot := runbooks.NewRunbookSnapshot(name, project.GetID(), runbook.GetID())
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

func DeleteTestRunbookSnapshot(t *testing.T, client *client.Client, runbookSnapshot *runbooks.RunbookSnapshot) {
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

	runbookSnapshot := CreateTestRunbookSnapshot(t, client, lifecycle, projectGroup, project, runbook)
	require.NotNil(t, runbookSnapshot)
	defer DeleteTestRunbookSnapshot(t, client, runbookSnapshot)

	runbookSnapshotTemplate, err := client.Runbooks.GetRunbookSnapshotTemplate(runbook)
	require.NoError(t, err)
	require.NotNil(t, runbookSnapshotTemplate)
}

func TestRunbookSnapshotServiceSnapshotVariablesByName(t *testing.T) {

	//Arrange
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	toggle, err := configuration.Get(octopusClient, &configuration.FeatureToggleConfigurationQuery{
		Name: "partial-updates-on-variables",
	})
	if err != nil {
		t.Skip("Could not get feature toggle configuration")
	} else if len(toggle.FeatureToggles) == 0 {
		t.Skip("PartialUpdatesOnVariables feature toggle is not present")
	} else if !toggle.FeatureToggles[0].IsEnabled {
		t.Skip("PartialUpdatesOnVariables feature toggle is not enabled")
	}

	space := GetDefaultSpace(t, octopusClient)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, octopusClient)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, octopusClient, lifecycle)

	projectGroup := CreateTestProjectGroup(t, octopusClient)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, octopusClient, projectGroup)

	project := CreateTestProject(t, octopusClient, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, octopusClient, project)

	variable := CreateTestVariable(t, project.ID, internal.GetRandomName())
	require.NotNil(t, variable)

	variable.Value = "oldValue"
	_, err = variables.UpdateSingle(octopusClient, space.ID, project.ID, variable)
	require.NoError(t, err)

	runbook := CreateTestRunbook(t, octopusClient, lifecycle, projectGroup, project)
	require.NotNil(t, runbook)
	defer DeleteTestRunbook(t, octopusClient, runbook)

	runbookSnapshot := CreateTestRunbookSnapshot(t, octopusClient, lifecycle, projectGroup, project, runbook)
	require.NotNil(t, runbookSnapshot)
	defer DeleteTestRunbookSnapshot(t, octopusClient, runbookSnapshot)

	oldProjectSnapshotId := runbookSnapshot.ProjectVariableSetSnapshotID

	// Act
	variable.Value = "newValue"
	_, err = variables.UpdateSingle(octopusClient, space.ID, project.ID, variable)
	require.NoError(t, err)

	variableIdentifier := core.VariableIdentifier{Name: variable.Name, OwnerID: project.ID}
	variableIdentifiers := []core.VariableIdentifier{variableIdentifier}

	updatedRunbookSnapshot, err := runbooks.SnapshotVariablesByName(octopusClient, runbookSnapshot, variableIdentifiers)
	require.NoError(t, err)
	require.NotNil(t, updatedRunbookSnapshot)

	// Assert
	assert.NotEqual(t, oldProjectSnapshotId, updatedRunbookSnapshot.ProjectVariableSetSnapshotID)
}

func TestRunbookSnapshotServiceSnapshotVariablesByNameConcurrency(t *testing.T) {

	//Arrange
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	toggle, err := configuration.Get(octopusClient, &configuration.FeatureToggleConfigurationQuery{
		Name: "partial-updates-on-variables",
	})
	if err != nil {
		t.Skip("Could not get feature toggle configuration")
	} else if len(toggle.FeatureToggles) == 0 {
		t.Skip("PartialUpdatesOnVariables feature toggle is not present")
	} else if !toggle.FeatureToggles[0].IsEnabled {
		t.Skip("PartialUpdatesOnVariables feature toggle is not enabled")
	}

	space := GetDefaultSpace(t, octopusClient)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, octopusClient)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, octopusClient, lifecycle)

	projectGroup := CreateTestProjectGroup(t, octopusClient)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, octopusClient, projectGroup)

	project := CreateTestProject(t, octopusClient, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, octopusClient, project)

	variableA := CreateTestVariable(t, project.ID, internal.GetRandomName())
	require.NotNil(t, variableA)
	variableB := CreateTestVariable(t, project.ID, internal.GetRandomName())
	require.NotNil(t, variableB)

	runbook := CreateTestRunbook(t, octopusClient, lifecycle, projectGroup, project)
	require.NotNil(t, runbook)
	defer DeleteTestRunbook(t, octopusClient, runbook)

	runbookSnapshot := CreateTestRunbookSnapshot(t, octopusClient, lifecycle, projectGroup, project, runbook)
	require.NotNil(t, runbookSnapshot)
	defer DeleteTestRunbookSnapshot(t, octopusClient, runbookSnapshot)

	staleToken := runbookSnapshot.VariableSnapshotConcurrencyToken

	variableA.Value = "updatedA"
	_, err = variables.UpdateSingle(octopusClient, space.ID, project.ID, variableA)
	require.NoError(t, err)

	// Act: the token captured right after creation is still valid for the first update
	firstUpdate, err := runbooks.SnapshotVariablesByName(octopusClient, runbookSnapshot, []core.VariableIdentifier{
		{Name: variableA.Name, OwnerID: project.ID},
	}, staleToken)
	require.NoError(t, err)
	require.NotNil(t, firstUpdate)

	// Assert: reusing the now-stale token is rejected with a conflict
	variableB.Value = "updatedB"
	_, err = variables.UpdateSingle(octopusClient, space.ID, project.ID, variableB)
	require.NoError(t, err)

	_, err = runbooks.SnapshotVariablesByName(octopusClient, runbookSnapshot, []core.VariableIdentifier{
		{Name: variableB.Name, OwnerID: project.ID},
	}, staleToken)
	require.Error(t, err)
	apiError, ok := err.(*core.APIError)
	require.True(t, ok)
	assert.Equal(t, http.StatusConflict, apiError.StatusCode)

	// Assert: supplying the current token succeeds
	secondUpdate, err := runbooks.SnapshotVariablesByName(octopusClient, firstUpdate, []core.VariableIdentifier{
		{Name: variableB.Name, OwnerID: project.ID},
	}, firstUpdate.VariableSnapshotConcurrencyToken)
	require.NoError(t, err)
	require.NotNil(t, secondUpdate)
}
