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
		service:            newService(ServiceServerStatusService, sling, uriTemplate),
	}
}

// Get returns the status of the server.
func (s serverStatusService) Get() (*ServerStatus, error) {
	path, err := getPath(s)
	if err != nil {
		return nil, err
	}

	response, err := apiGet(s.getClient(), new(ServerStatus), path)
	if err != nil {
		return nil, err
	}

	return response.(*ServerStatus), nil
}

var _ IService = &serverStatusService{}
