package e2e

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

const (
	TestVariableOctopusHost   string = "TEST_OCTOPUS_HOST"
	TestVariableOctopusApiKey string = "TEST_OCTOPUS_API_KEY"
)

func getOctopusClient() *client.Client {
	host := os.Getenv(TestVariableOctopusHost)
	apiKey := os.Getenv(TestVariableOctopusApiKey)

	if internal.IsEmpty(host) || internal.IsEmpty(apiKey) {
		log.Fatal("Please make sure to set the env variables '" + TestVariableOctopusHost + "' and '" + TestVariableOctopusApiKey + "' before running this test")
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

	octopusClient, err := client.NewClient(nil, apiURL, apiKey, "")
	if err != nil {
		log.Fatal(err)
	}

	return octopusClient
}
