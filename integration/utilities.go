package integration

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"

	uuid "github.com/google/uuid"
)

const emptyString string = ""

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func createRequiredParameterIsEmptyOrNilError(parameter string) error {
	return fmt.Errorf("the required parameter, %s is nil or empty", parameter)
}

func createRandomBoolean() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32() < 0.5
}

func IsEqualLinks(linksA map[string]string, linksB map[string]string) bool {
	return reflect.DeepEqual(linksA, linksB)
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

func getRandomAzureEnvironment() azureEnvironment {
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

func getRandomName() string {
	fullName := fmt.Sprintf("test-id %s", uuid.New())
	fullName = fullName[0:44]
	return fullName
}

func getShortRandomName() string {
	fullName := fmt.Sprintf("test-id %s", uuid.New())
	fullName = fullName[0:19]
	return fullName
}

func getRandomVarName() string {
	return fmt.Sprintf("go-octo-%v", time.Now().Unix())
}

func Bool(v bool) *bool       { return &v }
func Int(v int) *int          { return &v }
func Int64(v int64) *int64    { return &v }
func String(v string) *string { return &v }
