package releases

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/channels"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

type ReleaseService struct {
	services.CanDeleteService
}

func NewReleaseService(sling *sling.Sling, uriTemplate string) *ReleaseService {
	return &ReleaseService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceReleaseService, sling, uriTemplate),
		},
	}
}

// Add creates a new release.
func (s *ReleaseService) Add(release *Release) (*Release, error) {
	if IsNil(release) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterRelease)
	}

	path, err := services.GetAddPath(s, release)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), release, new(Release), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Release), nil
}

// Get returns a collection of releases based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s *ReleaseService) Get(releasesQuery ...ReleasesQuery) (*resources.Resources[*Release], error) {
	v, _ := query.Values(releasesQuery[0])
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Release]), path)
	if err != nil {
		return &resources.Resources[*Release]{}, err
	}

	return resp.(*resources.Resources[*Release]), nil
}

// GetByID returns the release that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *ReleaseService) GetByID(id string) (*Release, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Release), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Release), nil
}

func (s *ReleaseService) GetReleases(channel *channels.Channel, releaseQuery ...*ReleaseQuery) (*resources.Resources[*Release], error) {
	if channel == nil {
		return nil, internal.CreateInvalidParameterError("GetReleases", "channel")
	}

	uriTemplate, err := uritemplates.Parse(channel.GetLinks()[constants.LinkReleases])
	if err != nil {
		return &resources.Resources[*Release]{}, err
	}

	values := make(map[string]interface{})
	path, err := uriTemplate.Expand(values)
	if err != nil {
		return &resources.Resources[*Release]{}, err
	}

	if releaseQuery != nil {
		path, err = uriTemplate.Expand(releaseQuery[0])
		if err != nil {
			return &resources.Resources[*Release]{}, err
		}
	}

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Release]), path)
	if err != nil {
		return &resources.Resources[*Release]{}, err
	}

	return resp.(*resources.Resources[*Release]), nil
}

// ----- Experimental ---------------------------------------------------------

// ResourcesOrError carries either a resources response or error
type ResourcesOrError[T any] struct {
	Response *resources.Resources[T]
	Error    error
}

type Disposable interface {
	Dispose()
}

type BooleanDisposable interface {
	Disposable
	IsDisposed() bool
}

type anonymousDisposable struct {
	isDisposed bool
	dispose    func()
}

// Dispose conforms our anonymousDisposable to Disposable
func (a *anonymousDisposable) Dispose() {
	a.isDisposed = true
	if a.dispose != nil {
		a.dispose()
	}
}

// IsDisposed conforms our anonymousDisposable to BooleanDisposable
func (a *anonymousDisposable) IsDisposed() bool {
	return a.isDisposed
}

func getReleasesInProjectChannel(client newclient.Client, projectID string, channelID string, disposable BooleanDisposable, dispatch func(*resources.Resources[*Release], error)) {
	if client == nil {
		dispatch(nil, internal.CreateInvalidParameterError("GetReleasesForChannel", "client"))
		return
	}
	if projectID == "" {
		dispatch(nil, internal.CreateInvalidParameterError("GetReleasesForChannel", "project"))
		return
	}
	if channelID == "" {
		dispatch(nil, internal.CreateInvalidParameterError("GetReleasesForChannel", "channel"))
		return
	}
	if client.SpaceID() == "" {
		dispatch(nil, internal.CreateInvalidClientStateError("GetReleasesForChannel"))
		return
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.ReleasesByProjectAndChannel, map[string]any{
		"spaceId":   client.SpaceID(),
		"projectId": projectID,
		"channelId": channelID,
	})
	if err != nil {
		dispatch(nil, err)
		return
	}

	loadNextPage := true

	for loadNextPage {
		rawResp, err := api.ApiGet(client.Sling(), new(resources.Resources[*Release]), expandedUri)
		if disposable != nil && disposable.IsDisposed() {
			return
		}
		if err != nil {
			dispatch(nil, err)
			return
		}

		resp := rawResp.(*resources.Resources[*Release])
		dispatch(resp, nil)

		expandedUri, loadNextPage = services.LoadNextPage(resp.PagedResults)
	}
}

// GetReleasesInProjectChannel is EXPERIMENTAL
func GetReleasesInProjectChannel(client newclient.Client, projectID string, channelID string) ([]*Release, error) {
	var results []*Release
	var errorResult error

	getReleasesInProjectChannel(client, projectID, channelID, nil, func(pageOfResults *resources.Resources[*Release], err error) {
		if err != nil {
			errorResult = err
		} else if pageOfResults != nil {
			results = append(results, pageOfResults.Items...)
		} // else if both are nil we have all the results, nothing to do here
	})
	return results, errorResult
}

// GetReleasesInProjectChannelAsync is EXPERIMENTAL. This is basically rx
func GetReleasesInProjectChannelAsync(client newclient.Client, projectID string, channelID string) (chan ResourcesOrError[*Release], Disposable) {
	result := make(chan ResourcesOrError[*Release])

	disposable := &anonymousDisposable{
		dispose: func() {
			close(result)
		},
	}

	go func() {
		getReleasesInProjectChannel(client, projectID, channelID, disposable, func(pageOfResults *resources.Resources[*Release], err error) {
			if err != nil {
				result <- ResourcesOrError[*Release]{Error: err}
			} else if pageOfResults != nil {
				result <- ResourcesOrError[*Release]{Response: pageOfResults}
			} // else if both are nil we have all the results, nothing to do here
		})
		disposable.Dispose()
	}()
	return result, disposable
}

// GetReleaseInProject looks up a single release in the given project
func GetReleaseInProject(client newclient.Client, projectID string, releaseVersion string) (*Release, error) {
	if client == nil {
		return nil, internal.CreateInvalidParameterError("GetReleasesForChannel", "client")
	}
	if projectID == "" {
		return nil, internal.CreateInvalidParameterError("GetReleasesForChannel", "project")
	}
	if releaseVersion == "" {
		return nil, internal.CreateInvalidParameterError("GetReleasesForChannel", "channel")
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.ReleasesByProject, map[string]any{
		"spaceId":   client.SpaceID(),
		"projectId": projectID,
		"version":   releaseVersion,
	})
	if err != nil {
		return nil, err
	}

	rawResp, err := api.ApiGet(client.Sling(), new(Release), expandedUri)
	if err != nil {
		return nil, err
	}

	return rawResp.(*Release), nil
}
