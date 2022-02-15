package service

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

// eventService handles communication with event-related methods of the
// Octopus API.
type eventService struct {
	agentsPath        *uritemplates.UriTemplate
	categoriesPath    *uritemplates.UriTemplate
	documentTypesPath *uritemplates.UriTemplate
	groupsPath        *uritemplates.UriTemplate

	service
}

func newEventService(sling *sling.Sling, uriTemplate string, agentsPath string, categoriesPath string, documentTypesPath string, groupsPath string) *eventService {
	agents, _ := uritemplates.Parse(strings.TrimSpace(agentsPath))
	categories, _ := uritemplates.Parse(strings.TrimSpace(categoriesPath))
	documentTypes, _ := uritemplates.Parse(strings.TrimSpace(documentTypesPath))
	groups, _ := uritemplates.Parse(strings.TrimSpace(groupsPath))

	return &eventService{
		agentsPath:        agents,
		categoriesPath:    categories,
		documentTypesPath: documentTypes,
		groupsPath:        groups,
		service:           newService(ServiceEventService, sling, uriTemplate),
	}
}

// Get returns a collection of events based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s eventService) Get(query EventsQuery) (*Events, error) {
	path, err := s.getURITemplate().Expand(query)
	if err != nil {
		return &Events{}, err
	}

	response, err := apiGet(s.getClient(), new(Events), path)
	if err != nil {
		return &Events{}, err
	}

	return response.(*Events), nil
}

func (s eventService) GetAgents() (*[]EventAgent, error) {
	resp, err := apiGet(s.getClient(), new([]EventAgent), s.agentsPath.String())
	if err != nil {
		return nil, err
	}

	return resp.(*[]EventAgent), nil
}

func (s eventService) GetCategories(query EventCategoriesQuery) (*[]EventCategory, error) {
	path, err := s.categoriesPath.Expand(query)
	if err != nil {
		return &[]EventCategory{}, err
	}

	resp, err := apiGet(s.getClient(), new([]EventCategory), path)
	if err != nil {
		return nil, err
	}

	return resp.(*[]EventCategory), nil
}

func (s eventService) GetDocumentTypes() (*[]DocumentType, error) {
	resp, err := apiGet(s.getClient(), new([]DocumentType), s.documentTypesPath.String())
	if err != nil {
		return nil, err
	}

	return resp.(*[]DocumentType), nil
}

func (s eventService) GetGroups(query EventGroupsQuery) (*[]EventGroup, error) {
	path, err := s.groupsPath.Expand(query)
	if err != nil {
		return &[]EventGroup{}, err
	}

	resp, err := apiGet(s.getClient(), new([]EventGroup), path)
	if err != nil {
		return nil, err
	}

	return resp.(*[]EventGroup), nil
}
