package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type interruptionService struct {
	services.service
}

func newInterruptionService(sling *sling.Sling, uriTemplate string) *interruptionService {
	return &interruptionService{
		service: services.newService(ServiceInterruptionService, sling, uriTemplate),
	}
}

func (s interruptionService) getPagedResponse(path string) ([]*Interruption, error) {
	resources := []*Interruption{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Interruptions), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Interruptions)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// GetByID returns the interruption that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s interruptionService) GetByID(id string) (*Interruption, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Interruption), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Interruption), nil
}

// GetByIDs gets a list of interruptions that match the input IDs.
func (s interruptionService) GetByIDs(ids []string) ([]*Interruption, error) {
	if len(ids) == 0 {
		return []*Interruption{}, nil
	}

	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []*Interruption{}, err
	}

	return s.getPagedResponse(path)
}

// GetAll returns all interruptions. If none can be found or an error occurs,
// it returns an empty collection.
func (s interruptionService) GetAll() ([]*Interruption, error) {
	path, err := getPath(s)
	if err != nil {
		return []*Interruption{}, err
	}

	return s.getPagedResponse(path)
}

// Submit Submits a dictionary of form values for the interruption. Only the user with responsibility for this interruption can submit this form.
func (s interruptionService) Submit(resource *Interruption, r *InterruptionSubmitRequest) (*Interruption, error) {
	path := resource.Links[linkSubmit]

	resp, err := apiPost(s.getClient(), r, new(Interruption), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Interruption), nil
}

// GetResponsibility gets the User that is currently responsible for the Interruption.
func (s interruptionService) GetResponsibility(resource *Interruption) (*User, error) {
	path := resource.Links[linkResponsible]

	resp, err := apiGet(s.getClient(), new(User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*User), nil
}

// TakeResponsibility Allows the current user to take responsibility for this interruption. Only users in one of the responsible teams on this interruption can take responsibility for it.
func (s interruptionService) TakeResponsibility(resource *Interruption) (*User, error) {
	path := resource.Links[linkResponsible]

	resp, err := apiUpdate(s.getClient(), nil, new(User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*User), nil
}
