package service

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/dghubble/sling"
)

var version = "development"

type IClient interface {
	getSling() *sling.Sling
	getScopedBasePath() string
	getRequestingTool() *string
}

type client struct {
	sling                 *sling.Sling
	octopusServerEndpoint *octopusServerEndpoint
	scopedBasePath        string
	requestingTool        *string
	IClient
}

type IAdminClient interface {
	IClient
}

type spaceScopedClient struct {
	spaceID string
	IClient
}

type ISpaceScopedClient interface {
	IClient
	GetSpaceID() string
}

func newClient(octopusServerEndpoint *octopusServerEndpoint, requestingTool *string, pathScope string) IClient {
	httpClient := &http.Client{}
	scopedBasePath := octopusServerEndpoint.BaseURLWithAPI.String()
	if !internal.IsEmpty(pathScope) {
		scopedBasePath = fmt.Sprintf("%s/%s", scopedBasePath, pathScope)
	}
	base := sling.New().Client(httpClient).Base(scopedBasePath).Set(clientAPIKeyHTTPHeader, octopusServerEndpoint.ApiKey)
	base.Set("User-Agent", getUserAgentString(requestingTool))

	c := &client{
		sling:                 base,
		octopusServerEndpoint: octopusServerEndpoint,
		scopedBasePath:        scopedBasePath,
		requestingTool:        requestingTool,
	}
	return c
}

func (c client) getSling() *sling.Sling {
	return c.sling
}

func (c client) getScopedBasePath() string {
	return c.scopedBasePath
}

func (c client) getRequestingTool() *string {
	return c.requestingTool
}

// NewAdminClient returns a new Octopus API Client.
func NewAdminClient(octopusServerEndpoint *octopusServerEndpoint, requestingTool *string) (IAdminClient, error) {
	if octopusServerEndpoint == nil {
		return nil, internal.CreateInvalidParameterError("NewAdminClient", octopusdeploy.ParameterOctopusServerEndpoint)
	}

	c := newClient(octopusServerEndpoint, requestingTool, internal.Empty)

	return c, nil
}

func NewSpaceScopedClient(octopusServerEndpoint *octopusServerEndpoint, spaceID string, requestingTool *string) (ISpaceScopedClient, error) {
	if octopusServerEndpoint == nil {
		return nil, internal.CreateInvalidParameterError("NewSpaceScopedClient", octopusdeploy.ParameterOctopusServerEndpoint)
	}

	client := &spaceScopedClient{
		IClient: newClient(octopusServerEndpoint, requestingTool, spaceID),
		spaceID: spaceID,
	}

	return client, nil
}

func (s spaceScopedClient) GetSpaceID() string {
	return s.spaceID
}

// APIError is a generic structure for containing errors for API operations.
type APIError struct {
	Details         string   `json:"Details,omitempty"`
	ErrorMessage    string   `json:"ErrorMessage,omitempty"`
	Errors          []string `json:"Errors,omitempty"`
	FullException   string   `json:"FullException,omitempty"`
	HelpLink        string   `json:"HelpLink,omitempty"`
	HelpText        string   `json:"HelpText,omitempty"`
	ParsedHelpLinks []string `json:"ParsedHelpLinks,omitempty"`
	StatusCode      int
}

// Error creates a predefined error for Octopus API responses.
func (e APIError) Error() string {
	return fmt.Sprintf("Octopus API error: %v %+v %v", e.ErrorMessage, e.Errors, e.FullException)
}

// APIErrorChecker is a generic error handler for the OctopusDeploy API.
func APIErrorChecker(urlPath string, resp *http.Response, wantedResponseCode int, slingError error, octopusDeployError *APIError) error {
	if octopusDeployError.Errors != nil {
		return fmt.Errorf("octopus deploy api returned an error on endpoint %s - %s", urlPath, octopusDeployError.Errors)
	}

	if slingError != nil {
		return fmt.Errorf("cannot get endpoint %s from server. failure from http Client %v", urlPath, slingError)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		octopusDeployError.StatusCode = resp.StatusCode
		return octopusDeployError
	}

	if resp.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("bad request from endpoint %s. response from server %s", urlPath, resp.Status)
	}

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	if resp.StatusCode != wantedResponseCode {
		return octopusDeployError
	}

	return nil
}

func ApiGetMany[T resources.IResource](c IClient, pathRelativeToRoot string) (*IPagedResultsHandler[T], error) {
	resp, err := ApiGet[IPagedResultsHandler[T]](c, pathRelativeToRoot)
	return resp, err
}

func ApiGet[T any](c IClient, pathRelativeToRoot string) (*T, error) {
	getClient := c.getSling().New()
	if getClient == nil {
		return nil, internal.CreateClientInitializationError(OperationAPIGet)
	}

	resource := new(T)
	path := fmt.Sprintf("%s/%s", c.getScopedBasePath(), pathRelativeToRoot)
	getClient = getClient.Get(path)
	if getClient == nil {
		return nil, internal.CreateClientInitializationError(OperationAPIGet)
	}

	getClient.Set("User-Agent", getUserAgentString(c.getRequestingTool()))

	octopusDeployError := new(APIError)
	resp, err := getClient.Receive(resource, &octopusDeployError)
	// if err != nil {
	// 	return nil, err
	// }

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)
	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// Returns the User-Agent String "go-octopusdeploy/version (os; arch) go/version"
