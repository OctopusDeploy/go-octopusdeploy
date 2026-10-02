package machines

// AwsEcsClusterEndpoint represents the endpoint of an Amazon ECS cluster
// deployment target. The target type is supplied by the aws-ecs-target step
// package, but the server reports the endpoint with its own communication
// style and its settings as fields.
type AwsEcsClusterEndpoint struct {
	AccountID                        string `json:"AccountId,omitempty"`
	AssumedRoleARN                   string `json:"AssumedRoleArn,omitempty"`
	AssumedRoleSession               string `json:"AssumedRoleSession,omitempty"`
	AssumeRole                       bool   `json:"AssumeRole"`
	AssumeRoleExternalID             string `json:"AssumeRoleExternalId,omitempty"`
	AssumeRoleSessionDurationSeconds int    `json:"AssumeRoleSessionDurationSeconds,omitempty"`
	ClusterName                      string `json:"ClusterName,omitempty"`
	DefaultWorkerPoolID              string `json:"DefaultWorkerPoolId,omitempty"`
	Region                           string `json:"Region,omitempty"`
	UseInstanceRole                  bool   `json:"UseInstanceRole"`

	endpoint
}

// NewAwsEcsClusterEndpoint creates and initializes a new Amazon ECS cluster endpoint.
func NewAwsEcsClusterEndpoint(clusterName string, region string) *AwsEcsClusterEndpoint {
	return &AwsEcsClusterEndpoint{
		ClusterName: clusterName,
		Region:      region,
		endpoint:    *newEndpoint("AwsEcsCluster"),
	}
}

// GetDefaultWorkerPoolID returns the default worker pool ID of this endpoint.
func (e AwsEcsClusterEndpoint) GetDefaultWorkerPoolID() string {
	return e.DefaultWorkerPoolID
}

// SetDefaultWorkerPoolID sets the default worker pool ID of this endpoint.
func (e *AwsEcsClusterEndpoint) SetDefaultWorkerPoolID(defaultWorkerPoolID string) {
	e.DefaultWorkerPoolID = defaultWorkerPoolID
}

var _ IEndpoint = &AwsEcsClusterEndpoint{}
var _ IRunsOnAWorker = &AwsEcsClusterEndpoint{}
