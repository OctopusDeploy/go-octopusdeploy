package deployments

type DeploymentActionContainer struct {
	FeedID     string `json:"FeedId,omitempty"`
	Image      string `json:"Image,omitempty"`
	GitUrl     string `json:"GitUrl,omitempty"`
	Dockerfile string `json:"Dockerfile,omitempty"`
}

// NewDeploymentActionContainer creates and initializes a new Kubernetes endpoint.
func NewDeploymentActionContainer(feedID *string, image *string, gitUrl *string, dockerfile *string) *DeploymentActionContainer {
	deploymentActionContainer := &DeploymentActionContainer{}

	if feedID != nil && len(*feedID) > 0 {
		deploymentActionContainer.FeedID = *feedID
	}

	if image != nil && len(*image) > 0 {
		deploymentActionContainer.Image = *image
	}

	if gitUrl != nil && len(*gitUrl) > 0 {
		deploymentActionContainer.GitUrl = *gitUrl
	}

	if dockerfile != nil && len(*dockerfile) > 0 {
		deploymentActionContainer.Dockerfile = *dockerfile
	}

	return deploymentActionContainer
}
