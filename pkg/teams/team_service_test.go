package teams

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createTeamService(t *testing.T) *TeamService {
	service := NewTeamService(nil, constants.TestURITeams)
	services.NewServiceTests(t, service, constants.TestURITeams, constants.ServiceTeamService)
	return service
}

func TestTeamSetAddGetDelete(t *testing.T) {
	service := createTeamService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterTeam))
	require.Nil(t, resource)

	resource, err = service.Add(&Team{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestTeamServiceAdd(t *testing.T) {
	service := createTeamService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterTeam))
	require.Nil(t, resource)

	resource, err = service.Add(&Team{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestTeamServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", ""},
		{"Whitespace", " "},
		{"InvalidID", internal.GetRandomName()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createTeamService(t)
			require.NotNil(t, service)

			if internal.IsEmpty(tc.parameter) {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
				require.Nil(t, resource)

				resourceList, err := service.GetByPartialName(tc.parameter)
				require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName), err)
				require.NotNil(t, resourceList)

				err = service.DeleteByID(tc.parameter)
				require.Error(t, err)
				require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID))
			} else {
				resource, err := service.GetByID(tc.parameter)
				require.Error(t, err)
				require.Nil(t, resource)

				err = service.DeleteByID(tc.parameter)
				require.Error(t, err)
			}
		})
	}
}

func TestTeamServiceNew(t *testing.T) {
	ServiceFunction := NewTeamService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceTeamService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *TeamService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, ""},
		{"URITemplateWithWhitespace", ServiceFunction, client, " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}
