package environments

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

// Get returns a collection of environments based on the criteria defined by
// its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func Get(client newclient.Client, spaceID string, environmentsQuery EnvironmentQuery) (*resources.Resources[*Environment], error) {
	return newclient.GetByQuery[Environment](client, v2.Template, spaceID, environmentsQuery)
}

// GetByID returns the environment that matches the input ID. If one cannot be
// found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, id string) (*Environment, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	expandedUri, err := client.URITemplateCache().Expand(v2.GetById, map[string]any{
		"spaceId": spaceID,
		"id":      id,
	})

	resp, err := newclient.Get[Environment](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetByName returns the environment that matches the input name. If one cannot be
// found, it returns nil and an error.
func GetByName(client newclient.Client, spaceID string, name string) (*Environment, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName)
	}

	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	projects, err := Get(client, spaceID, EnvironmentQuery{
		PartialName: name,
	})
	if err != nil {
		return nil, err
	}

	for _, project := range projects.Items {
		if strings.EqualFold(project.Name, name) {
			return project, nil
		}
	}

	return nil, services.ErrItemNotFound
}

// GetByIdentifier returns the environment that matches environment based on ID or Name. If one cannot be
// found, it returns nil and an error.
func GetByIdentifier(client newclient.Client, spaceID string, identifier string) (*Environment, error) {
	environment, err := GetByID(client, spaceID, identifier)
	if err != nil {
		apiError, ok := err.(*core.APIError)
		if ok && apiError.StatusCode != 404 {
			return nil, err
		}
	} else {
		if environment != nil {
			return environment, nil
		}
	}

	return GetByName(client, spaceID, identifier)
}
