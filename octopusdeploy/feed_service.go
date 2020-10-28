package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
	"github.com/jinzhu/copier"
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

func toFeed(feedResource *FeedResource) (IFeed, error) {
	if isNil(feedResource) {
		return nil, createInvalidParameterError("toFeed", "feedResource")
	}

	var feed IFeed
	var err error
	switch feedResource.GetFeedType() {
	case FeedTypeAwsElasticContainerRegistry:
		feed = NewAwsElasticContainerRegistry(feedResource.GetName(), feedResource.AccessKey, feedResource.SecretKey, feedResource.Region)
	case FeedTypeBuiltIn:
		feed = NewBuiltInFeed(feedResource.GetName(), feedResource.FeedURI)
	case FeedTypeDocker:
		feed = NewDockerContainerRegistry(feedResource.GetName())
	case FeedTypeGitHub:
		feed = NewGitHubRepositoryFeed(feedResource.GetName())
	case FeedTypeHelm:
		feed = NewHelmFeed(feedResource.GetName())
	case FeedTypeMaven:
		feed = NewMavenFeed(feedResource.GetName())
	case FeedTypeNuGet:
		feed = NewNuGetFeed(feedResource.GetName(), feedResource.FeedURI)
	case FeedTypeOctopusProject:
		feed = NewOctopusProjectFeed(feedResource.GetName(), feedResource.FeedURI)
	}

	err = copier.Copy(feed, feedResource)
	if err != nil {
		return nil, err
	}

	return feed, nil
}

func toFeeds(feedResources *FeedResources) *Feeds {
	return &Feeds{
		Items:        toFeedArray(feedResources.Items),
		PagedResults: feedResources.PagedResults,
	}
}

func toFeedResource(feed IFeed) (*FeedResource, error) {
	if isNil(feed) {
		return nil, createInvalidParameterError("toFeedResource", ParameterFeed)
	}

	feedResource := NewFeedResource(feed.GetName(), feed.GetFeedType())

	err := copier.Copy(&feedResource, feed)
	if err != nil {
		return nil, err
	}

	return feedResource, nil
}

func toFeedArray(feedResources []*FeedResource) []IFeed {
	items := []IFeed{}
	for _, feedResource := range feedResources {
		feed, err := toFeed(feedResource)
		if err != nil {
			return nil
		}
		items = append(items, feed)
	}
	return items
}

// Add creates a new feed.
func (s feedService) Add(feed IFeed) (IFeed, error) {
	if feed == nil {
		return nil, createInvalidParameterError(OperationAdd, ParameterFeed)
	}

	feedResource, err := toFeedResource(feed)
	if err != nil {
		return nil, err
	}

	response, err := apiAdd(s.getClient(), feedResource, new(FeedResource), s.BasePath)
	if err != nil {
		return nil, err
	}

	return toFeed(response.(*FeedResource))
}

// Get returns a collection of feeds based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s feedService) Get(feedsQuery FeedsQuery) (*Feeds, error) {
	v, _ := query.Values(feedsQuery)
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := apiGet(s.getClient(), new(FeedResources), path)
	if err != nil {
		return &Feeds{}, err
	}

	return toFeeds(response.(*FeedResources)), nil
}

// GetAll returns all feeds. If none can be found or an error occurs, it
// returns an empty collection.
func (s feedService) GetAll() ([]IFeed, error) {
	items := []*FeedResource{}
	path := s.BasePath + "/all"

	_, err := apiGet(s.getClient(), &items, path)
	return toFeedArray(items), err
}

// GetByID returns the feed that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s feedService) GetByID(id string) (IFeed, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError(OperationGetByID, ParameterID)
	}

	path := s.BasePath + "/" + id
	resp, err := apiGet(s.getClient(), new(FeedResource), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(IFeed), nil
}

// GetBuiltInFeedStatistics returns statistics for the built-in feeds.
func (s feedService) GetBuiltInFeedStatistics() (*BuiltInFeedStatistics, error) {
	path := s.builtInFeedStats
	resp, err := apiGet(s.getClient(), new(BuiltInFeedStatistics), path)
	if err != nil {
		return nil, err
	}

	return resp.(*BuiltInFeedStatistics), nil
}

func (s feedService) SearchPackages(feed IFeed, searchPackagesQuery ...SearchPackagesQuery) (*PackageDescriptions, error) {
	if feed == nil {
		return nil, createInvalidParameterError(OperationSearchPackages, ParameterFeed)
	}

	uriTemplate, err := uritemplates.Parse(feed.GetLinks()[linkSearchPackagesTemplate])
	if err != nil {
		return &PackageDescriptions{}, err
	}

	values := make(map[string]interface{})
	path, err := uriTemplate.Expand(values)
	if err != nil {
		return &PackageDescriptions{}, err
	}

	if searchPackagesQuery != nil {
		path, err = uriTemplate.Expand(searchPackagesQuery[0])
		if err != nil {
			return &PackageDescriptions{}, err
		}
	}

	resp, err := apiGet(s.getClient(), new(PackageDescriptions), path)
	if err != nil {
		return &PackageDescriptions{}, err
	}

	return resp.(*PackageDescriptions), nil
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

	feedResource, err := toFeedResource(feed)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), feedResource, new(FeedResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(IFeed), nil
}
