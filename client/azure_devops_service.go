package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dghubble/sling"
)

type AzureDevOpsService struct {
	sling *sling.Sling
	path  string
}

func NewAzureDevOpsService(sling *sling.Sling) *AzureDevOpsService {
	if sling == nil {
		fmt.Println(fmt.Errorf("AzureDevOpsService: input parameter (sling) is nil"))
		return nil
	}

	return &AzureDevOpsService{
		sling: sling,
		path:  "azuredevopsissuetracker/connectivitycheck",
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
