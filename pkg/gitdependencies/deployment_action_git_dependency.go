package gitdependencies

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type DeploymentActionGitDependency struct {
	DeploymentActionSlug string `json:"DeploymentActionSlug" validate:"required,notblank"`
	GitDependencyName    string `json:"GitDependencyName"`
}

// Validate checks the state of the deployment action git dependency and returns an error if invalid.
func (d DeploymentActionGitDependency) Validate() error {
	v := validator.New()
	if err := v.RegisterValidation("notblank", validators.NotBlank); err != nil {
		return err
	}
	return v.Struct(d)
}
