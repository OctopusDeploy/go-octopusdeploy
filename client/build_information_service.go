package client

import (
	"strings"

	"github.com/dghubble/sling"
)

type BuildInformationService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewBuildInformationService(sling *sling.Sling, uriTemplate string) *BuildInformationService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &BuildInformationService{
		name:  "BuildInformationService",
		path:  path,
		sling: sling,
	}
}

func (s *BuildInformationService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &BuildInformationService{}
