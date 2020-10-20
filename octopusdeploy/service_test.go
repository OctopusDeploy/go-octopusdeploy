package octopusdeploy

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/stretchr/testify/require"
)

func getClient() *Client {
	octopusURL := os.Getenv("OCTOPUS_URL")
	octopusAPIKey := os.Getenv("OCTOPUS_APIKEY")

	if isEmpty(octopusURL) || isEmpty(octopusAPIKey) {
		log.Fatal("Please make sure to set the env variables 'OCTOPUS_URL' and 'OCTOPUS_APIKEY' before running this test")
	}

	// NOTE: You can direct traffic through a proxy trace like Fiddler
	// Everywhere by preconfiguring the client to route traffic through a
	// proxy.

	// proxyStr := "http://127.0.0.1:5555"
	// proxyURL, err := url.Parse(proxyStr)
	// if err != nil {
	// 	log.Println(err)
	// }

	// tr := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyURL),
	// }
	// httpClient := http.Client{Transport: tr}

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return nil
	}

	octopusClient, err := NewClient(nil, apiURL, octopusAPIKey, emptyString)
	if err != nil {
		log.Fatal(err)
	}

	return octopusClient
}

func testNewService(t *testing.T, service IService, uriTemplate string, serviceName string) {
	require.NotNil(t, service)
	require.NotNil(t, service.getClient())

	template, err := uritemplates.Parse(uriTemplate)
	require.NoError(t, err)
	require.Equal(t, service.getURITemplate(), template)
	require.Equal(t, service.getName(), serviceName)
}
