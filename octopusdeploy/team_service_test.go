package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createTeamService(t *testing.T) *teamService {
	service := newTeamService(nil, TestURITeams)
	testNewService(t, service, TestURITeams, serviceTeamService)
	return service
}

func TestTeamSetAddGetDelete(t *testing.T) {
	service := createTeamService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterResource))
	require.Nil(t, resource)

	resource, err = service.Add(&Team{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestTeamServiceAdd(t *testing.T) {
	service := createTeamService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterResource))
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
		{"Empty", emptyString},
		{"Whitespace", whitespaceString},
		{"InvalidID", getRandomName()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createTeamService(t)
			require.NotNil(t, service)

			if isEmpty(tc.parameter) {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
				require.Nil(t, resource)

				resourceList, err := service.GetByPartialName(tc.parameter)
				require.Equal(t, createInvalidParameterError(operationGetByPartialName, parameterName), err)
				require.NotNil(t, resourceList)

				err = service.DeleteByID(tc.parameter)
				require.Error(t, err)
				require.Equal(t, err, createInvalidParameterError(operationDeleteByID, parameterID))
			} else {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, createResourceNotFoundError(serviceTeamService, "ID", tc.parameter))
				require.Nil(t, resource)

				err = service.DeleteByID(tc.parameter)
				require.Error(t, err)
				require.Equal(t, err, createResourceNotFoundError(serviceTeamService, "ID", tc.parameter))
			}
		})
	}
}

func TestTeamServiceNew(t *testing.T) {
	serviceFunction := newTeamService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceTeamService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *teamService
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
