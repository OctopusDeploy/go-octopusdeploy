package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dghubble/sling"
)

type AzureDevOpsService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewAzureDevOpsService(sling *sling.Sling, uriTemplate string) *AzureDevOpsService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &AzureDevOpsService{
		sling: sling,
		path:  path,
	}
}

func (s *AzureDevOpsService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("AzureDevOpsService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("AzureDevOpsService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &AzureDevOpsService{}
