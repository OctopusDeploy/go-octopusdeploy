package environments

import (
	"errors"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

// Get returns a collection of environments based on the criteria defined by
// its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func Get(client newclient.Client, spaceID string, environmentsQuery EnvironmentQuery) (*EnvironmentResponse, error) {
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	values, success := uritemplates.Struct2map(environmentsQuery)
	if success == false {
		return nil, errors.New("failed to convert query")
	}

	if values == nil {
		values = map[string]any{}
	}
	values["spaceId"] = spaceID

	path, err := client.URITemplateCache().Expand(v2.Template, values)
	if err != nil {
		return nil, err
	}

	res, err := newclient.Get[EnvironmentResponse](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return res, nil
}
