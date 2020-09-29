package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type feedService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newFeedService(sling *sling.Sling, uriTemplate string) *feedService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &feedService{
		name:        serviceFeedService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s feedService) getClient() *sling.Sling {
	return s.sling
}

func (s feedService) getName() string {
	return s.name
}

func (s feedService) getPagedResponse(path string) ([]model.Feed, error) {
	resources := []model.Feed{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Feeds), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.Feeds)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

func (s feedService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// Add creates a new feed.
func (s feedService) Add(resource model.Feed) (*model.Feed, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.Feed), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Feed), nil
}

// DeleteByID deletes the feed that matches the input ID.
func (s feedService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

// GetByID returns the feed that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s feedService) GetByID(id string) (*model.Feed, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Feed), path)
	if err != nil {
		return nil, createResourceNotFoundError("feed", "ID", id)
	}

	return resp.(*model.Feed), nil
}

// GetAll returns all feeds. If none can be found or an error occurs, it
// returns an empty collection.
func (s feedService) GetAll() ([]model.Feed, error) {
	items := []model.Feed{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByPartialName performs a lookup and returns the Feed with a matching name.
func (s feedService) GetByPartialName(name string) ([]model.Feed, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []model.Feed{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a feed based on the one provided as input.
func (s feedService) Update(resource model.Feed) (*model.Feed, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.Feed), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Feed), nil
}

var _ ServiceInterface = &feedService{}
