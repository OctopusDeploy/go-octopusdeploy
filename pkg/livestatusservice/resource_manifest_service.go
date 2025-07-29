package livestatusservice

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

const (
	resourceManifestUntenantedTemplate = "/api/{spaceId}/projects/{projectId}/environments/{environmentId}/untenanted/machines/{machineId}/resources/{desiredOrKubernetesMonitoredResourceId}/manifest"
	resourceManifestTenantedTemplate   = "/api/{spaceId}/projects/{projectId}/environments/{environmentId}/tenants/{tenantId}/machines/{machineId}/resources/{desiredOrKubernetesMonitoredResourceId}/manifest"
)

// ResourceManifestService handles operations related to resource manifests
type ResourceManifestService struct {
	services.Service
}

// NewResourceManifestService creates a new ResourceManifestService
func NewResourceManifestService(sling *sling.Sling, uriTemplate string) *ResourceManifestService {
	return &ResourceManifestService{
		Service: services.NewService("ResourceManifestService", sling, uriTemplate),
	}
}

// GetResourceManifest retrieves a resource manifest for untenanted resources
func (s *ResourceManifestService) GetResourceManifest(request *GetResourceManifestRequest) (*GetResourceManifestResponse, error) {
	if request == nil {
		return nil, internal.CreateInvalidParameterError("GetResourceManifest", "request")
	}

	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	var templateStr string
	var pathVars map[string]interface{}

	if request.IsUntenanted() {
		templateStr = resourceManifestUntenantedTemplate
		pathVars = map[string]interface{}{
			"spaceId":                                request.SpaceID,
			"projectId":                              request.ProjectID,
			"environmentId":                          request.EnvironmentID,
			"machineId":                              request.MachineID,
			"desiredOrKubernetesMonitoredResourceId": request.DesiredOrKubernetesMonitoredResourceID,
		}
	} else {
		templateStr = resourceManifestTenantedTemplate
		pathVars = map[string]interface{}{
			"spaceId":                                request.SpaceID,
			"projectId":                              request.ProjectID,
			"environmentId":                          request.EnvironmentID,
			"tenantId":                               request.TenantID,
			"machineId":                              request.MachineID,
			"desiredOrKubernetesMonitoredResourceId": request.DesiredOrKubernetesMonitoredResourceID,
		}
	}

	template, err := uritemplates.Parse(templateStr)
	if err != nil {
		return nil, err
	}

	path, err := template.Expand(pathVars)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(GetResourceManifestResponse), path)
	if err != nil {
		return nil, err
	}

	return resp.(*GetResourceManifestResponse), nil
}

// New

// GetResourceManifestWithClient retrieves a resource manifest using the new client implementation
func GetResourceManifestWithClient(client newclient.Client, request *GetResourceManifestRequest) (*GetResourceManifestResponse, error) {
	if request == nil {
		return nil, internal.CreateInvalidParameterError("GetResourceManifest", "request")
	}

	spaceID, err := internal.GetSpaceID(request.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	var templateStr string
	var pathVars map[string]interface{}

	if request.IsUntenanted() {
		templateStr = resourceManifestUntenantedTemplate
		pathVars = map[string]interface{}{
			"spaceId":                                spaceID,
			"projectId":                              request.ProjectID,
			"environmentId":                          request.EnvironmentID,
			"machineId":                              request.MachineID,
			"desiredOrKubernetesMonitoredResourceId": request.DesiredOrKubernetesMonitoredResourceID,
		}
	} else {
		templateStr = resourceManifestTenantedTemplate
		pathVars = map[string]interface{}{
			"spaceId":                                spaceID,
			"projectId":                              request.ProjectID,
			"environmentId":                          request.EnvironmentID,
			"tenantId":                               request.TenantID,
			"machineId":                              request.MachineID,
			"desiredOrKubernetesMonitoredResourceId": request.DesiredOrKubernetesMonitoredResourceID,
		}
	}

	expandedUri, err := client.URITemplateCache().Expand(templateStr, pathVars)
	if err != nil {
		return nil, err
	}

	resp, err := newclient.Get[GetResourceManifestResponse](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
