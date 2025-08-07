package observability

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const (
	containerLogsTemplate = "/api/{spaceId}/observability/logs/sessions{/sessionId}"
)

// GetContainerLogsWithClient retrieves container logs using the new client implementation
func GetContainerLogsWithClient(client newclient.Client, request *GetContainerLogsRequest) (*GetContainerLogsResponse, error) {
	if request == nil {
		return nil, internal.CreateInvalidParameterError("GetContainerLogs", "request")
	}

	spaceID, err := internal.GetSpaceID(request.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	pathVars := map[string]interface{}{
		"spaceId":   spaceID,
		"sessionId": string(request.SessionID),
	}

	expandedUri, err := client.URITemplateCache().Expand(containerLogsTemplate, pathVars)
	if err != nil {
		return nil, err
	}

	resp, err := newclient.Get[GetContainerLogsResponse](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// BeginContainerLogsSessionWithClient begins a container logs session using the new client implementation
func BeginContainerLogsSessionWithClient(client newclient.Client, command *BeginContainerLogsSessionCommand) (*BeginContainerLogsSessionResponse, error) {
	if command == nil {
		return nil, internal.CreateInvalidParameterError("BeginContainerLogsSession", "command")
	}

	spaceID, err := internal.GetSpaceID(command.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	pathVars := map[string]interface{}{
		"spaceId": spaceID,
	}

	expandedUri, err := client.URITemplateCache().Expand(containerLogsTemplate, pathVars)
	if err != nil {
		return nil, err
	}

	resp, err := newclient.Post[BeginContainerLogsSessionResponse](client.HttpSession(), expandedUri, command)
	if err != nil {
		return nil, err
	}

	return resp, nil
}