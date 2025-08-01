package observability

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const (
	resourceManifestUntenantedTemplate = "/api/{spaceId}/projects/{projectId}/environments/{environmentId}/untenanted/machines/{machineId}/resources/{desiredOrKubernetesMonitoredResourceId}/manifest"
	resourceManifestTenantedTemplate   = "/api/{spaceId}/projects/{projectId}/environments/{environmentId}/tenants/{tenantId}/machines/{machineId}/resources/{desiredOrKubernetesMonitoredResourceId}/manifest"
)

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

	if !request.IsTenanted() {
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
