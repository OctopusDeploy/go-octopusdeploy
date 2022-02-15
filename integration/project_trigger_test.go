package integration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/service"
	"testing"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualProjectTriggers(t *testing.T, expected *service.ProjectTrigger, actual *service.ProjectTrigger) {
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

	// ProjectTrigger
	assert.Equal(t, expected.Action, actual.Action)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Filter, actual.Filter)
	assert.Equal(t, expected.IsDisabled, actual.IsDisabled)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.ProjectID, actual.ProjectID)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
}

func CreateTestProjectTrigger(t *testing.T, client *octopusdeploy.client, project *service.Project, environment *service.Environment) *service.ProjectTrigger {
	require.NotNil(t, project)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.UTC)
	// days := []octopusdeploy.Weekday{octopusdeploy.Sunday, octopusdeploy.Monday}

	action := service.NewAutoDeployAction(createRandomBoolean())
	// action := octopusdeploy.NewDeployLatestReleaseAction(getRandomName(), createRandomBoolean(), []string{getRandomName()}, "")
	// action := octopusdeploy.NewDeployNewReleaseAction(environment.GetID(), "", nil)
	// filter := octopusdeploy.NewDeploymentTargetFilter([]string{}, []string{}, []string{"MachineAvailableForDeployment"}, []string{})

	// OnceDailyScheduledTriggerFilter
	// filter := octopusdeploy.NewOnceDailyScheduledTriggerFilter(days, start)

	filter := service.NewDayOfWeekScheduledTriggerFilter("1", service.Tuesday, start)

	projectTrigger := service.NewProjectTrigger(getRandomName(), getRandomName(), createRandomBoolean(), project.GetID(), action, filter)
	require.NotNil(t, projectTrigger)
	require.NoError(t, projectTrigger.Validate())

	createdProjectTrigger, err := client.ProjectTriggers.Add(projectTrigger)
	require.NoError(t, err)
	require.NotNil(t, createdProjectTrigger)
	require.NotEmpty(t, createdProjectTrigger.GetID())

	// verify the add operation was successful
	projectTriggerToCompare, err := client.ProjectTriggers.GetByID(createdProjectTrigger.GetID())
	require.NoError(t, err)
	require.NotNil(t, projectTriggerToCompare)
	AssertEqualProjectTriggers(t, createdProjectTrigger, projectTriggerToCompare)

	return createdProjectTrigger
}

func DeleteTestProjectTrigger(t *testing.T, client *octopusdeploy.client, projectTrigger *service.ProjectTrigger) {
	require.NotNil(t, projectTrigger)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.ProjectTriggers.DeleteByID(projectTrigger.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedProjectTrigger, err := client.ProjectTriggers.GetByID(projectTrigger.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedProjectTrigger)
}

func TestProjectTriggerAddGetAndDelete(t *testing.T) {
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

	environment := CreateTestEnvironment(t, client)
	require.NotNil(t, environment)
	defer DeleteTestEnvironment(t, client, environment)

	projectTrigger := CreateTestProjectTrigger(t, client, project, environment)
	require.NotNil(t, lifecycle)
	defer DeleteTestProjectTrigger(t, client, projectTrigger)
}

func TestProjectTriggerGetAll(t *testing.T) {
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

	projectTriggers, err := client.ProjectTriggers.GetAll()
	require.NoError(t, err)
	require.NotNil(t, projectTriggers)
}
