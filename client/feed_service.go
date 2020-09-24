package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type feedService struct {
	name        string                    `validate:"required"`
	path        string                    `validate:"required"`
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
		path:        strings.TrimSpace(uriTemplate),
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

func (s feedService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

func (s feedService) GetByID(id string) (*model.Feed, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Feed), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Feed), nil
}

// GetAll returns all instances of a Feed. If none can be found or an error occurs, it returns an empty collection.
func (s feedService) GetAll() ([]model.Feed, error) {
	items := new([]model.Feed)
	path, err := getAllPath(s)
	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.getClient(), items, path)
	return *items, err
}

// GetByName performs a lookup and returns the Feed with a matching name.
func (s feedService) GetByName(name string) ([]model.Feed, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []model.Feed{}, err
	}

	return s.getPagedResponse(path)
}

// Add creates a new Feed.
func (s feedService) Add(feed model.Feed) (*model.Feed, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	err = feed.Validate()

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)

	resp, err := apiAdd(s.getClient(), &feed, new(model.Feed), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Feed), nil
}

func (s feedService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

func (s feedService) Update(feed model.Feed) (*model.Feed, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	err = feed.Validate()

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)
	path = fmt.Sprintf(path+"/%s", feed.ID)

	resp, err := apiUpdate(s.getClient(), feed, new(model.Feed), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Feed), nil
}

func (s feedService) getPagedResponse(path string) ([]model.Feed, error) {
	items := []model.Feed{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Feeds), path)
		if err != nil {
			return nil, err
		}

		responseList := resp.(*model.Feeds)
		items = append(items, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return items, nil
}

var _ ServiceInterface = &feedService{}
