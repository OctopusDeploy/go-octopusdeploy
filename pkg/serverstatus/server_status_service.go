package serverstatus

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
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

	response, err := api.ApiGet(s.GetClient(), new(ServerStatus), path)
	if err != nil {
		return nil, err
	}

	return response.(*ServerStatus), nil
}

var _ services.IService = &ServerStatusService{}

const (
	serverStatusTemplate   = "/api/serverstatus"
	healthStatusTemplate   = "/api/serverstatus/health"
	timezonesTemplate      = "/api/serverstatus/timezones"
	documentCountsTemplate = "/api/serverstatus/counts"
)

// GetServerStatus returns the status of the server.
func GetServerStatus(client newclient.Client) (*ServerStatus, error) {
	return newclient.Get[ServerStatus](client.HttpSession(), serverStatusTemplate)
}

// GetHealthStatus returns the health of the server cluster.
func GetHealthStatus(client newclient.Client) (*ServerHealthStatus, error) {
	return newclient.Get[ServerHealthStatus](client.HttpSession(), healthStatusTemplate)
}

// GetTimezones returns the time zones supported by the server.
func GetTimezones(client newclient.Client) ([]*Timezone, error) {
	timezones, err := newclient.Get[[]*Timezone](client.HttpSession(), timezonesTemplate)
	if err != nil {
		return nil, err
	}

	return *timezones, nil
}

// GetDocumentCounts returns the number of documents held by the server, grouped by area.
func GetDocumentCounts(client newclient.Client) (*DocumentCounts, error) {
	return newclient.Get[DocumentCounts](client.HttpSession(), documentCountsTemplate)
}
