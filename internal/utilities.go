package internal

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	Empty             = ""
	Whitespace string = " "
	Tab               = "\t"
)

func IsNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

// func isNil(i interface{}) bool {
// 	var ret bool

// 	switch i.(type) {
// 	case *accounts.AccountResource:
// 		v := i.(*accounts.AccountResource)
// 		ret = v == nil
// 	case *ActionTemplate:
// 		v := i.(*ActionTemplate)
// 		ret = v == nil
// 	case *ActionTemplateParameter:
// 		v := i.(*ActionTemplateParameter)
// 		ret = v == nil
// 	case *accounts.AmazonWebServicesAccount:
// 		v := i.(*accounts.AmazonWebServicesAccount)
// 		ret = v == nil
// 	case *APIKey:
// 		v := i.(*APIKey)
// 		ret = v == nil
// 	case *Artifact:
// 		v := i.(*Artifact)
// 		ret = v == nil
// 	case *Authentication:
// 		v := i.(*Authentication)
// 		ret = v == nil
// 	case *AzureCloudServiceEndpoint:
// 		v := i.(*AzureCloudServiceEndpoint)
// 		ret = v == nil
// 	case *AzureServiceFabricEndpoint:
// 		v := i.(*AzureServiceFabricEndpoint)
// 		ret = v == nil
// 	case *AzureServicePrincipalAccount:
// 		v := i.(*AzureServicePrincipalAccount)
// 		ret = v == nil
// 	case *accounts.AzureSubscriptionAccount:
// 		v := i.(*accounts.AzureSubscriptionAccount)
// 		ret = v == nil
// 	case *AzureWebAppEndpoint:
// 		v := i.(*AzureWebAppEndpoint)
// 		ret = v == nil
// 	case *CertificateResource:
// 		v := i.(*CertificateResource)
// 		ret = v == nil
// 	case *Channel:
// 		v := i.(*Channel)
// 		ret = v == nil
// 	case *CommunityActionTemplate:
// 		v := i.(*CommunityActionTemplate)
// 		ret = v == nil
// 	case *ConfigurationSection:
// 		v := i.(*ConfigurationSection)
// 		ret = v == nil
// 	case *DeploymentProcess:
// 		v := i.(*DeploymentProcess)
// 		ret = v == nil
// 	case *DeploymentStep:
// 		v := i.(*DeploymentStep)
// 		ret = v == nil
// 	case *DeploymentTarget:
// 		v := i.(*DeploymentTarget)
// 		ret = v == nil
// 	case *Environment:
// 		v := i.(*Environment)
// 		ret = v == nil
// 	case *Feed:
// 		v := i.(*Feed)
// 		ret = v == nil
// 	case *GoogleCloudPlatformAccount:
// 		v := i.(*GoogleCloudPlatformAccount)
// 		ret = v == nil
// 	case *Interruption:
// 		v := i.(*Interruption)
// 		ret = v == nil
// 	case *KubernetesEndpoint:
// 		v := i.(*KubernetesEndpoint)
// 		ret = v == nil
// 	case *LibraryVariableSetUsageEntry:
// 		v := i.(*LibraryVariableSetUsageEntry)
// 		ret = v == nil
// 	case *LibraryVariableSet:
// 		v := i.(*LibraryVariableSet)
// 		ret = v == nil
// 	case *Lifecycle:
// 		v := i.(*Lifecycle)
// 		ret = v == nil
// 	case *MachineConnectionStatus:
// 		v := i.(*MachineConnectionStatus)
// 		ret = v == nil
// 	case *MachinePolicy:
// 		v := i.(*MachinePolicy)
// 		ret = v == nil
// 	case *Package:
// 		v := i.(*Package)
// 		ret = v == nil
// 	case *ProjectGroup:
// 		v := i.(*ProjectGroup)
// 		ret = v == nil
// 	case *ProjectTrigger:
// 		v := i.(*ProjectTrigger)
// 		ret = v == nil
// 	case *Project:
// 		v := i.(*Project)
// 		ret = v == nil
// 	case *Release:
// 		v := i.(*Release)
// 		ret = v == nil
// 	case *ReleaseQuery:
// 		v := i.(*ReleaseQuery)
// 		ret = v == nil
// 	case *RootResource:
// 		v := i.(*RootResource)
// 		ret = v == nil
// 	case *Runbook:
// 		v := i.(*Runbook)
// 		ret = v == nil
// 	case *ScriptModule:
// 		v := i.(*ScriptModule)
// 		ret = v == nil
// 	case *Space:
// 		v := i.(*Space)
// 		ret = v == nil
// 	case *TagSet:
// 		v := i.(*TagSet)
// 		ret = v == nil
// 	case *access_management.Team:
// 		v := i.(*access_management.Team)
// 		ret = v == nil
// 	case *Tenant:
// 		v := i.(*Tenant)
// 		ret = v == nil
// 	case *User:
// 		v := i.(*User)
// 		ret = v == nil
// 	}

// 	return ret
// }

func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func CreateInvalidParameterError(methodName string, ParameterName string) error {
	return fmt.Errorf("%s: the input parameter (%s) is invalid", methodName, ParameterName)
}

func CreateRequiredParameterIsEmptyOrNilError(parameter string) error {
	return fmt.Errorf("the required parameter, %s is nil or Empty", parameter)
}

func CreateBuiltInTeamsCannotDeleteError() error {
	return fmt.Errorf("The built-in teams cannot be deleted.")
}

func CreateInvalidClientStateError(ServiceName string) error {
	return fmt.Errorf("%s: the state of the internal Client is invalid", ServiceName)
}

func CreateInvalidPathError(ServiceName string) error {
	return fmt.Errorf("%s: the internal path is not set", ServiceName)
}

func CreateItemNotFoundError(ServiceName string, methodName string, name string) error {
	return fmt.Errorf("%s: the item (%s) via %s was not found", ServiceName, name, methodName)
}

func CreateClientInitializationError(methodName string) error {
	return fmt.Errorf("%s: unable to initialize internal Client", methodName)
}

func CreateResourceNotFoundError(name string, identifier string, value string) error {
	return fmt.Errorf("the service, %s could not find the %s (%s)", name, identifier, value)
}

func CreateValidationFailureError(methodName string, err error) error {
	return fmt.Errorf("validation failure in %s; %v", methodName, err)
}

func Bool(v bool) *bool       { return &v }
func Int(v int) *int          { return &v }
func Int64(v int64) *int64    { return &v }
func String(v string) *string { return &v }