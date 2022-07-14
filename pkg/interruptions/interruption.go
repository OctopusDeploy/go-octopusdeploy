package interruptions

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type Interruptions struct {
	Items []*Interruption `json:"Items"`
	resources.PagedResults
}

type Interruption struct {
	CanTakeResponsibility       bool      `json:"CanTakeResponsibility,omitempty"`
	CorrelationID               string    `json:"CorrelationId,omitempty"`
	Created                     time.Time `json:"Created,omitempty"`
	Form                        *Form     `json:"Form,omitempty"`
	HasResponsibility           bool      `json:"HasResponsibility"`
	IsLinkedToOtherInterruption bool      `json:"IsLinkedToOtherInterruption"`
	IsPending                   bool      `json:"IsPending"`
	RelatedDocumentIDs          []string  `json:"RelatedDocumentIds"`
	ResponsibleTeamIDs          []string  `json:"ResponsibleTeamIds"`
	ResponsibleUserID           string    `json:"ResponsibleUserId,omitempty"`
	SpaceID                     string    `json:"SpaceId,omitempty"`
	TaskID                      string    `json:"TaskId,omitempty"`
	Title                       string    `json:"Title,omitempty"`

	resources.Resource
}

func NewInterruption() *Interruption {
	return &Interruption{
		Resource: *resources.NewResource(),
	}
}

const ManualInterverventionApprove = "Proceed"
const ManualInterventionDecline = "Abort"
