package integration

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/model"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	uuid "github.com/google/uuid"
)

var (
	octopusURL    string = os.Getenv("OCTOPUS_URL")
	octopusAPIKey string = os.Getenv("OCTOPUS_APIKEY")
	octopusClient *client.Client
)

func initTest() *client.Client {

	if isEmpty(octopusURL) || isEmpty(octopusAPIKey) {
		log.Fatal("Please make sure to set the env variables 'OCTOPUS_URL' and 'OCTOPUS_APIKEY' before running this test")
	}

	octopusClient, err := client.NewClient(&http.Client{}, octopusURL, octopusAPIKey, nil)

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

	httpClient := http.Client{}
	octopusClient, err = client.NewClient(&httpClient, octopusURL, octopusAPIKey, nil)
	if err != nil {
		log.Fatal(err)
	}

	return octopusClient
}

func getRandomName() string {
	fullName := fmt.Sprintf("go-octopusdeploy %s", uuid.New())
	fullName = fullName[0:49] //Some names in Octopus have a max limit of 50 characters (such as Environment Name)
	return fullName
}

func getRandomVarName() string {
	return fmt.Sprintf("go-octo-%v", time.Now().Unix())
}

func generateSensitiveValue() model.SensitiveValue {
	sensitiveValue, err := model.NewSensitiveValue(getRandomName())

	if err != nil {
		return model.SensitiveValue{}
	}

	return *sensitiveValue
}
