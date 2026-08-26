package platformhubgithubconnections

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/githubconnections"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const (
	connectionsTemplate  = "/api/platformhub/githubconnections/connections{?skip,take}"
	connectionTemplate   = "/api/platformhub/githubconnections/connections/{id}"
	repositoriesTemplate = "/api/platformhub/githubconnections/connections/{connectionId}/repositories"
)

// ConnectionDetails represents a single Platform Hub GitHub App connection, including the
// repositories it grants access to which githubconnections.Connection doesn't include.
type ConnectionDetails struct {
	ID                  string                                 `json:"Id"`
	Status              githubconnections.ConnectionStatus     `json:"Status"`
	StatusUserMessage   string                                 `json:"StatusUserMessage,omitempty"`
	Installation        *githubconnections.Installation        `json:"Installation,omitempty"`
	Repositories        []*githubconnections.Repository        `json:"Repositories"`
	UnknownRepositories []*githubconnections.UnknownRepository `json:"UnknownRepositories"`
}

// ConnectionsQuery represents the query parameters for listing connections. Both
// skip and take are required by the server contract, so neither is omitted when empty.
type ConnectionsQuery struct {
	Skip int `uri:"skip" json:"skip"`
	Take int `uri:"take" json:"take"`
}

// ConnectionsQueryResult is a paginated collection of Platform Hub GitHub App connections.
type ConnectionsQueryResult struct {
	Connections   []*githubconnections.Connection `json:"Connections"`
	ItemsPerPage  int                             `json:"ItemsPerPage"`
	NumberOfPages int                             `json:"NumberOfPages"`
	TotalResults  int                             `json:"TotalResults"`
}

// List returns a single page of Platform Hub GitHub App connections.
func List(client newclient.Client, skip int, take int) (*ConnectionsQueryResult, error) {
	path, err := client.URITemplateCache().Expand(connectionsTemplate, ConnectionsQuery{Skip: skip, Take: take})
	if err != nil {
		return nil, err
	}

	return newclient.Get[ConnectionsQueryResult](client.HttpSession(), path)
}

// GetByID returns the Platform Hub GitHub App connection with the given ID, along with the
// repositories it grants access to.
func GetByID(client newclient.Client, id string) (*ConnectionDetails, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError("GetByID", "id")
	}

	path, err := client.URITemplateCache().Expand(connectionTemplate, map[string]any{"id": id})
	if err != nil {
		return nil, err
	}

	return newclient.Get[ConnectionDetails](client.HttpSession(), path)
}

type repositoriesResponse struct {
	Repositories []*githubconnections.Repository `json:"Repositories"`
}

// GetRepositories returns the GitHub repositories reachable through the given connection.
func GetRepositories(client newclient.Client, connectionID string) ([]*githubconnections.Repository, error) {
	if internal.IsEmpty(connectionID) {
		return nil, internal.CreateInvalidParameterError("GetRepositories", "connectionID")
	}

	path, err := client.URITemplateCache().Expand(repositoriesTemplate, map[string]any{"connectionId": connectionID})
	if err != nil {
		return nil, err
	}

	response, err := newclient.Get[repositoriesResponse](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return response.Repositories, nil
}
