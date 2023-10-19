package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/packages"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

// FeedService handles communication with feed-related methods of the Octopus
// API.
type FeedService struct {
	builtInFeedStats string

	services.CanDeleteService
}

// NewFeedService returns an feed service with a preconfigured client.
func NewFeedService(sling *sling.Sling, uriTemplate string, builtInFeedStats string) *FeedService {
	return &FeedService{
		builtInFeedStats: builtInFeedStats,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceFeedService, sling, uriTemplate),
		},
	}
}

// Add creates a new feed.
//
// Deprecated: use feeds.Add
func (s *FeedService) Add(feed IFeed) (IFeed, error) {
	if IsNil(feed) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterFeed)
	}

	response, err := services.ApiAdd(s.GetClient(), feed, feed, s.BasePath)
	if err != nil {
		return nil, err
	}

	return response.(IFeed), nil
}

// Get returns a collection of feeds based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
//
// Deprecated: use feeds.Get
func (s *FeedService) Get(feedsQuery FeedsQuery) (*Feeds, error) {
	// TODO this method is wired for /api/Spaces-1/feeds?ids=feeds-builtin
	// but the server also supports a simpler single-value at /api/Spaces-1/feeds/feeds-builtin
	// we should support that too.

	v, _ := query.Values(feedsQuery)
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*FeedResource]), path)
	if err != nil {
		return &Feeds{}, err
	}

	return ToFeeds(response.(*resources.Resources[*FeedResource])), nil
}

// GetAll returns all feeds. If none can be found or an error occurs, it
// returns an empty collection.
func (s *FeedService) GetAll() ([]IFeed, error) {
	items := []*FeedResource{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return ToFeedArray(items), err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return ToFeedArray(items), err
}

// GetByID returns the feed that matches the input ID. If one cannot be found,
// it returns nil and an error.
//
// Deprecated: use feeds.GetByID
func (s *FeedService) GetByID(id string) (IFeed, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, "id")
	}

	path := s.BasePath + "/" + id
	resp, err := api.ApiGet(s.GetClient(), new(FeedResource), path)
	if err != nil {
		return nil, err
	}

	return ToFeed(resp.(*FeedResource))
}

// GetBuiltInFeedStatistics returns statistics for the built-in feeds.
func (s *FeedService) GetBuiltInFeedStatistics() (*BuiltInFeedStatistics, error) {
	path := s.builtInFeedStats
	resp, err := api.ApiGet(s.GetClient(), new(BuiltInFeedStatistics), path)
	if err != nil {
		return nil, err
	}

	return resp.(*BuiltInFeedStatistics), nil
}

// TODO remove or rename this method in API Client v3; the first parameter wants to be an IFeed, not a PackageDescription
func (s *FeedService) SearchPackageVersions(packageDescription *packages.PackageDescription, searchPackageVersionsQuery SearchPackageVersionsQuery) (*resources.Resources[*packages.PackageVersion], error) {
	if packageDescription == nil {
		return nil, internal.CreateInvalidParameterError("SearchPackageVersions", "packageDescription")
	}

	uriTemplate, err := uritemplates.Parse(packageDescription.GetLinks()[constants.LinkSearchPackageVersionsTemplate])
	if err != nil {
		return &resources.Resources[*packages.PackageVersion]{}, err
	}

	path, err := uriTemplate.Expand(searchPackageVersionsQuery)
	if err != nil {
		return &resources.Resources[*packages.PackageVersion]{}, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*packages.PackageVersion]), path)
	if err != nil {
		return &resources.Resources[*packages.PackageVersion]{}, err
	}

	return resp.(*resources.Resources[*packages.PackageVersion]), nil
}

// TODO this method should be called SearchFeedPackageVersions for consistency, but that would be a breaking change in v2 of the client; defer to v3
func (s *FeedService) SearchFeedPackageVersions(feed IFeed, searchPackageVersionsQuery SearchPackageVersionsQuery) (*resources.Resources[*packages.PackageVersion], error) {
	if feed == nil {
		return nil, internal.CreateInvalidParameterError("SearchFeedPackageVersions", "feed")
	}

	uriTemplate, err := uritemplates.Parse(feed.GetLinks()[constants.LinkSearchPackageVersionsTemplate])
	if err != nil {
		return nil, err
	}

	path, err := uriTemplate.Expand(searchPackageVersionsQuery)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*packages.PackageVersion]), path)
	if err != nil {
		return nil, err
	}

	return resp.(*resources.Resources[*packages.PackageVersion]), nil
}

