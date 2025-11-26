package platformhubversioncontrolsettings

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const (
	path = "/api/platformhub/versioncontrol"
)

// Get returns the Platform Hub version control settings.
func Get(client newclient.Client) (*Resource, error) {
	return newclient.Get[Resource](client.HttpSession(), path)
}

// Update modifies the Platform Hub version control settings.
func Update(client newclient.Client, resource *Resource) (*Resource, error) {
	return newclient.Put[Resource](client.HttpSession(), path, resource)
}
