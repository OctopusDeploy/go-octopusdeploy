package machines

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// An Amazon ECS cluster target as the server returns it, taken verbatim from
// /api/Spaces-1/machines. Both authentication shapes are represented: an
// instance role, and an account.
const ecsClusterTargetsAsJSON = `[{
  "Id": "Machines-63",
  "Name": "aws ecs",
  "EnvironmentIds": ["Environments-1"],
  "Roles": ["ecs"],
  "HealthStatus": "Healthy",
  "IsDisabled": false,
  "Endpoint": {
    "CommunicationStyle": "AwsEcsCluster",
    "DefaultWorkerPoolId": "WorkerPools-1",
    "ClusterName": "repro-604-cluster",
    "Region": "ap-southeast-2",
    "AccountId": "",
    "UseInstanceRole": true,
    "AssumeRole": false,
    "AssumedRoleArn": null,
    "AssumedRoleSession": null,
    "AssumeRoleSessionDurationSeconds": null,
    "AssumeRoleExternalId": null,
    "Id": null,
    "LastModifiedOn": null,
    "LastModifiedBy": null,
    "Links": {}
  }
},{
  "Id": "Machines-66",
  "Name": "aws ecs (account auth)",
  "EnvironmentIds": ["Environments-1"],
  "Roles": ["ecs"],
  "HealthStatus": "Healthy",
  "IsDisabled": false,
  "Endpoint": {
    "CommunicationStyle": "AwsEcsCluster",
    "DefaultWorkerPoolId": "WorkerPools-1",
    "ClusterName": "repro-604-account-cluster",
    "Region": "us-east-1",
    "AccountId": "Accounts-1",
    "UseInstanceRole": false,
    "AssumeRole": false,
    "AssumedRoleArn": null,
    "AssumedRoleSession": null,
    "AssumeRoleSessionDurationSeconds": null,
    "AssumeRoleExternalId": null,
    "Id": null,
    "LastModifiedOn": null,
    "LastModifiedBy": null,
    "Links": {}
  }
}]`

func TestAwsEcsClusterEndpointNew(t *testing.T) {
	endpoint := NewAwsEcsClusterEndpoint("my-cluster", "us-east-1")

	require.NoError(t, endpoint.Validate())
	assert.Equal(t, "AwsEcsCluster", endpoint.GetCommunicationStyle())
}

// The endpoint used to be discarded, leaving DeploymentTarget.Endpoint nil and
// panicking every consumer that read it (OctopusDeploy/cli#604).
func TestDeploymentTargetUnmarshalAwsEcsClusterEndpoint(t *testing.T) {
	var targets []*DeploymentTarget
	require.NoError(t, json.Unmarshal([]byte(ecsClusterTargetsAsJSON), &targets))
	require.Len(t, targets, 2)

	instanceRole, account := targets[0], targets[1]

	require.False(t, IsNil(instanceRole.Endpoint))
	assert.Equal(t, "AwsEcsCluster", instanceRole.Endpoint.GetCommunicationStyle())

	endpoint, ok := instanceRole.Endpoint.(*AwsEcsClusterEndpoint)
	require.True(t, ok)
	assert.Equal(t, "repro-604-cluster", endpoint.ClusterName)
	assert.Equal(t, "ap-southeast-2", endpoint.Region)
	assert.Equal(t, "WorkerPools-1", endpoint.DefaultWorkerPoolID)
	assert.True(t, endpoint.UseInstanceRole)
	assert.Empty(t, endpoint.AccountID)

	accountEndpoint, ok := account.Endpoint.(*AwsEcsClusterEndpoint)
	require.True(t, ok)
	assert.Equal(t, "Accounts-1", accountEndpoint.AccountID)
	assert.False(t, accountEndpoint.UseInstanceRole)
}

// The server sends null for the assume-role fields whenever it is not in use.
func TestAwsEcsClusterEndpointToleratesNulls(t *testing.T) {
	var targets []*DeploymentTarget
	require.NoError(t, json.Unmarshal([]byte(ecsClusterTargetsAsJSON), &targets))

	endpoint := targets[0].Endpoint.(*AwsEcsClusterEndpoint)

	assert.Empty(t, endpoint.AssumedRoleARN)
	assert.Empty(t, endpoint.AssumedRoleSession)
	assert.Empty(t, endpoint.AssumeRoleExternalID)
	assert.Zero(t, endpoint.AssumeRoleSessionDurationSeconds)
}

func TestAwsEcsClusterEndpointRunsOnAWorker(t *testing.T) {
	endpoint := NewAwsEcsClusterEndpoint("my-cluster", "us-east-1")
	endpoint.SetDefaultWorkerPoolID("WorkerPools-1")

	runsOnAWorker, ok := any(endpoint).(IRunsOnAWorker)
	require.True(t, ok)
	assert.Equal(t, "WorkerPools-1", runsOnAWorker.GetDefaultWorkerPoolID())
}

// ToEndpoint left endpoint nil for a style it had no case for and then called
// SetLinks on it, panicking inside the library.
func TestToEndpointUnmodelledCommunicationStyle(t *testing.T) {
	resource := NewEndpointResource("Ftp")
	resource.ClusterURL = &url.URL{Scheme: "https", Host: "cluster"}
	resource.Thumbprint = "thumbprint"
	resource.URI = &url.URL{Scheme: "https", Host: "host"}

	var endpoint IEndpoint
	var err error
	require.NotPanics(t, func() {
		endpoint, err = ToEndpoint(resource)
	})

	assert.Nil(t, endpoint)
	assert.Equal(t, internal.CreateInvalidParameterError("ToEndpoint", "endpointResource.CommunicationStyle"), err)
}

func TestIsNilRecognisesTypedNilEndpoints(t *testing.T) {
	var awsEcsCluster *AwsEcsClusterEndpoint
	var kubernetesTentacle *KubernetesTentacleEndpoint

	assert.True(t, IsNil(awsEcsCluster))
	assert.True(t, IsNil(kubernetesTentacle))
}
