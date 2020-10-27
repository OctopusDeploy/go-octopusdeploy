package integration

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"

	uuid "github.com/google/uuid"
)

const (
	emptyString      string = ""
	whitespaceString string = " "
)

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func createInvalidParameterError(methodName string, ParameterName string) error {
	return fmt.Errorf("%s: the input parameter (%s) is invalid", methodName, ParameterName)
}

func createInvalidClientStateError(ServiceName string) error {
	return fmt.Errorf("%s: the state of the internal client is invalid", ServiceName)
}

func createInvalidPathError(ServiceName string) error {
	return fmt.Errorf("%s: the internal path is not set", ServiceName)
}

func createItemNotFoundError(ServiceName string, methodName string, name string) error {
	return fmt.Errorf("%s: the item (%s) via %s was not found", ServiceName, name, methodName)
}

func createRequiredParameterIsEmptyOrNilError(parameter string) error {
	return fmt.Errorf("the required parameter, %s is nil or empty", parameter)
}

func createClientInitializationError(methodName string) error {
	return fmt.Errorf("%s: unable to initialize internal client", methodName)
}

func createRandomBoolean() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32() < 0.5
}

func createResourceNotFoundError(name string, identifier string, value string) error {
	return fmt.Errorf("the service, %s could not find the %s (%s)", name, identifier, value)
}

func createValidationFailureError(methodName string, err error) error {
	return fmt.Errorf("validation failure in %s; %v", methodName, err)
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

func getRandomDuration() time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", rand.Int63n(1000)))
	return duration
}

func Bool(v bool) *bool       { return &v }
func Int(v int) *int          { return &v }
func Int64(v int64) *int64    { return &v }
func String(v string) *string { return &v }
