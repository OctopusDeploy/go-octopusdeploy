package client

import (
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
)

// RootResource -
type RootResource struct {
	Application          string            `json:"Application" validate:"required"`
	Version              string            `json:"Version" validate:"required"`
	APIVersion           string            `json:"ApiVersion" validate:"required"`
	InstallationID       *uuid.UUID        `json:"InstallationId" validate:"required"`
	IsEarlyAccessProgram bool              `json:"IsEarlyAccessProgram"`
	HasLongTermSupport   bool              `json:"HasLongTermSupport" validate:"required"`
	Links                map[string]string `json:"Links" validate:"required"`
}

// Validate checks the state of the root resource and returns an error if
// invalid.
func (r *RootResource) Validate() error {
	return validator.New().Struct(r)
}
