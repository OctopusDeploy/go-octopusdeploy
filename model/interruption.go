package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Interruptions struct {
	Items []Interruption `json:"Items"`
	PagedResults
}

type Interruption struct {
	CanTakeResponsibility       bool      `json:"CanTakeResponsibility,omitempty"`
	CorrelationID               string    `json:"CorrelationId,omitempty"`
	Created                     time.Time `json:"Created,omitempty"`
	Form                        *Form     `json:"Form,omitempty"`
	HasResponsibility           bool      `json:"HasResponsibility"`
	IsLinkedToOtherInterruption bool      `json:"IsLinkedToOtherInterruption"`
	IsPending                   bool      `json:"IsPending"`
	RelatedDocumentIds          []string  `json:"RelatedDocumentIds"`
	ResponsibleTeamIds          []string  `json:"ResponsibleTeamIds"`
	ResponsibleUserID           string    `json:"ResponsibleUserId,omitempty"`
	SpaceID                     string    `json:"SpaceId,omitempty"`
	TaskID                      string    `json:"TaskId,omitempty"`
	Title                       string    `json:"Title,omitempty"`

	Resource
}

// GetID returns the ID value of the Interruption.
func (resource Interruption) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Interruption.
func (resource Interruption) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Interruption was changed.
func (resource Interruption) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Interruption.
func (resource Interruption) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the Interruption and returns an error if invalid.
func (resource Interruption) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

type InterruptionSubmitRequest struct {
	Instructions string `json:"Instructions"`
	Notes        string `json:"Notes"`
	Result       string `json:"Result"`
}

const ManualInterverventionApprove = "Proceed"
const ManualInterventionDecline = "Abort"

var _ ResourceInterface = &Interruption{}
