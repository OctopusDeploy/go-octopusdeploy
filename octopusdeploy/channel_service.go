package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type channelService struct {
	versionRuleTestPath string

	canDeleteService
}

func newChannelService(sling *sling.Sling, uriTemplate string, versionRuleTestPath string) *channelService {
	channelService := &channelService{
		versionRuleTestPath: versionRuleTestPath,
	}
	channelService.service = newService(ServiceChannelService, sling, uriTemplate)

	return channelService
}

func (s channelService) getPagedResponse(path string) ([]*Channel, error) {
	resources := []*Channel{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Channels), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Channels)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new channel.
func (s channelService) Add(resource *Channel) (*Channel, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(Channel), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Channel), nil
}

// Get returns a collection of channels based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s channelService) Get(channelsQuery ChannelsQuery) (*Channels, error) {
	path, err := s.getURITemplate().Expand(channelsQuery)
	if err != nil {
		return &Channels{}, err
	}

	response, err := apiGet(s.getClient(), new(Channels), path)
	if err != nil {
		return &Channels{}, err
	}

	return response.(*Channels), nil
}

// GetAll returns all channels. If none can be found or an error occurs, it
// returns an empty collection.
func (s channelService) GetAll() ([]*Channel, error) {
	items := []*Channel{}
	path := s.BasePath + "/all"

	_, err := apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the channel that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s channelService) GetByID(id string) (*Channel, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError(OperationGetByID, ParameterID)
	}

	path := s.BasePath + "/" + id
	resp, err := apiGet(s.getClient(), new(Channel), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*Channel), nil
}

func (s channelService) GetProject(channel *Channel) (*Project, error) {
	if channel == nil {
		return nil, createInvalidParameterError(OperationGetProject, ParameterChannel)
	}

	path := channel.GetLinks()[linkProjects]
	resp, err := apiGet(s.getClient(), new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}

func (s channelService) GetReleases(channel *Channel, releaseQuery ...*ReleaseQuery) (*Releases, error) {
	if channel == nil {
		return nil, createInvalidParameterError(OperationGetReleases, ParameterChannel)
	}

	uriTemplate, err := uritemplates.Parse(channel.GetLinks()[linkReleases])
	if err != nil {
		return &Releases{}, err
	}

	values := make(map[string]interface{})
	path, err := uriTemplate.Expand(values)
	if err != nil {
		return &Releases{}, err
	}

	if releaseQuery != nil {
		path, err = uriTemplate.Expand(releaseQuery[0])
		if err != nil {
			return &Releases{}, err
		}
	}

	resp, err := apiGet(s.getClient(), new(Releases), path)
	if err != nil {
		return &Releases{}, err
	}

	return resp.(*Releases), nil
}

// Update modifies an Channel based on the one provided as input.
func (s channelService) Update(resource Channel) (*Channel, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(Channel), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Channel), nil
}
