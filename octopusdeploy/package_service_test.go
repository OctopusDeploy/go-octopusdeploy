package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestPackage(t *testing.T, service *packageService) *Package {
	if service == nil {
		service = createPackageService(t)
	}
	require.NotNil(t, service)

	octopusPackage := NewPackage()

	require.NoError(t, octopusPackage.Validate())

	resource, err := service.Add(octopusPackage)
	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource
}

func createPackageService(t *testing.T) *packageService {
	service := newPackageService(nil, TestURIPackages, TestURIPackageDeltaSignature, TestURIPackageDeltaUpload, TestURIPackageNotesList, TestURIPackagesBulk, TestURIPackageUpload)
	testNewService(t, service, TestURIPackages, ServicePackageService)
	return service
}

func DeleteTestPackage(t *testing.T, service *packageService, octopusPackage *Package) error {
	if service == nil {
		service = createPackageService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(octopusPackage.GetID())
}

func IsEqualPackages(t *testing.T, expected *Package, actual *Package) {
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

	// TODO: add package comparisons
}

func UpdatePackage(t *testing.T, service *packageService, octopusPackage *Package) *Package {
	if service == nil {
		service = createPackageService(t)
	}
	require.NotNil(t, service)

	updatedPackage, err := service.Update(octopusPackage)
	assert.NoError(t, err)
	require.NotNil(t, updatedPackage)

	return updatedPackage
}

func TestPackageServiceAdd(t *testing.T) {
	service := createPackageService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterResource))
	require.Nil(t, resource)

	resource, err = service.Add(&Package{})
	require.Error(t, err)
	require.Nil(t, resource)

	octopusPackage := CreateTestPackage(t, service)
	err = DeleteTestPackage(t, service, octopusPackage)
	require.NoError(t, err)
}

func TestPackageServiceDeleteAll(t *testing.T) {
	service := createPackageService(t)
	require.NotNil(t, service)

	packages, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, packages)

	for _, octopusPackage := range packages {
		err = DeleteTestPackage(t, service, octopusPackage)
		require.NoError(t, err)
	}
}

func TestPackageServiceGetAll(t *testing.T) {
	service := createPackageService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		require.NotNil(t, resource)
		assert.NotEmpty(t, resource.GetID())
	}
}

func TestPackageServiceGetByID(t *testing.T) {
	service := createPackageService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	require.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	require.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	require.Nil(t, resource)

	id := getRandomName()
	resource, err = service.GetByID(id)
	require.Equal(t, createResourceNotFoundError(service.getName(), "ID", id), err)
	require.Nil(t, resource)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		resourceToCompare, err := service.GetByID(resource.GetID())
		require.NoError(t, err)
		IsEqualPackages(t, resource, resourceToCompare)
	}
}

func TestPackageServiceNew(t *testing.T) {
	ServiceFunction := newPackageService
	client := &sling.Sling{}
	uriTemplate := emptyString
	ServiceName := ServicePackageService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string, string, string, string, string, string) *packageService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, TestURIPackageDeltaSignature, TestURIPackageDeltaUpload, TestURIPackageNotesList, TestURIPackagesBulk, TestURIPackageUpload)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestPackageServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", emptyString},
		{"Whitespace", whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createPackageService(t)
			require.NotNil(t, service)

			resource, err := service.GetByID(tc.parameter)
			require.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
			require.Nil(t, resource)

			err = service.DeleteByID(tc.parameter)
			require.Error(t, err)
			require.Equal(t, err, createInvalidParameterError(OperationDeleteByID, ParameterID))
		})
	}
}

func TestPackageServiceUpdateWithEmptyPackage(t *testing.T) {
	service := createPackageService(t)
	require.NotNil(t, service)

	updatedPackage, err := service.Update(nil)
	require.Error(t, err)
	require.Nil(t, updatedPackage)

	updatedPackage, err = service.Update(&Package{})
	require.Error(t, err)
	require.Nil(t, updatedPackage)
}

func TestPackageServiceUpdate(t *testing.T) {
	service := createPackageService(t)
	require.NotNil(t, service)

	expected := CreateTestPackage(t, service)
	expected.Title = getRandomName()
	actual := UpdatePackage(t, service, expected)
	IsEqualPackages(t, expected, actual)
	err := DeleteTestPackage(t, service, expected)
	require.NoError(t, err)
}
