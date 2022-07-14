package actiontemplates

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createActionTemplate(t *testing.T) *ActionTemplate {
	resource := NewActionTemplate(internal.GetRandomName(), constants.ActionTypeOctopusScript)
	require.NotNil(t, resource)

	resource.Properties = map[string]core.PropertyValue{}
	resource.Properties[constants.ActionTypeOctopusActionScriptBody] = core.NewPropertyValue(internal.GetRandomName(), false)

	return resource
}

func createActionTemplateService(t *testing.T) *ActionTemplateService {
	serviceName := constants.ServiceActionTemplateService
	uriTemplate := constants.TestURIActionTemplates
	categoriesPath := constants.TestURIActionTemplatesCategories
	logoPath := constants.TestURIActionTemplatesLogo
	searchPath := constants.TestURIActionTemplatesSearch
	versionedLogoPath := constants.TestURIActionTemplateVersionedLogo

	service := NewActionTemplateService(nil, uriTemplate, categoriesPath, logoPath, searchPath, versionedLogoPath)

	require.NotNil(t, service)
	require.NotNil(t, service.GetClient())

	template, err := uritemplates.Parse(uriTemplate)
	require.NoError(t, err)
	require.Equal(t, service.GetURITemplate(), template)
	require.Equal(t, service.GetName(), serviceName)

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
	assert.True(t, internal.IsLinksEqual(expected.GetLinks(), actual.GetLinks()))

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
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterActionTemplate))
	require.Nil(t, resource)

	invalidResource := &ActionTemplate{}
	resource, err = service.Add(invalidResource)
	require.Equal(t, internal.CreateValidationFailureError(constants.OperationAdd, invalidResource.Validate()), err)
	require.Nil(t, resource)

	resource = createActionTemplate(t)
	require.NotNil(t, resource)
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

	resource, err := service.GetByID("")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(" ")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID), err)
	require.Nil(t, resource)

	id := internal.GetRandomName()
	resource, err = service.GetByID(id)
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestActionTemplateServiceNew(t *testing.T) {
	ServiceFunction := NewActionTemplateService
	client := &sling.Sling{}
	uriTemplate := ""
	serviceName := constants.ServiceActionTemplateService
	categoriesPath := ""
	logoPath := ""
	searchPath := ""
	versionedLogoPath := ""

	testCases := []struct {
		name              string
		f                 func(*sling.Sling, string, string, string, string, string) *ActionTemplateService
		client            *sling.Sling
		uriTemplate       string
		categoriesPath    string
		logoPath          string
		searchPath        string
		versionedLogoPath string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, categoriesPath, logoPath, searchPath, versionedLogoPath},
		{"EmptyURITemplate", ServiceFunction, client, "", categoriesPath, logoPath, searchPath, versionedLogoPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", categoriesPath, logoPath, searchPath, versionedLogoPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.categoriesPath, tc.logoPath, tc.searchPath, tc.versionedLogoPath)
			services.NewServiceTests(t, service, uriTemplate, serviceName)
		})
	}
}
