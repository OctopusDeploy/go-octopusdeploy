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

func TestMatchesScopeWithEmptyVariableScope(t *testing.T) {
	variableScope := VariableScope{
		Environments:  nil,
		Machines:      nil,
		Actions:       nil,
		Roles:         nil,
		Channels:      nil,
		TenantTags:    nil,
		ProcessOwners: nil,
	}
	definedScope := &VariableScope{
		Environments:  []string{"env-1"},
		Machines:      []string{"machine-1"},
		Actions:       []string{"action-1"},
		Roles:         []string{"role-1"},
		Channels:      []string{"channel-1"},
		TenantTags:    []string{"tenant-1"},
		ProcessOwners: []string{"process-1"},
	}

	match, matchedScope, err := MatchesScope(variableScope, definedScope)
	require.Nil(t, err)
	require.False(t, match)
	require.Nil(t, matchedScope.Environments)
	require.Nil(t, matchedScope.Machines)
	require.Nil(t, matchedScope.Actions)
	require.Nil(t, matchedScope.Roles)
	require.Nil(t, matchedScope.Channels)
	require.Nil(t, matchedScope.TenantTags)
	require.Nil(t, matchedScope.ProcessOwners)
}

func TestMatchesScopeWithNonMatchingVariableScope(t *testing.T) {
	variableScope := VariableScope{
		Environments:  []string{"env-1"},
		Machines:      []string{"machine-1"},
		Actions:       []string{"action-1"},
		Roles:         []string{"role-1"},
		Channels:      []string{"channel-1"},
		TenantTags:    []string{"tenant-1"},
		ProcessOwners: []string{"process-1"},
	}
	definedScope := &VariableScope{
		Environments:  []string{"env-2"},
		Machines:      []string{"machine-2"},
		Actions:       []string{"action-2"},
		Roles:         []string{"role-2"},
		Channels:      []string{"channel-2"},
		TenantTags:    []string{"tenant-2"},
		ProcessOwners: []string{"process-2"},
	}

	match, matchedScope, err := MatchesScope(variableScope, definedScope)
	require.Nil(t, err)
	require.False(t, match)
	require.Nil(t, matchedScope.Environments)
	require.Nil(t, matchedScope.Machines)
	require.Nil(t, matchedScope.Actions)
	require.Nil(t, matchedScope.Roles)
	require.Nil(t, matchedScope.Channels)
	require.Nil(t, matchedScope.TenantTags)
	require.Nil(t, matchedScope.ProcessOwners)
}

func TestMatchesScopeWithMatchingVariableScope(t *testing.T) {
	variableScope := VariableScope{
		Environments:  []string{"env-1"},
		Machines:      []string{"machine-1"},
		Actions:       []string{"action-1"},
		Roles:         []string{"role-1"},
		Channels:      []string{"channel-1"},
		TenantTags:    []string{"tenant-1"},
		ProcessOwners: []string{"process-1"},
	}
	definedScope := &VariableScope{
		Environments:  []string{"env-1"},
		Machines:      []string{"machine-1"},
		Actions:       []string{"action-1"},
		Roles:         []string{"role-1"},
		Channels:      []string{"channel-1"},
		TenantTags:    []string{"tenant-1"},
		ProcessOwners: []string{"process-1"},
	}

	match, matchedScope, err := MatchesScope(variableScope, definedScope)
	require.Nil(t, err)
	require.True(t, match)
	require.ElementsMatch(t, matchedScope.Environments, variableScope.Environments)
	require.ElementsMatch(t, matchedScope.Machines, variableScope.Machines)
	require.ElementsMatch(t, matchedScope.Actions, variableScope.Actions)
	require.ElementsMatch(t, matchedScope.Roles, variableScope.Roles)
	require.ElementsMatch(t, matchedScope.Channels, variableScope.Channels)
	require.ElementsMatch(t, matchedScope.TenantTags, variableScope.TenantTags)
	require.ElementsMatch(t, matchedScope.ProcessOwners, variableScope.ProcessOwners)
}
