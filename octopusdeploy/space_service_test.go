package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createSpaceService(t *testing.T) *spaceService {
	service := newSpaceService(nil, TestURISpaces, TestURISpaceHome)
	testNewService(t, service, TestURISpaces, ServiceSpaceService)
	return service
}

func TestSpaceSetAddGetDelete(t *testing.T) {
	service := createSpaceService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(OperationAdd, "space"))
	require.Nil(t, resource)

	resource, err = service.Add(&Space{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestSpaceServiceAdd(t *testing.T) {
	service := createSpaceService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(OperationAdd, "space"))
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
		{"Empty", emptyString},
		{"Whitespace", whitespaceString},
		{"InvalidID", getRandomName()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createSpaceService(t)
			require.NotNil(t, service)

			if isEmpty(tc.parameter) {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
				require.Nil(t, resource)

				resourceList, err := service.GetByPartialName(tc.parameter)
				require.Equal(t, createInvalidParameterError(OperationGetByPartialName, ParameterName), err)
				require.NotNil(t, resourceList)

				err = service.DeleteByID(tc.parameter)
				require.Error(t, err)
				require.Equal(t, err, createInvalidParameterError(OperationDeleteByID, ParameterID))
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
	ServiceFunction := newSpaceService
	client := &sling.Sling{}
	uriTemplate := emptyString
	homePath := emptyString
	ServiceName := ServiceSpaceService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string, string) *spaceService
		client      *sling.Sling
		uriTemplate string
		homePath    string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, homePath},
		{"EmptyURITemplate", ServiceFunction, client, emptyString, homePath},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString, homePath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.homePath)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
