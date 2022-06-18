package serverstatus

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type ServerStatusService struct {
	extensionStatsPath string
	healthStatusPath   string
	timezonesPath      string

	services.Service
}

func NewServerStatusService(sling *sling.Sling, uriTemplate string, extensionStatsPath string, healthStatusPath string, timezonesPath string) *ServerStatusService {
	return &ServerStatusService{
		extensionStatsPath: extensionStatsPath,
		healthStatusPath:   healthStatusPath,
		timezonesPath:      timezonesPath,
		Service:            services.NewService(constants.ServiceServerStatusService, sling, uriTemplate),
	}
}

// Get returns the status of the server.
func (s *ServerStatusService) Get() (*ServerStatus, error) {
	path, err := services.GetPath(s)
	if err != nil {
		return nil, err
	}

	response, err := services.ApiGet(s.GetClient(), new(ServerStatus), path)
	if err != nil {
		return nil, err
	}

	return response.(*ServerStatus), nil
}

var _ services.IService = &ServerStatusService{}
