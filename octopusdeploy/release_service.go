package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

type releaseService struct {
	services.canDeleteService
}

func newReleaseService(sling *sling.Sling, uriTemplate string) *releaseService {
	releaseService := &releaseService{}
	releaseService.service = services.newService(ServiceReleaseService, sling, uriTemplate)

	return releaseService
}

// Add creates a new release.
func (s releaseService) Add(release *Release) (*Release, error) {
	path, err := getAddPath(s, release)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), release, new(Release), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Release), nil
}

// Get returns a collection of releases based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s releaseService) Get(releasesQuery ...ReleasesQuery) (*Releases, error) {
	v, _ := query.Values(releasesQuery[0])
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	resp, err := apiGet(s.getClient(), new(Releases), path)
	if err != nil {
		return &Releases{}, err
	}

	return resp.(*Releases), nil
}

// GetByID returns the release that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s releaseService) GetByID(id string) (*Release, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Release), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Release), nil
}

func (s deploymentService) GetDeployments(release *Release, deploymentQuery ...*DeploymentQuery) (*Deployments, error) {
	if release == nil {
		return nil, createInvalidParameterError(OperationGetDeployments, ParameterRelease)
	}

	uriTemplate, err := uritemplates.Parse(release.GetLinks()[linkDeployments])
	if err != nil {
		return &Deployments{}, err
	}

	values := make(map[string]interface{})
	path, err := uriTemplate.Expand(values)
	if err != nil {
		return &Deployments{}, err
	}

	if deploymentQuery != nil {
		path, err = uriTemplate.Expand(deploymentQuery[0])
		if err != nil {
			return &Deployments{}, err
		}
	}

	resp, err := apiGet(s.getClient(), new(Deployments), path)
	if err != nil {
		return &Deployments{}, err
	}

	return resp.(*Deployments), nil
}

func (s deploymentService) GetProgression(release *Release) (*Progression, error) {
	if release == nil {
		return nil, createInvalidParameterError(OperationGetDeployments, ParameterRelease)
	}

	path := release.GetLinks()[linkProgression]
	resp, err := apiGet(s.getClient(), new(Progression), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Progression), nil
}
