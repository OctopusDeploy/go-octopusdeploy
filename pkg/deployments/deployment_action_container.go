package deployments

type DeploymentActionContainer struct {
	FeedID string `json:"FeedId,omitempty"`
	Image  string `json:"Image,omitempty"`
}

// NewDeploymentActionContainer creates and initializes a new Kubernetes endpoint.
func NewDeploymentActionContainer(feedID *string, image *string) *DeploymentActionContainer {
	deploymentActionContainer := &DeploymentActionContainer{}

	if feedID != nil && len(*feedID) > 0 {
		deploymentActionContainer.FeedID = *feedID
	}

	if image != nil && len(*image) > 0 {
		deploymentActionContainer.Image = *image
	}

	return deploymentActionContainer
}
