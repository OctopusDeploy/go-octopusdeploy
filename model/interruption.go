package model

import "time"

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

type Interruptions struct {
	Items []Interruption `json:"Items"`
	PagedResults
}

type InterruptionSubmitRequest struct {
	Instructions string `json:"Instructions"`
	Notes        string `json:"Notes"`
	Result       string `json:"Result"`
}

const ManualInterverventionApprove = "Proceed"
const ManualInterventionDecline = "Abort"
