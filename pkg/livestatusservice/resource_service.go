package livestatusservice

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const (
	getResourceUntenantedTemplate = "/api/{spaceId}/projects/{projectId}/environments/{environmentId}/untenanted/machines/{machineId}/resources/{desiredOrKubernetesMonitoredResourceId}"
	getResourceTenantedTemplate   = "/api/{spaceId}/projects/{projectId}/environments/{environmentId}/tenants/{tenantId}/machines/{machineId}/resources/{desiredOrKubernetesMonitoredResourceId}"
)

// GetResourceWithClient retrieves detailed summary of a live kubernetes resource using the new client implementation
func GetResourceWithClient(client newclient.Client, request *GetResourceRequest) (*GetResourceResponse, error) {
	if request == nil {
		return nil, internal.CreateInvalidParameterError("GetResource", "request")
	}

	spaceID, err := internal.GetSpaceID(request.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	var templateStr string
	var pathVars map[string]interface{}

	if !request.IsTenanted() {
		templateStr = getResourceUntenantedTemplate
		pathVars = map[string]interface{}{
			"spaceId":                                spaceID,
			"projectId":                              request.ProjectID,
			"environmentId":                          request.EnvironmentID,
			"machineId":                              request.MachineID,
			"desiredOrKubernetesMonitoredResourceId": request.DesiredOrKubernetesMonitoredResourceID,
		}
	} else {
		templateStr = getResourceTenantedTemplate
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

	resp, err := newclient.Get[GetResourceResponse](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}

	return resp, nil
}