package releases

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/channels"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
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
func (s *ReleaseService) Get(releasesQuery ...ReleasesQuery) (*Releases, error) {
	v, _ := query.Values(releasesQuery[0])
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	resp, err := services.ApiGet(s.GetClient(), new(Releases), path)
	if err != nil {
		return &Releases{}, err
	}

	return resp.(*Releases), nil
}

func (s *ReleaseService) CreateV1(createReleaseV1 *CreateReleaseV1) (*CreateReleaseResponseV1, error) {
	if createReleaseV1 == nil {
		return nil, internal.CreateInvalidParameterError("CreateV1", "createReleaseV1")
	}
	resp, err := services.ApiPost(s.GetClient(), createReleaseV1, new(CreateReleaseResponseV1), s.GetBasePath()+"/create/v1")
	if err != nil {
		return nil, err
	}

	return resp.(*CreateReleaseResponseV1), nil
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

	resp, err := services.ApiGet(s.GetClient(), new(Release), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Release), nil
}

func (s *ReleaseService) GetReleases(channel *channels.Channel, releaseQuery ...*ReleaseQuery) (*Releases, error) {
	if channel == nil {
		return nil, internal.CreateInvalidParameterError("GetReleases", "channel")
	}

	uriTemplate, err := uritemplates.Parse(channel.GetLinks()[constants.LinkReleases])
	if err != nil {
		return &Releases{}, err
	}

	values := make(map[string]interface{})
	path, err := uriTemplate.Expand(values)
	if err != nil {
		return &Releases{}, err
	}

	if releaseQuery != nil {
		path, err = uriTemplate.Expand(releaseQuery[0])
		if err != nil {
			return &Releases{}, err
		}
	}

	resp, err := services.ApiGet(s.GetClient(), new(Releases), path)
	if err != nil {
		return &Releases{}, err
	}

	return resp.(*Releases), nil
}
