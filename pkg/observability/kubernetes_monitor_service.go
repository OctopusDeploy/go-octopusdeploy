package observability

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const (
	kubernetesMonitorsTemplate = "/api/{spaceId}/observability/kubernetes-monitors"
)

// RegisterKubernetesMonitorWithClient registers a new Kubernetes Monitor using the new client implementation
func RegisterKubernetesMonitorWithClient(client newclient.Client, command *RegisterKubernetesMonitorCommand) (*RegisterKubernetesMonitorResponse, error) {
	if command == nil {
		return nil, internal.CreateInvalidParameterError("RegisterKubernetesMonitor", "command")
	}

	spaceID, err := internal.GetSpaceID(command.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	pathVars := map[string]interface{}{
		"spaceId": spaceID,
	}

	expandedUri, err := client.URITemplateCache().Expand(kubernetesMonitorsTemplate, pathVars)
	if err != nil {
		return nil, err
	}

	resp, err := newclient.Post[RegisterKubernetesMonitorResponse](client.HttpSession(), expandedUri, command)
	if err != nil {
		return nil, err
	}

	return resp, nil
}