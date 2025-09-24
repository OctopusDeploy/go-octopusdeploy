package interruptions

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/users"
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

// GetByID returns the interruption that matches the input ID. If one cannot be
// found, it returns nil and an error.
//
// Deprecated: use interruptions.GetByID()
func (s *InterruptionService) GetByID(id string) (*Interruption, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Interruption), path)
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

	return services.GetPagedResponse[Interruption](s, path)
}

// GetAll returns all interruptions. If none can be found or an error occurs,
// it returns an empty collection.
//
// Deprecated: use interruptions.GetAll()
func (s *InterruptionService) GetAll() ([]*Interruption, error) {
	items := []*Interruption{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// Submit Submits a dictionary of form values for the interruption. Only the user with responsibility for this interruption can submit this form.
//
// Deprecated: use interruptions.Submit()
func (s *InterruptionService) Submit(resource *Interruption, r *InterruptionSubmitRequest) (*Interruption, error) {
	path := resource.Links[constants.LinkSubmit]

	resp, err := services.ApiPost(s.GetClient(), r, new(Interruption), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Interruption), nil
}

// GetResponsibility gets the User that is currently responsible for the Interruption.
//
// Deprecated: use interruptions.GetResponsibility()
func (s InterruptionService) GetResponsibility(resource *Interruption) (*users.User, error) {
	path := resource.Links[constants.LinkResponsible]

	resp, err := api.ApiGet(s.GetClient(), new(users.User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*users.User), nil
}

// TakeResponsibility Allows the current user to take responsibility for this interruption. Only users in one of the responsible teams on this interruption can take responsibility for it.
//
// Deprecated: use interruptions.TakeResponsibility()
func (s InterruptionService) TakeResponsibility(resource *Interruption) (*users.User, error) {
	path := resource.Links[constants.LinkResponsible]

	resp, err := services.ApiUpdate(s.GetClient(), nil, new(users.User), path)
	if err != nil {
		return nil, err
	}
	return resp.(*users.User), nil
}

// --- new ---

const template = "/api/{spaceId}/interruptions{/id}{?skip,take,regarding,pendingOnly,ids}"
const submitTemplate = "/api/{spaceId}/interruptions/{id}/submit"
const responsibleTemplate = "/api/{spaceId}/interruptions/{id}/responsible"

// GetByID returns the interruption that matches the input ID. If one cannot be
// found, it returns nil and an error.
func GetByID(client newclient.Client, spaceId string, id string) (*Interruption, error) {
	return newclient.GetByID[Interruption](client, template, spaceId, id)
}

// GetAll returns all interruptions. If none can be found or an error occurs,
// it returns an empty collection.
func GetAll(client newclient.Client, spaceId string) ([]*Interruption, error) {
	return newclient.GetAll[Interruption](client, template, spaceId)
}

// GetResponsibility gets the User that is currently responsible for the Interruption.
func GetResponsibility(client newclient.Client, interruption *Interruption) (*users.User, error) {
	return newclient.GetByID[users.User](client, responsibleTemplate, interruption.SpaceID, interruption.ID)
}

// TakeResponsibility Allows the current user to take responsibility for this interruption. Only users in one of the responsible teams on this interruption can take responsibility for it.
func TakeResponsibility(client newclient.Client, interruption *Interruption) (*users.User, error) {
	return newclient.Update[users.User](client, responsibleTemplate, interruption.SpaceID, interruption.ID, new(users.User))
}

// Submit Submits a dictionary of form values for the interruption. Only the user with responsibility for this interruption can submit this form.
func Submit(client newclient.Client, interruption *Interruption, interruptionSubmitRequest *InterruptionSubmitRequest) (*Interruption, error) {
	if interruptionSubmitRequest == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAPIPost, constants.ParameterResource)
	}

	spaceID, err := internal.GetSpaceID(interruption.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	path, err := client.URITemplateCache().Expand(submitTemplate, map[string]any{
		"spaceId": spaceID,
		"id":      interruption.ID,
	})
	if err != nil {
		return nil, err
	}

	res, err := newclient.Post[Interruption](client.HttpSession(), path, interruptionSubmitRequest)
	if err != nil {
		return nil, err
	}

	return res, nil
}
