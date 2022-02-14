package resources

type DeploymentActionContainer struct {
	FeedID string `json:"FeedId,omitempty"`
	Image  string `json:"Image,omitempty"`
}

// NewKubernetesEndpoint creates and initializes a new Kubernetes endpoint.
func NewDeploymentActionContainer(feedID *string, image *string) *DeploymentActionContainer {
	deploymentActionContainer := &DeploymentActionContainer{}

	if len(*feedID) > 0 {
		deploymentActionContainer.FeedID = *feedID
	}

	if len(*image) > 0 {
		deploymentActionContainer.Image = *image
	}

	return deploymentActionContainer
}
