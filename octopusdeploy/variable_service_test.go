package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func TestNewVariableService(t *testing.T) {
	ServiceFunction := newVariableService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	namesPath := services.emptyString
	previewPath := services.emptyString
	ServiceName := ServiceVariableService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string, string, string) *variableService
		client      *sling.Sling
		uriTemplate string
		namesPath   string
		previewPath string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, namesPath, previewPath},
		{"EmptyURITemplate", ServiceFunction, client, services.emptyString, namesPath, previewPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, services.whitespaceString, namesPath, previewPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.namesPath, tc.previewPath)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestVariableServiceGetAllWithEmptyID(t *testing.T) {
	service := newVariableService(nil, TestURIVariables, TestURIVariableNames, TestURIVariablePreview)

	variableSet, err := service.GetAll(services.emptyString)
	require.Error(t, err)
	require.Len(t, variableSet.Variables, 0)

	variableSet, err = service.GetAll(services.whitespaceString)
	require.Error(t, err)
	require.Len(t, variableSet.Variables, 0)
}

func TestVariableServiceGetByIDWithEmptyID(t *testing.T) {
	service := newVariableService(nil, TestURIVariables, TestURIVariableNames, TestURIVariablePreview)

	variable, err := service.GetByID(services.emptyString, services.emptyString)
	require.Error(t, err)
	require.Nil(t, variable)

	variable, err = service.GetByID(services.whitespaceString, services.emptyString)
	require.Error(t, err)
	require.Nil(t, variable)

	variable, err = service.GetByID(services.emptyString, services.whitespaceString)
	require.Error(t, err)
	require.Nil(t, variable)

	variable, err = service.GetByID(services.whitespaceString, services.whitespaceString)
	require.Error(t, err)
	require.Nil(t, variable)
}

func TestVariableServiceDeleteSingleWithEmptyID(t *testing.T) {
	service := newVariableService(nil, TestURIVariables, TestURIVariableNames, TestURIVariablePreview)

	variableSet, err := service.DeleteSingle(services.emptyString, services.emptyString)
	require.Error(t, err)
	require.NotNil(t, variableSet)

	variableSet, err = service.DeleteSingle(services.whitespaceString, services.emptyString)
	require.Error(t, err)
	require.NotNil(t, variableSet)

	variableSet, err = service.DeleteSingle(services.emptyString, services.whitespaceString)
	require.Error(t, err)
	require.NotNil(t, variableSet)

	variableSet, err = service.DeleteSingle(services.whitespaceString, services.whitespaceString)
	require.Error(t, err)
	require.NotNil(t, variableSet)
}
