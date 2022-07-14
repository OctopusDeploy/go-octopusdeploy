package spaces

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createSpaceService(t *testing.T) *SpaceService {
	service := NewSpaceService(nil, constants.TestURISpaces, constants.TestURISpaceHome)
	services.NewServiceTests(t, service, constants.TestURISpaces, constants.ServiceSpaceService)
	return service
}

func TestSpaceSetAddGetDelete(t *testing.T) {
	service := createSpaceService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, "space"))
	require.Nil(t, resource)

	resource, err = service.Add(&Space{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestSpaceServiceAdd(t *testing.T) {
	service := createSpaceService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, "space"))
	require.Nil(t, resource)

	resource, err = service.Add(&Space{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestSpaceServiceParameters(t *testing.T) {
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
			service := createSpaceService(t)
			require.NotNil(t, service)

			if internal.IsEmpty(tc.parameter) {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
				require.Nil(t, resource)

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

func TestSpaceServiceNew(t *testing.T) {
	ServiceFunction := NewSpaceService
	client := &sling.Sling{}
	uriTemplate := ""
	homePath := ""
	ServiceName := constants.ServiceSpaceService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string, string) *SpaceService
		client      *sling.Sling
		uriTemplate string
		homePath    string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, homePath},
		{"EmptyURITemplate", ServiceFunction, client, "", homePath},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", homePath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.homePath)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}
