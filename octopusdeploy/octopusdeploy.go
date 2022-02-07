package octopusdeploy

import (
	"errors"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
)

var version = "development"

type Client struct {
	sling                 *sling.Sling
	octopusServerEndpoint *octopusServerEndpoint
	scopedBasePath        string
	requestingTool        string
}

type AdminClient struct {
	Client
}

type SpaceScopedClient struct {
	spaceID string
	Client
}

func newClient(octopusServerEndpoint *octopusServerEndpoint, requestingTool string, pathScope string) *Client {
	httpClient := &http.Client{}
	scopedBasePath := octopusServerEndpoint.BaseURLWithAPI.String()
	if !IsEmpty(pathScope) {
		scopedBasePath = fmt.Sprintf("%s/%s", scopedBasePath, pathScope)
	}
	base := sling.New().Client(httpClient).Base(scopedBasePath).Set(clientAPIKeyHTTPHeader, octopusServerEndpoint.ApiKey)
	base.Set("User-Agent", getUserAgentString(requestingTool))

	c := &Client{
		sling:                 base,
		octopusServerEndpoint: octopusServerEndpoint,
		scopedBasePath:        scopedBasePath,
		requestingTool:        requestingTool,
	}
	return c
}

// NewAdminClient returns a new Octopus API Client.
func NewAdminClient(octopusServerEndpoint *octopusServerEndpoint, requestingTool string) (*AdminClient, error) {
	if octopusServerEndpoint == nil {
		return nil, CreateInvalidParameterError("NewAdminClient", ParameterOctopusServerEndpoint)
	}

	c := AdminClient{
		Client: *newClient(octopusServerEndpoint, requestingTool, empty),
	}

	return &c, nil
}

func NewSpaceScopedClient(octopusServerEndpoint *octopusServerEndpoint, spaceID string, requestingTool string) (*SpaceScopedClient, error) {
	if octopusServerEndpoint == nil {
		return nil, CreateInvalidParameterError("NewSpaceScopedClient", ParameterOctopusServerEndpoint)
	}

	httpClient := &http.Client{}
	base := sling.New().Client(httpClient).Base(octopusServerEndpoint.BaseURLWithAPI.String()).Set(clientAPIKeyHTTPHeader, octopusServerEndpoint.ApiKey)
	base.Set("User-Agent", getUserAgentString(requestingTool))

	client := &SpaceScopedClient{
		Client:  *newClient(octopusServerEndpoint, requestingTool, spaceID),
		spaceID: spaceID,
	}

	return client, nil
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

func pluralize(resource interface{}) string {
	paged := resource.(*PagedResults)
	if paged != nil {
		return strings.ToLower(reflect.TypeOf(resource).Name())
	}
	resourceName := strings.ToLower(strings.ReplaceAll(reflect.TypeOf(resource).Name(), "Resource", ""))
	return fmt.Sprintf("%ss", resourceName)
}

func (c Client) getBaseUrlForResourceType(resource interface{}) string {
	return fmt.Sprintf("%s/%s", c.scopedBasePath, pluralize(resource))
}

func ApiGetByID[T Resource](c Client, id string) (*T, error) {
	getClient := c.sling.New()
	if getClient == nil {
		return nil, createClientInitializationError(OperationAPIGetByID)
	}

	resource := new(T)
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
	if !IsEmpty(requestingTool) {
		requestingTool = fmt.Sprintf(" %s", requestingTool)
	}
	return fmt.Sprintf("%s/%s (%s; %s) go/%s%s", "go-octopusdeploy", version, runtime.GOOS, runtime.GOARCH, runtime.Version(), requestingTool)
}

// Generic OctopusDeploy API Add Function. Expects a 201 response.
func ApiAdd[T Resource](c Client, inputStruct *T) (*T, error) {
	postClient := c.sling.New()
	if postClient == nil {
		return nil, createClientInitializationError(OperationAPIAdd)
	}

	resource := new(T)
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
func (c Client) apiAddWithResponseStatus(inputStruct interface{}, resource interface{}, path string, httpStatus int) (interface{}, error) {
	if IsEmpty(path) {
		return nil, CreateInvalidParameterError(OperationAPIAdd, ParameterPath)
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
func (c Client) apiPost(inputStruct interface{}, resource interface{}, path string) (interface{}, error) {
	if IsEmpty(path) {
		return nil, CreateInvalidParameterError(OperationAPIPost, ParameterPath)
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
func ApiUpdate[T Resource](c Client, inputStruct *T) (*T, error) {
	putClient := c.sling.New()
	if putClient == nil {
		return nil, createClientInitializationError(OperationAPIUpdate)
	}

	resource := new(T)
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
func ApiDelete[T Resource](c Client, id string) error {
	if IsEmpty(id) {
		return CreateInvalidParameterError(OperationAPIDelete, ParameterID)
	}

	resource := new(T)
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
