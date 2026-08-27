package githubconnections

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const settingsPath = "/api/github/app/settings"

// Settings represents the server-wide GitHub App settings.
type Settings struct {
	CanUseGitHubApp   bool `json:"CanUseGitHubApp"`
	CanUseTrustedFlow bool `json:"CanUseTrustedFlow"`
}

// GetSettings returns the server's GitHub App settings.
func GetSettings(client newclient.Client) (*Settings, error) {
	return newclient.Get[Settings](client.HttpSession(), settingsPath)
}
