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

func TestListeningTentacleEndpointMarshalJSON(t *testing.T) {
	lastModifiedOn, _ := time.Parse(time.RFC3339, "2020-10-02T00:44:11.284Z")
	links := map[string]string{
		"Self":   "/api/Spaces-1/fake/fake-1",
		"Fake-2": "/api/Spaces-1/fake/fake-1/fake-2",
	}
	url, _ := url.Parse("https://example.com")
	version := "1.2.3"

	listeningTentacleEndpoint := NewListeningTentacleEndpoint(url, "thumbprint")
	listeningTentacleEndpoint.CertificateSignatureAlgorithm = "fake-algorithm-1"
	listeningTentacleEndpoint.ID = "endpoint-123"
	listeningTentacleEndpoint.ModifiedBy = "john.smith@example.com"
	listeningTentacleEndpoint.ModifiedOn = &lastModifiedOn
	listeningTentacleEndpoint.Links = links
	listeningTentacleEndpoint.ProxyID = "fake-proxy-id"
	listeningTentacleEndpoint.TentacleVersionDetails = NewTentacleVersionDetails(&version, true, true, false)
	listeningTentacleEndpoint.Thumbprint = "0E80576D3BD8D40854802A4F8340A2E23AE961FD"

	jsonEncoding, err := json.Marshal(listeningTentacleEndpoint)
	require.NoError(t, err)
	require.NotNil(t, jsonEncoding)

	actual := string(jsonEncoding)

	jsonassert.New(t).Assertf(actual, listeningTentacleEndpointAsJSON)

}

func TestListeningTentacleEndpointUnmarshalJSON(t *testing.T) {
	var listeningTentacleEndpoint ListeningTentacleEndpoint
	err := json.Unmarshal([]byte(listeningTentacleEndpointAsJSON), &listeningTentacleEndpoint)
	require.NoError(t, err)
	require.NotNil(t, listeningTentacleEndpoint)

	lastModifiedOn, _ := time.Parse(time.RFC3339, "2020-10-02T00:44:11.284Z")
	links := map[string]string{
		"Self":   "/api/Spaces-1/fake/fake-1",
		"Fake-2": "/api/Spaces-1/fake/fake-1/fake-2",
	}

	version := "1.2.3"
	tentacleVersionDetails := NewTentacleVersionDetails(&version, true, true, false)

	url, _ := url.Parse("https://example.com")

	assert.Equal(t, "fake-algorithm-1", listeningTentacleEndpoint.CertificateSignatureAlgorithm)
	assert.Equal(t, "endpoint-123", listeningTentacleEndpoint.GetID())
	assert.Equal(t, "john.smith@example.com", listeningTentacleEndpoint.GetModifiedBy())
	assert.Equal(t, &lastModifiedOn, listeningTentacleEndpoint.GetModifiedOn())
	assert.Equal(t, links, listeningTentacleEndpoint.Links)
	assert.Equal(t, "fake-proxy-id", listeningTentacleEndpoint.ProxyID)
	assert.Equal(t, tentacleVersionDetails, listeningTentacleEndpoint.TentacleVersionDetails)
	assert.Equal(t, "0E80576D3BD8D40854802A4F8340A2E23AE961FD", listeningTentacleEndpoint.Thumbprint)
	assert.Equal(t, tentacleVersionDetails, listeningTentacleEndpoint.TentacleVersionDetails)
	assert.Equal(t, url, listeningTentacleEndpoint.URI)
}

const listeningTentacleEndpointAsJSON string = `{
  "CertificateSignatureAlgorithm": "fake-algorithm-1",
  "CommunicationStyle": "TentaclePassive",
  "Id": "endpoint-123",
  "LastModifiedBy": "john.smith@example.com",
  "LastModifiedOn": "2020-10-02T00:44:11.284Z",
  "Links": {
    "Self": "/api/Spaces-1/fake/fake-1",
    "Fake-2": "/api/Spaces-1/fake/fake-1/fake-2"
  },
  "ProxyId": "fake-proxy-id",
  "TentacleVersionDetails": {
	"UpgradeLocked": true,
	"UpgradeRequired": false,
	"UpgradeSuggested": true,
	"Version": "1.2.3"
  },
  "Thumbprint": "0E80576D3BD8D40854802A4F8340A2E23AE961FD",
  "Uri": "https://example.com"
}`
