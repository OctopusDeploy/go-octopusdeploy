package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createCommunityActionTemplateService(t *testing.T) *communityActionTemplateService {
	service := newCommunityActionTemplateService(nil, TestURICommunityActionTemplates)
	testNewService(t, service, TestURICommunityActionTemplates, serviceCommunityActionTemplateService)
	return service
}

func TestCommunityActionTemplateService(t *testing.T) {
	t.Run("GetAll", TestCommunityActionTemplateServiceGetAll)
	t.Run("GetByID", TestCommunityActionTemplateServiceGetByID)
	t.Run("GetByName", TestCommunityActionTemplateServiceGetByName)
	t.Run("Install", TestCommunityActionTemplateServiceInstall)
	t.Run("New", TestCommunityActionTemplateServiceNew)
	t.Run("Parameters", TestCommunityActionTemplateServiceParameters)
}

func TestCommunityActionTemplateServiceGetByID(t *testing.T) {
	assert := assert.New(t)

	service := createCommunityActionTemplateService(t)
	assert.NotNil(service)
	if service == nil {
		return
	}

	resourceList, err := service.GetAll()
	assert.NoError(err)
	assert.NotNil(resourceList)

	if len(resourceList) > 0 {
		resourceToCompare, err := service.GetByID(resourceList[0].ID)
		assert.NoError(err)
		assert.EqualValues(resourceList[0], *resourceToCompare)
	}

	value := getRandomName()
	resource, err := service.GetByID(value)

	assert.Equal(err, createResourceNotFoundError("community action template", "ID", value))
	assert.Nil(resource)
}

func TestCommunityActionTemplateServiceGetByName(t *testing.T) {
	assert := assert.New(t)

	service := createCommunityActionTemplateService(t)
	assert.NotNil(service)
	if service == nil {
		return
	}

	resourceList, err := service.GetAll()
	assert.NoError(err)
	assert.NotNil(resourceList)

	if len(resourceList) > 0 {
		resourceToCompare, err := service.GetByName(resourceList[0].Name)
		assert.NoError(err)
		assert.EqualValues(*resourceToCompare, resourceList[0])
	}
}

func TestCommunityActionTemplateServiceGetAll(t *testing.T) {
	assert := assert.New(t)

	service := createCommunityActionTemplateService(t)
	assert.NotNil(service)
	if service == nil {
		return
	}

	resourceList, err := service.GetAll()
	assert.NoError(err)
	assert.NotNil(resourceList)
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
			assert := assert.New(t)

			assert.NotNil(service)
			if service == nil {
				return
			}

			resource, err := service.GetByID(tc.parameter)

			assert.Equal(err, createInvalidParameterError(operationGetByID, parameterID))
			assert.Nil(resource)

			err = service.DeleteByID(tc.parameter)

			assert.Error(err)
			assert.Equal(err, createInvalidParameterError(operationDeleteByID, parameterID))
		})
	}
}

func TestCommunityActionTemplateServiceGetByIDs(t *testing.T) {
	assert := assert.New(t)

	service := createCommunityActionTemplateService(t)
	assert.NotNil(service)
	if service == nil {
		return
	}

	resourceList, err := service.GetAll()
	assert.NoError(err)
	assert.NotNil(resourceList)

	idList := []string{}
	for i := 0; i < 3; i++ {
		idList = append(idList, resourceList[i].ID)
	}

	resourceListToCompare, err := service.GetByIDs(idList)

	assert.NoError(err)
	assert.NotNil(resourceListToCompare)
}

func TestCommunityActionTemplateServiceInstall(t *testing.T) {
	assert := assert.New(t)

	service := createCommunityActionTemplateService(t)
	assert.NotNil(service)
	if service == nil {
		return
	}

	resource, err := service.Install(model.CommunityActionTemplate{})
	assert.Error(err)
	assert.Nil(resource)

	resource, err = model.NewCommunityActionTemplate(getRandomName())
	require.NoError(t, err)
	require.NotNil(t, resource)
}
