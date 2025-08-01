package livestatusservice

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const (
	resourceEventsTemplate = "/api/{spaceId}/observability/events/sessions/{sessionId}"
)

// GetResourceEventsWithClient retrieves resource events using the new client implementation
func GetResourceEventsWithClient(client newclient.Client, request *GetResourceEventsRequest) (*GetResourceEventsResponse, error) {
	if request == nil {
		return nil, internal.CreateInvalidParameterError("GetResourceEvents", "request")
	}

	spaceID, err := internal.GetSpaceID(request.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	pathVars := map[string]interface{}{
		"spaceId":   spaceID,
		"sessionId": string(request.SessionID),
	}

	expandedUri, err := client.URITemplateCache().Expand(resourceEventsTemplate, pathVars)
	if err != nil {
		return nil, err
	}

	resp, err := newclient.Get[GetResourceEventsResponse](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}

	return resp, nil
}