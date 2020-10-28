package integration

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func getOctopusClient() *octopusdeploy.Client {
	octopusURL := os.Getenv("OCTOPUS_URL")
	apiKey := os.Getenv("OCTOPUS_APIKEY")

	if isEmpty(octopusURL) || isEmpty(apiKey) {
		log.Fatal("Please make sure to set the env variables 'OCTOPUS_URL' and 'OCTOPUS_APIKEY' before running this test")
	}

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return nil
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

	octopusClient, err := octopusdeploy.NewClient(nil, apiURL, apiKey, emptyString)
	if err != nil {
		log.Fatal(err)
	}

	return octopusClient
}

func generateSensitiveValue() *octopusdeploy.SensitiveValue {
	return octopusdeploy.NewSensitiveValue(getRandomName())
}