func getUserAgentString(requestingTool *string) string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			// TODO: use go 1.18 dependency check
			if dep.Path == "github.com/OctopusDeploy/go-octopusdeploy" {
				if dep.Version != "" {
					version = dep.Version
				}
			}
		}
	}
	requestingToolPart := ""
	if requestingTool != nil {
		// prepend a space to the requesting tool so the user agent parsing works correctly later
		requestingToolPart = fmt.Sprintf(" %s", requestingTool)
	}
	return fmt.Sprintf("%s/%s (%s; %s) go/%s%s", "go-octopusdeploy", version, runtime.GOOS, runtime.GOARCH, runtime.Version(), requestingToolPart)
}

// Generic OctopusDeploy API Add Function. Expects a 201 response.
func ApiAdd[T resources.IResource](c IClient, inputStruct *T, path string) (*T, error) {
	return apiAddWithResponseStatus[T](c, inputStruct, path, http.StatusCreated)
}

// apiAddWithResponseStatus function with defined response.
func apiAddWithResponseStatus[T any](c IClient, inputStruct *T, path string, httpStatus int) (*T, error) {
	if internal.IsEmpty(path) {
		return nil, internal.CreateInvalidParameterError(OperationAPIAdd, octopusdeploy.ParameterPath)
	}

	postClient := c.getSling().New()
	if postClient == nil {
		return nil, internal.CreateClientInitializationError(OperationAPIAdd)
	}

	postClient = postClient.Post(path)
	if postClient == nil {
		return nil, internal.CreateClientInitializationError(OperationAPIAdd)
	}

	postClient.Set("User-Agent", getUserAgentString(c.getRequestingTool()))

	request := postClient.BodyJSON(inputStruct)
	if request == nil {
		return nil, internal.CreateClientInitializationError(OperationAPIAdd)
	}

	octopusDeployError := new(APIError)
	resource := new(T)
	resp, err := request.Receive(resource, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, httpStatus, err, octopusDeployError)
	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// apiPost post to octopus and expect a 200 response code.
func (c client) apiPost(inputStruct interface{}, resource interface{}, path string) (interface{}, error) {
	if internal.IsEmpty(path) {
		return nil, internal.CreateInvalidParameterError(OperationAPIPost, octopusdeploy.ParameterPath)
	}

	postClient := c.sling.New()
	if postClient == nil {
		return nil, internal.CreateClientInitializationError(OperationAPIPost)
	}

	postClient = postClient.Post(path)
	if postClient == nil {
		return nil, internal.CreateClientInitializationError(OperationAPIPost)
	}

	postClient.Set("User-Agent", getUserAgentString(c.requestingTool))

	request := postClient.BodyJSON(inputStruct)
	if request == nil {
		return nil, internal.CreateClientInitializationError(OperationAPIPost)
	}

	octopusDeployError := new(APIError)
	resp, err := request.Receive(resource, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)
	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// Generic OctopusDeploy API Update Function.
func ApiUpdate[T resources.IResource](c IClient, inputStruct *T, path string) (*T, error) {
	putClient := c.getSling().New()
	if putClient == nil {
		return nil, internal.CreateClientInitializationError(OperationAPIUpdate)
	}

	resource := new(T)
	putClient = putClient.Put(path)
	if putClient == nil {
		return nil, internal.CreateClientInitializationError(OperationAPIUpdate)
	}

	putClient.Set("User-Agent", getUserAgentString(c.getRequestingTool()))

	request := putClient.BodyJSON(inputStruct)
	if request == nil {
		return nil, internal.CreateClientInitializationError(OperationAPIUpdate)
	}

	octopusDeployError := new(APIError)
	resp, err := request.Receive(resource, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)
	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// Generic OctopusDeploy API Delete Function.
func ApiDelete[T resources.IResource](c IClient, id string, path string) error {
	if internal.IsEmpty(id) {
		return internal.CreateInvalidParameterError(OperationAPIDelete, octopusdeploy.ParameterID)
	}

	deleteClient := c.getSling().New()
	if deleteClient == nil {
		return internal.CreateClientInitializationError(OperationAPIDelete)
	}

	deleteClient = deleteClient.Delete(path)
	if deleteClient == nil {
		return internal.CreateClientInitializationError(OperationAPIDelete)
	}

	octopusDeployError := new(APIError)
	resp, err := deleteClient.Receive(nil, &octopusDeployError)

	return APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)
}

// ErrItemNotFound is an OctopusDeploy error returned an item cannot be found.
var ErrItemNotFound = errors.New("cannot find the item")
