package e2e

import (
	"testing"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/filters"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/triggers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualProjectTriggers(t *testing.T, expected *triggers.ProjectTrigger, actual *triggers.ProjectTrigger) {
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

	// ProjectTrigger
	assert.Equal(t, expected.Action, actual.Action)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Filter, actual.Filter)
	assert.Equal(t, expected.IsDisabled, actual.IsDisabled)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.ProjectID, actual.ProjectID)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
}

func CreateTestProjectTrigger(t *testing.T, client *client.Client, project *projects.Project) *triggers.ProjectTrigger {
	require.NotNil(t, project)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.UTC)
	// days := []filters.Weekday{filters.Sunday, filters.Monday}

	action := actions.NewAutoDeployAction(createRandomBoolean())
	// action := actions.NewDeployLatestReleaseAction(internal.GetRandomName(), createRandomBoolean(), []string{internal.GetRandomName()}, "")
	// action := actions.NewDeployNewReleaseAction(environment.GetID(), "", nil)
	// filter := filters.NewDeploymentTargetFilter([]string{}, []string{}, []string{"MachineAvailableForDeployment"}, []string{})

	// OnceDailyScheduledTriggerFilter
	// filter := filters.NewOnceDailyScheduledTriggerFilter(days, start)

	filter := filters.NewOnceDailyScheduledTriggerFilter([]filters.Weekday{filters.Tuesday}, start)

	projectTrigger := triggers.NewProjectTrigger(internal.GetRandomName(), internal.GetRandomName(), createRandomBoolean(), project, action, filter)
	require.NotNil(t, projectTrigger)
	require.NoError(t, projectTrigger.Validate())

	createdProjectTrigger, err := triggers.Add(client, projectTrigger)
	require.NoError(t, err)
	require.NotNil(t, createdProjectTrigger)
	require.NotEmpty(t, createdProjectTrigger.GetID())

	// verify the add operation was successful
	projectTriggerToCompare, err := triggers.GetById(client, project.SpaceID, createdProjectTrigger.GetID())
	require.NoError(t, err)
	require.NotNil(t, projectTriggerToCompare)
	AssertEqualProjectTriggers(t, createdProjectTrigger, projectTriggerToCompare)

	return createdProjectTrigger
}

func TestProjectScheduledRunbookTrigger(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	space := GetDefaultSpace(t, octopusClient)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, octopusClient)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, octopusClient, lifecycle)

	environment := CreateTestEnvironment(t, octopusClient)
	require.NotNil(t, environment)
	defer DeleteTestEnvironment(t, octopusClient, environment)

	projectGroup := CreateTestProjectGroup(t, octopusClient)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, octopusClient, projectGroup)

	project := CreateTestProject(t, octopusClient, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, octopusClient, project)

	runbook := CreateTestRunbook(t, octopusClient, lifecycle, projectGroup, project)
	require.NotNil(t, runbook)
	defer DeleteTestRunbook(t, octopusClient, runbook)

	action := actions.NewRunRunbookAction()
	action.Runbook = runbook.GetID()
	action.Environments = []string{environment.GetID()}

	password := getShortRandomName()
	secret := core.SensitiveValue{HasValue: true, NewValue: &password}

	filter := filters.NewWebhookTriggerFilter(secret)

	projectTrigger := triggers.NewProjectTrigger(internal.GetRandomName(), internal.GetRandomName(), createRandomBoolean(), project, action, filter)
	require.NotNil(t, projectTrigger)
	require.NoError(t, projectTrigger.Validate())

	createdProjectTrigger, err := triggers.Add(octopusClient, projectTrigger)
	require.NoError(t, err)
	require.NotNil(t, createdProjectTrigger)
	require.NotEmpty(t, createdProjectTrigger.GetID())

	// verify the add operation was successful
	projectTriggerToCompare, err := triggers.GetById(octopusClient, project.SpaceID, createdProjectTrigger.GetID())
	require.NoError(t, err)
	require.NotNil(t, projectTriggerToCompare)
	AssertEqualProjectTriggers(t, createdProjectTrigger, projectTriggerToCompare)
}

func DeleteTestProjectTrigger(t *testing.T, client *client.Client, projectTrigger *triggers.ProjectTrigger) {
	require.NotNil(t, projectTrigger)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := triggers.DeleteById(client, projectTrigger.SpaceID, projectTrigger.ID)
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedProjectTrigger, err := triggers.GetById(client, projectTrigger.SpaceID, projectTrigger.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedProjectTrigger)
}

func UpdateTestProjectTrigger(t *testing.T, client *client.Client, projectTrigger *triggers.ProjectTrigger) *triggers.ProjectTrigger {
	require.NotNil(t, projectTrigger)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	projectTrigger, err := client.ProjectTriggers.Update(projectTrigger)
	assert.NoError(t, err)
	assert.NotNil(t, projectTrigger)

	return projectTrigger
}

func TestProjectTriggerAddGetUpdateDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

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

	environment := CreateTestEnvironment(t, octopusClient)
	require.NotNil(t, environment)
	defer DeleteTestEnvironment(t, octopusClient, environment)

	projectTrigger := CreateTestProjectTrigger(t, octopusClient, project)
	require.NotNil(t, lifecycle)
	defer DeleteTestProjectTrigger(t, octopusClient, projectTrigger)

	projectTrigger.Name = GetRandomName()
	updatedProjectTrigger := UpdateTestProjectTrigger(t, octopusClient, projectTrigger)
	require.NotNil(t, updatedProjectTrigger)
}

func TestProjectTriggerGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

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

	projectTriggers, err := octopusClient.ProjectTriggers.GetAll()
	require.NoError(t, err)
	require.NotNil(t, projectTriggers)
}
