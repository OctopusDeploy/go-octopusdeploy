package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// AwsElasticContainerRegistry represents an Amazon Web Services (AWS) Elastic
// Container Registry (ECR).
type AwsElasticContainerRegistry struct {
	AccessKey string          `json:"AccessKey,omitempty"`
	Region    string          `json:"Region"`
	SecretKey *SensitiveValue `json:"SecretKey,omitempty"`

	Feed
}

// NewAwsElasticContainerRegistry creates and initializes an Amazon Web
// Services (AWS) Elastic Container Registry (ECR).
func NewAwsElasticContainerRegistry(name string, accessKey string, secretKey *SensitiveValue, region string) *AwsElasticContainerRegistry {
	return &AwsElasticContainerRegistry{
		AccessKey: accessKey,
		Region:    region,
		SecretKey: secretKey,
		Feed:      *newFeed(name, FeedTypeAwsElasticContainerRegistry),
	}
}

// GetFeedType returns the feed type of this Amazon Web Services (AWS) Elastic
// Container Registry (ECR).
func (a *AwsElasticContainerRegistry) GetFeedType() FeedType {
	return a.FeedType
}

// Validate checks the state of this Amazon Web Services (AWS) Elastic
// Container Registry (ECR) and returns an error if invalid.
func (a *AwsElasticContainerRegistry) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(a)
}

var _ IFeed = &AwsElasticContainerRegistry{}
