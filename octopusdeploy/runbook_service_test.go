package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRunbookService(t *testing.T) *runbookService {
	service := newRunbookService(nil, TestURIRunbooks)
	testNewService(t, service, TestURIRunbooks, serviceRunbookService)
	return service
}

func CreateTestRunbook(t *testing.T, service *runbookService, lifecycle *Lifecycle, projectGroup *ProjectGroup, project *Project) *Runbook {
	if service == nil {
		service = createRunbookService(t)
	}
	require.NotNil(t, service)

	if lifecycle == nil {
		lifecycle = CreateTestLifecycle(t, nil)
		require.NotNil(t, lifecycle)
	}

	if projectGroup == nil {
		projectGroup = CreateTestProjectGroup(t, nil)
		require.NotNil(t, projectGroup)
	}

	if project == nil {
		project = CreateTestProject(t, nil, lifecycle, projectGroup)
		require.NotNil(t, project)
	}

	name := getRandomName()

	runbook := NewRunbook(name, project.GetID())
	require.NotNil(t, runbook)

	createdRunbook, err := service.Add(runbook)
	require.NoError(t, err)
	require.NotNil(t, createdRunbook)

	return createdRunbook
}

func DeleteTestRunbook(t *testing.T, service *runbookService, runbook *Runbook) error {
	require.NotNil(t, runbook)

	if service == nil {
		service = createRunbookService(t)
	}
	require.NotNil(t, service)

	err := service.DeleteByID(runbook.GetID())
	assert.NoError(t, err)

	return err
}

func TestRunbookServiceDeleteAll(t *testing.T) {
	runbookService := createRunbookService(t)
	require.NotNil(t, runbookService)

	lifecycle := CreateTestLifecycle(t, nil)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, nil, lifecycle)

	projectGroup := CreateTestProjectGroup(t, nil)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, nil, projectGroup)

	project := CreateTestProject(t, nil, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, nil, project)

	runbooks := []Runbook{}

	// create 30 test runbooks (to be deleted)
	for i := 0; i < 30; i++ {
		runbook := CreateTestRunbook(t, runbookService, lifecycle, projectGroup, project)
		require.NotNil(t, runbook)
		runbooks = append(runbooks, *runbook)
	}

	allRunbooks, err := runbookService.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allRunbooks)
	require.True(t, len(allRunbooks) >= 30)

	for _, runbook := range runbooks {
		require.NotNil(t, runbook)
		require.NotEmpty(t, runbook.GetID())
		err = DeleteTestRunbook(t, runbookService, &runbook)
		require.NoError(t, err)
	}
}

func TestRunbookServiceAddGetDelete(t *testing.T) {
	runbookService := createRunbookService(t)
	require.NotNil(t, runbookService)

	resource, err := runbookService.Add(nil)
	assert.Equal(t, err, createInvalidParameterError(operationAdd, parameterRunbook))
	assert.Nil(t, resource)

	invalidResource := &Runbook{}
	resource, err = runbookService.Add(invalidResource)
	assert.Equal(t, createValidationFailureError(operationAdd, invalidResource.Validate()), err)
	assert.Nil(t, resource)

	lifecycle := CreateTestLifecycle(t, nil)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, nil, lifecycle)

	projectGroup := CreateTestProjectGroup(t, nil)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, nil, projectGroup)

	project := CreateTestProject(t, nil, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, nil, project)

	runbook := CreateTestRunbook(t, runbookService, lifecycle, projectGroup, project)
	require.NotNil(t, runbook)
	defer DeleteTestRunbook(t, runbookService, runbook)

	createdRunbook, err := runbookService.Add(runbook)
	require.Error(t, err)
	require.Nil(t, createdRunbook)
}

func TestRunbookServiceGetByID(t *testing.T) {
	service := createRunbookService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		resourceToCompare, err := service.GetByID(resource.GetID())
		require.NoError(t, err)
		assert.EqualValues(t, resource, resourceToCompare)
	}
}

func TestRunbookServiceNew(t *testing.T) {
	serviceFunction := newRunbookService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceRunbookService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *runbookService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate},
		{"EmptyURITemplate", serviceFunction, client, emptyString},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}
