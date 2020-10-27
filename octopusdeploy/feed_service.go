package octopusdeploy

import (
	"github.com/dghubble/sling"
)

// feedService handles communication with feed-related methods of the Octopus
// API.
type feedService struct {
	builtInFeedStats string

	canDeleteService
}

// newFeedService returns an feed service with a preconfigured client.
func newFeedService(sling *sling.Sling, uriTemplate string, builtInFeedStats string) *feedService {
	feedService := &feedService{
		builtInFeedStats: builtInFeedStats,
	}
	feedService.service = newService(ServiceFeedService, sling, uriTemplate)

	return feedService
}

func toFeedArray(feeds []*Feed) []IFeed {
	items := []IFeed{}
	for _, feed := range feeds {
		items = append(items, feed)
	}
	return items
}

func (s feedService) getPagedResponse(path string) ([]IFeed, error) {
	resources := []*Feed{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Feeds), path)
		if err != nil {
			return toFeedArray(resources), err
		}

		responseList := resp.(*Feeds)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return toFeedArray(resources), nil
}

// Add creates a new feed.
func (s feedService) Add(feed IFeed) (IFeed, error) {
	if feed == nil {
		return nil, createInvalidParameterError(OperationAdd, ParameterFeed)
	}

	if isEmpty(feed.GetFeedType()) {
		return nil, createInvalidParameterError(OperationAdd, ParameterFeed)
	}

	path, err := getAddPath(s, feed)
	if err != nil {
		return nil, err
	}

	var feedResource interface{}
	switch feed.GetFeedType() {
	case feedAwsElasticContainerRegistry:
		feedResource = new(AwsElasticContainerRegistry)
	case feedBuiltIn:
		feedResource = new(BuiltInFeed)
	case feedDocker:
		feedResource = new(DockerContainerRegistry)
	case feedGitHub:
		feedResource = new(GitHubRepositoryFeed)
	case feedHelm:
		feedResource = new(HelmFeed)
	case feedMaven:
		feedResource = new(MavenFeed)
	case feedNuGet:
		feedResource = new(NuGetFeed)
	case feedOctopusProject:
		feedResource = new(OctopusProjectFeed)
	}

	resp, err := apiAdd(s.getClient(), feed, feedResource, path)
	if err != nil {
		return nil, err
	}

	return resp.(IFeed), nil
}

// GetAll returns all feeds. If none can be found or an error occurs, it
// returns an empty collection.
func (s feedService) GetAll() ([]IFeed, error) {
	path, err := getAllPath(s)
	if err != nil {
		return []IFeed{}, err
	}

	var feeds []*Feed
	_, err = apiGet(s.getClient(), &feeds, path)

	return toFeedArray(feeds), err
}

// GetByID returns the feed that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s feedService) GetByID(id string) (IFeed, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Feed), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(IFeed), nil
}

// GetByPartialName performs a lookup and returns instances of an feed with a
// matching partial name.
func (s feedService) GetByPartialName(name string) ([]IFeed, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []IFeed{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a feed based on the one provided as input.
func (s feedService) Update(feed IFeed) (IFeed, error) {
	if feed == nil {
		return nil, createInvalidParameterError(OperationUpdate, ParameterFeed)
	}

	path, err := getUpdatePath(s, feed)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), feed, new(Feed), path)
	if err != nil {
		return nil, err
	}

	return resp.(IFeed), nil
}
