package ephemeralenvironments

import (
	"math"

	v2 "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2/environments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

func GetAll(client newclient.Client, spaceID string) (*resources.Resources[*EphemeralEnvironment], error) {
	query := &environments.EnvironmentQuery{
		Skip: 0,
		Take: math.MaxInt32,
		Type: []string{"Ephemeral"},
	}

	return newclient.GetByQuery[EphemeralEnvironment](client, v2.Template, spaceID, query)
}

type CreateEnvironmentResponse struct {
	Id string `json:"Id"`
}

type CreateEnvironmentCommand struct {
	EnvironmentName string `json:"EnvironmentName"`
	SpaceID         string `uri:"spaceId"`
	ProjectID       string `uri:"projectId"`
}

type DeprovisionEphemeralEnvironmentProjectCommand struct {
}

type DeprovisionEphemeralEnvironmentProjectResponse struct {
	DeprovisioningRun DeprovisioningRunbookRun `json:"DeprovisioningRunbookRuns"`
}

type DeprovisioningRunbookRun struct {
	RunbookRunID string `json:"RunbookRunId"`
	TaskId       string `json:"TaskId"`
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
