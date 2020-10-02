package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// EndpointInterface defines the interface for all endpoints.
type EndpointInterface interface {
	GetCommunicationStyle() string

	ResourceInterface
}

// endpoint is the base definition of an endpoint and describes its
// communication style (SSH, Kubernetes, etc.)
type endpoint struct {
	CommunicationStyle string

	Resource
}

func (resource endpoint) GetCommunicationStyle() string {
	return resource.CommunicationStyle
}

// GetID returns the ID value of the endpoint.
func (resource endpoint) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of
// this endpoint.
func (resource endpoint) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this endpoint
// was changed.
func (resource endpoint) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this endpoint.
func (resource endpoint) GetLinks() map[string]string {
	return resource.Links
}

func (resource endpoint) SetID(id string) {
	resource.ID = id
}

func (resource endpoint) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

func (resource endpoint) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the endpoint and returns an error if invalid.
func (resource endpoint) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}

		return err
	}

	return nil
}

var _ EndpointInterface = &endpoint{}
