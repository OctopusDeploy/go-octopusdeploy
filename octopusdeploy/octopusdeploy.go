package octopusdeploy

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dghubble/sling"
)

// Client is a Twitter client for making Twitter API requests.
type Client struct {
	sling *sling.Sling
	// Octopus Deploy API Services
	DeploymentProcess *DeploymentProcessService
	ProjectGroup      *ProjectGroupService
	Projects          *ProjectService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, octopusURL, octopusAPIKey string) *Client {
	baseURLWithAPI := strings.TrimRight(octopusURL, "/")
	baseURLWithAPI = fmt.Sprintf("%s/api/", baseURLWithAPI)
	fmt.Println(baseURLWithAPI)
	base := sling.New().Client(httpClient).Base(baseURLWithAPI).Set("X-Octopus-ApiKey", octopusAPIKey)
	return &Client{
		sling:             base,
		DeploymentProcess: NewDeploymentProcessService(base.New()),
		ProjectGroup:      NewProjectGroupService(base.New()),
		Projects:          NewProjectService(base.New()),
	}
}
