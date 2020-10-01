package client

import (
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createActionTemplate(t *testing.T) (*model.ActionTemplate, error) {
	resource, err := model.NewActionTemplate(getRandomName(), model.ActionTypeOctopusScript)
	assert.NoError(t, err)
	assert.NotNil(t, resource)

	resource.Properties = map[string]model.PropertyValue{}
	resource.Properties[model.ActionTypeOctopusActionScriptBody] = model.PropertyValue(getRandomName())

	return resource, err
}

func createActionTemplateService(t *testing.T) *actionTemplateService {
	categoriesURL, _ := url.Parse(TestURIActionTemplatesCategories)
	searchURL, _ := url.Parse(TestURIActionTemplatesSearch)
	versionedLogoURL, _ := url.Parse(TestURIActionTemplateVersionedLogo)

	service := newActionTemplateService(nil, TestURIActionTemplates, *categoriesURL, *searchURL, *versionedLogoURL)
	testNewService(t, service, TestURIActionTemplates, serviceActionTemplateService)
	return service
}

func TestActionTemplateService(t *testing.T) {
	t.Run("Add", TestActionTemplateServiceAdd)
	t.Run("Delete", TestActionTemplateServiceDelete)
	t.Run("GetByID", TestActionTemplateServiceGetByID)
	t.Run("GetCategories", TestActionTemplateServiceGetCategories)
	t.Run("New", TestActionTemplateServiceNew)
	t.Run("Search", TestActionTemplateServiceSearch)
}

func TestActionTemplateServiceAdd(t *testing.T) {
	assert := assert.New(t)

	service := createActionTemplateService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	assert.Equal(err, createInvalidParameterError(operationAdd, parameterResource))
	assert.Nil(resource)

	invalidResource := &model.ActionTemplate{}
	resource, err = service.Add(invalidResource)
	assert.Equal(createValidationFailureError("Add", invalidResource.Validate()), err)
	assert.Nil(resource)

	resource, err = createActionTemplate(t)
	assert.NoError(err)
	assert.NotNil(resource)

	if err != nil {
		return
	}

	resource, err = service.Add(resource)
	assert.NoError(err)
	assert.NotNil(resource)

	err = service.DeleteByID(resource.ID)
	assert.NoError(err)
}

func TestActionTemplateServiceDelete(t *testing.T) {
	assert := assert.New(t)

	service := createActionTemplateService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(emptyString)
	assert.Equal(createInvalidParameterError(operationDeleteByID, parameterID), err)

	err = service.DeleteByID(whitespaceString)
	assert.Equal(createInvalidParameterError(operationDeleteByID, parameterID), err)

	id := getRandomName()
	err = service.DeleteByID(id)
	assert.Equal(createResourceNotFoundError("action template", "ID", id), err)
}

func TestActionTemplateServiceGetCategories(t *testing.T) {
	assert := assert.New(t)

	service := createActionTemplateService(t)
	require.NotNil(t, service)

	resource, err := service.GetCategories()
	assert.NoError(err)
	assert.NotEmpty(resource)
}

func TestActionTemplateServiceGetByID(t *testing.T) {
	assert := assert.New(t)

	service := createActionTemplateService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	assert.Equal(createInvalidParameterError(operationGetByID, parameterID), err)
	assert.Nil(resource)

	resource, err = service.GetByID(whitespaceString)
	assert.Equal(createInvalidParameterError(operationGetByID, parameterID), err)
	assert.Nil(resource)

	id := getRandomName()
	resource, err = service.GetByID(id)
	assert.Equal(createResourceNotFoundError("action template", "ID", id), err)
	assert.Nil(resource)

	resources, err := service.GetAll()
	assert.NoError(err)
	assert.NotNil(resources)

	if len(resources) > 0 {
		resourceToCompare, err := service.GetByID(resources[0].ID)
		assert.NoError(err)
		assert.EqualValues(resources[0], *resourceToCompare)
	}
}

func TestActionTemplateServiceNew(t *testing.T) {
	serviceFunction := newActionTemplateService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceActionTemplateService
	categoriesURL := url.URL{}
	searchURL := url.URL{}
	versionedLogoURL := url.URL{}

	testCases := []struct {
		name             string
		f                func(*sling.Sling, string, url.URL, url.URL, url.URL) *actionTemplateService
		client           *sling.Sling
		uriTemplate      string
		categoriesURL    url.URL
		searchURL        url.URL
		versionedLogoURL url.URL
	}{
		{"NilClient", serviceFunction, nil, uriTemplate, categoriesURL, searchURL, versionedLogoURL},
		{"EmptyURITemplate", serviceFunction, client, emptyString, categoriesURL, searchURL, versionedLogoURL},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString, categoriesURL, searchURL, versionedLogoURL},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.categoriesURL, tc.searchURL, tc.versionedLogoURL)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestActionTemplateServiceSearch(t *testing.T) {
	assert := assert.New(t)

	service := createActionTemplateService(t)
	require.NotNil(t, service)

	resource, err := service.Search()
	assert.NoError(err)
	assert.NotEmpty(resource)
}
