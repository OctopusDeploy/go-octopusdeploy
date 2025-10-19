package ephemeralenvironments

import (
	"math"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2"
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
