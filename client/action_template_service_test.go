package client

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestActionTemplateService(t *testing.T) {
	t.Run("New", TestNewActionTemplateService)
}

func TestNewActionTemplateService(t *testing.T) {
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

func TestActionTemplateServiceOperationsWithStringParameter(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"EmptyParameter", emptyString},
		{"WhitespaceParameter", whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createActionTemplateService(t)

			assert.NotNil(t, service)
			if service == nil {
				return
			}

			resourceList, err := service.GetAll()

			assert.NoError(t, err)
			assert.NotNil(t, resourceList)

			for _, actionTemplate := range resourceList {
				fmt.Println(actionTemplate.GetID())

				resourceToCompare, err := service.GetByID(actionTemplate.GetID())

				assert.NoError(t, err)
				assert.NotNil(t, resourceToCompare)
			}

			resource, err := service.GetByID(tc.parameter)

			assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
			assert.Nil(t, resource)

			resourceList, err = service.GetByName(tc.parameter)

			assert.Equal(t, err, createInvalidParameterError(operationGetByName, parameterName))
			assert.NotNil(t, resourceList)

			err = service.DeleteByID(tc.parameter)

			assert.Error(t, err)
			assert.Equal(t, err, createInvalidParameterError(operationDeleteByID, parameterID))
		})
	}
}

func TestActionTemplateServiceAddWithNilActionTemplate(t *testing.T) {
	service := createActionTemplateService(t)

	resource, err := service.Add(nil)

	assert.Equal(t, err, createInvalidParameterError(operationAdd, parameterResource))
	assert.Nil(t, resource)
}

func TestActionTemplateServiceAddWithInvalidActionTemplate(t *testing.T) {
	service := createActionTemplateService(t)

	resource, err := service.Add(&model.ActionTemplate{})

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestActionTemplateServiceGetCategories(t *testing.T) {
	service := createActionTemplateService(t)

	resource, err := service.GetCategories()

	assert.NoError(t, err)
	assert.NotEmpty(t, resource)
}

func TestActionTemplateServiceSearch(t *testing.T) {
	service := createActionTemplateService(t)

	resource, err := service.Search()

	assert.NoError(t, err)
	assert.NotEmpty(t, resource)
}

func createActionTemplateService(t *testing.T) *actionTemplateService {
	categoriesURL, _ := url.Parse(TestURIActionTemplatesCategories)
	searchURL, _ := url.Parse(TestURIActionTemplatesSearch)
	versionedLogoURL, _ := url.Parse(TestURIActionTemplateVersionedLogo)

	service := newActionTemplateService(nil, TestURIActionTemplates, *categoriesURL, *searchURL, *versionedLogoURL)
	testNewService(t, service, TestURIActionTemplates, serviceActionTemplateService)
	return service
}
