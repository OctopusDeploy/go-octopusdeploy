package octopusdeploy

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestSSHEndpoint(t *testing.T) *SSHEndpoint {
	accountID := "Accounts-316"
	dotNetCorePlatform := "linux-x64"
	fingerprint := "22:22:22:22:22:22:22:22:22:22:22:22:22:22:22"
	host := "example.com"
	id := "endpoint-id"
	lastModifiedBy := "john.smith@example.com"
	lastModifiedOn, _ := time.Parse(time.RFC3339, "2020-10-02T00:44:11.284Z")
	links := map[string]string{
		"Self": "/api/foo/bar/quux",
		"test": "/api/xyzzy",
	}
	proxyID := "proxy-id"
	port := 22

	sshEndpoint := NewSSHEndpoint(host, port, fingerprint)
	require.NoError(t, sshEndpoint.Validate())

	sshEndpoint.AccountID = accountID
	sshEndpoint.DotNetCorePlatform = dotNetCorePlatform
	sshEndpoint.ID = id
	sshEndpoint.ModifiedBy = lastModifiedBy
	sshEndpoint.ModifiedOn = &lastModifiedOn
	sshEndpoint.Links = links
	sshEndpoint.ProxyID = proxyID

	require.NoError(t, sshEndpoint.Validate())

	return sshEndpoint
}

func TestSSHEndpointNew(t *testing.T) {
	resource := NewSSHEndpoint("example.com", 22, "fingerprint-1")
	require.NoError(t, resource.Validate())
}

func TestSSHEndpointMarshalJSON(t *testing.T) {
	endpoint := CreateTestSSHEndpoint(t)

	jsonEncoding, err := json.Marshal(endpoint)
	require.NoError(t, err)
	require.NotNil(t, jsonEncoding)

	jsonassert.New(t).Assertf(string(jsonEncoding), sshEndpointAsJSON)
}

func TestSSHEndpointUnmarshalJSON(t *testing.T) {
	var resource *SSHEndpoint
	err := json.Unmarshal([]byte(sshEndpointAsJSON), &resource)

	require.NoError(t, err)
	require.NotNil(t, resource)

	endpoint := CreateTestSSHEndpoint(t)
	assert.Equal(t, endpoint, resource)
}

const sshEndpointAsJSON string = `{
	"CommunicationStyle": "Ssh",
	"AccountId": "Accounts-316",
	"Host": "example.com",
	"Port": 22,
	"Fingerprint": "22:22:22:22:22:22:22:22:22:22:22:22:22:22:22",
	"Uri": "ssh://example.com:22/",
	"ProxyId": "proxy-id",
	"DotNetCorePlatform": "linux-x64",
	"Id": "endpoint-id",
	"LastModifiedOn": "2020-10-02T00:44:11.284Z",
	"LastModifiedBy": "john.smith@example.com",
	"Links": {
		"Self": "/api/foo/bar/quux",
		"test": "/api/xyzzy"
	}
}`
