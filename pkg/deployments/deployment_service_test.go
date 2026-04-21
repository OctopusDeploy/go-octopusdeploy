package deployments_test

import (
	"encoding/json"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeploymentServiceGetByIDs(t *testing.T) {
	service := createDeploymentService(t)
	require.NotNil(t, service)

	ids := []string{"Accounts-285", "Accounts-286"}
	resources, err := service.GetByIDs(ids)

	assert.NoError(t, err)
	assert.NotNil(t, resources)
}

func createDeploymentService(t *testing.T) *deployments.DeploymentService {
	service := deployments.NewDeploymentService(nil, constants.TestURIDeployments)
	services.NewServiceTests(t, service, constants.TestURIDeployments, constants.ServiceDeploymentService)
	return service
}

func TestCreateExecutionAbstractCommandV1_MarshalJSON(t *testing.T) {
	cmd := deployments.CreateExecutionAbstractCommandV1{
		SpaceID:                "Spaces-1",
		ProjectIDOrName:        "MyProject",
		SpecificMachineNames:   []string{"web1", "web2"},
		ExcludedMachineNames:   []string{"web3"},
		SpecificTargetTagNames: []string{"Role/WebServer", "Environment/Production"},
		ExcludedTargetTagNames: []string{"Role/Database"},
		SkipStepNames:          []string{"Skip Step 1"},
	}

	jsonBytes, err := json.Marshal(cmd)
	require.NoError(t, err)

	expectedJSON := `{
		"spaceId": "Spaces-1",
		"projectName": "MyProject",
		"specificMachineNames": ["web1", "web2"],
		"excludedMachineNames": ["web3"],
		"specificTargetTagNames": ["Role/WebServer", "Environment/Production"],
		"excludedTargetTagNames": ["Role/Database"],
		"skipStepNames": ["Skip Step 1"],
		"forcePackageDownload": false,
		"useGuidedFailure": null
	}`

	require.JSONEq(t, expectedJSON, string(jsonBytes))
}

func TestCreateDeploymentTenantedCommandV1_MarshalJSON(t *testing.T) {
	cmd := deployments.CreateDeploymentTenantedCommandV1{
		ReleaseVersion:  "1.0.0",
		EnvironmentName: "Production",
		Tenants:         []string{"Tenant-1"},
		CreateExecutionAbstractCommandV1: deployments.CreateExecutionAbstractCommandV1{
			SpaceID:                "Spaces-1",
			ProjectIDOrName:        "MyProject",
			SpecificTargetTagNames: []string{"Role/WebServer"},
			ExcludedTargetTagNames: []string{"Role/Database"},
		},
	}

	jsonBytes, err := json.Marshal(cmd)
	require.NoError(t, err)

	expectedJSON := `{
		"releaseVersion": "1.0.0",
		"environmentName": "Production",
		"tenants": ["Tenant-1"],
		"spaceId": "Spaces-1",
		"projectName": "MyProject",
		"specificTargetTagNames": ["Role/WebServer"],
		"excludedTargetTagNames": ["Role/Database"],
		"forcePackageDownload": false,
		"forcePackageRedeployment": false,
		"updateVariableSnapshot": false,
		"useGuidedFailure": null
	}`

	require.JSONEq(t, expectedJSON, string(jsonBytes))
}

func TestCreateDeploymentUntenantedCommandV1_MarshalJSON(t *testing.T) {
	cmd := deployments.CreateDeploymentUntenantedCommandV1{
		ReleaseVersion:   "1.0.0",
		EnvironmentNames: []string{"Production", "Staging"},
		CreateExecutionAbstractCommandV1: deployments.CreateExecutionAbstractCommandV1{
			SpaceID:                "Spaces-1",
			ProjectIDOrName:        "MyProject",
			SpecificTargetTagNames: []string{"Role/WebServer", "Environment/Production"},
			ExcludedTargetTagNames: []string{"Role/Database"},
		},
	}

	jsonBytes, err := json.Marshal(cmd)
	require.NoError(t, err)

	expectedJSON := `{
		"releaseVersion": "1.0.0",
		"environmentNames": ["Production", "Staging"],
		"spaceId": "Spaces-1",
		"projectName": "MyProject",
		"specificTargetTagNames": ["Role/WebServer", "Environment/Production"],
		"excludedTargetTagNames": ["Role/Database"],
		"forcePackageDownload": false,
		"forcePackageRedeployment": false,
		"updateVariableSnapshot": false,
		"useGuidedFailure": null
	}`

	require.JSONEq(t, expectedJSON, string(jsonBytes))
}

func TestDeployment_MarshalJSON(t *testing.T) {
	deployment := deployments.Deployment{
		ReleaseID:            "Releases-1",
		EnvironmentID:        "Environments-1",
		SpecificMachineIDs:   []string{"Machines-1", "Machines-2"},
		ExcludedMachineIDs:   []string{"Machines-3"},
		SpecificTargetTagIds: []string{"TagSets-1/Tags-1", "TagSets-2/Tags-2"},
		ExcludedTargetTagIds: []string{"TagSets-1/Tags-3"},
	}

	jsonBytes, err := json.Marshal(deployment)
	require.NoError(t, err)

	expectedJSON := `{
		"Changes": null,
		"DeployedToMachineIds": null,
		"ReleaseId": "Releases-1",
		"EnvironmentId": "Environments-1",
		"SpecificMachineIds": ["Machines-1", "Machines-2"],
		"ExcludedMachineIds": ["Machines-3"],
		"SpecificTargetTagIds": ["TagSets-1/Tags-1", "TagSets-2/Tags-2"],
		"ExcludedTargetTagIds": ["TagSets-1/Tags-3"],
		"FailureEncountered": false,
		"ForcePackageDownload": false,
		"ForcePackageRedeployment": false,
		"SkipActions": null,
		"UseGuidedFailure": false
	}`

	require.JSONEq(t, expectedJSON, string(jsonBytes))
}

func TestDeploymentTemplateStep_MarshalJSON(t *testing.T) {
	step := deployments.DeploymentTemplateStep{
		ActionID:     "Actions-1",
		ActionName:   "Deploy Package",
		ActionNumber: "1",
		Roles:        []string{"web-server", "app-server"},
		AvailableTagSets: []*deployments.TagSetPreview{
			{
				TagSetID:   "TagSets-1",
				TagSetName: "Role",
				TagSetType: "MultiSelect",
				SortOrder:  1,
				AvailableTags: []*deployments.TargetTagPreview{
					{
						TagID:     "Tags-1",
						TagName:   "WebServer",
						SortOrder: 1,
					},
					{
						TagID:     "Tags-2",
						TagName:   "Database",
						SortOrder: 2,
					},
				},
			},
			{
				TagSetID:   "TagSets-2",
				TagSetName: "Environment",
				TagSetType: "SingleSelect",
				SortOrder:  2,
				AvailableTags: []*deployments.TargetTagPreview{
					{
						TagID:     "Tags-3",
						TagName:   "Production",
						SortOrder: 1,
					},
				},
			},
		},
		MachineNames:            []string{"web-01", "web-02"},
		CanBeSkipped:            true,
		IsDisabled:              false,
		HasNoApplicableMachines: false,
	}

	jsonBytes, err := json.Marshal(step)
	require.NoError(t, err)

	expectedJSON := `{
		"ActionId": "Actions-1",
		"ActionName": "Deploy Package",
		"ActionNumber": "1",
		"Roles": ["web-server", "app-server"],
		"AvailableTagSets": [
			{
				"TagSetId": "TagSets-1",
				"TagSetName": "Role",
				"TagSetType": "MultiSelect",
				"SortOrder": 1,
				"AvailableTags": [
					{
						"TagId": "Tags-1",
						"TagName": "WebServer",
						"SortOrder": 1
					},
					{
						"TagId": "Tags-2",
						"TagName": "Database",
						"SortOrder": 2
					}
				]
			},
			{
				"TagSetId": "TagSets-2",
				"TagSetName": "Environment",
				"TagSetType": "SingleSelect",
				"SortOrder": 2,
				"AvailableTags": [
					{
						"TagId": "Tags-3",
						"TagName": "Production",
						"SortOrder": 1
					}
				]
			}
		],
		"MachineNames": ["web-01", "web-02"],
		"CanBeSkipped": true,
		"IsDisabled": false,
		"HasNoApplicableMachines": false
	}`

	require.JSONEq(t, expectedJSON, string(jsonBytes))
}

func TestDeploymentPreview_MarshalJSON(t *testing.T) {
	preview := deployments.DeploymentPreview{
		Form: &deployments.Form{
			Values:   map[string]string{"Variable1": "Value1"},
			Elements: []*deployments.Element{},
		},
		StepsToExecute: []*deployments.DeploymentTemplateStep{
			{
				ActionID:     "Actions-1",
				ActionName:   "Deploy Package",
				ActionNumber: "1",
				AvailableTagSets: []*deployments.TagSetPreview{
					{
						TagSetID:   "TagSets-1",
						TagSetName: "Role",
						AvailableTags: []*deployments.TargetTagPreview{
							{
								TagID:   "Tags-1",
								TagName: "WebServer",
							},
						},
					},
				},
				MachineNames: []string{"web-01"},
			},
		},
		UseGuidedFailureModeByDefault: true,
	}

	jsonBytes, err := json.Marshal(preview)
	require.NoError(t, err)

	expectedJSON := `{
		"Form": {
			"Values": {"Variable1": "Value1"},
			"Elements": []
		},
		"StepsToExecute": [
			{
				"ActionId": "Actions-1",
				"ActionName": "Deploy Package",
				"ActionNumber": "1",
				"AvailableTagSets": [
					{
						"TagSetId": "TagSets-1",
						"TagSetName": "Role",
						"AvailableTags": [
							{
								"TagId": "Tags-1",
								"TagName": "WebServer"
							}
						]
					}
				],
				"MachineNames": ["web-01"],
				"CanBeSkipped": false,
				"IsDisabled": false,
				"HasNoApplicableMachines": false
			}
		],
		"UseGuidedFailureModeByDefault": true
	}`

	require.JSONEq(t, expectedJSON, string(jsonBytes))
}
