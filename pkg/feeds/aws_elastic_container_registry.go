package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// AwsElasticContainerRegistry represents an Amazon Web Services (AWS) Elastic
// Container Registry (ECR).
type AwsElasticContainerRegistry struct {
	AccessKey string               `json:"AccessKey" validate:"required"`
	Region    string               `json:"Region" validate:"required"`
	SecretKey *core.SensitiveValue `json:"SecretKey" validate:"required"`

	feed
}

// NewAwsElasticContainerRegistry creates and initializes an Amazon Web
// Services (AWS) Elastic Container Registry (ECR).
func NewAwsElasticContainerRegistry(name string, accessKey string, secretKey *core.SensitiveValue, region string) (*AwsElasticContainerRegistry, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	if internal.IsEmpty(accessKey) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("accessKey")
	}

	if secretKey == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("secretKey")
	}

	if internal.IsEmpty(region) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("region")
	}

	feed := AwsElasticContainerRegistry{
		AccessKey: accessKey,
		Region:    region,
		SecretKey: secretKey,
		feed:      *newFeed(name, FeedTypeAwsElasticContainerRegistry),
	}

	// validate to ensure that all expectations are met
	if err := feed.Validate(); err != nil {
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
