package v2

import (
	"errors"
	"math"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/ephemeralenvironments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

const templateV2 = "/api/{spaceId}/environments/v2{/id}{?name,skip,ids,take,partialName,type}"

// GetV2 returns a collection of environments based on the criteria defined by
// its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func GetV2(client newclient.Client, spaceID string, environmentsV2Query EnvironmentV2Query) (*resources.Resources[*EnvironmentV2], error) {
	return newclient.GetByQuery[EnvironmentV2](client, templateV2, spaceID, environmentsV2Query)
}

// GetAllEphemeralEnvironments returns a response containing all ephemeral environments. If an error occurs, it returns nil.
func GetAllEphemeralEnvironments(client newclient.Client, spaceID string) (*ephemeralenvironments.EphemeralEnvironmentResponse, error) {
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	query := &EnvironmentV2Query{
		Skip: 0,
		Take: math.MaxInt32,
		Type: []string{"Ephemeral"},
	}

	values, success := uritemplates.Struct2map(query)
	if success == false {
		return nil, errors.New("failed to convert query")
	}

	if values == nil {
		values = map[string]any{}
	}
	values["spaceId"] = spaceID

	path, err := client.URITemplateCache().Expand(templateV2, values)
	if err != nil {
		return nil, err
	}

	res, err := newclient.Get[ephemeralenvironments.EphemeralEnvironmentResponse](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return res, nil
}
