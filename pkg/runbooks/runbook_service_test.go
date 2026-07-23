package runbooks

import (
	"encoding/json"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRunbookService(t *testing.T) *RunbookService {
	service := NewRunbookService(nil, constants.TestURIRunbooks)
	services.NewServiceTests(t, service, constants.TestURIRunbooks, constants.ServiceRunbookService)
	return service
}

func TestRunbookServiceAddGetDelete(t *testing.T) {
	runbookService := createRunbookService(t)
	require.NotNil(t, runbookService)

	resource, err := runbookService.Add(nil)
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterRunbook))
	assert.Nil(t, resource)

	invalidResource := &Runbook{}
	resource, err = runbookService.Add(invalidResource)
	assert.Equal(t, internal.CreateValidationFailureError(constants.OperationAdd, invalidResource.Validate()), err)
	assert.Nil(t, resource)
}

func TestRunbookServiceNew(t *testing.T) {
	ServiceFunction := NewRunbookService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceRunbookService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *RunbookService
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

func TestRunbookRunCommandV1_MarshalJSON(t *testing.T) {
	cmd := RunbookRunCommandV1{
		RunbookName:      "My Runbook",
		EnvironmentNames: []string{"Production", "Staging"},
		Tenants:          []string{"Tenant-1"},
		TenantTags:       []string{"Region/US", "Tier/Premium"},
		Snapshot:         "Snapshot-1",
		CreateExecutionAbstractCommandV1: deployments.CreateExecutionAbstractCommandV1{
			SpaceID:                "Spaces-1",
			ProjectIDOrName:        "MyProject",
			SpecificMachineNames:   []string{"runbook-server1"},
			ExcludedMachineNames:   []string{"maintenance-server"},
			SpecificTargetTagNames: []string{"Role/RunbookServer", "Environment/Production"},
			ExcludedTargetTagNames: []string{"Role/Database", "Maintenance/True"},
			SkipStepNames:          []string{"Optional Step"},
		},
	}

	jsonBytes, err := json.Marshal(cmd)
	require.NoError(t, err)

	expectedJSON := `{
		"runbookName": "My Runbook",
		"environmentNames": ["Production", "Staging"],
		"tenants": ["Tenant-1"],
		"tenantTags": ["Region/US", "Tier/Premium"],
		"snapshot": "Snapshot-1",
		"spaceId": "Spaces-1",
		"projectName": "MyProject",
		"specificMachineNames": ["runbook-server1"],
		"excludedMachineNames": ["maintenance-server"],
		"specificTargetTagNames": ["Role/RunbookServer", "Environment/Production"],
		"excludedTargetTagNames": ["Role/Database", "Maintenance/True"],
		"skipStepNames": ["Optional Step"],
		"forcePackageDownload": false,
		"useGuidedFailure": null
	}`

	require.JSONEq(t, expectedJSON, string(jsonBytes))
}
