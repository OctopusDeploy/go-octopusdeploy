package octopusdeploy

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkersNew(t *testing.T) {
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)
	require.NotNil(t, url)

	kubernetesEndpoint := NewKubernetesEndpoint(url)
	require.NotNil(t, kubernetesEndpoint)

	name := getRandomName()

	worker := NewWorker(name, kubernetesEndpoint)
	require.NotNil(t, worker)
	require.NoError(t, worker.Validate())
}

func TestWorkersUnmarshalJSON(t *testing.T) {
	var tentaclePassiveWorker Worker
	err := json.Unmarshal([]byte(tentaclePassiveWorkerAsJSON), &tentaclePassiveWorker)
	require.NoError(t, err)
	require.NotNil(t, tentaclePassiveWorker)

	endpointLastModifiedOn, _ := time.Parse(time.RFC3339, "2020-10-02T00:44:11.284Z")
	lastModifiedOn, _ := time.Parse(time.RFC3339, "2019-10-02T00:44:11.284Z")
	links := map[string]string{
		"Self":       "/api/Spaces-1/workers/Workers-1",
		"Connection": "/api/Spaces-1/workers/Workers-1/connection",
	}
	url, _ := url.Parse("https://example.com")
	version := "1.2.3"

	endpoint := NewListeningTentacleEndpoint(url, "0E80575D3BD8D30854802A4F8340A2E23AE961FD")
	endpoint.ID = "endpoint-123"
	endpoint.ModifiedBy = "john.smith@example.com"
	endpoint.ModifiedOn = &endpointLastModifiedOn
	endpoint.ProxyID = "fake-proxy-id"
	endpoint.TentacleVersionDetails = NewTentacleVersionDetails(&version, true, true, false)

	listeningTentacleEndpoint := tentaclePassiveWorker.Endpoint.(*ListeningTentacleEndpoint)

	assert.Equal(t, endpoint.CertificateSignatureAlgorithm, listeningTentacleEndpoint.CertificateSignatureAlgorithm)
	assert.Equal(t, endpoint.GetID(), listeningTentacleEndpoint.GetID())
	assert.Equal(t, endpoint.GetModifiedBy(), listeningTentacleEndpoint.GetModifiedBy())
	assert.Equal(t, endpoint.GetModifiedOn(), listeningTentacleEndpoint.GetModifiedOn())
	assert.Equal(t, endpoint.Links, listeningTentacleEndpoint.Links)
	assert.Equal(t, endpoint.ProxyID, listeningTentacleEndpoint.ProxyID)
	assert.Equal(t, endpoint.resource, listeningTentacleEndpoint.resource)
	assert.Equal(t, endpoint.TentacleVersionDetails, listeningTentacleEndpoint.TentacleVersionDetails)
	assert.Equal(t, endpoint.Thumbprint, listeningTentacleEndpoint.Thumbprint)
	assert.Equal(t, endpoint.URI, listeningTentacleEndpoint.URI)
	assert.Equal(t, "alice.smith@example.com", tentaclePassiveWorker.GetModifiedBy())
	assert.Equal(t, &lastModifiedOn, tentaclePassiveWorker.GetModifiedOn())
	assert.Equal(t, links, tentaclePassiveWorker.Links)
	assert.False(t, tentaclePassiveWorker.HasLatestCalamari)
	assert.Equal(t, "Unknown", tentaclePassiveWorker.HealthStatus)
	assert.Equal(t, "Workers-1", tentaclePassiveWorker.GetID())
	assert.False(t, tentaclePassiveWorker.IsDisabled)
	assert.False(t, tentaclePassiveWorker.IsInProcess)
	assert.Equal(t, "MachinePolicies-1", tentaclePassiveWorker.MachinePolicyID)
	assert.Equal(t, "sdfsdfsd", tentaclePassiveWorker.Name)
	assert.Equal(t, "Unknown", tentaclePassiveWorker.ShellName)
	assert.Equal(t, "Unknown", tentaclePassiveWorker.ShellVersion)
	assert.Equal(t, "Spaces-1", tentaclePassiveWorker.SpaceID)
	assert.Equal(t, "Unknown", tentaclePassiveWorker.Status)
	assert.Equal(t, "Unknown", tentaclePassiveWorker.OperatingSystem)
	assert.Equal(t, "This machine was recently added. Please perform a health check.", tentaclePassiveWorker.StatusSummary)
	assert.Equal(t, "0E80575D3BD8D30854802A4E8340A2E23AE961FD", tentaclePassiveWorker.Thumbprint)
	assert.Equal(t, "https://example.com/", tentaclePassiveWorker.URI)
	assert.Equal(t, []string{"WorkerPools-83"}, tentaclePassiveWorker.WorkerPoolIDs)

	var sshWorker Worker
	err = json.Unmarshal([]byte(sshWorkerAsJSON), &sshWorker)
	require.NoError(t, err)
	require.NotNil(t, sshWorker)
}

