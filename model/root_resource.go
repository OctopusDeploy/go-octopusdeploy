package model

import (
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
)

type RootResource struct {
	Application          string     `json:"Application" validate:"required"`
	Version              string     `json:"Version" validate:"required"`
	APIVersion           string     `json:"ApiVersion" validate:"required"`
	InstallationID       *uuid.UUID `json:"InstallationId" validate:"required"`
	IsEarlyAccessProgram bool       `json:"IsEarlyAccessProgram"`
	HasLongTermSupport   bool       `json:"HasLongTermSupport"`
	Resource
}

func (r *RootResource) GetID() string {
	return r.ID
}

func (r *RootResource) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)

	if err != nil {
		return err
	}

	validations := []error{
		ValidateSemanticVersion("Version", r.Version),
		ValidateSemanticVersion("APIVersion", r.APIVersion),
		ValidateRequiredUUID("InstallationID", r.InstallationID),
	}

	return ValidateMultipleProperties(validations)
}

var _ ResourceInterface = &RootResource{}
