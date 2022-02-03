package octopusdeploy

import (
	"errors"
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
	"net/http"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
)

var version = "development"

type Client interface {
	apiGetByID(resource interface{}, id string) (interface{}, error)
	apiQuery(inputStruct interface{}, template *uritemplates.UriTemplate) (interface{}, error)
	apiAdd(inputStruct interface{}, resource interface{}) (interface{}, error)
	apiUpdate(inputStruct interface{}, resource interface{}) (interface{}, error)
	apiDelete(resource interface{}, id string) error
}

type AdminClient interface {
	Client
}

type SpaceScopedClient interface {
	GetSpaceID() string
	Client
}

type client struct {
	sling                 *sling.Sling
	octopusServerEndpoint *octopusServerEndpoint
	scopedBasePath        string
	requestingTool        string
	Client
}

// Client is an OctopusDeploy for making Octopus API requests.
type adminClient struct {
	client
	AdminClient
}

func newClient(octopusServerEndpoint *octopusServerEndpoint, requestingTool string, pathScope string) client {
	httpClient := &http.Client{}
	scopedBasePath := octopusServerEndpoint.BaseURLWithAPI.String()
	if !isEmpty(pathScope) {
		scopedBasePath = fmt.Sprintf("%s/%s", scopedBasePath, pathScope)
	}
	base := sling.New().Client(httpClient).Base(scopedBasePath).Set(clientAPIKeyHTTPHeader, octopusServerEndpoint.ApiKey)
	base.Set("User-Agent", getUserAgentString(requestingTool))

	c := client{
		sling:                 base,
		octopusServerEndpoint: octopusServerEndpoint,
		scopedBasePath:        scopedBasePath,
		requestingTool:        requestingTool,
	}
	return c
}

// NewAdminClient returns a new Octopus API client.
func NewAdminClient(octopusServerEndpoint *octopusServerEndpoint, requestingTool string) (AdminClient, error) {
	if octopusServerEndpoint == nil {
		return nil, createInvalidParameterError("NewAdminClient", ParameterOctopusServerEndpoint)
	}

	c := adminClient{
		client: newClient(octopusServerEndpoint, requestingTool, emptyString),
	}

	return &c, nil
}

type spaceScopedClient struct {
	spaceID string
	client
	SpaceScopedClient
}

func NewSpaceScopedClient(octopusServerEndpoint *octopusServerEndpoint, spaceID string, requestingTool string) (SpaceScopedClient, error) {
	if octopusServerEndpoint == nil {
		return nil, createInvalidParameterError("NewSpaceScopedClient", ParameterOctopusServerEndpoint)
	}

	httpClient := &http.Client{}
	base := sling.New().Client(httpClient).Base(octopusServerEndpoint.BaseURLWithAPI.String()).Set(clientAPIKeyHTTPHeader, octopusServerEndpoint.ApiKey)
	base.Set("User-Agent", getUserAgentString(requestingTool))

	client := spaceScopedClient{
		client:  newClient(octopusServerEndpoint, requestingTool, spaceID),
		spaceID: spaceID,
	}

	return &client, nil
}

func (c spaceScopedClient) GetSpaceID() string {
	return c.spaceID
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
		return fmt.Errorf("cannot get endpoint %s from server. failure from http client %v", urlPath, slingError)
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

// LoadNextPage checks if the next page should be loaded from the API. Returns
// the new path and a bool if the next page should be checked.
func LoadNextPage(pagedResults PagedResults) (string, bool) {
	if pagedResults.Links.PageNext != emptyString {
		return pagedResults.Links.PageNext, true
	}

	return emptyString, false
}

func pluralize(resource interface{}) string {
	paged := resource.(*PagedResults)
	if paged != nil {
		return strings.ToLower(reflect.TypeOf(resource).Name())
	}
	resourceName := strings.ToLower(strings.ReplaceAll(reflect.TypeOf(resource).Name(), "Resource", ""))
	return fmt.Sprintf("%ss", resourceName)
}

func (c client) getBaseUrlForResourceType(resource interface{}) string {
	return fmt.Sprintf("%s/%s", c.scopedBasePath, pluralize(resource))
}

// Generic OctopusDeploy API Get Function.
func (c client) apiGetByID(resource interface{}, id string) (interface{}, error) {
	getClient := c.sling.New()
	if getClient == nil {
		return nil, createClientInitializationError(OperationAPIGetByID)
	}

	path := fmt.Sprintf("%s/%s", c.getBaseUrlForResourceType(resource), id)
	getClient = getClient.Get(path)
	if getClient == nil {
		return nil, createClientInitializationError(OperationAPIGetByID)
	}

	getClient.Set("User-Agent", getUserAgentString(c.requestingTool))

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
func getUserAgentString(requestingTool string) string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			if dep.Path == "github.com/OctopusDeploy/go-octopusdeploy" {
				if dep.Version != "" {
					version = dep.Version
				}
			}
		}
	}
	if !isEmpty(requestingTool) {
		requestingTool = fmt.Sprintf(" %s", requestingTool)
	}
	return fmt.Sprintf("%s/%s (%s; %s) go/%s%s", "go-octopusdeploy", version, runtime.GOOS, runtime.GOARCH, runtime.Version(), requestingTool)
}

