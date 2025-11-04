package environments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

// Get returns a collection of environments based on the criteria defined by
// its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func Get(client newclient.Client, spaceID string, environmentsQuery EnvironmentQuery) (*resources.Resources[*Environment], error) {
	return newclient.GetByQuery[Environment](client, v2.Template, spaceID, environmentsQuery)
}
