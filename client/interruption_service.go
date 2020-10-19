package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type interruptionService struct {
	service
}

func newInterruptionService(sling *sling.Sling, uriTemplate string) *interruptionService {
	return &interruptionService{
		service: newService(serviceInterruptionService, sling, uriTemplate, new(model.Interruption)),
	}
}

func (s interruptionService) getPagedResponse(path string) ([]*model.Interruption, error) {
	resources := []*model.Interruption{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Interruptions), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.Interruptions)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// GetByID returns the interruption that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s interruptionService) GetByID(id string) (*model.Interruption, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Interruption), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.Interruption), nil
}

// GetByIDs gets a list of interruptions that match the input IDs.
func (s interruptionService) GetByIDs(ids []string) ([]*model.Interruption, error) {
	if len(ids) == 0 {
		return []*model.Interruption{}, nil
	}

	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []*model.Interruption{}, err
	}

	return s.getPagedResponse(path)
}

// GetAll returns all interruptions. If none can be found or an error occurs,
// it returns an empty collection.
func (s interruptionService) GetAll() ([]*model.Interruption, error) {
	path, err := getPath(s)
	if err != nil {
		return []*model.Interruption{}, err
	}

	return s.getPagedResponse(path)
}

// Submit Submits a dictionary of form values for the interruption. Only the user with responsibility for this interruption can submit this form.
func (s interruptionService) Submit(resource *model.Interruption, r *model.InterruptionSubmitRequest) (*model.Interruption, error) {
	path := resource.Links[linkSubmit]

	resp, err := apiPost(s.getClient(), r, new(model.Interruption), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Interruption), nil
}

// GetResponsibility gets the User that is currently responsible for the Interruption.
func (s interruptionService) GetResponsibility(resource *model.Interruption) (*model.User, error) {
	path := resource.Links[linkResponsible]

	resp, err := apiGet(s.getClient(), new(model.User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*model.User), nil
}

// TakeResponsibility Allows the current user to take responsibility for this interruption. Only users in one of the responsible teams on this interruption can take responsibility for it.
func (s interruptionService) TakeResponsibility(resource *model.Interruption) (*model.User, error) {
	path := resource.Links[linkResponsible]

	resp, err := apiUpdate(s.getClient(), nil, new(model.User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*model.User), nil
}
