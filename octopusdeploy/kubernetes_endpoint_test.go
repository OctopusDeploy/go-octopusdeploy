package octopusdeploy

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestKubernetesEndpoint(t *testing.T) *KubernetesEndpoint {
	authentication := &EndpointAuthentication{}
	clusterCertificate := ""
	defaultWorkerPoolID := "default-worker-pool-id"
	lastModifiedBy := "john.smith@example.com"
	lastModifiedOn, _ := time.Parse(time.RFC3339, "2020-10-02T00:44:11.284Z")
	links := map[string]string{
		"Self": "/api/foo/bar/quux",
		"test": "/api/xyzzy",
	}
	id := "endpoint-id"
	proxyID := "proxy-id"
	url, _ := url.Parse("https://example.com/")

	kubernetesEndpoint := NewKubernetesEndpoint(*url)
	require.NoError(t, kubernetesEndpoint.Validate())

	kubernetesEndpoint.DefaultWorkerPoolID = defaultWorkerPoolID
	kubernetesEndpoint.ClusterCertificate = clusterCertificate
	kubernetesEndpoint.Authentication = *authentication
	kubernetesEndpoint.ID = id
	kubernetesEndpoint.LastModifiedBy = lastModifiedBy
	kubernetesEndpoint.LastModifiedOn = &lastModifiedOn
	kubernetesEndpoint.Links = links
	kubernetesEndpoint.ProxyID = proxyID

	require.NoError(t, kubernetesEndpoint.Validate())

	return kubernetesEndpoint
}

func TestKubernetesEndpointNew(t *testing.T) {
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)
	require.NotNil(t, url)

	resource := NewKubernetesEndpoint(*url)
	assert.NoError(t, resource.Validate())
}

func TestKubernetesEndpointMarshalJSON(t *testing.T) {
	feedID := "feed-id"
	lastModifiedOn, _ := time.Parse(time.RFC3339, "2020-10-02T00:44:11.284Z")
	links := map[string]string{
		"Self": "/api/foo/bar/quux",
		"test": "/api/xyzzy",
	}
	id := "endpoint-id"
	url, _ := url.Parse("https://example.com/")

	resource := NewKubernetesEndpoint(*url)
	resource.ClusterCertificate = "cluster-certificate"
	resource.Container.FeedID = &feedID
	resource.DefaultWorkerPoolID = "default-worker-pool-id"
	resource.ID = id
	resource.LastModifiedBy = "john.smith@example.com"
	resource.LastModifiedOn = &lastModifiedOn
	resource.Links = links
	resource.Namespace = "namespace-test"
	resource.SkipTLSVerification = true
	resource.ProxyID = "proxy-id"

	require.NoError(t, resource.Validate())

	jsonEncoding, err := json.Marshal(resource)
	require.NoError(t, err)
	require.NotNil(t, jsonEncoding)

	actual := string(jsonEncoding)

	expected := `{
		"Authentication": {},
		"ClusterCertificate": "cluster-certificate",
		"ClusterUrl": "https://example.com/",
		"CommunicationStyle": "Kubernetes",
		"Container": {
			"Image": null,
			"FeedId": "feed-id"
		},
		"DefaultWorkerPoolId": "default-worker-pool-id",
		"Namespace": "namespace-test",
		"ProxyId": "proxy-id",
		"RunningInContainer": false,
		"SkipTlsVerification": "True",
        "Id": "endpoint-id",
        "LastModifiedOn": "2020-10-02T00:44:11.284Z",
        "LastModifiedBy": "john.smith@example.com",
		"Links": {
			"Self": "/api/foo/bar/quux",
			"test": "/api/xyzzy"
		}
	}`

	jsonassert.New(t).Assertf(actual, expected)
}

func TestKubernetesEndpointUnmarshalJSON(t *testing.T) {
	lastModifiedOn, _ := time.Parse(time.RFC3339, "2020-10-02T00:44:11.284Z")
	links := map[string]string{
		"Self": "/api/foo/bar/quux",
		"test": "/api/xyzzy",
	}
	url, _ := url.Parse("https://kubernetes.example.com")

	var resource KubernetesEndpoint
	err := json.Unmarshal([]byte(kubernetesEndpointAsJSON), &resource)

	require.NoError(t, err)
	require.NotNil(t, resource)
	require.NotNil(t, resource.Authentication)
	require.NotNil(t, resource.Container)

	// Authentication field
	assert.Equal(t, "Accounts-392", resource.Authentication.AccountID)
	assert.Equal(t, "KubernetesStandard", resource.Authentication.AuthenticationType)
	assert.Equal(t, "client-certificate", resource.Authentication.ClientCertificate)

	// Container field
	assert.Equal(t, "image", *resource.Container.Image)
	assert.Equal(t, "feed-id", *resource.Container.FeedID)

	// basic fields
	assert.Equal(t, "Certificates-22-r-BY2FT", resource.ClusterCertificate)
	assert.Equal(t, url, resource.ClusterURL)
	assert.Equal(t, "default-worker-pool-id", resource.DefaultWorkerPoolID)
	assert.Equal(t, "default", resource.Namespace)
	assert.Equal(t, "proxy-id", resource.ProxyID)
	assert.True(t, resource.RunningInContainer)
	assert.False(t, resource.SkipTLSVerification)

	// resource
	assert.Equal(t, "endpoint-1", resource.ID)
	assert.Equal(t, "john.smith@example.com", resource.LastModifiedBy)
	assert.Equal(t, &lastModifiedOn, resource.LastModifiedOn)
	assert.Equal(t, links, resource.Links)
}

const kubernetesEndpointAsJSON string = `{
	"Authentication": {
		"AccountId": "Accounts-392",
		"AuthenticationType": "KubernetesStandard",
		"ClientCertificate": "client-certificate"
	},
	"ClusterCertificate": "Certificates-22-r-BY2FT",
	"ClusterUrl": "https://kubernetes.example.com",
	"CommunicationStyle": "Kubernetes",
	"Container": {
		"Image": "image",
		"FeedId": "feed-id"
	},
	"DefaultWorkerPoolId": "default-worker-pool-id",
	"Namespace": "default",
	"ProxyId": "proxy-id",
	"RunningInContainer": true,
	"SkipTlsVerification": "False",
	"Id": "endpoint-1",
	"LastModifiedOn": "2020-10-02T00:44:11.284Z",
	"LastModifiedBy": "john.smith@example.com",
	"Links": {
		"Self": "/api/foo/bar/quux",
		"test": "/api/xyzzy"
	}
}`