const tentaclePassiveWorkerAsJSON string = `{
	"Id": "Workers-1",
	"WorkerPoolIds": [
	  "WorkerPools-83"
	],
	"SpaceId": "Spaces-1",
	"Name": "sdfsdfsd",
	"Thumbprint": "0E80575D3BD8D30854802A4E8340A2E23AE961FD",
	"Uri": "https://example.com/",
	"IsDisabled": false,
	"MachinePolicyId": "MachinePolicies-1",
	"Status": "Unknown",
	"HealthStatus": "Unknown",
	"HasLatestCalamari": false,
	"StatusSummary": "This machine was recently added. Please perform a health check.",
	"IsInProcess": false,
	"Endpoint": {
	  "CommunicationStyle": "TentaclePassive",
	  "Uri": "https://example.com",
	  "ProxyId": "fake-proxy-id",
	  "Thumbprint": "0E80575D3BD8D30854802A4F8340A2E23AE961FD",
	  "TentacleVersionDetails": {
		"UpgradeLocked": true,
		"Version": "1.2.3",
		"UpgradeSuggested": true,
		"UpgradeRequired": false
	  },
	  "CertificateSignatureAlgorithm": null,
	  "Id": "endpoint-123",
	  "LastModifiedOn": "2020-10-02T00:44:11.284Z",
	  "LastModifiedBy": "john.smith@example.com",
	  "Links": {}
	},
	"OperatingSystem": "Unknown",
	"ShellName": "Unknown",
	"ShellVersion": "Unknown",
	"LastModifiedOn": "2019-10-02T00:44:11.284Z",
	"LastModifiedBy": "alice.smith@example.com",
	"Links": {
	  "Self": "/api/Spaces-1/workers/Workers-1",
	  "Connection": "/api/Spaces-1/workers/Workers-1/connection"
	}
  }`

const sshWorkerAsJSON string = `{
	"Id": "Workers-2",
	"WorkerPoolIds": [
	  "WorkerPools-83"
	],
	"SpaceId": "Spaces-1",
	"Name": "SSH Connection",
	"Thumbprint": null,
	"Uri": null,
	"IsDisabled": false,
	"MachinePolicyId": "MachinePolicies-1",
	"Status": "CalamariNeedsUpgrade",
	"HealthStatus": "Unhealthy",
	"HasLatestCalamari": false,
	"StatusSummary": "There was a problem communicating with this machine (last checked: Saturday, 03 October 2020 8:19:56 AM +00:00)",
	"IsInProcess": false,
	"Endpoint": {
	  "CommunicationStyle": "Ssh",
	  "AccountId": "Accounts-352",
	  "Host": "hostname.com",
	  "Port": 22,
	  "Fingerprint": "66:66:66:66:66:66:66:66:66:66:66:66:66:66",
	  "Uri": "ssh://hostname.com:22/",
	  "ProxyId": null,
	  "DotNetCorePlatform": "linux-x64",
	  "Id": "endpoint-123",
	  "LastModifiedOn": "2020-10-02T00:44:11.284Z",
	  "LastModifiedBy": "john.smith@example.com",
	  "Links": {}
	},
	"OperatingSystem": "Unknown",
	"ShellName": "Unknown",
	"ShellVersion": "Unknown",
	"Links": {
	  "Self": "/api/Spaces-1/workers/Workers-2",
	  "Connection": "/api/Spaces-1/workers/Workers-2/connection"
	}
  }`
