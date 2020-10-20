package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createCommunityActionTemplateService(t *testing.T) *communityActionTemplateService {
	service := newCommunityActionTemplateService(nil, TestURICommunityActionTemplates)
	testNewService(t, service, TestURICommunityActionTemplates, serviceCommunityActionTemplateService)
	return service
}

func TestCommunityActionTemplateServiceGetByID(t *testing.T) {
	service := createCommunityActionTemplateService(t)
	require.NotNil(t, service)

	resourceList, err := service.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, resourceList)

	if len(resourceList) > 0 {
		resourceToCompare, err := service.GetByID(resourceList[0].ID)
		assert.NoError(t, err)
		assert.EqualValues(t, resourceList[0], *resourceToCompare)
	}

	value := getRandomName()
	resource, err := service.GetByID(value)

	assert.Equal(t, err, createResourceNotFoundError(service.getName(), "ID", value))
	assert.Nil(t, resource)
}

func TestCommunityActionTemplateServiceGetByName(t *testing.T) {
	service := createCommunityActionTemplateService(t)
	require.NotNil(t, service)

	resourceList, err := service.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, resourceList)

	if len(resourceList) > 0 {
		resourceToCompare, err := service.GetByName(resourceList[0].Name)
		assert.NoError(t, err)
		assert.EqualValues(t, *resourceToCompare, resourceList[0])
	}
}

func TestCommunityActionTemplateServiceGetAll(t *testing.T) {
	service := createCommunityActionTemplateService(t)
	require.NotNil(t, service)

	resourceList, err := service.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, resourceList)
}

func TestCommunityActionTemplateServiceNew(t *testing.T) {
	serviceFunction := newCommunityActionTemplateService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceCommunityActionTemplateService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *communityActionTemplateService
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

func TestCommunityActionTemplateServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", emptyString},
		{"Whitespace", whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			service := createCommunityActionTemplateService(t)
			require.NotNil(t, service)

			resource, err := service.GetByID(tc.parameter)
			assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
			assert.Nil(t, resource)
		})
	}
}

func TestCommunityActionTemplateServiceGetByIDs(t *testing.T) {
	service := createCommunityActionTemplateService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	ids := []string{}
	for _, resource := range resources {
		ids = append(ids, resource.ID)
	}

	// no need to test if ID collection size is less than 2
	if len(ids) < 2 {
		return
	}

	resourceListToCompare, err := service.GetByIDs(ids[0:2])
	assert.NoError(t, err)
	assert.NotNil(t, resourceListToCompare)
	assert.Equal(t, 2, len(resourceListToCompare))
}

func TestCommunityActionTemplateServiceInstall(t *testing.T) {
	service := createCommunityActionTemplateService(t)
	require.NotNil(t, service)

	resource, err := service.Install(CommunityActionTemplate{})
	require.Error(t, err)
	require.Nil(t, resource)

	resource = NewCommunityActionTemplate(getRandomName())
	require.NotNil(t, resource)
}
