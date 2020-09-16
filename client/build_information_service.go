package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dghubble/sling"
)

type BuildInformationService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewBuildInformationService(sling *sling.Sling, uriTemplate string) *BuildInformationService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &BuildInformationService{
		sling: sling,
		path:  path,
	}
}

func (s *BuildInformationService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("BuildInformationService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("BuildInformationService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &BuildInformationService{}
