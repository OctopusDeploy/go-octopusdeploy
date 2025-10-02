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

// GetMissingPackages returns missing built-in feed packages associated with the release
func GetMissingPackages(client newclient.Client, release *Release) ([]MissingPackageInfo, error) {
	spaceID, err := internal.GetSpaceID(release.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	path, err := client.URITemplateCache().Expand(uritemplates.MissingPackagesForRelease, map[string]any{
		"spaceId":   spaceID,
		"releaseId": release.ID,
	})

	if err != nil {
		return nil, err
	}

	res, err := newclient.Get[MissingPackages](client.HttpSession(), path)
	if err != nil {
		return []MissingPackageInfo{}, err
	}

	return res.Packages, nil
}

func GetReleaseDeploymentTemplate(client newclient.Client, spaceID string, releaseID string) (*ReleaseDeploymentTemplate, error) {
	if client == nil {
		return nil, internal.CreateInvalidParameterError("GetReleaseDeploymentTemplate", "client")
	}
	if spaceID == "" {
		return nil, internal.CreateInvalidParameterError("GetReleaseDeploymentTemplate", "spaceID")
	}
	if releaseID == "" {
		return nil, internal.CreateInvalidParameterError("GetReleaseDeploymentTemplate", "releaseID")
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.ReleaseDeploymentTemplate, map[string]any{
		"spaceId":   spaceID,
		"releaseId": releaseID,
	})

	if err != nil {
		return nil, err
	}

	res, err := newclient.Get[ReleaseDeploymentTemplate](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ----- Experimental ---------------------------------------------------------

// GetReleasesInProjectChannel is EXPERIMENTAL
func GetReleasesInProjectChannel(client newclient.Client, spaceID string, projectID string, channelID string) ([]*Release, error) {
	if client == nil {
		return nil, internal.CreateInvalidParameterError("GetReleasesInProjectChannel", "client")
	}
	if projectID == "" {
		return nil, internal.CreateInvalidParameterError("GetReleasesInProjectChannel", "project")
	}
	if channelID == "" {
		return nil, internal.CreateInvalidParameterError("GetReleasesInProjectChannel", "channel")
	}
	if spaceID == "" {
		return nil, internal.CreateInvalidParameterError("GetReleasesInProjectChannel", "spaceID")
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.ReleasesByProjectAndChannel, map[string]any{
		"spaceId":   spaceID,
		"projectId": projectID,
		"channelId": channelID,
	})
	if err != nil {
		return nil, err
	}

	loadNextPage := true

	var allResults []*Release
	for loadNextPage { // can't stop the loop
		resp, err := newclient.Get[resources.Resources[*Release]](client.HttpSession(), expandedUri)
		if err != nil {
			return nil, err
		}

		allResults = append(allResults, resp.Items...)
		expandedUri, loadNextPage = services.LoadNextPage(resp.PagedResults)
	}

	return allResults, nil
}

// GetReleaseInProject looks up a single release in the given project
func GetReleaseInProject(client newclient.Client, spaceID string, projectID string, releaseVersion string) (*Release, error) {
	if client == nil {
		return nil, internal.CreateInvalidParameterError("GetReleasesForChannel", "client")
	}
	if spaceID == "" {
		return nil, internal.CreateInvalidParameterError("GetReleasesForChannel", "project")
	}
	if projectID == "" {
		return nil, internal.CreateInvalidParameterError("GetReleasesForChannel", "project")
	}
	if releaseVersion == "" {
		return nil, internal.CreateInvalidParameterError("GetReleasesForChannel", "channel")
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.ReleasesByProject, map[string]any{
		"spaceId":   spaceID,
		"projectId": projectID,
		"version":   releaseVersion,
	})
	if err != nil {
		return nil, err
	}
	return newclient.Get[Release](client.HttpSession(), expandedUri)
}
