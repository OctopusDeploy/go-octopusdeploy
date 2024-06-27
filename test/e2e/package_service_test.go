package e2e

import (
	"archive/zip"
	"bytes"
	"fmt"
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

func UploadRandomTestPackage(t *testing.T, client *client.Client) (*packages.PackageUploadResponse, bool) {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	var fileName = fmt.Sprintf("%s.1.0.0.zip", internal.GetRandomString(8))

	// octopus expects packages to be zip files, best we make a zip
	var zipBuffer bytes.Buffer
	zipWriter := zip.NewWriter(&zipBuffer)
	fileWriter, err := zipWriter.Create("content.txt")
	require.NoError(t, err)

	_, err = fileWriter.Write([]byte("This is a text file inside a test zip package"))
	require.NoError(t, err)

	err = zipWriter.Close()
	require.NoError(t, err)

	resource, createdNewFile, err := packages.Upload(client, client.GetSpaceID(), fileName, bytes.NewReader(zipBuffer.Bytes()), packages.OverwriteModeOverwriteExisting)

	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource, createdNewFile
}

func DeleteTestPackage(t *testing.T, client *client.Client, packageID string) {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Packages.DeleteByID(packageID)
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedPackage, err := client.Packages.GetByID(packageID)
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

func TestPackageServiceAddGetAndDelete(t *testing.T) {
	octopus := getOctopusClient()
	require.NotNil(t, octopus)

	uploaded, createdNewPackage := UploadRandomTestPackage(t, octopus)
	assert.True(t, createdNewPackage)
	require.NotNil(t, uploaded)
	packageID := uploaded.GetID()

	defer DeleteTestPackage(t, octopus, packageID)

	fetched, err := octopus.Packages.GetByID(packageID)
	require.NoError(t, err)
	require.NotNil(t, fetched)

	assert.Equal(t, uploaded.NuGetFeedId, "feeds-builtin") // possible future brittle test if the space changes, this will be Feeds-NNN

	assert.Equal(t, uploaded.NuGetPackageId, fetched.NuGetPackageID)
	assert.Equal(t, uploaded.Description, fetched.Description)
}

func TestPackageServiceGetAll(t *testing.T) {
	octopus := getOctopusClient()
	require.NotNil(t, octopus)

	uploaded1, _ := UploadRandomTestPackage(t, octopus)
	defer DeleteTestPackage(t, octopus, uploaded1.GetID())
	uploaded2, _ := UploadRandomTestPackage(t, octopus)
	defer DeleteTestPackage(t, octopus, uploaded2.GetID())

	resources, err := octopus.Packages.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	allPackageIds := make([]string, 0)
	for _, resource := range resources {
		require.NotNil(t, resource)
		assert.NotEmpty(t, resource.GetID())
		allPackageIds = append(allPackageIds, resource.GetID())
	}

	assert.Containsf(t, allPackageIds, uploaded1.GetID(), "resources does not contain uploaded1")
	assert.Containsf(t, allPackageIds, uploaded2.GetID(), "resources does not contain uploaded2")
}

func TestPackageServiceUpdate(t *testing.T) {
	octopus := getOctopusClient()
	require.NotNil(t, octopus)

	//expected, _ := UploadNewTestPackage(t, octopus)
	//expected.Title = internal.GetRandomName()
	//actual := UpdatePackage(t, octopus, expected)
	//AssertEqualPackages(t, expected, actual)
	//defer DeleteTestPackage(t, octopus, expected)
}
