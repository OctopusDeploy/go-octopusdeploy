package serverstatus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
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
	systemInfoTemplate     = "/api/serverstatus/system-info"
	systemReportTemplate   = "/api/serverstatus/system-report"
	recentLogsTemplate     = "/api/serverstatus/logs{?skip,take,includeDetail}"
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

// GetSystemInfo returns diagnostic information about the running server. Requires the
// AdministerSystem permission.
func GetSystemInfo(client newclient.Client) (*SystemInfo, error) {
	return newclient.Get[SystemInfo](client.HttpSession(), systemInfoTemplate)
}

// GetRecentLogs returns recent entries from the server log. Requires the AdministerSystem
// permission.
func GetRecentLogs(client newclient.Client, logsQuery LogsQuery) ([]*LogEntry, error) {
	values, ok := uritemplates.Struct2map(logsQuery)
	if !ok {
		values = map[string]any{}
	}

	path, err := client.URITemplateCache().Expand(recentLogsTemplate, values)
	if err != nil {
		return nil, err
	}

	logEntries, err := newclient.Get[[]*LogEntry](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return *logEntries, nil
}

// GetSystemReport returns the server's diagnostic report, a zip archive. The caller is
// responsible for closing the returned reader. Requires the AdministerSystem permission.
func GetSystemReport(client newclient.Client) (io.ReadCloser, error) {
	path, err := url.Parse(systemReportTemplate)
	if err != nil {
		return nil, err
	}

	request := &http.Request{
		Method: http.MethodGet,
		URL:    path,
		Header: make(http.Header),
	}

	response, err := client.HttpSession().DoRawRequest(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		defer newclient.CloseResponse(response)

		apiError := new(core.APIError)
		if err := json.NewDecoder(response.Body).Decode(apiError); err != nil {
			return nil, fmt.Errorf("cannot get system report from server. response from server %s", response.Status)
		}
		apiError.StatusCode = response.StatusCode

		return nil, apiError
	}

	return response.Body, nil
}
