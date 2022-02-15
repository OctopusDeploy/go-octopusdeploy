package internal

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func CreateRandomBoolean() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32() < 0.5
}

type azureEnvironment struct {
	AuthenticationEndpoint  string
	Name                    string
	DisplayName             string
	GraphEndpoint           string
	ManagementEndpoint      string
	ResourceManagerEndpoint string
	StorageEndpointSuffix   string
}

func GetRandomAzureEnvironment() azureEnvironment {
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(5) {
	case 0:
		return azureEnvironment{
			Name:                    "AzureCloud",
			DisplayName:             "Azure Global Cloud (default)",
			AuthenticationEndpoint:  "https://login.microsoftonline.com/",
			ResourceManagerEndpoint: "https://management.azure.com/",
			GraphEndpoint:           "https://graph.windows.net/",
			ManagementEndpoint:      "https://management.core.windows.net/",
			StorageEndpointSuffix:   "core.windows.net",
		}
	case 1:
		return azureEnvironment{
			Name:                    "AzureChinaCloud",
			DisplayName:             "Azure China Cloud",
			AuthenticationEndpoint:  "https://login.chinacloudapi.cn/",
			ResourceManagerEndpoint: "https://management.chinacloudapi.cn/",
			GraphEndpoint:           "https://graph.chinacloudapi.cn/",
			ManagementEndpoint:      "https://management.core.chinacloudapi.cn/",
			StorageEndpointSuffix:   "core.chinacloudapi.cn",
		}
	case 2:
		return azureEnvironment{
			Name:                    "AzureUSGovernment",
			DisplayName:             "Azure US Government",
			AuthenticationEndpoint:  "https://login.microsoftonline.us/",
			ResourceManagerEndpoint: "https://management.usgovcloudapi.net/",
			GraphEndpoint:           "https://graph.windows.net/",
			ManagementEndpoint:      "https://management.core.usgovcloudapi.net",
			StorageEndpointSuffix:   "core.usgovcloudapi.net",
		}
	case 3:
		return azureEnvironment{
			Name:                    "AzureGermanCloud",
			DisplayName:             "Azure German Cloud",
			AuthenticationEndpoint:  "https://login.microsoftonline.de/",
			ResourceManagerEndpoint: "https://management.microsoftazure.de/",
			GraphEndpoint:           "https://graph.cloudapi.de/",
			ManagementEndpoint:      "https://management.core.cloudapi.de",
			StorageEndpointSuffix:   "core.cloudapi.de",
		}
	}

	return azureEnvironment{}
}

func GetShortRandomName() string {
	fullName := fmt.Sprintf("test-id %s", uuid.New())
	fullName = fullName[0:19]
	return fullName
}

func GetRandomVarName() string {
	return fmt.Sprintf("go-octo-%v", time.Now().Unix())
}

func GetRandomDuration(mininum time.Duration) time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", rand.Int63n(1000)))
	duration += mininum
	return duration
}