package machines

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newEcsEndpointResource() *EndpointResource {
	resource := NewEndpointResource("AwsEcsCluster")
	resource.AccountID = "Accounts-1"
	resource.AssumedRoleARN = "arn:aws:iam::123456789012:role/OctopusEcsDeploy"
	resource.AssumedRoleSession = "octopus-cli-604-session"
	resource.AssumeRole = true
	resource.AssumeRoleExternalID = "external-id-604"
	resource.AssumeRoleSessionDurationSeconds = 3600
	resource.ClusterName = "repro-604-assumerole-cluster"
	resource.DefaultWorkerPoolID = "WorkerPools-1"
	resource.Region = "eu-west-1"
	resource.UseInstanceRole = false

	return resource
}

// The conversion was unreachable for an ECS resource while ClusterUrl,
// Thumbprint and Uri were required of every communication style.
func TestToEndpointAwsEcsCluster(t *testing.T) {
	converted, err := ToEndpoint(newEcsEndpointResource())
	require.NoError(t, err)

	endpoint, ok := converted.(*AwsEcsClusterEndpoint)
	require.True(t, ok)
	assert.Equal(t, "repro-604-assumerole-cluster", endpoint.ClusterName)
	assert.Equal(t, "eu-west-1", endpoint.Region)
	assert.Equal(t, "Accounts-1", endpoint.AccountID)
	assert.Equal(t, "WorkerPools-1", endpoint.DefaultWorkerPoolID)
	assert.True(t, endpoint.AssumeRole)
	assert.Equal(t, "arn:aws:iam::123456789012:role/OctopusEcsDeploy", endpoint.AssumedRoleARN)
	assert.Equal(t, "octopus-cli-604-session", endpoint.AssumedRoleSession)
	assert.Equal(t, 3600, endpoint.AssumeRoleSessionDurationSeconds)
	assert.Equal(t, "external-id-604", endpoint.AssumeRoleExternalID)
	assert.False(t, endpoint.UseInstanceRole)
}

// Every field survives resource -> endpoint -> resource, which is the point of
// the pair being symmetric.
func TestAwsEcsClusterSurvivesARoundTripThroughBothConversions(t *testing.T) {
	original := newEcsEndpointResource()

	endpoint, err := ToEndpoint(original)
	require.NoError(t, err)

	returned, err := ToEndpointResource(endpoint)
	require.NoError(t, err)

	assert.Equal(t, original.CommunicationStyle, returned.CommunicationStyle)
	assert.Equal(t, original.AccountID, returned.AccountID)
	assert.Equal(t, original.AssumedRoleARN, returned.AssumedRoleARN)
	assert.Equal(t, original.AssumedRoleSession, returned.AssumedRoleSession)
	assert.Equal(t, original.AssumeRole, returned.AssumeRole)
	assert.Equal(t, original.AssumeRoleExternalID, returned.AssumeRoleExternalID)
	assert.Equal(t, original.AssumeRoleSessionDurationSeconds, returned.AssumeRoleSessionDurationSeconds)
	assert.Equal(t, original.ClusterName, returned.ClusterName)
	assert.Equal(t, original.DefaultWorkerPoolID, returned.DefaultWorkerPoolID)
	assert.Equal(t, original.Region, returned.Region)
	assert.Equal(t, original.UseInstanceRole, returned.UseInstanceRole)
}

// Validation now asks each style only for the fields it uses.
func TestEndpointResourceValidateIsStyleAware(t *testing.T) {
	uri := &url.URL{Scheme: "https", Host: "tentacle:10933"}

	t.Run("a tentacle still needs its uri and thumbprint", func(t *testing.T) {
		bare := NewEndpointResource("TentaclePassive")
		require.Error(t, bare.Validate())

		complete := NewEndpointResource("TentaclePassive")
		complete.URI = uri
		complete.Thumbprint = "thumbprint"
		require.NoError(t, complete.Validate())
	})

	t.Run("kubernetes still needs its cluster url", func(t *testing.T) {
		bare := NewEndpointResource("Kubernetes")
		require.Error(t, bare.Validate())

		complete := NewEndpointResource("Kubernetes")
		complete.ClusterURL = &url.URL{Scheme: "https", Host: "cluster"}
		require.NoError(t, complete.Validate())
	})

	// These used to fail for want of a cluster url, thumbprint and uri that
	// none of them has.
	for _, style := range []string{"AwsEcsCluster", "Ssh", "None", "OfflineDrop", "AzureWebApp"} {
		t.Run(style+" no longer needs fields it does not use", func(t *testing.T) {
			require.NoError(t, NewEndpointResource(style).Validate())
		})
	}
}
