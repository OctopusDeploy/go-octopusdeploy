package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func TestNewVariableService(t *testing.T) {
	ServiceFunction := newVariableService
	client := &sling.Sling{}
	uriTemplate := emptyString
	namesPath := emptyString
	previewPath := emptyString
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
		{"EmptyURITemplate", ServiceFunction, client, emptyString, namesPath, previewPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString, namesPath, previewPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.namesPath, tc.previewPath)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestVariableServiceGetAllWithEmptyID(t *testing.T) {
	service := newVariableService(nil, TestURIVariables, TestURIVariableNames, TestURIVariablePreview)

	variableSet, err := service.GetAll(emptyString)
	require.Error(t, err)
	require.Len(t, variableSet.Variables, 0)

	variableSet, err = service.GetAll(whitespaceString)
	require.Error(t, err)
	require.Len(t, variableSet.Variables, 0)
}

func TestVariableServiceGetByIDWithEmptyID(t *testing.T) {
	service := newVariableService(nil, TestURIVariables, TestURIVariableNames, TestURIVariablePreview)

	variable, err := service.GetByID(emptyString, emptyString)
	require.Error(t, err)
	require.Nil(t, variable)

	variable, err = service.GetByID(whitespaceString, emptyString)
	require.Error(t, err)
	require.Nil(t, variable)

	variable, err = service.GetByID(emptyString, whitespaceString)
	require.Error(t, err)
	require.Nil(t, variable)

	variable, err = service.GetByID(whitespaceString, whitespaceString)
	require.Error(t, err)
	require.Nil(t, variable)
}

func TestVariableServiceDeleteSingleWithEmptyID(t *testing.T) {
	service := newVariableService(nil, TestURIVariables, TestURIVariableNames, TestURIVariablePreview)

	variableSet, err := service.DeleteSingle(emptyString, emptyString)
	require.Error(t, err)
	require.NotNil(t, variableSet)

	variableSet, err = service.DeleteSingle(whitespaceString, emptyString)
	require.Error(t, err)
	require.NotNil(t, variableSet)

	variableSet, err = service.DeleteSingle(emptyString, whitespaceString)
	require.Error(t, err)
	require.NotNil(t, variableSet)

	variableSet, err = service.DeleteSingle(whitespaceString, whitespaceString)
	require.Error(t, err)
	require.NotNil(t, variableSet)
}
