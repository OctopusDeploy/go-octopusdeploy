package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createArtifactService(t *testing.T) *artifactService {
	service := newArtifactService(nil, TestURIArtifacts)
	testNewService(t, service, TestURIArtifacts, serviceArtifactService)
	return service
}

func TestArtifactService(t *testing.T) {
	t.Run("Delete", TestArtifactServiceDelete)
	t.Run("GetByID", TestArtifactServiceGetByID)
	t.Run("New", TestArtifactServiceNew)
}

func TestArtifactServiceDelete(t *testing.T) {
	assert := assert.New(t)

	service := createArtifactService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(emptyString)
	assert.Equal(createInvalidParameterError(operationDeleteByID, parameterID), err)

	err = service.DeleteByID(whitespaceString)
	assert.Equal(createInvalidParameterError(operationDeleteByID, parameterID), err)

	id := getRandomName()
	err = service.DeleteByID(id)
	assert.Equal(createResourceNotFoundError("artifact", "ID", id), err)
}

func TestArtifactServiceGetByID(t *testing.T) {
	assert := assert.New(t)

	service := createArtifactService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	assert.Equal(createInvalidParameterError(operationGetByID, parameterID), err)
	assert.Nil(resource)

	resource, err = service.GetByID(whitespaceString)
	assert.Equal(createInvalidParameterError(operationGetByID, parameterID), err)
	assert.Nil(resource)

	id := getRandomName()
	resource, err = service.GetByID(id)
	assert.Equal(createResourceNotFoundError("artifact", "ID", id), err)
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

func TestArtifactServiceGetByPartialName(t *testing.T) {
	assert := assert.New(t)

	service := createArtifactService(t)
	require.NotNil(t, service)

	resources, err := service.GetByPartialName(emptyString)
	assert.Equal(err, createInvalidParameterError(operationGetByPartialName, parameterName))
	assert.NotNil(resources)
	assert.Len(resources, 0)

	resources, err = service.GetByPartialName(whitespaceString)
	assert.Equal(err, createInvalidParameterError(operationGetByPartialName, parameterName))
	assert.NotNil(resources)
	assert.Len(resources, 0)

	resources, err = service.GetAll()
	assert.NoError(err)
	assert.NotNil(resources)

	// TODO

	// if len(resources) > 0 {
	// 	resourcesToCompare, err := service.GetByPartialName(resources[0].Name)
	// 	assert.NoError(err)
	// 	assert.EqualValues(resourcesToCompare[0], resources[0])
	// }
}

func TestArtifactServiceNew(t *testing.T) {
	serviceFunction := newArtifactService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceArtifactService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *artifactService
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
