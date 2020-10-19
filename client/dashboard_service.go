package client

import "github.com/dghubble/sling"

type dashboardService struct {
	dashboardDynamicPath string

	service
}

func newDashboardService(sling *sling.Sling, uriTemplate string, dashboardDynamicPath string) *dashboardService {
	return &dashboardService{
		dashboardDynamicPath: dashboardDynamicPath,
		service:              newService(serviceDashboardService, sling, uriTemplate, nil),
	}
}
