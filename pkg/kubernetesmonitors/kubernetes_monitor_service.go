package kubernetesmonitors

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const template = "/api/{spaceId}/observability/kubernetes-monitors{/id}"

// Register registers the given Kubernetes Monitor parameters with the Octopus Deploy server.
func Register(
	client newclient.Client, command *RegisterKubernetesMonitorCommand,
) (*RegisterKubernetesMonitorResponse, error) {
	if command == nil {
		return nil, internal.CreateInvalidParameterError("Register", "command")
	}

	spaceID, err := internal.GetSpaceID(command.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{
		"spaceId": spaceID,
	})
	if err != nil {
		return nil, err
	}

	res, err := newclient.Post[RegisterKubernetesMonitorResponse](client.HttpSession(), path, command)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID returns the Kubernetes Monitor that matches the input ID. If one cannot be found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, ID string) (*GetKubernetesMonitorResponse, error) {
	return newclient.GetByID[GetKubernetesMonitorResponse](client, template, spaceID, ID)
}

// DeleteByID deletes a Kubernetes Monitor based on the provided ID.
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}
