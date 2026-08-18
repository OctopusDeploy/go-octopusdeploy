package dashboard

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

type DashboardService struct {
	dashboardDynamicPath string

	services.Service
}

func NewDashboardService(sling *sling.Sling, uriTemplate string, dashboardDynamicPath string) *DashboardService {
	return &DashboardService{
		dashboardDynamicPath: dashboardDynamicPath,
		Service:              services.NewService(constants.ServiceDashboardService, sling, uriTemplate),
	}
}

// GetDynamicDashboard returns the release currently deployed to each
// environment, optionally narrowed to the given projects and environments, and
// optionally including the deployment preceding the current one.
//
// Passing a zero-valued query returns every project and environment the caller
// can see, subject to the server's project limit.
//
// The query filters on IDs, not names. The server does not reject a name, it
// simply matches nothing, so a name-filtered call returns an empty dashboard
// that is indistinguishable from nothing being deployed.
func (s *DashboardService) GetDynamicDashboard(query DashboardDynamicQuery) (*Dashboard, error) {
	path, err := s.getDynamicDashboardPath(query)
	if err != nil {
		return nil, err
	}

	response, err := api.ApiGet(s.GetClient(), new(Dashboard), path)
	if err != nil {
		return nil, err
	}

	return response.(*Dashboard), nil
}

// getDynamicDashboardPath expands the dynamic dashboard link template with the
// query. The dynamic dashboard has its own link rather than living under the
// service's URI template, so it is parsed here instead of via GetURITemplate.
func (s *DashboardService) getDynamicDashboardPath(query DashboardDynamicQuery) (string, error) {
	if internal.IsEmpty(s.dashboardDynamicPath) {
		return "", internal.CreateInvalidParameterError(constants.OperationGet, "dashboardDynamicPath")
	}

	template, err := uritemplates.Parse(s.dashboardDynamicPath)
	if err != nil {
		return "", err
	}

	return template.Expand(query)
}
