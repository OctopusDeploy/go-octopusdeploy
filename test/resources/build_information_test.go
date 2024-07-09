package resources

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/buildinformation"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestBuildInformationCommand(t *testing.T) {
	var spaceId string
	var packageId string
	var version string
	var buildInformation buildinformation.OctopusBuildInformation

	buildInfo := buildinformation.NewCreateBuildInformationCommand(spaceId, packageId, version, buildInformation)
	require.NotNil(t, buildInfo)
	require.Equal(t, packageId, buildInfo.PackageId)
	require.Equal(t, version, buildInfo.Version)
}

func TestCreateBuildInformationCommandMarshalJSON(t *testing.T) {
	spaceId := "Spaces-1"
	packageId := "ThePackage"
	version := "1.0.0"
	buildInformation := &buildinformation.OctopusBuildInformation{
		BuildEnvironment: "GitHub",
		BuildNumber:      "1",
		BuildUrl:         "https://github.com/user/repo/actions/run/1",
		Branch:           "main",
		VcsType:          "Git",
		VcsRoot:          "https://github.com/user/repo",
		VcsCommitNumber:  "abc123",
		Commits: []*buildinformation.Commit{
			{
				Id:      "abc123",
				Comment: "The comment",
			},
		},
	}

	buildInformationAsJSON, err := json.Marshal(buildInformation)
	require.NoError(t, err)
	require.NotNil(t, buildInformation)

	expectedJson := fmt.Sprintf(`{
		"SpaceId": "%s",
		"PackageId": "%s",
		"Version": "%s",
		"OctopusBuildInformation": %s
	}`, spaceId, packageId, version, buildInformationAsJSON)

	createBuildInformationCommand := buildinformation.NewCreateBuildInformationCommand(spaceId, packageId, version, *buildInformation)
	commandAsJSON, err := json.Marshal(createBuildInformationCommand)
	require.NoError(t, err)
	require.NotNil(t, commandAsJSON)
	jsonassert.New(t).Assertf(string(commandAsJSON), expectedJson)
}

func TestCreateBuildInformationCommandUnmarshalJSON(t *testing.T) {
	spaceId := "Spaces-1"
	packageId := "ThePackage"
	version := "1.0.0"
	buildInformation := &buildinformation.OctopusBuildInformation{
		BuildEnvironment: "GitHub",
		BuildNumber:      "1",
		BuildUrl:         "https://github.com/user/repo/actions/run/1",
		Branch:           "main",
		VcsType:          "Git",
		VcsRoot:          "https://github.com/user/repo",
		VcsCommitNumber:  "abc123",
		Commits: []*buildinformation.Commit{
			{
				Id:      "abc123",
				Comment: "The comment",
			},
		},
	}

	buildInformationAsJSON, err := json.Marshal(buildInformation)
	require.NoError(t, err)
	require.NotNil(t, buildInformation)

	inputJSON := fmt.Sprintf(`{
		"SpaceId": "%s",
		"PackageId": "%s",
		"Version": "%s",
		"OctopusBuildInformation": %s
	}`, spaceId, packageId, version, buildInformationAsJSON)

	var createBuildInformationCommand buildinformation.CreateBuildInformationCommand
	err = json.Unmarshal([]byte(inputJSON), &createBuildInformationCommand)
	require.NoError(t, err)
	require.NotNil(t, createBuildInformationCommand)
	require.Equal(t, spaceId, createBuildInformationCommand.SpaceId)
	require.Equal(t, packageId, createBuildInformationCommand.PackageId)
	require.Equal(t, version, createBuildInformationCommand.Version)
	require.Equal(t, buildInformation.BuildEnvironment, createBuildInformationCommand.OctopusBuildInformation.BuildEnvironment)
	require.Len(t, createBuildInformationCommand.OctopusBuildInformation.Commits, 1)
	require.Equal(t, "abc123", createBuildInformationCommand.OctopusBuildInformation.Commits[0].Id)
}
