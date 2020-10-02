package model

import (
	"time"

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

// GetID returns the ID value of the RootResource.
func (resource RootResource) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this RootResource.
func (resource RootResource) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this RootResource was changed.
func (resource RootResource) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this RootResource.
func (resource RootResource) GetLinks() map[string]string {
	return resource.Links
}

func (resource RootResource) SetID(id string) {
	resource.ID = id
}

func (resource RootResource) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

func (resource RootResource) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the RootResource and returns an error if invalid.
func (resource RootResource) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	validations := []error{
		ValidateSemanticVersion("Version", resource.Version),
		ValidateSemanticVersion("APIVersion", resource.APIVersion),
		ValidateRequiredUUID("InstallationID", resource.InstallationID),
	}

	return ValidateMultipleProperties(validations)
}

var _ ResourceInterface = &RootResource{}
