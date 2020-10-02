package model

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKubernetesEndpoints(t *testing.T) {
	t.Run("CommunicationStyles", TestKubernetesEndpointCommunicationStyles)
	t.Run("AsJSON", TestKubernetesEndpointAsJSON)
}

func TestKubernetesEndpointCommunicationStyles(t *testing.T) {
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)
	require.NotNil(t, url)

	resource := NewKubernetesEndpoint(*url, "Pools-1")
	assert.NoError(t, resource.Validate())

	resource.CommunicationStyle = "None"
	assert.Error(t, resource.Validate())

	resource.CommunicationStyle = "none"
	assert.Error(t, resource.Validate())

	resource.CommunicationStyle = "TentaclePassive"
	assert.Error(t, resource.Validate())

	resource.CommunicationStyle = "TentacleActive"
	assert.Error(t, resource.Validate())

	resource.CommunicationStyle = "Ssh"
	assert.Error(t, resource.Validate())

	resource.CommunicationStyle = "SshOfflineDrop"
	assert.Error(t, resource.Validate())

	resource.CommunicationStyle = "AzureWebApp"
	assert.Error(t, resource.Validate())

	resource.CommunicationStyle = "Ftp"
	assert.Error(t, resource.Validate())

	resource.CommunicationStyle = "AzureCloudService"
	assert.Error(t, resource.Validate())

	resource.CommunicationStyle = "AzureServiceFabricCluster"
	assert.Error(t, resource.Validate())

	resource.CommunicationStyle = "Kubernetes"
	assert.NoError(t, resource.Validate())

	resource.CommunicationStyle = "kubernetes"
	assert.Error(t, resource.Validate())
}

func TestKubernetesEndpointAsJSON(t *testing.T) {
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)
	require.NotNil(t, url)

	resource := NewKubernetesEndpoint(*url, "Pools-1")
	require.NoError(t, resource.Validate())

	jsonEncoding, err := json.Marshal(resource)
	require.NoError(t, err)
	require.NotNil(t, jsonEncoding)

	actual := string(jsonEncoding)

	expected := `{
		"Authentication": {},
		"ClusterCertificate": "",
		"ClusterUrl": "https://example.com/",
		"CommunicationStyle": "Kubernetes",
		"Container": {
			"Image": null,
			"FeedId": null
		},
		"DefaultWorkerPoolId": "Pools-1",
		"Namespace": "",
		"ProxyId": "",
		"RunningInContainer": false,
		"SkipTlsVerification": ""
	}`

	jsonassert.New(t).Assertf(actual, expected)

	expected = `{
		"Authentication": {
			"AccountId": "Accounts-392",
			"AuthenticationType": "KubernetesStandard",
			"ClientCertificate": "fjasd"
		},
		"ClusterCertificate": "Certificates-22-r-BY2FT",
		"ClusterUrl": "https://kubernetes.example.com",
		"CommunicationStyle": "Kubernetes",
		"Container": {
			"Image": "image",
			"FeedId": "feed-id"
		},
		"DefaultWorkerPoolId": "new-worker-pool-1",
		"Namespace": "default",
		"ProxyId": "proxy-id",
		"RunningInContainer": true,
		"SkipTlsVerification": "True",
		"Id": "asd",
		"LastModifiedBy": "john.smith@example.com"

	}`

	url, _ = url.Parse("https://kubernetes.example.com")

	var unmarshalledEndpoint KubernetesEndpoint
	err = json.Unmarshal([]byte(expected), &unmarshalledEndpoint)
	require.NoError(t, err)
	require.NotNil(t, unmarshalledEndpoint)
	require.NotNil(t, unmarshalledEndpoint.Authentication)
	require.NotNil(t, unmarshalledEndpoint.Container)
	assert.Equal(t, unmarshalledEndpoint.Authentication.AccountID, "Accounts-392")
	assert.Equal(t, unmarshalledEndpoint.Authentication.AuthenticationType, "KubernetesStandard")
	assert.Equal(t, unmarshalledEndpoint.Authentication.ClientCertificate, "fjasd")
	assert.Equal(t, unmarshalledEndpoint.ClusterCertificate, "Certificates-22-r-BY2FT")
	assert.Equal(t, unmarshalledEndpoint.ClusterURL, *url)
	assert.Equal(t, "Kubernetes", unmarshalledEndpoint.CommunicationStyle)
	assert.Equal(t, "image", *unmarshalledEndpoint.Container.Image)
	assert.Equal(t, "feed-id", *unmarshalledEndpoint.Container.FeedID)
	assert.Equal(t, unmarshalledEndpoint.DefaultWorkerPoolID, "new-worker-pool-1")
	assert.Equal(t, unmarshalledEndpoint.ID, "asd")
	assert.Equal(t, unmarshalledEndpoint.Namespace, "default")
	assert.Equal(t, unmarshalledEndpoint.ProxyID, "proxy-id")
	assert.True(t, unmarshalledEndpoint.RunningInContainer)
	assert.Equal(t, unmarshalledEndpoint.SkipTLSVerification, "True")
	assert.Equal(t, unmarshalledEndpoint.LastModifiedBy, "john.smith@example.com")
}
