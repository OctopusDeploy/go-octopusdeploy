package e2e

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
)

func getOctopusClient() *client.Client {
	// host := os.Getenv(constants.EnvironmentVariableOctopusHost)
	// apiKey := os.Getenv(constants.EnvironmentVariableOctopusApiKey)
	spaceId := os.Getenv(constants.EnvironmentVariableOctopusSpace)

	host := "http://localhost:8066"
	apiKey := "API-4KM7F0OLXQBRBYLYHTWQSJYWVSFVB"

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
