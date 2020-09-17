package client

import (
	"strings"

	"github.com/dghubble/sling"
)

type AzureDevOpsService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewAzureDevOpsService(sling *sling.Sling, uriTemplate string) *AzureDevOpsService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &AzureDevOpsService{
		name:  "AzureDevOpsService",
		path:  path,
		sling: sling,
	}
}

func (s *AzureDevOpsService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &AzureDevOpsService{}
