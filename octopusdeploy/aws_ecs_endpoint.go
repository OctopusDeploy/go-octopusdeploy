package octopusdeploy

type StepPackageInputs struct {
	ClusterName  string `json:"clusterName" validate:"required"`
	Region       string `json:"region" validate:"required"`
	AwsAccountID string `json:"awsAccount" validate:"required"`
}

// AmazonECSEndpoint represents an Amazon ECS endpoint.
type AmazonECSEndpoint struct {
	DefaultWorkerPoolID    string             `json:"DefaultWorkerPoolId"`
	DeploymentTargetTypeId string             `json:"DeploymentTargetTypeId" validate:"required"`
	StepPackageId          string             `json:"StepPackageId" validate:"required"`
	StepPackageVersion     string             `json:"StepPackageVersion" validate:"required"`
	Inputs                 *StepPackageInputs `json:"Inputs" validate:"required"`
	endpoint
}

// NewAmazonECSEndpoint creates a new endpoint for Amazon ECS.
func NewAmazonECSEndpoint() *AmazonECSEndpoint {
	appEndpoint := &AmazonECSEndpoint{
		endpoint:               *newEndpoint("StepPackage"),
		StepPackageId:          "aws-ecs-target",
		DeploymentTargetTypeId: "aws-ecs-target",
	}

	return appEndpoint
}

// GetAccountID returns the account ID associated with this endpoint.
func (endpoint AmazonECSEndpoint) GetAccountID() string {
	return endpoint.Inputs.AwsAccountID
}

// GetDefaultWorkerPoolID returns the default worker pool ID of this endpoint.
func (endpoint AmazonECSEndpoint) GetDefaultWorkerPoolID() string {
	return endpoint.DefaultWorkerPoolID
}

// SetDefaultWorkerPoolID sets the default worker pool ID of this endpoint.
func (endpoint AmazonECSEndpoint) SetDefaultWorkerPoolID(defaultWorkerPoolID string) {
	endpoint.DefaultWorkerPoolID = defaultWorkerPoolID
}

var _ IResource = &AmazonECSEndpoint{}
var _ IEndpoint = &AmazonECSEndpoint{}
var _ IEndpointWithAccount = &AmazonECSEndpoint{}
var _ IRunsOnAWorker = &AmazonECSEndpoint{}
