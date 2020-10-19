package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// feedService handles communication with feed-related methods of the Octopus
// API.
type feedService struct {
	builtInFeedStats string

	service
}

// newFeedService returns an feed service with a preconfigured client.
func newFeedService(sling *sling.Sling, uriTemplate string, builtInFeedStats string) *feedService {
	feedService := &feedService{
		builtInFeedStats: builtInFeedStats,
	}
	feedService.service = newService(serviceFeedService, sling, uriTemplate, new(model.Feed))

	return feedService
}

func toFeedArray(feeds []*model.Feed) []model.IFeed {
	items := []model.IFeed{}
	for _, feed := range feeds {
		items = append(items, feed)
	}
	return items
}

func (s feedService) getPagedResponse(path string) ([]model.IFeed, error) {
	resources := []*model.Feed{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Feeds), path)
		if err != nil {
			return toFeedArray(resources), err
		}

		responseList := resp.(*model.Feeds)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return toFeedArray(resources), nil
}

// Add creates a new feed.
func (s feedService) Add(feed model.IFeed) (model.IFeed, error) {
	if feed == nil {
		return nil, createInvalidParameterError(operationAdd, parameterFeed)
	}

	if isEmpty(feed.GetFeedType()) {
		return nil, createInvalidParameterError(operationAdd, parameterFeed)
	}

	path, err := getAddPath(s, feed)
	if err != nil {
		return nil, err
	}

	var feedResource interface{}
	switch feed.GetFeedType() {
	case feedAwsElasticContainerRegistry:
		feedResource = new(model.AwsElasticContainerRegistry)
	case feedBuiltIn:
		feedResource = new(model.BuiltInFeed)
	case feedDocker:
		feedResource = new(model.DockerContainerRegistry)
	case feedGitHub:
		feedResource = new(model.GitHubRepositoryFeed)
	case feedHelm:
		feedResource = new(model.HelmFeed)
	case feedMaven:
		feedResource = new(model.MavenFeed)
	case feedNuGet:
		feedResource = new(model.NuGetFeed)
	case feedOctopusProject:
		feedResource = new(model.OctopusProjectFeed)
	}

	resp, err := apiAdd(s.getClient(), feed, feedResource, path)
	if err != nil {
		return nil, err
	}

	return resp.(model.IFeed), nil
}

// GetAll returns all feeds. If none can be found or an error occurs, it
// returns an empty collection.
func (s feedService) GetAll() ([]model.IFeed, error) {
	path, err := getAllPath(s)
	if err != nil {
		return []model.IFeed{}, err
	}

	var feeds []*model.Feed
	_, err = apiGet(s.getClient(), &feeds, path)

	return toFeedArray(feeds), err
}

// GetByID returns the feed that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s feedService) GetByID(id string) (model.IFeed, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Feed), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(model.IFeed), nil
}

// GetByPartialName performs a lookup and returns instances of an feed with a
// matching partial name.
func (s feedService) GetByPartialName(name string) ([]model.IFeed, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []model.IFeed{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a feed based on the one provided as input.
func (s feedService) Update(feed model.IFeed) (model.IFeed, error) {
	if feed == nil {
		return nil, createInvalidParameterError(operationUpdate, parameterFeed)
	}

	path, err := getUpdatePath(s, feed)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), feed, new(model.Feed), path)
	if err != nil {
		return nil, err
	}

	return resp.(model.IFeed), nil
}
