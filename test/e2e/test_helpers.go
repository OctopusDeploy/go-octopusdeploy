package e2e

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/machines"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/stretchr/testify/require"
)

func getOctopusClient() *client.Client {
	host := os.Getenv(constants.EnvironmentVariableOctopusHost)
	apiKey := os.Getenv(constants.EnvironmentVariableOctopusApiKey)
	spaceId := os.Getenv(constants.EnvironmentVariableOctopusSpace)

	if len(host) == 0 {
		host = os.Getenv(constants.ClientURLEnvironmentVariable)
	}

	if len(apiKey) == 0 {
		apiKey = os.Getenv(constants.ClientAPIKeyEnvironmentVariable)
	}

	if internal.IsEmpty(host) || internal.IsEmpty(apiKey) {
		log.Fatal("Please make sure to set the env variables 'OCTOPUS_HOST' and 'OCTOPUS_API_KEY' before running this test")
	}

	apiURL, err := url.Parse(host)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return nil
	}

	// NOTE: You can direct traffic through a proxy trace like Fiddler
	// Everywhere by preconfiguring the client to route traffic through a
	// proxy.

	// proxyStr := "http://127.0.0.1:8866"
	// proxyURL, err := url.Parse(proxyStr)
	// if err != nil {
	// 	log.Println(err)
	// }

	// tr := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyURL),
	// }
	// httpClient := http.Client{Transport: tr}

	octopusClient, err := client.NewClient(nil, apiURL, apiKey, spaceId)
	if err != nil {
		log.Fatal(err)
	}

	return octopusClient
}

// setupLiveStatusClientForTest sets up the Octopus client and new client for livestatus service testing
func setupLiveStatusClientForTest(t *testing.T) (*client.Client, newclient.Client) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient, "octopusClient should not be nil - check environment variables")

	// Validate the client has required methods
	httpSession := octopusClient.HttpSession()
	require.NotNil(t, httpSession, "HttpSession should not be nil")

	clientSpaceID := octopusClient.GetSpaceID()
	require.NotEmpty(t, clientSpaceID, "SpaceID should not be empty")

	// Create a new client instance for the livestatus service
	newClient := newclient.NewClientS(httpSession, clientSpaceID)
	require.NotNil(t, newClient, "newClient should not be nil")

	return octopusClient, newClient
}

// CleanLiveStatusDeploymentTarget cleans up the deployment target used for livestatus service testing
func CleanLiveStatusDeploymentTarget(t *testing.T, client *client.Client, deploymentTarget *machines.DeploymentTarget) {
	if client == nil || deploymentTarget == nil {
		return
	}

	err := client.Machines.DeleteByID(deploymentTarget.GetID())
	require.NoError(t, err)
}