func (s *FeedService) SearchPackages(feed IFeed, searchPackagesQuery SearchPackagesQuery) (*resources.Resources[*packages.PackageDescription], error) {
	if feed == nil {
		return nil, internal.CreateInvalidParameterError("SearchPackages", "feed")
	}

	uriTemplate, err := uritemplates.Parse(feed.GetLinks()[constants.LinkSearchPackagesTemplate])
	if err != nil {
		return &resources.Resources[*packages.PackageDescription]{}, err
	}

	path, err := uriTemplate.Expand(searchPackagesQuery)
	if err != nil {
		return &resources.Resources[*packages.PackageDescription]{}, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*packages.PackageDescription]), path)
	if err != nil {
		return &resources.Resources[*packages.PackageDescription]{}, err
	}

	return resp.(*resources.Resources[*packages.PackageDescription]), nil
}

// Update modifies a feed based on the one provided as input.
//
// Deprecated: use feeds.Update
func (s *FeedService) Update(feed IFeed) (IFeed, error) {
	if feed == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, "feed")
	}

	path, err := services.GetUpdatePath(s, feed)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), feed, feed, path)
	if err != nil {
		return nil, err
	}

	return resp.(IFeed), nil
}

// --- new ---

const template = "/api/{spaceId}/feeds{/id}{?skip,take,ids,partialName,feedType,name}"

// note, the FeedID here has to be a real ID. You can't supply "feeds-builtin", you need to lookup the ID as "Feeds-101" etc, and use that instead
func SearchPackageVersions(client newclient.Client, spaceID string, feedID string, packageID string, filter string, limit int) (*resources.Resources[*packages.PackageVersion], error) {
	if spaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	if feedID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("feedID")
	}
	if packageID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("packageID")
	}
	templateParams := map[string]any{"spaceId": spaceID, "feedId": feedID, "packageId": packageID}
	if filter != "" {
		templateParams["filter"] = filter
	}
	if limit > 0 {
		templateParams["take"] = limit
	}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.FeedSearchPackageVersions, templateParams)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return newclient.Get[resources.Resources[*packages.PackageVersion]](client.HttpSession(), expandedUri)
}

// Add creates a new feed.
func Add(client newclient.Client, feed IFeed) (IFeed, error) {
	res, err := newclient.Add[FeedResource](client, template, feed.GetSpaceID(), feed)
	if err != nil {
		return nil, err
	}
	return ToFeed(res)
}

// Get returns a collection of feeds based on the criteria defined by its
// input query parameter.
func Get(client newclient.Client, spaceID string, feedsQuery FeedsQuery) (*Feeds, error) {
	// TODO this method is wired for /api/Spaces-1/feeds?ids=feeds-builtin
	// but the server also supports a simpler single-value at /api/Spaces-1/feeds/feeds-builtin
	// we should support that too.
	res, err := newclient.GetByQuery[FeedResource](client, template, spaceID, feedsQuery)
	if err != nil {
		return nil, err
	}
	return ToFeeds(res), nil
}

// GetByID returns the feed that matches the input ID. If one cannot be found,
// it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, id string) (IFeed, error) {
	res, err := newclient.GetByID[FeedResource](client, template, spaceID, id)
	if err != nil {
		return nil, err
	}
	return ToFeed(res)
}

// Update modifies a feed based on the one provided as input.
func Update(client newclient.Client, feed IFeed) (IFeed, error) {
	feedResource, err := ToFeedResource(feed)
	if err != nil {
		return nil, err
	}

	res, err := newclient.Add[FeedResource](client, template, feed.GetSpaceID(), feedResource)
	if err != nil {
		return nil, err
	}

	return ToFeed(res)
}

// DeleteByID will delete a account with the provided id.
func DeleteByID(client newclient.Client, spaceID string, id string) error {
	return newclient.DeleteByID(client, template, spaceID, id)
}

// GetAll returns all feeds. If an error occurs, it returns nil.
func GetAll(client newclient.Client, spaceID string) ([]IFeed, error) {
	items, err := newclient.GetAll[FeedResource](client, template, spaceID)
	return ToFeedArray(items), err
}
