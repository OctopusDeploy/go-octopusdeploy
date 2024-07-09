package e2e

import (
	"strings"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/buildinformation"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/spaces"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestSpaceWithCurrentUserAsSpaceManager(t *testing.T, client *client.Client) *spaces.Space {
	space := CreateTestSpace(t, client)

	me, err := client.Users.GetMe()
	require.NoError(t, err)
	require.NotNil(t, me)

	space.SpaceManagersTeamMembers = append(space.SpaceManagersTeamMembers, me.GetID())
	space, err = spaces.Update(client, space)
	require.NoError(t, err)
	require.Contains(t, space.SpaceManagersTeamMembers, me.GetID())

	return space
}

func CreateTestBuildInformation(t *testing.T, client *client.Client, space *spaces.Space, packageId string, version string) *buildinformation.BuildInformation {
	buildInfo := &buildinformation.OctopusBuildInformation{
		BuildEnvironment: "GitHub",
		BuildNumber:      "1",
		BuildUrl:         "https://github.com/user/repo/actions/run/1",
		Branch:           "main",
		VcsType:          "Git",
		VcsRoot:          "https://github.com/user/repo",
		VcsCommitNumber:  strings.ReplaceAll(uuid.New().String(), "-", ""),
		Commits: []*buildinformation.Commit{
			{
				Id:      strings.ReplaceAll(uuid.New().String(), "-", ""),
				Comment: "This is the comment",
			},
		},
	}

	createBuildInformationCommand := buildinformation.NewCreateBuildInformationCommand(space.GetID(), packageId, version, *buildInfo)

	createdBuildInformation, err := buildinformation.Add(client, createBuildInformationCommand)
	require.NoError(t, err)
	require.NotNil(t, createdBuildInformation)

	return createdBuildInformation
}

func IsEqualBuildInformation(t *testing.T, expected *buildinformation.BuildInformation, actual *buildinformation.BuildInformation) {
	// IResource
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.True(t, internal.IsLinksEqual(expected.GetLinks(), actual.GetLinks()))

	// build information
	assert.Equal(t, expected.PackageID, actual.PackageID)
	assert.Equal(t, expected.Version, actual.Version)
	assert.Equal(t, expected.BuildEnvironment, actual.BuildEnvironment)
	assert.Equal(t, expected.BuildNumber, actual.BuildNumber)
	assert.Equal(t, expected.BuildURL, actual.BuildURL)
	assert.Equal(t, expected.Branch, actual.Branch)
	assert.Equal(t, expected.VcsType, actual.VcsType)
	assert.Equal(t, expected.VcsRoot, actual.VcsRoot)
	assert.Equal(t, expected.VcsCommitNumber, actual.VcsCommitNumber)
	assert.Equal(t, expected.VcsCommitURL, actual.VcsCommitURL)
	assert.ElementsMatch(t, expected.Commits, actual.Commits)
}

func TestBuildInformationCreate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := CreateTestSpaceWithCurrentUserAsSpaceManager(t, client)
	require.NotNil(t, space)

	packageId := "TestPackage"
	version := internal.GetRandomVersion()

	createdBuildInfo := CreateTestBuildInformation(t, client, space, packageId, version)
	require.NotNil(t, createdBuildInfo)

	buildInfo, err := buildinformation.GetById(client, space.GetID(), createdBuildInfo.ID)
	require.NoError(t, err)
	require.NotNil(t, buildInfo)
	IsEqualBuildInformation(t, createdBuildInfo, buildInfo)

	t.Cleanup(func() { DeleteTestSpace(t, client, space) })
}
