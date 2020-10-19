package client

import (
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type channelService struct {
	versionRuleTestPath string

	service
}

func newChannelService(sling *sling.Sling, uriTemplate string, versionRuleTestPath string) *channelService {
	channelService := &channelService{
		versionRuleTestPath: versionRuleTestPath,
	}
	channelService.service = newService(serviceChannelService, sling, uriTemplate, new(model.Channel))

	return channelService
}

func (s channelService) getPagedResponse(path string) ([]*model.Channel, error) {
	resources := []*model.Channel{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Channels), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.Channels)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new channel.
func (s channelService) Add(resource *model.Channel) (*model.Channel, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Channel), nil
}

// GetAll returns all channels. If none can be found or an error occurs, it
// returns an empty collection.
func (s channelService) GetAll() ([]*model.Channel, error) {
	path, err := getPath(s)
	if err != nil {
		return []*model.Channel{}, err
	}

	return s.getPagedResponse(path)
}

// GetByID returns the channel that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s channelService) GetByID(id string) (*model.Channel, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Channel), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.Channel), nil
}

// GetByPartialName performs a lookup and returns instances of a channel with a matching partial name.
func (s channelService) GetByPartialName(name string) ([]*model.Channel, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*model.Channel{}, err
	}

	return s.getPagedResponse(path)
}

func (s channelService) GetProject(channel *model.Channel) (*model.Project, error) {
	if channel == nil {
		return nil, createInvalidParameterError(operationGetReleases, parameterChannel)
	}

	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := channel.Links["Project"]
	resp, err := apiGet(s.getClient(), new(model.Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Project), nil
}

func (s channelService) GetReleases(channel *model.Channel) ([]*model.Release, error) {
	if channel == nil {
		return nil, createInvalidParameterError(operationGetReleases, parameterChannel)
	}

	releases := []*model.Release{}

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
		resp, err := apiGet(s.getClient(), new(model.Releases), path)

		if err != nil {
			return releases, err
		}

		r := resp.(*model.Releases)
		releases = append(releases, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return releases, nil
}

// Update modifies an Channel based on the one provided as input.
func (s channelService) Update(resource model.Channel) (*model.Channel, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Channel), nil
}
