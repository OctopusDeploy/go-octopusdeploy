package dashboard

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

type DashboardService struct {
	dashboardDynamicUriTemplate *uritemplates.UriTemplate

	services.Service
}

func NewDashboardService(sling *sling.Sling, uriTemplate string, dashboardDynamicUriTemplate string) *DashboardService {
	dynamicTemplate, _ := uritemplates.Parse(strings.TrimSpace(dashboardDynamicUriTemplate))
	return &DashboardService{
		Service:                     services.NewService(constants.ServiceDashboardService, sling, uriTemplate),
		dashboardDynamicUriTemplate: dynamicTemplate,
	}
}

func (s *DashboardService) GetDashboard() (*Dashboard, error) {
	values := make(map[string]interface{})
	path, err := s.GetURITemplate().Expand(values)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(*Dashboard), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Dashboard), nil
}

func (s *DashboardService) GetDynamicDashboard(query *DashboardDynamicQuery) (*Dashboard, error) {
	path, err := s.dashboardDynamicUriTemplate.Expand(query)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(*Dashboard), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Dashboard), nil
}
