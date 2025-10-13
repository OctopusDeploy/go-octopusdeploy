package ephemeralenvironments

import (
	"errors"
	"math"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2/environments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

// GetAllEphemeralEnvironments returns a response containing all ephemeral environments. If an error occurs, it returns nil.
func GetAllEphemeralEnvironments(client newclient.Client, spaceID string) (*EphemeralEnvironmentResponse, error) {
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	query := &environments.EnvironmentQuery{
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

	path, err := client.URITemplateCache().Expand(v2.Template, values)
	if err != nil {
		return nil, err
	}

	res, err := newclient.Get[EphemeralEnvironmentResponse](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return res, nil
}
