package teamV1

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/service"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createTeamService(t *testing.T) *teamService {
	service := newTeamService(nil, octopusdeploy.TestURITeams)
	octopusdeploy.testNewService(t, service, octopusdeploy.TestURITeams, service.ServiceTeamService)
	return service
}

func TestTeamSetAddGetDelete(t *testing.T) {
	service := createTeamService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, octopusdeploy.createInvalidParameterError(service.OperationAdd, service.ParameterResource))
	require.Nil(t, resource)

	resource, err = service.Add(&Team{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestTeamServiceAdd(t *testing.T) {
	service := createTeamService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, octopusdeploy.createInvalidParameterError(service.OperationAdd, service.ParameterResource))
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
		{"Empty", octopusdeploy.emptyString},
		{"Whitespace", octopusdeploy.whitespaceString},
		{"InvalidID", octopusdeploy.getRandomName()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createTeamService(t)
			require.NotNil(t, service)

			if octopusdeploy.isEmpty(tc.parameter) {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, octopusdeploy.createInvalidParameterError(service.OperationGetByID, service.ParameterID))
				require.Nil(t, resource)

				resourceList, err := service.GetByPartialName(tc.parameter)
				require.Equal(t, octopusdeploy.createInvalidParameterError(service.OperationGetByPartialName, service.ParameterName), err)
				require.NotNil(t, resourceList)

				err = service.DeleteByID(tc.parameter)
				require.Error(t, err)
				require.Equal(t, err, octopusdeploy.createInvalidParameterError(service.OperationDeleteByID, service.ParameterID))
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
	ServiceFunction := newTeamService
	client := &sling.Sling{}
	uriTemplate := octopusdeploy.emptyString
	ServiceName := service.ServiceTeamService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *teamService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, octopusdeploy.emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, octopusdeploy.whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			octopusdeploy.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
