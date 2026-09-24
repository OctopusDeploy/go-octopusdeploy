package credentials_test

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/stretchr/testify/require"
)

func createClient() newclient.Client {
	return newclient.NewClient(&newclient.HttpSession{})
}

func TestAddV2NilCredential(t *testing.T) {
	response, err := credentials.AddV2(createClient(), nil)
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterGitCredential), err)
	require.Nil(t, response)
}

func TestGetByIDV2EmptyID(t *testing.T) {
	resource, err := credentials.GetByIDV2(createClient(), "Spaces-1", "")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID), err)
	require.Nil(t, resource)

	resource, err = credentials.GetByIDV2(createClient(), "Spaces-1", " ")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID), err)
	require.Nil(t, resource)
}

func TestUpdateV2NilCredential(t *testing.T) {
	err := credentials.UpdateV2(createClient(), nil)
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterGitCredential), err)
}
