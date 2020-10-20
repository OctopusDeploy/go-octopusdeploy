package octopusdeploy

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

func CreateTestArtifact(t *testing.T, service *artifactService) IResource {
	if service == nil {
		service = createArtifactService(t)
	}
	require.NotNil(t, service)

	filename := getRandomName()

	artifact := NewArtifact(filename)
	require.NotNil(t, artifact)

	createdArtifact, err := service.Add(artifact)
	require.NoError(t, err)
	require.NotNil(t, createdArtifact)
	require.NotEmpty(t, createdArtifact.GetID())

	return createdArtifact
}

func DeleteTestArtifact(t *testing.T, service *artifactService, artifact *Artifact) error {
	if service == nil {
		service = createArtifactService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(artifact.GetID())
}

func TestArtifactServiceGetAll(t *testing.T) {
	service := createArtifactService(t)
	require.NotNil(t, service)

	// // create 30 test artifacts (to be deleted)
	// for i := 0; i < 30; i++ {
	// 	artifact := CreateTestArtifact(t, service)
	// 	require.NotNil(t, artifact)
	// }

	artifacts, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, artifacts)

	for _, artifact := range artifacts {
		require.NotNil(t, artifact)
		assert.NotEmpty(t, artifact.GetID())
		err = DeleteTestArtifact(t, service, artifact)
		assert.NoError(t, err)
	}
}

func TestArtifactServiceGetByID(t *testing.T) {
	service := createArtifactService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	if len(resources) > 0 {
		resourceToCompare, err := service.GetByID(resources[0].GetID())
		require.NoError(t, err)
		assert.EqualValues(t, resources[0], resourceToCompare)
	}
}

func TestArtifactServiceGetByPartialName(t *testing.T) {
	service := createArtifactService(t)
	require.NotNil(t, service)

	resources, err := service.GetByPartialName(emptyString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
	assert.NotNil(t, resources)
	assert.Len(t, resources, 0)

	resources, err = service.GetByPartialName(whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
	assert.NotNil(t, resources)
	assert.Len(t, resources, 0)

	resources, err = service.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, resources)

	// TODO

	// if len(resources) > 0 {
	// 	resourcesToCompare, err := service.GetByPartialName(resources[0].Name)
	// 	assert.NoError(t, err)
	// 	assert.EqualValues(t, resourcesToCompare[0], resources[0])
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
