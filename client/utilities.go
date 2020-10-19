package client

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

func isNil(i model.IResource) bool {
	var ret bool

	switch i.(type) {
	case *model.ActionTemplate:
		v := i.(*model.ActionTemplate)
		ret = v == nil
	case *model.ActionTemplateParameter:
		v := i.(*model.ActionTemplateParameter)
		ret = v == nil
	case *model.APIKey:
		v := i.(*model.APIKey)
		ret = v == nil
	case *model.Artifact:
		v := i.(*model.Artifact)
		ret = v == nil
	case *model.Authentication:
		v := i.(*model.Authentication)
		ret = v == nil
	case *model.Certificate:
		v := i.(*model.Certificate)
		ret = v == nil
	case *model.Channel:
		v := i.(*model.Channel)
		ret = v == nil
	case *model.CommunityActionTemplate:
		v := i.(*model.CommunityActionTemplate)
		ret = v == nil
	case *model.ConfigurationSection:
		v := i.(*model.ConfigurationSection)
		ret = v == nil
	case *model.DeploymentProcess:
		v := i.(*model.DeploymentProcess)
		ret = v == nil
	case *model.DeploymentStep:
		v := i.(*model.DeploymentStep)
		ret = v == nil
	case *model.DeploymentTarget:
		v := i.(*model.DeploymentTarget)
		ret = v == nil
	case *model.Environment:
		v := i.(*model.Environment)
		ret = v == nil
	case *model.Feed:
		v := i.(*model.Feed)
		ret = v == nil
	case *model.Interruption:
		v := i.(*model.Interruption)
		ret = v == nil
	case *model.KubernetesEndpoint:
		v := i.(*model.KubernetesEndpoint)
		ret = v == nil
	case *model.LibraryVariableSetUsageEntry:
		v := i.(*model.LibraryVariableSetUsageEntry)
		ret = v == nil
	case *model.LibraryVariableSet:
		v := i.(*model.LibraryVariableSet)
		ret = v == nil
	case *model.Lifecycle:
		v := i.(*model.Lifecycle)
		ret = v == nil
	case *model.MachineConnectionStatus:
		v := i.(*model.MachineConnectionStatus)
		ret = v == nil
	case *model.MachinePolicy:
		v := i.(*model.MachinePolicy)
		ret = v == nil
	case *model.ProjectGroup:
		v := i.(*model.ProjectGroup)
		ret = v == nil
	case *model.ProjectTrigger:
		v := i.(*model.ProjectTrigger)
		ret = v == nil
	case *model.Project:
		v := i.(*model.Project)
		ret = v == nil
	case *model.Release:
		v := i.(*model.Release)
		ret = v == nil
	case *model.RootResource:
		v := i.(*model.RootResource)
		ret = v == nil
	case *model.Runbook:
		v := i.(*model.Runbook)
		ret = v == nil
	case *model.Space:
		v := i.(*model.Space)
		ret = v == nil
	case *model.TagSet:
		v := i.(*model.TagSet)
		ret = v == nil
	case *model.Team:
		v := i.(*model.Team)
		ret = v == nil
	case *model.Tenant:
		v := i.(*model.Tenant)
		ret = v == nil
	case *model.User:
		v := i.(*model.User)
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
