package resources

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListeningTentacleEndpointMarshalJSON(t *testing.T) {
	url, _ := url.Parse("https://example.com")
	version := "1.2.3"

	listeningTentacleEndpoint := NewListeningTentacleEndpoint(url, "thumbprint")
	listeningTentacleEndpoint.CertificateSignatureAlgorithm = "fake-algorithm-1"
	listeningTentacleEndpoint.ID = "endpoint-123"
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

	version := "1.2.3"
	tentacleVersionDetails := NewTentacleVersionDetails(&version, true, true, false)

	url, _ := url.Parse("https://example.com")

	assert.Equal(t, "fake-algorithm-1", listeningTentacleEndpoint.CertificateSignatureAlgorithm)
	assert.Equal(t, "endpoint-123", listeningTentacleEndpoint.GetID())
	assert.Equal(t, "fake-proxy-id", listeningTentacleEndpoint.ProxyID)
	assert.Equal(t, tentacleVersionDetails, listeningTentacleEndpoint.TentacleVersionDetails)
	assert.Equal(t, "0E80576D3BD8D40854802A4F8340A2E23AE961FD", listeningTentacleEndpoint.Thumbprint)
	assert.Equal(t, tentacleVersionDetails, listeningTentacleEndpoint.TentacleVersionDetails)
	assert.Equal(t, url, listeningTentacleEndpoint.URI)
}

const listeningTentacleEndpointAsJSON string = `{
  "CommunicationStyle": "TentaclePassive",
  "Uri": "https://example.com",
  "ProxyId": "fake-proxy-id",
  "Thumbprint": "0E80576D3BD8D40854802A4F8340A2E23AE961FD",
  "TentacleVersionDetails": {
	"UpgradeLocked": true,
	"Version": "1.2.3",
	"UpgradeSuggested": true,
	"UpgradeRequired": false
  },
  "CertificateSignatureAlgorithm": "fake-algorithm-1",
  "Id": "endpoint-123",
}`
