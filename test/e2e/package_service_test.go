package e2e

import (
	"archive/zip"
	"bytes"
	cryptoRand "crypto/rand"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/packages"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// UploadRandomTestPackage uses the V1 Upload method in go-octopusdeploy. It does not support delta compression
func UploadRandomTestPackage(t *testing.T, client *client.Client) (*packages.PackageUploadResponse, bool) {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	var fileName = fmt.Sprintf("%s.1.0.0.zip", internal.GetRandomString(8))

	// octopus expects packages to be zip files, best we make a zip
	zipFileBytes := createZip(t, zippedFile{
		name:    "content.txt",
		content: []byte("This is a text file inside a test zip package"),
	})

	resource, createdNewFile, err := packages.Upload(client, client.GetSpaceID(), fileName, bytes.NewReader(zipFileBytes), packages.OverwriteModeOverwriteExisting)

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

func TestPackageServiceUploadDelta_NotEfficient(t *testing.T) {
	octopus := getOctopusClient()
	require.NotNil(t, octopus)

	pkgName := "delta_" + internal.GetRandomString(8)
	var fileName = fmt.Sprintf("%s.1.0.0.zip", pkgName)

	zipFileBytes1 := createZip(t, zippedFile{
		name:    "content.txt",
		content: []byte("This is a text file inside a test zip package"),
	})

	initialResponse, err := packages.UploadV2(
		octopus,
		octopus.GetSpaceID(),
		fileName,
		bytes.NewReader(zipFileBytes1),
		packages.OverwriteModeOverwriteExisting,
		true)

	require.NoError(t, err)
	require.NotNil(t, initialResponse)
	defer DeleteTestPackage(t, octopus, initialResponse.GetID())

	assert.Equal(t, packages.DeltaBehaviourNoPreviousFile, initialResponse.UploadInfo.DeltaBehaviour)

	var fileName2 = fmt.Sprintf("%s.2.0.0.zip", pkgName)

	// because the zip files are tiny here, the delta ends up larger than the zip so should produce the "not efficient"
	// outcome. We need bigger files for Delta to work
	zipFileBytes2 := createZip(t, zippedFile{
		name:    "content.txt",
		content: []byte("This is a text file inside a test zip package"),
	}, zippedFile{
		name:    "content2.txt",
		content: []byte("This is a second text file inside a test zip package"),
	})

	deltaResponse, err := packages.UploadV2(
		octopus,
		octopus.GetSpaceID(),
		fileName2,
		bytes.NewReader(zipFileBytes2),
		packages.OverwriteModeOverwriteExisting,
		true)

	require.NoError(t, err)
	require.NotNil(t, deltaResponse)
	defer DeleteTestPackage(t, octopus, deltaResponse.GetID())

	// it tells us if it did delta via response.UploadInfo
	assert.Equal(t, packages.DeltaBehaviourNotEfficient, deltaResponse.UploadInfo.DeltaBehaviour)

	// Orion: Note These are the values when run on my machine, but because the data is random, and
	// due to platform differences they may not always be the same; commented out and left for
	// explanatory purposes only.
	// assert.Equal(t, int64(355), deltaResponse.UploadInfo.FileSize)
	// assert.Equal(t, int64(406), deltaResponse.UploadInfo.DeltaSize)
}

func TestPackageServiceUploadDelta_UploadedDelta(t *testing.T) {
	octopus := getOctopusClient()
	require.NotNil(t, octopus)

	pkgName := "delta_" + internal.GetRandomString(8)
	var fileName = fmt.Sprintf("%s.1.0.0.zip", pkgName)

	// we need a bigger file to make the delta worthwhile. Random so it shouldn't compress too much
	randomContent1 := make([]byte, 512*1024)
	_, err := cryptoRand.Read(randomContent1)
	require.NoError(t, err)

	zipFileBytes1 := createZip(t, zippedFile{
		name:    "content.txt",
		content: randomContent1,
	})

	initialResponse, err := packages.UploadV2(
		octopus,
		octopus.GetSpaceID(),
		fileName,
		bytes.NewReader(zipFileBytes1),
		packages.OverwriteModeOverwriteExisting,
		true)

	require.NoError(t, err)
	require.NotNil(t, initialResponse)
	defer DeleteTestPackage(t, octopus, initialResponse.GetID())

	assert.Equal(t, packages.DeltaBehaviourNoPreviousFile, initialResponse.UploadInfo.DeltaBehaviour)

	var fileName2 = fmt.Sprintf("%s.2.0.0.zip", pkgName)

	randomContent2 := make([]byte, 512*1024)
	_, err = cryptoRand.Read(randomContent2)
	require.NoError(t, err)

	zipFileBytes2 := createZip(t, zippedFile{
		name:    "content.txt",
		content: randomContent1,
	}, zippedFile{
		name:    "content2.txt",
		content: randomContent2,
	})

	deltaResponse, err := packages.UploadV2(
		octopus,
		octopus.GetSpaceID(),
		fileName2,
		bytes.NewReader(zipFileBytes2),
		packages.OverwriteModeOverwriteExisting,
		true)

	require.NoError(t, err)
	require.NotNil(t, deltaResponse)
	defer DeleteTestPackage(t, octopus, deltaResponse.GetID())

	// it tells us if it did delta via response.UploadInfo
	// Orion: Note These are the values when run on my machine, but because the data is random, and
	// due to platform differences they may not always be the same; commented out and left for
	// explanatory purposes only.
	assert.Equal(t, int64(1049158), deltaResponse.UploadInfo.FileSize)
	assert.Equal(t, int64(524938), deltaResponse.UploadInfo.DeltaSize)
	assert.Equal(t, packages.DeltaBehaviourUploadedDeltaFile, deltaResponse.UploadInfo.DeltaBehaviour)
}

type zippedFile struct {
	name    string
	content []byte
}

func createZip(t *testing.T, files ...zippedFile) []byte {
	var zipBuffer bytes.Buffer
	zipWriter := zip.NewWriter(&zipBuffer)

	for _, file := range files {
		fileWriter, err := zipWriter.Create(file.name)
		require.NoError(t, err)

		_, err = fileWriter.Write(file.content)
		require.NoError(t, err)
	}

	err := zipWriter.Close()
	require.NoError(t, err)

	return zipBuffer.Bytes()
}
