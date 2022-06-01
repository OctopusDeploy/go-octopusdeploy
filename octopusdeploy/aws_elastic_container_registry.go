package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// AwsElasticContainerRegistry represents an Amazon Web Services (AWS) Elastic
// Container Registry (ECR).
type AwsElasticContainerRegistry struct {
	AccessKey string          `json:"AccessKey" validate:"required"`
	Region    string          `json:"Region" validate:"required"`
	SecretKey *SensitiveValue `json:"SecretKey" validate:"required"`

	feed
}

// NewAwsElasticContainerRegistry creates and initializes an Amazon Web
// Services (AWS) Elastic Container Registry (ECR).
func NewAwsElasticContainerRegistry(name string, accessKey string, secretKey *SensitiveValue, region string) (*AwsElasticContainerRegistry, error) {
	if isEmpty(name) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	if isEmpty(accessKey) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterAccessKey)
	}

	if secretKey == nil {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterSecretKey)
	}

	if isEmpty(region) {
		return nil, createRequiredParameterIsEmptyOrNilError("region")
	}

	feed := AwsElasticContainerRegistry{
		AccessKey: accessKey,
		Region:    region,
		SecretKey: secretKey,
		feed:      *newFeed(name, FeedTypeAwsElasticContainerRegistry),
	}

	// validate to ensure that all expectations are met
	err := feed.Validate()
	if err != nil {
		return nil, err
	}

	return &feed, nil
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
