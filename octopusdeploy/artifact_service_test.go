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

func CreateTestArtifact(t *testing.T, service *artifactService) *Artifact {
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

	artifacts := []Artifact{}

	// create 30 test artifacts (to be deleted)
	for i := 0; i < 30; i++ {
		artifact := CreateTestArtifact(t, service)
		require.NotNil(t, artifact)
		artifacts = append(artifacts, *artifact)
	}

	allArtifacts, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allArtifacts)
	require.True(t, len(allArtifacts) >= 30)

	for _, artifact := range artifacts {
		require.NotNil(t, artifact)
		require.NotEmpty(t, artifact.GetID())
		err = DeleteTestArtifact(t, service, &artifact)
		require.NoError(t, err)
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
