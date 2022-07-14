package variables

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func TestNewVariableService(t *testing.T) {
	ServiceFunction := NewVariableService
	client := &sling.Sling{}
	uriTemplate := ""
	namesPath := ""
	previewPath := ""
	ServiceName := constants.ServiceVariableService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string, string, string) *VariableService
		client      *sling.Sling
		uriTemplate string
		namesPath   string
		previewPath string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, namesPath, previewPath},
		{"EmptyURITemplate", ServiceFunction, client, "", namesPath, previewPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", namesPath, previewPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.namesPath, tc.previewPath)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestVariableServiceGetAllWithEmptyID(t *testing.T) {
	service := NewVariableService(nil, constants.TestURIVariables, constants.TestURIVariableNames, constants.TestURIVariablePreview)

	variableSet, err := service.GetAll("")
	require.Error(t, err)
	require.Len(t, variableSet.Variables, 0)

	variableSet, err = service.GetAll(" ")
	require.Error(t, err)
	require.Len(t, variableSet.Variables, 0)
}

func TestVariableServiceGetByIDWithEmptyID(t *testing.T) {
	service := NewVariableService(nil, constants.TestURIVariables, constants.TestURIVariableNames, constants.TestURIVariablePreview)

	variable, err := service.GetByID("", "")
	require.Error(t, err)
	require.Nil(t, variable)

	variable, err = service.GetByID(" ", "")
	require.Error(t, err)
	require.Nil(t, variable)

	variable, err = service.GetByID("", " ")
	require.Error(t, err)
	require.Nil(t, variable)

	variable, err = service.GetByID(" ", " ")
	require.Error(t, err)
	require.Nil(t, variable)
}

func TestVariableServiceDeleteSingleWithEmptyID(t *testing.T) {
	service := NewVariableService(nil, constants.TestURIVariables, constants.TestURIVariableNames, constants.TestURIVariablePreview)

	variableSet, err := service.DeleteSingle("", "")
	require.Error(t, err)
	require.NotNil(t, variableSet)

	variableSet, err = service.DeleteSingle(" ", "")
	require.Error(t, err)
	require.NotNil(t, variableSet)

	variableSet, err = service.DeleteSingle("", " ")
	require.Error(t, err)
	require.NotNil(t, variableSet)

	variableSet, err = service.DeleteSingle(" ", " ")
	require.Error(t, err)
	require.NotNil(t, variableSet)
}