// Generic OctopusDeploy API Add Function. Expects a 201 response.
func (c client) apiAdd(inputStruct interface{}, resource interface{}) (interface{}, error) {
	postClient := c.sling.New()
	if postClient == nil {
		return nil, createClientInitializationError(OperationAPIAdd)
	}

	path := c.getBaseUrlForResourceType(resource)

	postClient = postClient.Post(path)
	if postClient == nil {
		return nil, createClientInitializationError(OperationAPIAdd)
	}

	postClient.Set("User-Agent", getUserAgentString(c.requestingTool))

	request := postClient.BodyJSON(inputStruct)
	if request == nil {
		return nil, createClientInitializationError(OperationAPIAdd)
	}

	octopusDeployError := new(APIError)
	resp, err := request.Receive(resource, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusCreated, err, octopusDeployError)
	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// apiAddWithResponseStatus function with defined response.
func (c client) apiAddWithResponseStatus(inputStruct interface{}, resource interface{}, path string, httpStatus int) (interface{}, error) {
	if isEmpty(path) {
		return nil, createInvalidParameterError(OperationAPIAdd, ParameterPath)
	}

	postClient := c.sling.New()
	if postClient == nil {
		return nil, createClientInitializationError(OperationAPIAdd)
	}

	postClient = postClient.Post(path)
	if postClient == nil {
		return nil, createClientInitializationError(OperationAPIAdd)
	}

	postClient.Set("User-Agent", getUserAgentString(c.requestingTool))

	request := postClient.BodyJSON(inputStruct)
	if request == nil {
		return nil, createClientInitializationError(OperationAPIAdd)
	}

	octopusDeployError := new(APIError)
	resp, err := request.Receive(resource, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, httpStatus, err, octopusDeployError)
	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// apiPost post to octopus and expect a 200 response code.
func (c client) apiPost(inputStruct interface{}, resource interface{}, path string) (interface{}, error) {
	if isEmpty(path) {
		return nil, createInvalidParameterError(OperationAPIPost, ParameterPath)
	}

	postClient := c.sling.New()
	if postClient == nil {
		return nil, createClientInitializationError(OperationAPIPost)
	}

	postClient = postClient.Post(path)
	if postClient == nil {
		return nil, createClientInitializationError(OperationAPIPost)
	}

	postClient.Set("User-Agent", getUserAgentString(c.requestingTool))

	request := postClient.BodyJSON(inputStruct)
	if request == nil {
		return nil, createClientInitializationError(OperationAPIPost)
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
func (c client) apiUpdate(inputStruct interface{}, resource interface{}) (interface{}, error) {
	putClient := c.sling.New()
	if putClient == nil {
		return nil, createClientInitializationError(OperationAPIUpdate)
	}

	path := fmt.Sprintf("%s/%s", c.scopedBasePath, pluralize(inputStruct))

	putClient = putClient.Put(path)
	if putClient == nil {
		return nil, createClientInitializationError(OperationAPIUpdate)
	}

	putClient.Set("User-Agent", getUserAgentString(c.requestingTool))

	request := putClient.BodyJSON(inputStruct)
	if request == nil {
		return nil, createClientInitializationError(OperationAPIUpdate)
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
func (c client) apiDelete(resource interface{}, id string) error {
	if isEmpty(id) {
		return createInvalidParameterError(OperationAPIDelete, ParameterID)
	}

	path := fmt.Sprintf("%s/%s", c.getBaseUrlForResourceType(resource), id)

	deleteClient := c.sling.New()
	if deleteClient == nil {
		return createClientInitializationError(OperationAPIDelete)
	}

	deleteClient = deleteClient.Delete(path)
	if deleteClient == nil {
		return createClientInitializationError(OperationAPIDelete)
	}

	deleteClient.Set("User-Agent", getUserAgentString(c.requestingTool))

	octopusDeployError := new(APIError)
	resp, err := deleteClient.Receive(nil, &octopusDeployError)

	return APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)
}

// ErrItemNotFound is an OctopusDeploy error returned an item cannot be found.
var ErrItemNotFound = errors.New("cannot find the item")
