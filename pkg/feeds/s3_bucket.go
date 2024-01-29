package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// AwsS3Bucket represents an Amazon Web Services (AWS) S3 Bucket
type AwsS3Bucket struct {
	AccessKey string               `json:"AccessKey" validate:"required"`
	SecretKey *core.SensitiveValue `json:"SecretKey" validate:"required"`

	feed
}

// NewAwsS3Bucket creates and initializes an Amazon S3 Bucket Feed
func NewAwsS3Bucket(name string, accessKey string, secretKey *core.SensitiveValue) (*AwsS3Bucket, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	if internal.IsEmpty(accessKey) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("accessKey")
	}

	if secretKey == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("secretKey")
	}

	feed := AwsS3Bucket{
		AccessKey: accessKey,
		SecretKey: secretKey,
		feed:      *newFeed(name, FeedTypeAwsS3Bucket),
	}

	// validate to ensure that all expectations are met
	if err := feed.Validate(); err != nil {
		return nil, err
	}

	return &feed, nil
}

// Validate checks the state of this Amazon Web Services (AWS) S3 Bucket
func (a *AwsS3Bucket,) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
