package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dghubble/sling"
)

type BuildInformationService struct {
	sling *sling.Sling
	path  string
}

func NewBuildInformationService(sling *sling.Sling) *BuildInformationService {
	if sling == nil {
		fmt.Println(fmt.Errorf("BuildInformationService: input parameter (sling) is nil"))
		return nil
	}

	return &BuildInformationService{
		sling: sling,
		path:  "build-information",
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
