package octopusdeploy

import "github.com/dghubble/sling"

type serverStatusService struct {
	extensionStatsPath string
	healthStatusPath   string
	timezonesPath      string

	service
}

func newServerStatusService(sling *sling.Sling, uriTemplate string, extensionStatsPath string, healthStatusPath string, timezonesPath string) *serverStatusService {
	return &serverStatusService{
		extensionStatsPath: extensionStatsPath,
		healthStatusPath:   healthStatusPath,
		timezonesPath:      timezonesPath,
		service:            newService(ServiceServerStatuService, sling, uriTemplate),
	}
}
