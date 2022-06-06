package octopusdeploy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/dghubble/sling"
)

func isNil(i interface{}) bool {
	switch v := i.(type) {
	case *AccountResource:
		return v == nil
	case *ActionTemplate:
		return v == nil
	case *ActionTemplateParameter:
		return v == nil
	case *AmazonWebServicesAccount:
		return v == nil
	case *APIKey:
		return v == nil
	case *Artifact:
		return v == nil
	case *Authentication:
		return v == nil
	case *AzureCloudServiceEndpoint:
		return v == nil
	case *AzureServiceFabricEndpoint:
		return v == nil
	case *AzureServicePrincipalAccount:
		return v == nil
	case *AzureSubscriptionAccount:
		return v == nil
	case *AzureWebAppEndpoint:
		return v == nil
	case *CertificateResource:
		return v == nil
	case *Channel:
		return v == nil
	case *CommunityActionTemplate:
		return v == nil
	case *ConfigurationSection:
		return v == nil
	case *DeploymentProcess:
		return v == nil
	case *DeploymentStep:
		return v == nil
	case *DeploymentTarget:
		return v == nil
	case *Environment:
		return v == nil
	case *EndpointResource:
		return v == nil
	case *FeedResource:
		return v == nil
	case *GoogleCloudPlatformAccount:
		return v == nil
	case *Interruption:
		return v == nil
	case *KubernetesEndpoint:
		return v == nil
	case *LibraryVariableSetUsageEntry:
		return v == nil
	case *LibraryVariableSet:
		return v == nil
	case *Lifecycle:
		return v == nil
	case *MachineConnectionStatus:
		return v == nil
	case *MachinePolicy:
		return v == nil
	case *Package:
		return v == nil
	case *ProjectGroup:
		return v == nil
	case *ProjectTrigger:
		return v == nil
	case *Project:
		return v == nil
	case *Release:
		return v == nil
	case *ReleaseQuery:
		return v == nil
	case *RootResource:
		return v == nil
	case *Runbook:
		return v == nil
	case *ScriptModule:
		return v == nil
	case *Space:
		return v == nil
	case *SSHKeyAccount:
		return v == nil
	case *TagSet:
		return v == nil
	case *Team:
		return v == nil
	case *Tenant:
		return v == nil
	case *TokenAccount:
		return v == nil
	case *User:
		return v == nil
	case *UsernamePasswordAccount:
		return v == nil
	case *WorkerPoolResource:
		return v == nil
	default:
		return v == nil
	}
}

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func isAPIKey(apiKey string) bool {
	if len(apiKey) < 5 {
		return false
	}

	var expression = regexp.MustCompile(`^(API-)([A-Z\d])+$`)
	return expression.MatchString(apiKey)
}

func PrettyJSON(data interface{}) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	encoder.Encode(data)
	return buffer.String()
}

func getDefaultClient() *sling.Sling {
	octopusURL := os.Getenv(clientURLEnvironmentVariable)
	octopusAPIKey := os.Getenv(clientAPIKeyEnvironmentVariable)

	// NOTE: You can direct traffic through a proxy trace like Fiddler
	// Everywhere by preconfiguring the client to route traffic through a
	// proxy.

	// proxyStr := "http://127.0.0.1:5555"
	// proxyURL, _ := url.Parse(proxyStr)

	// tr := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyURL),
	// }
	// httpClient := http.Client{Transport: tr}

	return sling.New().Client(nil).Base(octopusURL).Set(clientAPIKeyHTTPHeader, octopusAPIKey)
}

func trimTemplate(uri string) string {
	return strings.Split(uri, "{")[0]
}

func createBuiltInTeamsCannotDeleteError() error {
	return fmt.Errorf("the built-in teams cannot be deleted")
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

func createClientInitializationError(methodName string) error {
	return fmt.Errorf("%s: unable to initialize internal client", methodName)
}

func createRequiredParameterIsEmptyOrNilError(parameter string) error {
	return fmt.Errorf("the required parameter, %s is nil or empty", parameter)
}

func createResourceNotFoundError(name string, identifier string, value string) error {
	return fmt.Errorf("the service, %s could not find the %s (%s)", name, identifier, value)
}

func createValidationFailureError(methodName string, err error) error {
	return fmt.Errorf("validation failure in %s; %v", methodName, err)
}

func Bool(v bool) *bool       { return &v }
func Int(v int) *int          { return &v }
func Int64(v int64) *int64    { return &v }
func String(v string) *string { return &v }
