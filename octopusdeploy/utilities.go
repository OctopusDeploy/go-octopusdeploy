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

func isNil(i IResource) bool {
	var ret bool

	switch i.(type) {
	case *ActionTemplate:
		v := i.(*ActionTemplate)
		ret = v == nil
	case *ActionTemplateParameter:
		v := i.(*ActionTemplateParameter)
		ret = v == nil
	case *APIKey:
		v := i.(*APIKey)
		ret = v == nil
	case *Artifact:
		v := i.(*Artifact)
		ret = v == nil
	case *Authentication:
		v := i.(*Authentication)
		ret = v == nil
	case *Certificate:
		v := i.(*Certificate)
		ret = v == nil
	case *Channel:
		v := i.(*Channel)
		ret = v == nil
	case *CommunityActionTemplate:
		v := i.(*CommunityActionTemplate)
		ret = v == nil
	case *ConfigurationSection:
		v := i.(*ConfigurationSection)
		ret = v == nil
	case *DeploymentProcess:
		v := i.(*DeploymentProcess)
		ret = v == nil
	case *DeploymentStep:
		v := i.(*DeploymentStep)
		ret = v == nil
	case *DeploymentTarget:
		v := i.(*DeploymentTarget)
		ret = v == nil
	case *Environment:
		v := i.(*Environment)
		ret = v == nil
	case *Feed:
		v := i.(*Feed)
		ret = v == nil
	case *Interruption:
		v := i.(*Interruption)
		ret = v == nil
	case *KubernetesEndpoint:
		v := i.(*KubernetesEndpoint)
		ret = v == nil
	case *LibraryVariableSetUsageEntry:
		v := i.(*LibraryVariableSetUsageEntry)
		ret = v == nil
	case *LibraryVariableSet:
		v := i.(*LibraryVariableSet)
		ret = v == nil
	case *Lifecycle:
		v := i.(*Lifecycle)
		ret = v == nil
	case *MachineConnectionStatus:
		v := i.(*MachineConnectionStatus)
		ret = v == nil
	case *MachinePolicy:
		v := i.(*MachinePolicy)
		ret = v == nil
	case *ProjectGroup:
		v := i.(*ProjectGroup)
		ret = v == nil
	case *ProjectTrigger:
		v := i.(*ProjectTrigger)
		ret = v == nil
	case *Project:
		v := i.(*Project)
		ret = v == nil
	case *Release:
		v := i.(*Release)
		ret = v == nil
	case *RootResource:
		v := i.(*RootResource)
		ret = v == nil
	case *Runbook:
		v := i.(*Runbook)
		ret = v == nil
	case *Space:
		v := i.(*Space)
		ret = v == nil
	case *TagSet:
		v := i.(*TagSet)
		ret = v == nil
	case *Team:
		v := i.(*Team)
		ret = v == nil
	case *Tenant:
		v := i.(*Tenant)
		ret = v == nil
	case *User:
		v := i.(*User)
		ret = v == nil
	}

	return ret
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

func createInvalidParameterError(methodName string, parameterName string) error {
	return fmt.Errorf("%s: the input parameter (%s) is invalid", methodName, parameterName)
}

func createInvalidClientStateError(serviceName string) error {
	return fmt.Errorf("%s: the state of the internal client is invalid", serviceName)
}

func createInvalidPathError(serviceName string) error {
	return fmt.Errorf("%s: the internal path is not set", serviceName)
}

func createItemNotFoundError(serviceName string, methodName string, name string) error {
	return fmt.Errorf("%s: the item (%s) via %s was not found", serviceName, name, methodName)
}

func createClientInitializationError(methodName string) error {
	return fmt.Errorf("%s: unable to initialize internal client", methodName)
}

func createResourceNotFoundError(name string, identifier string, value string) error {
	return fmt.Errorf("the service, %s could not find the %s (%s)", name, identifier, value)
}

func createValidationFailureError(methodName string, err error) error {
	return fmt.Errorf("validation failure in %s; %v", methodName, err)
}
