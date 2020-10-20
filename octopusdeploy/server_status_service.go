package octopusdeploy

import "github.com/dghubble/sling"

type serverStatuService struct {
	extensionStatsPath string
	healthStatusPath   string
	timezonesPath      string

	service
}

func newServerStatuService(sling *sling.Sling, uriTemplate string, extensionStatsPath string, healthStatusPath string, timezonesPath string) *serverStatuService {
	return &serverStatuService{
		extensionStatsPath: extensionStatsPath,
		healthStatusPath:   healthStatusPath,
		timezonesPath:      timezonesPath,
		service:            newService(serviceServerStatuService, sling, uriTemplate, nil),
	}
}
