package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type RootService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewRootService(sling *sling.Sling, uriTemplate string) *RootService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &RootService{
		sling: sling,
		path:  path,
	}
}

func (s *RootService) Get() (*model.RootResource, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new(model.RootResource), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.RootResource), nil
}

func (s *RootService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("RootService: the internal client is nil")
	}

	return nil
}

var _ ServiceInterface = &RootService{}
