package ephemeralenvironments

import (
	"math"

	v2 "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2/environments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

func GetAll(client newclient.Client, spaceID string) (*resources.Resources[*environments.Environment], error) {
	query := &environments.EnvironmentQuery{
		Skip: 0,
		Take: math.MaxInt32,
		Type: []string{"Ephemeral"},
	}

	return newclient.GetByQuery[environments.Environment](client, v2.Template, spaceID, query)
}

type CreateEnvironmentResponse struct {
	Id string `json:"Id"`
}

type CreateEnvironmentCommand struct {
	EnvironmentName string `json:"EnvironmentName"`
	SpaceID         string `uri:"spaceId"`
	ProjectID       string `uri:"projectId"`
}

func Create(client newclient.Client, spaceID string, projectID string, environmentName string) (*CreateEnvironmentResponse, error) {
	body := &CreateEnvironmentCommand{
		EnvironmentName: environmentName,
		SpaceID:         spaceID,
		ProjectID:       projectID,
	}

	path, err := client.URITemplateCache().Expand(v2.CreateTemplate, map[string]any{
		"projectId": projectID,
		"spaceId":   spaceID,
	})

	if err != nil {
		return nil, err
	}

	return newclient.Add[CreateEnvironmentResponse](client, path, spaceID, body)
}
