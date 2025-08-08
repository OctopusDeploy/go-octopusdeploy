package observability

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const (
	liveStatusUntenantedTemplate = "/api/spaces/{spaceId}/projects/{projectId}/environments/{environmentId}/untenanted/livestatus"
	liveStatusTenantedTemplate   = "/api/spaces/{spaceId}/projects/{projectId}/environments/{environmentId}/tenants/{tenantId}/livestatus"
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

	expandedUri, err := client.URITemplateCache().Expand(template, pathVars)
	if err != nil {
		return nil, err
	}

	// Add query parameters if needed
	queryParams := map[string]string{}
	if request.SummaryOnly {
		queryParams["summaryOnly"] = "true"
	}

	if len(queryParams) > 0 {
		expandedUri = expandedUri + "?"
		first := true
		for key, value := range queryParams {
			if !first {
				expandedUri = expandedUri + "&"
			}
			expandedUri = expandedUri + key + "=" + value
			first = false
		}
	}

	resp, err := newclient.Get[GetLiveStatusResponse](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}

	return resp, nil
}