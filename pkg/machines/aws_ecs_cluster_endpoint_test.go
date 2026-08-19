package machines

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

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

// The same target with account credentials and an assumed role, every field the
// server can populate set. Captured from /api/Spaces-1/machines/Machines-89.
const assumeRoleEndpointAsJSON = `{
  "CommunicationStyle": "AwsEcsCluster",
  "DefaultWorkerPoolId": "WorkerPools-1",
  "ClusterName": "repro-604-assumerole-cluster",
  "Region": "eu-west-1",
  "AccountId": "Accounts-1",
  "UseInstanceRole": false,
  "AssumeRole": true,
  "AssumedRoleArn": "arn:aws:iam::123456789012:role/OctopusEcsDeploy",
  "AssumedRoleSession": "octopus-cli-604-session",
  "AssumeRoleSessionDurationSeconds": 3600,
  "AssumeRoleExternalId": "external-id-604",
  "Id": null,
  "LastModifiedOn": null,
  "LastModifiedBy": null,
  "Links": {}
}`

func TestAwsEcsClusterEndpointAssumeRole(t *testing.T) {
	var endpoint AwsEcsClusterEndpoint
	require.NoError(t, json.Unmarshal([]byte(assumeRoleEndpointAsJSON), &endpoint))

	assert.Equal(t, "repro-604-assumerole-cluster", endpoint.ClusterName)
	assert.Equal(t, "eu-west-1", endpoint.Region)
	assert.Equal(t, "Accounts-1", endpoint.AccountID)
	assert.False(t, endpoint.UseInstanceRole)
	assert.True(t, endpoint.AssumeRole)
	assert.Equal(t, "arn:aws:iam::123456789012:role/OctopusEcsDeploy", endpoint.AssumedRoleARN)
	assert.Equal(t, "octopus-cli-604-session", endpoint.AssumedRoleSession)
	assert.Equal(t, 3600, endpoint.AssumeRoleSessionDurationSeconds)
	assert.Equal(t, "external-id-604", endpoint.AssumeRoleExternalID)
}

// Guards field coverage: if the server starts sending a key this endpoint does
// not model, this fails rather than silently dropping it. The payloads above are
// captured from a server, so a Server-side addition shows up here on capture.
func TestAwsEcsClusterEndpointModelsEveryFieldTheServerSends(t *testing.T) {
	// Populated so that no omitempty field drops out of the comparison.
	modelled := &AwsEcsClusterEndpoint{
		AccountID:                        "Accounts-1",
		AssumedRoleARN:                   "arn",
		AssumedRoleSession:               "session",
		AssumeRole:                       true,
		AssumeRoleExternalID:             "external",
		AssumeRoleSessionDurationSeconds: 3600,
		ClusterName:                      "cluster",
		DefaultWorkerPoolID:              "WorkerPools-1",
		Region:                           "us-east-1",
		UseInstanceRole:                  true,
		endpoint:                         *newEndpoint("AwsEcsCluster"),
	}
	modelled.SetID("Machines-89")
	modelled.SetModifiedBy("nick")
	now := time.Now()
	modelled.SetModifiedOn(&now)
	modelled.SetLinks(map[string]string{"Self": "/api/Spaces-1/machines/Machines-89"})

	encoded, err := json.Marshal(modelled)
	require.NoError(t, err)

	var known map[string]any
	require.NoError(t, json.Unmarshal(encoded, &known))

	for _, payload := range []string{ecsClusterTargetsAsJSON, assumeRoleEndpointAsJSON} {
		for _, sent := range endpointsIn(t, payload) {
			for key := range sent {
				assert.Contains(t, known, key, "the server sends %q but AwsEcsClusterEndpoint does not model it", key)
			}
		}
	}
}

// endpointsIn returns the endpoint objects in a captured payload, which is
// either a target, a list of targets, or a bare endpoint.
func endpointsIn(t *testing.T, payload string) []map[string]any {
	t.Helper()

	var list []map[string]any
	if err := json.Unmarshal([]byte(payload), &list); err == nil {
		endpoints := []map[string]any{}
		for _, target := range list {
			endpoints = append(endpoints, target["Endpoint"].(map[string]any))
		}
		return endpoints
	}

	var single map[string]any
	require.NoError(t, json.Unmarshal([]byte(payload), &single))
	if endpoint, ok := single["Endpoint"].(map[string]any); ok {
		return []map[string]any{endpoint}
	}

	return []map[string]any{single}
}
