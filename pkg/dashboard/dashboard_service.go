package dashboard

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

type DashboardService struct {
	dashboardDynamicPath string

	services.Service
}

const (
	dashboardDynamicTemplate = "/api/{spaceId}/dashboard/dynamic{?projects,environments,includePrevious}"
)


func NewDashboardService(sling *sling.Sling, uriTemplate string, dashboardDynamicPath string) *DashboardService {
	return &DashboardService{
		dashboardDynamicPath: dashboardDynamicPath,
		Service:              services.NewService(constants.ServiceDashboardService, sling, uriTemplate),
	}
}

func GetDynamicDashboardItem(client newclient.Client,spaceID string, dashboardQuery DashboardQuery) (*resources.Resources[*DashboardItem], error) {
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	values, _ := uritemplates.Struct2map(dashboardQuery)
	if values == nil {
		values = map[string]any{}
	}
	values["spaceId"] = spaceID

	expandedUri, err := client.URITemplateCache().Expand(dashboardDynamicTemplate, values)
	if err != nil {
		return nil, err
	}

	resp, err := newclient.Get[resources.Resources[*DashboardItem]](client.HttpSession(), expandedUri)
	if err != nil {
		return &resources.Resources[*DashboardItem]{}, err
	}

	return resp, nil
}
