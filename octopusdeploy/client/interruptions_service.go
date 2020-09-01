package client

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

type InterruptionsService struct {
	sling *sling.Sling
	path  string
}

func NewInterruptionService(sling *sling.Sling) *InterruptionsService {
	return &InterruptionsService{
		sling: sling,
		path:  "interruptions",
	}
}

// Get returns the interruption matching the id
func (s *InterruptionsService) Get(id string) (*model.Interruption, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Interruption), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Interruption), nil
}

// GetAll returns all interruptions in Octopus Deploy
func (s *InterruptionsService) GetAll() (*[]model.Interruption, error) {
	var p []model.Interruption
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.Interruptions), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.Interruptions)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

// Submit Submits a dictionary of form values for the interruption. Only the user with responsibility for this interruption can submit this form.
func (s *InterruptionsService) Submit(i *model.Interruption, r *model.InterruptionSubmitRequest) (*model.Interruption, error) {
	path := i.Links["Submit"]

	resp, err := apiPost(s.sling, r, new(model.Interruption), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Interruption), nil
}

// GetResponsability Gets the user that is currently responsible for this interruption.
func (s *InterruptionsService) GetResponsability(i *model.Interruption) (*model.User, error) {
	path := i.Links["Responsible"]

	resp, err := apiGet(s.sling, new(model.User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*model.User), nil
}

// TakeResponsability Allows the current user to take responsibility for this interruption. Only users in one of the responsible teams on this interruption can take responsibility for it.
func (s *InterruptionsService) TakeResponsability(i *model.Interruption) (*model.User, error) {
	path := i.Links["Responsible"]

	resp, err := apiUpdate(s.sling, nil, new(model.User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*model.User), nil
}
