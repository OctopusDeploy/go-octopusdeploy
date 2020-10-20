package octopusdeploy

import (
	"net/url"
	"strings"

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
	channelService.service = newService(serviceChannelService, sling, uriTemplate, new(Channel))

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

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*Channel), nil
}

// GetAll returns all channels. If none can be found or an error occurs, it
// returns an empty collection.
func (s channelService) GetAll() ([]*Channel, error) {
	path, err := getPath(s)
	if err != nil {
		return []*Channel{}, err
	}

	return s.getPagedResponse(path)
}

// GetByID returns the channel that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s channelService) GetByID(id string) (*Channel, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Channel), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*Channel), nil
}

// GetByPartialName performs a lookup and returns instances of a channel with a matching partial name.
func (s channelService) GetByPartialName(name string) ([]*Channel, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*Channel{}, err
	}

	return s.getPagedResponse(path)
}

func (s channelService) GetProject(channel *Channel) (*Project, error) {
	if channel == nil {
		return nil, createInvalidParameterError(operationGetReleases, parameterChannel)
	}

	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := channel.Links["Project"]
	resp, err := apiGet(s.getClient(), new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}

func (s channelService) GetReleases(channel *Channel) ([]*Release, error) {
	if channel == nil {
		return nil, createInvalidParameterError(operationGetReleases, parameterChannel)
	}

	releases := []*Release{}

	err := validateInternalState(s)
	if err != nil {
		return releases, err
	}

	url, err := url.Parse(channel.Links[linkReleases])
	if err != nil {
		return releases, err
	}

	path := strings.Split(url.Path, "{")[0]

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Releases), path)

		if err != nil {
			return releases, err
		}

		r := resp.(*Releases)
		releases = append(releases, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return releases, nil
}

// Update modifies an Channel based on the one provided as input.
func (s channelService) Update(resource Channel) (*Channel, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*Channel), nil
}
