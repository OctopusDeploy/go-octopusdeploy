package octopusdeploy

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeploymentTargetUnmarshalJSON(t *testing.T) {
	var azureCloudService DeploymentTarget
	err := json.Unmarshal([]byte(azureCloudServiceAsJSON), &azureCloudService)
	require.NoError(t, err)
	require.NotNil(t, azureCloudService)

	var offlinePackageDrop DeploymentTarget
	err = json.Unmarshal([]byte(offlinePackageDropAsJSON), &offlinePackageDrop)
	require.NoError(t, err)
	require.NotNil(t, offlinePackageDrop)

	var kubernetes DeploymentTarget
	err = json.Unmarshal([]byte(kubernetesAsJSON), &kubernetes)
	require.NoError(t, err)
	require.NotNil(t, kubernetes)

	var cloudRegionDeploymentTarget DeploymentTarget
	err = json.Unmarshal([]byte(cloudRegionDeploymentTargetAsJSON), &cloudRegionDeploymentTarget)
	require.NoError(t, err)
	require.NotNil(t, cloudRegionDeploymentTarget)
}

const azureCloudServiceAsJSON string = `{
	"Id": "Machines-2254",
	"EnvironmentIds": [
	  "Environments-183"
	],
	"Roles": [
	  "Prod"
	],
	"TenantedDeploymentParticipation": "Untenanted",
	"TenantIds": [
		"tenant-id-1",
		"tenant-id-2"
	],
	"TenantTags": [
		"tenant-tag-1",
		"tenant-tag-2"
	],
	"SpaceId": "Spaces-1",
	"Name": "Azure Cloud Service",
	"Thumbprint": "thumbprint",
	"Uri": "http://www.example.com/",
	"IsDisabled": false,
	"MachinePolicyId": "MachinePolicies-1",
	"Status": "CalamariNeedsUpgrade",
	"HealthStatus": "Unhealthy",
	"HasLatestCalamari": true,
	"StatusSummary": "There was a problem communicating with this machine (last checked: Friday, 02 October 2020 8:28:11 PM +00:00)",
	"IsInProcess": true,
	"Endpoint": {
	  "CommunicationStyle": "AzureCloudService",
	  "AccountId": "Accounts-285",
	  "CloudServiceName": "Azure Cloud Service Name",
	  "StorageAccountName": "Azure Storage Account Name",
	  "Slot": "Staging",
	  "SwapIfPossible": true,
	  "UseCurrentInstanceCount": true,
	  "DefaultWorkerPoolId": "",
	  "Id": "endpoint-id",
	  "LastModifiedOn": null,
	  "LastModifiedBy": "john.smith@example.com",
	  "Links": {}
	},
	"OperatingSystem": "foo",
	"ShellName": "bar",
	"ShellVersion": "quux",
	"Links": {
	  "Self": "/api/Spaces-1/machines/Machines-2254",
	  "Connection": "/api/Spaces-1/machines/Machines-2254/connection",
	  "TasksTemplate": "/api/Spaces-1/machines/Machines-2254/tasks{?skip,take}"
	}
  }`

const cloudRegionDeploymentTargetAsJSON string = `{
	"Id": "Machines-2646",
	"EnvironmentIds": [
	  "Environments-183"
	],
	"Roles": [
	  "Prod"
	],
	"TenantedDeploymentParticipation": "Tenanted",
	"TenantIds": [
	  "Tenants-1922"
	],
	"TenantTags": [],
	"SpaceId": "Spaces-1",
	"Name": "Cloud Region Deployment Target",
	"Thumbprint": null,
	"Uri": null,
	"IsDisabled": false,
	"MachinePolicyId": "MachinePolicies-1",
	"Status": "Online",
	"HealthStatus": "Healthy",
	"HasLatestCalamari": true,
	"StatusSummary": "This target is enabled.",
	"IsInProcess": true,
	"Endpoint": {
	  "CommunicationStyle": "None",
	  "DefaultWorkerPoolId": "WorkerPools-2",
	  "Id": null,
	  "LastModifiedOn": null,
	  "LastModifiedBy": null,
	  "Links": {}
	},
	"OperatingSystem": "Unknown",
	"ShellName": "Unknown",
	"ShellVersion": "Unknown",
	"Links": {
	  "Self": "/api/Spaces-1/machines/Machines-2646",
	  "Connection": "/api/Spaces-1/machines/Machines-2646/connection",
	  "TasksTemplate": "/api/Spaces-1/machines/Machines-2646/tasks{?skip,take}"
	}
  }`

