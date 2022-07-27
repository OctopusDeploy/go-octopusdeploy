package events

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type Event struct {
	Category                string            `json:"Category,omitempty"`
	ChangeDetails           *ChangeDetails    `json:"ChangeDetails,omitempty"`
	Comments                string            `json:"Comments,omitempty"`
	Details                 string            `json:"Details,omitempty"`
	IdentityEstablishedWith string            `json:"IdentityEstablishedWith,omitempty"`
	IsService               bool              `json:"IsService,omitempty"`
	Message                 string            `json:"Message,omitempty"`
	MessageHTML             string            `json:"MessageHtml,omitempty"`
	MessageReferences       []*EventReference `json:"MessageReferences"`
	Occurred                time.Time         `json:"Occurred,omitempty"`
	RelatedDocumentIds      []string          `json:"RelatedDocumentIds"`
	SpaceID                 string            `json:"SpaceId,omitempty"`
	UserAgent               string            `json:"UserAgent,omitempty"`
	UserID                  string            `json:"UserId,omitempty"`
	Username                string            `json:"Username,omitempty"`

	resources.Resource
}
