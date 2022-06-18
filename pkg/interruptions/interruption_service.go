package interruptions

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/users"
	"github.com/dghubble/sling"
)

type InterruptionService struct {
	services.Service
}

func NewInterruptionService(sling *sling.Sling, uriTemplate string) *InterruptionService {
	return &InterruptionService{
		Service: services.NewService(constants.ServiceInterruptionService, sling, uriTemplate),
	}
}

func (s InterruptionService) getPagedResponse(path string) ([]*Interruption, error) {
	resources := []*Interruption{}
	loadNextPage := true

	for loadNextPage {
		resp, err := services.ApiGet(s.GetClient(), new(Interruptions), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Interruptions)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = services.LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// GetByID returns the interruption that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *InterruptionService) GetByID(id string) (*Interruption, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(Interruption), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Interruption), nil
}

// GetByIDs gets a list of interruptions that match the input IDs.
func (s *InterruptionService) GetByIDs(ids []string) ([]*Interruption, error) {
	if len(ids) == 0 {
		return []*Interruption{}, nil
	}

	path, err := services.GetByIDsPath(s, ids)
	if err != nil {
		return []*Interruption{}, err
	}

	return s.getPagedResponse(path)
}

// GetAll returns all interruptions. If none can be found or an error occurs,
// it returns an empty collection.
func (s *InterruptionService) GetAll() ([]*Interruption, error) {
	items := []*Interruption{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = services.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// Submit Submits a dictionary of form values for the interruption. Only the user with responsibility for this interruption can submit this form.
func (s *InterruptionService) Submit(resource *Interruption, r *InterruptionSubmitRequest) (*Interruption, error) {
	path := resource.Links[constants.LinkSubmit]

	resp, err := services.ApiPost(s.GetClient(), r, new(Interruption), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Interruption), nil
}

// GetResponsibility gets the User that is currently responsible for the Interruption.
func (s InterruptionService) GetResponsibility(resource *Interruption) (*users.User, error) {
	path := resource.Links[constants.LinkResponsible]

	resp, err := services.ApiGet(s.GetClient(), new(users.User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*users.User), nil
}

// TakeResponsibility Allows the current user to take responsibility for this interruption. Only users in one of the responsible teams on this interruption can take responsibility for it.
func (s InterruptionService) TakeResponsibility(resource *Interruption) (*users.User, error) {
	path := resource.Links[constants.LinkResponsible]

	resp, err := services.ApiUpdate(s.GetClient(), nil, new(users.User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*users.User), nil
}
