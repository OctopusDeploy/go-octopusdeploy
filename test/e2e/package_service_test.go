package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/packages"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualPackages(t *testing.T, expected *packages.Package, actual *packages.Package) {
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

	// TODO: add package comparisons
}

func CreateTestPackage(t *testing.T, client *client.Client) *packages.Package {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	octopusPackage := packages.NewPackage()
	require.NotNil(t, octopusPackage)
	require.NoError(t, octopusPackage.Validate())

	// resource, err := client.Packages.Upload(packageBytes)
	// require.NoError(t, err)
	// require.NotNil(t, resource)

	return octopusPackage
}

func DeleteTestPackage(t *testing.T, client *client.Client, octopusPackage *packages.Package) {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Packages.DeleteByID(octopusPackage.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedPackage, err := client.Packages.GetByID(octopusPackage.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedPackage)
}

func UpdatePackage(t *testing.T, client *client.Client, octopusPackage *packages.Package) *packages.Package {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	updatedPackage, err := client.Packages.Update(octopusPackage)
	assert.NoError(t, err)
	require.NotNil(t, updatedPackage)

	return updatedPackage
}

// TODO: fix test
// func TestPackageServiceAdd(t *testing.T) {
// 	client := getOctopusClient()
// 	require.NotNil(t, client)

// 	octopusPackage := CreateTestPackage(t, client)
// 	require.NotNil(t, octopusPackage)
// 	defer DeleteTestPackage(t, client, octopusPackage)
// }

// TODO: fix test
// func TestPackageServiceDeleteAll(t *testing.T) {
// 	client := getOctopusClient()
// 	require.NotNil(t, client)

// 	packages, err := client.Packages.GetAll()
// 	require.NoError(t, err)
// 	require.NotNil(t, packages)

// 	for _, octopusPackage := range packages {
// 		defer DeleteTestPackage(t, client, octopusPackage)
// 	}
// }

// TODO: fix test
// func TestPackageServiceGetAll(t *testing.T) {
// 	client := getOctopusClient()
// 	require.NotNil(t, client)

// 	resources, err := client.Packages.GetAll()
// 	require.NoError(t, err)
// 	require.NotNil(t, resources)

// 	for _, resource := range resources {
// 		require.NotNil(t, resource)
// 		assert.NotEmpty(t, resource.GetID())
// 	}
// }

// TODO: fix test
// func TestPackageServiceGetByID(t *testing.T) {
// 	client := getOctopusClient()
// 	require.NotNil(t, client)

// 	id := getRandomName()
// 	resource, err := client.Packages.GetByID(id)
// 	require.Error(t, err)
// 	require.Nil(t, resource)

// 	resources, err := client.Packages.GetAll()
// 	require.NoError(t, err)
// 	require.NotNil(t, resources)

// 	for _, resource := range resources {
// 		resourceToCompare, err := client.Packages.GetByID(resource.GetID())
// 		require.NoError(t, err)
// 		AssertEqualPackages(t, resource, resourceToCompare)
// 	}
// }

// TODO: fix test
// func TestPackageServiceUpdate(t *testing.T) {
// 	client := getOctopusClient()
// 	require.NotNil(t, client)

// 	expected := CreateTestPackage(t, client)
// 	expected.Title = getRandomName()
// 	actual := UpdatePackage(t, client, expected)
// 	AssertEqualPackages(t, expected, actual)
// 	defer DeleteTestPackage(t, client, expected)
// }
