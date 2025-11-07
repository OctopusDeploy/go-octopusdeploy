package ephemeralenvironments

import (
	"math"

	v2 "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2/environments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

func GetByID(client newclient.Client, spaceID string, id string) (*EphemeralEnvironment, error) {
	query := &environments.EnvironmentQuery{
		Skip: 0,
		Take: 1,
		Type: []string{"Ephemeral"},
		IDs:  []string{id},
	}

	result, err := newclient.GetByQuery[EphemeralEnvironment](client, v2.Template, spaceID, query)
	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, nil
	}

	return result.Items[0], nil
}

func GetAll(client newclient.Client, spaceID string) (*resources.Resources[*EphemeralEnvironment], error) {
	query := &environments.EnvironmentQuery{
		Skip: 0,
		Take: math.MaxInt32,
		Type: []string{"Ephemeral"},
	}

	return newclient.GetByQuery[EphemeralEnvironment](client, v2.Template, spaceID, query)
}

func GetByPartialName(client newclient.Client, spaceID string, partialName string) (*resources.Resources[*EphemeralEnvironment], error) {
	query := &environments.EnvironmentQuery{
		Skip:        0,
		Take:        math.MaxInt32,
		PartialName: partialName,
		Type:        []string{"Ephemeral"},
	}

	return newclient.GetByQuery[EphemeralEnvironment](client, v2.Template, spaceID, query)
}

func Add(client newclient.Client, spaceID string, projectID string, environmentName string) (*CreateEnvironmentResponse, error) {
	body := &CreateEnvironmentCommand{
		EnvironmentName: environmentName,
		SpaceID:         spaceID,
		ProjectID:       projectID,
	}

	path, err := client.URITemplateCache().Expand(v2.CreateEphemeralEnvironmentTemplate, map[string]any{
		"projectId": projectID,
		"spaceId":   spaceID,
	})

	if err != nil {
		return nil, err
	}

	return newclient.Add[CreateEnvironmentResponse](client, path, spaceID, body)
}

func DeprovisionForProject(client newclient.Client, spaceID string, environmentId string, projectId string) (*DeprovisionEphemeralEnvironmentProjectResponse, error) {
	body := &DeprovisionEphemeralEnvironmentProjectCommand{}

	path, err := client.URITemplateCache().Expand(v2.DeprovisionEphemeralEnvironmentForProjectTemplate, map[string]any{
		"id":        environmentId,
		"projectId": projectId,
		"spaceId":   spaceID,
	})

	if err != nil {
		return nil, err
	}

	return newclient.Add[DeprovisionEphemeralEnvironmentProjectResponse](client, path, spaceID, body)
}

func Deprovision(client newclient.Client, spaceID string, environmentId string) (*DeprovisionEphemeralEnvironmentResponse, error) {
	body := &DeprovisionEphemeralEnvironmentCommand{}

	path, err := client.URITemplateCache().Expand(v2.DeprovisionEphemeralEnvironmentTemplate, map[string]any{
		"id":      environmentId,
		"spaceId": spaceID,
	})

	if err != nil {
		return nil, err
	}

	return newclient.Add[DeprovisionEphemeralEnvironmentResponse](client, path, spaceID, body)
}
