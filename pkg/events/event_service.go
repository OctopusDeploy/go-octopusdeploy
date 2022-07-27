package events

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

// EventService handles communication with event-related methods of the
// Octopus API.
type EventService struct {
	agentsPath        *uritemplates.UriTemplate
	categoriesPath    *uritemplates.UriTemplate
	documentTypesPath *uritemplates.UriTemplate
	groupsPath        *uritemplates.UriTemplate

	services.Service
}

func NewEventService(sling *sling.Sling, uriTemplate string, agentsPath string, categoriesPath string, documentTypesPath string, groupsPath string) *EventService {
	agents, _ := uritemplates.Parse(strings.TrimSpace(agentsPath))
	categories, _ := uritemplates.Parse(strings.TrimSpace(categoriesPath))
	documentTypes, _ := uritemplates.Parse(strings.TrimSpace(documentTypesPath))
	groups, _ := uritemplates.Parse(strings.TrimSpace(groupsPath))

	return &EventService{
		agentsPath:        agents,
		categoriesPath:    categories,
		documentTypesPath: documentTypes,
		groupsPath:        groups,
		Service:           services.NewService(constants.ServiceEventService, sling, uriTemplate),
	}
}

// Get returns a collection of events based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s EventService) Get(query EventsQuery) (*resources.Resources[Event], error) {
	path, err := s.GetURITemplate().Expand(query)
	if err != nil {
		return &resources.Resources[Event]{}, err
	}

	response, err := services.ApiGet(s.GetClient(), new(resources.Resources[Event]), path)
	if err != nil {
		return &resources.Resources[Event]{}, err
	}

	return response.(*resources.Resources[Event]), nil
}

func (s EventService) GetAgents() (*[]EventAgent, error) {
	resp, err := services.ApiGet(s.GetClient(), new([]EventAgent), s.agentsPath.String())
	if err != nil {
		return nil, err
	}

	return resp.(*[]EventAgent), nil
}

func (s EventService) GetCategories(query EventCategoriesQuery) (*[]EventCategory, error) {
	path, err := s.categoriesPath.Expand(query)
	if err != nil {
		return &[]EventCategory{}, err
	}

	resp, err := services.ApiGet(s.GetClient(), new([]EventCategory), path)
	if err != nil {
		return nil, err
	}

	return resp.(*[]EventCategory), nil
}

func (s EventService) GetDocumentTypes() (*[]DocumentType, error) {
	resp, err := services.ApiGet(s.GetClient(), new([]DocumentType), s.documentTypesPath.String())
	if err != nil {
		return nil, err
	}

	return resp.(*[]DocumentType), nil
}

func (s EventService) GetGroups(query EventGroupsQuery) (*[]EventGroup, error) {
	path, err := s.groupsPath.Expand(query)
	if err != nil {
		return &[]EventGroup{}, err
	}

	resp, err := services.ApiGet(s.GetClient(), new([]EventGroup), path)
	if err != nil {
		return nil, err
	}

	return resp.(*[]EventGroup), nil
}
