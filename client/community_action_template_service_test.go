package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestCommunityActionTemplateService(t *testing.T) {
	t.Run("Add", TestCommunityActionTemplateServiceAdd)
	t.Run("GetAll", TestCommunityActionTemplateServiceGetAll)
	t.Run("GetByID", TestCommunityActionTemplateServiceGetByID)
	t.Run("GetByName", TestCommunityActionTemplateServiceGetByName)
	t.Run("New", TestNewCommunityActionTemplateService)
	t.Run("Parameters", TestCommunityActionTemplateServiceParameters)
}

func TestNewCommunityActionTemplateService(t *testing.T) {
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

func TestCommunityActionTemplateServiceGetByID(t *testing.T) {
	service := createCommunityActionTemplateService(t)
	assert := assert.New(t)

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

	assert.Equal(err, createResourceNotFoundError("account", "ID", value))
	assert.Nil(resource)
}

func TestCommunityActionTemplateServiceGetByName(t *testing.T) {
	service := createCommunityActionTemplateService(t)
	assert := assert.New(t)

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
	service := createCommunityActionTemplateService(t)
	assert := assert.New(t)

	assert.NotNil(service)
	if service == nil {
		return
	}

	resourceList, err := service.GetAll()

	assert.NoError(err)
	assert.NotNil(resourceList)
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
	service := createCommunityActionTemplateService(t)
	assert := assert.New(t)

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

func TestCommunityActionTemplateServiceAdd(t *testing.T) {
	service := createCommunityActionTemplateService(t)
	assert := assert.New(t)

	resource, err := service.Install(model.CommunityActionTemplate{})

	assert.Error(err)
	assert.Nil(resource)

	resource, err = model.NewCommunityActionTemplate(getRandomName())
	resource.ID = "CommunityActionTemplates-126"

	assert.NoError(err)
	assert.NotNil(resource)

	if err != nil {
		return
	}

	resource, err = service.Install(*resource)

	assert.NoError(err)
	assert.NotNil(resource)

	err = service.DeleteByID(resource.ID)

	assert.NoError(err)
}

func TestCommunityActionTemplateServiceUpdate(t *testing.T) {
	service := createCommunityActionTemplateService(t)
	assert := assert.New(t)

	resource, err := model.NewCommunityActionTemplate(getRandomName())

	assert.NoError(err)
	assert.NotNil(resource)

	if err != nil {
		return
	}

	resourceToCompare, err := service.Install(*resource)

	assert.NoError(err)
	assert.NotNil(resourceToCompare)

	resourceToCompare.Name = getRandomName()

	updatedResource, err := service.Update(*resourceToCompare)

	assert.NoError(err)
	assert.Equal(resourceToCompare.Name, updatedResource.Name)
}

func createCommunityActionTemplateService(t *testing.T) *communityActionTemplateService {
	service := newCommunityActionTemplateService(nil, TestURICommunityActionTemplates)
	testNewService(t, service, TestURICommunityActionTemplates, serviceCommunityActionTemplateService)
	return service
}
