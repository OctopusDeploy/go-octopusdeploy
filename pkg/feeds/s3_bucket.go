package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// S3Feed represents an Amazon Web Services (AWS) S3 Bucket Feed
type S3Feed struct {
	AccessKey             string               `json:"AccessKey,omitempty"`
	SecretKey             *core.SensitiveValue `json:"SecretKey,omitempty"`
	UseMachineCredentials bool                 `json:"UseMachineCredentials"`
	feed
}

// NewS3Feed creates and initializes an Amazon S3 Bucket Feed
func NewS3Feed(name string, accessKey string, secretKey *core.SensitiveValue, useMachineCredentials bool) (*S3Feed, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}
	if !useMachineCredentials {
		if internal.IsEmpty(accessKey) {
			return nil, internal.CreateRequiredParameterIsEmptyOrNilError("accessKey")
		}

		if secretKey == nil {
			return nil, internal.CreateRequiredParameterIsEmptyOrNilError("secretKey")
		}
	}
	feed := S3Feed{
		AccessKey:             accessKey,
		SecretKey:             secretKey,
		UseMachineCredentials: useMachineCredentials,
		feed:                  *newFeed(name, FeedTypeS3),
	}

	// validate to ensure that all expectations are met
	if err := feed.Validate(); err != nil {
		return nil, err
	}

	return &feed, nil
}

// Validate checks the state of this Amazon Web Services (AWS) S3 Bucket Feed
func (a *S3Feed) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