const offlinePackageDropAsJSON string = `{
	"Id": "Machines-1864",
	"EnvironmentIds": [
	  "Environments-3721"
	],
	"Roles": [
	  "Prod"
	],
	"TenantedDeploymentParticipation": "Untenanted",
	"TenantIds": [],
	"TenantTags": [],
	"SpaceId": "Spaces-1",
	"Name": "1dd22c6a-f66a-49af-a111-127ee4e0",
	"Thumbprint": null,
	"Uri": null,
	"IsDisabled": true,
	"MachinePolicyId": "MachinePolicies-1",
	"Status": "Disabled",
	"HealthStatus": "Unknown",
	"HasLatestCalamari": false,
	"StatusSummary": "This machine has been disabled.",
	"IsInProcess": false,
	"Endpoint": {
	  "CommunicationStyle": "OfflineDrop",
	  "Destination": {
		"DestinationType": "Artifact",
		"DropFolderPath": null
	  },
	  "SensitiveVariablesEncryptionPassword": {
		"HasValue": false,
		"NewValue": null
	  },
	  "ApplicationsDirectory": "C:\\Applications",
	  "OctopusWorkingDirectory": "C:\\Octopus",
	  "DropFolderPath": null,
	  "Id": null,
	  "LastModifiedOn": null,
	  "LastModifiedBy": null,
	  "Links": {}
	},
	"OperatingSystem": "Unknown",
	"ShellName": "Unknown",
	"ShellVersion": "Unknown",
	"Links": {
	  "Self": "/api/Spaces-1/machines/Machines-1864",
	  "Connection": "/api/Spaces-1/machines/Machines-1864/connection",
	  "TasksTemplate": "/api/Spaces-1/machines/Machines-1864/tasks{?skip,take}"
	}
  }`

const kubernetesAsJSON string = `{
	"Id": "Machines-2216",
	"EnvironmentIds": [
	  "Environments-930"
	],
	"Roles": [
	  "Prod"
	],
	"TenantedDeploymentParticipation": "Untenanted",
	"TenantIds": [],
	"TenantTags": [],
	"SpaceId": "Spaces-1",
	"Name": "a",
	"Thumbprint": null,
	"Uri": null,
	"IsDisabled": false,
	"MachinePolicyId": "MachinePolicies-1",
	"Status": "CalamariNeedsUpgrade",
	"HealthStatus": "Unhealthy",
	"HasLatestCalamari": false,
	"StatusSummary": "There was a problem communicating with this machine (last checked: Friday, 02 October 2020 4:22:08 AM +00:00)",
	"IsInProcess": false,
	"Endpoint": {
	  "CommunicationStyle": "Kubernetes",
	  "ClusterCertificate": "cluster-certificate",
	  "ClusterUrl": "https://kubernetes.example.com",
	  "Namespace": "namespace",
	  "SkipTlsVerification": "True",
	  "ProxyId": "proxy-id",
	  "DefaultWorkerPoolId": "default-worker-pool-id",
	  "Container": {
		"Image": null,
		"FeedId": "feed-id"
	  },
	  "Authentication": {
		"AccountId": "Accounts-675",
		"AuthenticationType": "KubernetesStandard"
	  },
	  "Id": "kubernetes-endpoint",
	  "LastModifiedOn": null,
	  "LastModifiedBy": "alice.smith@example.com",
	  "Links": {}
	},
	"OperatingSystem": "Unknown",
	"ShellName": "Unknown",
	"ShellVersion": "Unknown",
	"Links": {
	  "Self": "/api/Spaces-1/machines/Machines-2216",
	  "Connection": "/api/Spaces-1/machines/Machines-2216/connection",
	  "TasksTemplate": "/api/Spaces-1/machines/Machines-2216/tasks{?skip,take}"
	}
  }`
