package services

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestIsAPIKey(t *testing.T) {
	testCases := []struct {
		name    string
		test    string
		isValid bool
	}{
		{"Empty", emptyString, false},
		{"EmptySpace", whitespaceString, false},
		{"StartWithAPI", "API-", false},
		{"Invalid", "API-?OBYCAMCZ7WWBKSTMXT66FCUDPS", false},
		{"Invalid", "API-OBYC👍AMCZ7WWBKSTMXT66FCUDPS", false},
		{"Invalid", "API-***************************", false},
		{"Valid", "API-EOAYCAFCZ7WWBKSTMVT66FCUDPS", true},
		{"Valid", "API-EOBYCAMCZ7WWBKSTMVT66FCUDPR", true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testResult := isAPIKey(tc.test)
			assert.Equal(t, tc.isValid, testResult)
		})
	}
}

func createRandomBoolean() bool {
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

func getShortRandomName() string {
	fullName := fmt.Sprintf("test-id %s", uuid.New())
	fullName = fullName[0:19]
	return fullName
}

func getRandomVarName() string {
	return fmt.Sprintf("go-octo-%v", time.Now().Unix())
}

func getRandomDuration(mininum time.Duration) time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", rand.Int63n(1000)))
	duration += mininum
	return duration
}

func generateSensitiveValue() *SensitiveValue {
	sensitiveValue := NewSensitiveValue(getRandomName())
	return sensitiveValue
}

func IsEqualLinks(linksA map[string]string, linksB map[string]string) bool {
	return reflect.DeepEqual(linksA, linksB)
}
