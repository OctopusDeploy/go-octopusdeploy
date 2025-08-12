package observability

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const (
	liveStatusUntenantedTemplate = "/api/spaces/{spaceId}/projects/{projectId}/environments/{environmentId}/untenanted/livestatus{?summaryOnly}"
	liveStatusTenantedTemplate   = "/api/spaces/{spaceId}/projects/{projectId}/environments/{environmentId}/tenants/{tenantId}/livestatus{?summaryOnly}"
)

// GetLiveStatusWithClient retrieves live status using the new client implementation
func GetLiveStatusWithClient(client newclient.Client, request *GetLiveStatusRequest) (*GetLiveStatusResponse, error) {
	if request == nil {
		return nil, internal.CreateInvalidParameterError("GetLiveStatus", "request")
	}

	spaceID, err := internal.GetSpaceID(request.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	pathVars := map[string]interface{}{
		"spaceId":       spaceID,
		"projectId":     request.ProjectID,
		"environmentId": request.EnvironmentID,
	}

	var template string
	if request.TenantID != nil {
		template = liveStatusTenantedTemplate
		pathVars["tenantId"] = *request.TenantID
	} else {
		template = liveStatusUntenantedTemplate
	}

	if request.SummaryOnly {
		pathVars["summaryOnly"] = "true"
	}

	expandedUri, err := client.URITemplateCache().Expand(template, pathVars)
	if err != nil {
		return nil, err
	}

	resp, err := newclient.Get[GetLiveStatusResponse](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
