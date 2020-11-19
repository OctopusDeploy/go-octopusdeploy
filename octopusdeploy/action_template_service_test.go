package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createActionTemplate(t *testing.T) *ActionTemplate {
	resource := NewActionTemplate(getRandomName(), ActionTypeOctopusScript)
	require.NotNil(t, resource)

	resource.Properties = map[string]PropertyValue{}
	resource.Properties[ActionTypeOctopusActionScriptBody] = PropertyValue(getRandomName())

	return resource
}

func createActionTemplateService(t *testing.T) *actionTemplateService {
	categoriesPath := TestURIActionTemplatesCategories
	logoPath := TestURIActionTemplatesLogo
	searchPath := TestURIActionTemplatesSearch
	versionedLogoPath := TestURIActionTemplateVersionedLogo

	service := newActionTemplateService(nil, TestURIActionTemplates, categoriesPath, logoPath, searchPath, versionedLogoPath)
	testNewService(t, service, TestURIActionTemplates, ServiceActionTemplateService)
	return service
}

func IsEqualActionTemplates(t *testing.T, expected *ActionTemplate, actual *ActionTemplate) {
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

	// ActionTemplate
	assert.Equal(t, expected.ActionType, actual.ActionType)
	assert.Equal(t, expected.CommunityActionTemplateID, actual.CommunityActionTemplateID)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Packages, actual.Packages)
	assert.Equal(t, expected.Parameters, actual.Parameters)
	assert.Equal(t, expected.Properties, actual.Properties)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
	assert.Equal(t, expected.Version, actual.Version)
}

func TestActionTemplateServiceAdd(t *testing.T) {
	service := createActionTemplateService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	assert.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterResource))
	assert.Nil(t, resource)

	invalidResource := &ActionTemplate{}
	resource, err = service.Add(invalidResource)
	assert.Equal(t, createValidationFailureError("Add", invalidResource.Validate()), err)
	assert.Nil(t, resource)

	resource = createActionTemplate(t)
	require.NotNil(t, resource)

	resource, err = service.Add(resource)
	require.NoError(t, err)
	require.NotNil(t, resource)

	err = service.DeleteByID(resource.GetID())
	assert.NoError(t, err)
}

func TestActionTemplateServiceGetCategories(t *testing.T) {
	service := createActionTemplateService(t)
	require.NotNil(t, service)

	resource, err := service.GetCategories()
	assert.NoError(t, err)
	assert.NotEmpty(t, resource)
}

func TestActionTemplateServiceGetByID(t *testing.T) {
	service := createActionTemplateService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	require.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	require.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	require.Nil(t, resource)

	id := getRandomName()
	resource, err = service.GetByID(id)
	require.Error(t, err)
	require.Nil(t, resource)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		resourceToCompare, err := service.GetByID(resource.GetID())
		require.NoError(t, err)
		IsEqualActionTemplates(t, resource, resourceToCompare)
	}
}

func TestActionTemplateServiceNew(t *testing.T) {
	ServiceFunction := newActionTemplateService
	client := &sling.Sling{}
	uriTemplate := emptyString
	ServiceName := ServiceActionTemplateService
	categoriesPath := emptyString
	logoPath := emptyString
	searchPath := emptyString
	versionedLogoPath := emptyString

	testCases := []struct {
		name              string
		f                 func(*sling.Sling, string, string, string, string, string) *actionTemplateService
		client            *sling.Sling
		uriTemplate       string
		categoriesPath    string
		logoPath          string
		searchPath        string
		versionedLogoPath string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, categoriesPath, logoPath, searchPath, versionedLogoPath},
		{"EmptyURITemplate", ServiceFunction, client, emptyString, categoriesPath, logoPath, searchPath, versionedLogoPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString, categoriesPath, logoPath, searchPath, versionedLogoPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.categoriesPath, tc.logoPath, tc.searchPath, tc.versionedLogoPath)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestActionTemplateServiceSearch(t *testing.T) {
	service := createActionTemplateService(t)
	require.NotNil(t, service)

	resource, err := service.Search()
	assert.NoError(t, err)
	assert.NotEmpty(t, resource)
}
