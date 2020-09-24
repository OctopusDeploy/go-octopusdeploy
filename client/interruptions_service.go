package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type interruptionsService struct {
	name        string                    `validate:"required"`
	path        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newInterruptionsService(sling *sling.Sling, uriTemplate string) *interruptionsService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &interruptionsService{
		name:        serviceInterruptionsService,
		path:        strings.TrimSpace(uriTemplate),
		sling:       sling,
		uriTemplate: template,
	}
}

func (s interruptionsService) getClient() *sling.Sling {
	return s.sling
}

func (s interruptionsService) getName() string {
	return s.name
}

func (s interruptionsService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// GetByID returns the interruption that matches the input ID. If one cannot be found, it returns nil and an error.
func (s interruptionsService) GetByID(id string) (*model.Interruption, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Interruption), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Interruption), nil
}

// GetByIDs gets a list of interruptions that match the input IDs.
func (s interruptionsService) GetByIDs(ids []string) ([]model.Interruption, error) {
	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []model.Interruption{}, err
	}

	return s.getPagedResponse(path)
}

// GetAll returns interruptions for user attention. The results will be sorted by date from most recently to least recently created.
func (s interruptionsService) GetAll() ([]model.Interruption, error) {
	path, err := getPath(s)
	if err != nil {
		return []model.Interruption{}, err
	}

	return s.getPagedResponse(path)
}

// Submit Submits a dictionary of form values for the interruption. Only the user with responsibility for this interruption can submit this form.
func (s interruptionsService) Submit(resource *model.Interruption, r *model.InterruptionSubmitRequest) (*model.Interruption, error) {
	path := resource.Links[linkSubmit]

	resp, err := apiPost(s.getClient(), r, new(model.Interruption), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Interruption), nil
}

// GetResponsibility gets the User that is currently responsible for the Interruption.
func (s interruptionsService) GetResponsibility(resource *model.Interruption) (*model.User, error) {
	path := resource.Links[linkResponsible]

	resp, err := apiGet(s.getClient(), new(model.User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*model.User), nil
}

// TakeResponsibility Allows the current user to take responsibility for this interruption. Only users in one of the responsible teams on this interruption can take responsibility for it.
func (s interruptionsService) TakeResponsibility(resource *model.Interruption) (*model.User, error) {
	path := resource.Links[linkResponsible]

	resp, err := apiUpdate(s.getClient(), nil, new(model.User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*model.User), nil
}

func (s interruptionsService) getPagedResponse(path string) ([]model.Interruption, error) {
	var resources []model.Interruption
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Interruptions), path)
		if err != nil {
			return nil, err
		}

		responseList := resp.(*model.Interruptions)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

var _ ServiceInterface = &interruptionsService{}
