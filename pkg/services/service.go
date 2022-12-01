package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

// IService defines the contract for all services that communicate with the
// Octopus API.
type IService interface {
	GetBasePath() string
	GetClient() *sling.Sling
	GetName() string
	GetPath() string
	GetURITemplate() *uritemplates.UriTemplate
}

type Service struct {
	BasePath    string
	Name        string
	Path        string
	Sling       *sling.Sling
	URITemplate *uritemplates.UriTemplate
}

type CanDeleteService struct {
	Service
}

func NewService(name string, sling *sling.Sling, uriTemplate string) Service {
	if sling == nil {
		sling = internal.GetDefaultClient()
	}

	template, _ := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	basePath, _ := template.Expand(make(map[string]interface{}))

	return Service{
		BasePath:    basePath,
		Name:        name,
		Path:        strings.TrimSpace(uriTemplate),
		Sling:       sling,
		URITemplate: template,
	}
}

func GetPagedResponse[T any](s IService, path string) ([]*T, error) {
	resourcesToReturn := []*T{}
	loadNextPage := true

	for loadNextPage {
		resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*T]), path)
		if err != nil {
			return resourcesToReturn, err
		}

		responseList := resp.(*resources.Resources[*T])
		resourcesToReturn = append(resourcesToReturn, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resourcesToReturn, nil
}

func (s *Service) GetBasePath() string {
	return s.BasePath
}

func (s *Service) GetClient() *sling.Sling {
	return s.Sling
}

func (s *Service) GetName() string {
	return s.Name
}

func (s *Service) GetPath() string {
	return s.Path
}

func (s *Service) GetURITemplate() *uritemplates.UriTemplate {
	return s.URITemplate
}

func GetAddPath(s IService, resource resources.IResource) (string, error) {
	if resource == nil || IsNil(resource) {
		return "", internal.CreateInvalidParameterError(constants.OperationGetAddPath, constants.ParameterResource)
	}

	if err := resource.Validate(); err != nil {
		return "", internal.CreateValidationFailureError(constants.OperationGetAddPath, err)
	}

	if err := ValidateInternalState(s); err != nil {
		return "", err
	}

	values := make(map[string]interface{})
	return s.GetURITemplate().Expand(values)
}

func GetPath(s IService) (string, error) {
	if err := ValidateInternalState(s); err != nil {
		return "", err
	}

	values := make(map[string]interface{})
	return s.GetURITemplate().Expand(values)
}

func GetAllPath(s IService) (string, error) {
	if err := ValidateInternalState(s); err != nil {
		return "", err
	}

	return s.GetBasePath() + "/all", nil
}

func GetByIDPath(s IService, id string) (string, error) {
	if internal.IsEmpty(id) {
		return "", internal.CreateInvalidParameterError(constants.OperationGetByIDPath, constants.ParameterID)
	}

	if err := ValidateInternalState(s); err != nil {
		return "", err
	}

	values := make(map[string]interface{})
	values[constants.ParameterID] = id

	return s.GetURITemplate().Expand(values)
}

func GetByIDsPath(s IService, ids []string) (string, error) {
	if len(ids) == 0 {
		return s.GetURITemplate().Expand(make(map[string]interface{}))
	}

	if err := ValidateInternalState(s); err != nil {
		return "", err
	}

	idValues := ""

	for i := 0; i < len(ids); i++ {
		idValues += ids[i]
		if i < len(ids)-1 {
			idValues += ","
		}
	}

	values := make(map[string]interface{})
	values["ids"] = idValues

	return s.GetURITemplate().Expand(values)
}

func GetByNamePath(s IService, name string) (string, error) {
	if internal.IsEmpty(name) {
		return "", internal.CreateInvalidParameterError(constants.OperationGetByNamePath, constants.ParameterName)
	}

	if err := ValidateInternalState(s); err != nil {
		return "", err
	}

	values := make(map[string]interface{})
	values["name"] = name

	return s.GetURITemplate().Expand(values)
}

func GetByPartialNamePath(s IService, name string) (string, error) {
	if internal.IsEmpty(name) {
		return "", internal.CreateInvalidParameterError(constants.OperationGetByPartialNamePath, constants.ParameterName)
	}

	if err := ValidateInternalState(s); err != nil {
		return "", err
	}

	values := make(map[string]interface{})
	values["partialName"] = name

	return s.GetURITemplate().Expand(values)
}

func GetUpdatePath(s IService, resource resources.IResource) (string, error) {
	if IsNil(resource) {
		return "", internal.CreateInvalidParameterError(constants.OperationUpdatePath, constants.ParameterResource)
	}

	if err := resource.Validate(); err != nil {
		return "", internal.CreateValidationFailureError(constants.OperationUpdatePath, err)
	}

	if err := ValidateInternalState(s); err != nil {
		return "", err
	}

	values := make(map[string]interface{})
	values["id"] = resource.GetID()
	return s.GetURITemplate().Expand(values)
}

func ValidateInternalState(s IService) error {
	if s.GetClient() == nil {
		return internal.CreateInvalidClientStateError(s.GetName())
	}

	values := make(map[string]interface{})
	path, err := s.GetURITemplate().Expand(values)

	if internal.IsEmpty(path) {
		return internal.CreateInvalidPathError(s.GetName())
	}

	return err
}

// DeleteByID deletes the resource that matches the input ID.
func (s *CanDeleteService) DeleteByID(id string) error {
	if internal.IsEmpty(id) {
		return internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID)
	}

	if err := ValidateInternalState(s); err != nil {
		return err
	}

	values := make(map[string]interface{})
	values["id"] = id

	path, err := s.GetURITemplate().Expand(values)
	if err != nil {
		return err
	}

	return ApiDelete(s.GetClient(), path)
}

var _ IService = &Service{}

// Generic OctopusDeploy API Add Function. Expects a 201 response.
func ApiAdd(sling *sling.Sling, inputStruct interface{}, resource interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAPIAdd, "sling")
	}

	if internal.IsEmpty(path) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAPIAdd, "path")
	}

	client := sling.New()
	if client == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIAdd)
	}

	client = client.Post(path)
	if client == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIAdd)
	}

	request := client.BodyJSON(inputStruct)
	if request == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIAdd)
	}

	octopusDeployError := new(core.APIError)
	resp, err := request.Receive(resource, &octopusDeployError)

	// workaround to account for API responses where it's either 200 or 201
	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		return resource, nil
	}

	if apiErrorCheck := core.APIErrorChecker(path, resp, http.StatusCreated, err, octopusDeployError); apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// ApiAddWithResponseStatus function with defined response.
func ApiAddWithResponseStatus(sling *sling.Sling, inputStruct interface{}, resource interface{}, path string, httpStatus int) (interface{}, error) {
	if sling == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationApiAddWithResponseStatus, "sling")
	}

	if internal.IsEmpty(path) {
		return nil, internal.CreateInvalidParameterError(constants.OperationApiAddWithResponseStatus, "path")
	}

	client := sling.New()
	if client == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationApiAddWithResponseStatus)
	}

	client = client.Post(path)
	if client == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationApiAddWithResponseStatus)
	}

	request := client.BodyJSON(inputStruct)
	if request == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationApiAddWithResponseStatus)
	}

	octopusDeployError := new(core.APIError)
	resp, err := request.Receive(resource, &octopusDeployError)

	if apiErrorCheck := core.APIErrorChecker(path, resp, httpStatus, err, octopusDeployError); apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// ApiPost post to octopus and expect a 200 response code.
func ApiPost(sling *sling.Sling, inputStruct interface{}, resource interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAPIPost, "sling")
	}

	if internal.IsEmpty(path) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAPIPost, "path")
	}

	client := sling.New()
	if client == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIPost)
	}

	client = client.Post(path)
	if client == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIPost)
	}

	request := client.BodyJSON(inputStruct)
	if request == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIPost)
	}

	octopusDeployError := new(core.APIError)
	resp, err := request.Receive(resource, &octopusDeployError)

	if apiErrorCheck := core.APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError); apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// ApiPost post some JSON to octopus and expect a 200 response code.
func ApiPostNew[TResponse any](httpClient *http.Client, absoluteUrl *url.URL, apiKey string, requestPayload any) (*TResponse, error) {
	if httpClient == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAPIPost, "httpClient")
	}

	if absoluteUrl == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAPIPost, "absoluteUrl")
	}

	var body io.Reader = nil
	if requestPayload != nil {
		payload, err := json.Marshal(requestPayload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(payload)
	}

	req, err := http.NewRequest(http.MethodPost, absoluteUrl.String(), body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", api.UserAgentString)
	}

	req.Header.Set(constants.ClientAPIKeyHTTPHeader, apiKey)
	return doRequestReturningJson[TResponse](httpClient, req)
}

func doRequestReturningJson[TResponse any](httpClient *http.Client, req *http.Request) (*TResponse, error) {
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	// behaviour copied from Sling
	defer func() {
		// The default HTTP client's Transport may not
		// reuse HTTP/1.x "keep-alive" TCP connections if the Body is
		// not read to completion and closed.
		// See: https://golang.org/pkg/net/http/#Response
		_, _ = io.Copy(io.Discard, resp.Body)

		// when err is nil, resp contains a non-nil resp.Body which must be closed
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusNoContent || resp.ContentLength == 0 {
		// Potential gotcha: If someone calls this with TResponse of string, int or other non-nullable primitive,
		// then this may panic. But why are you using a non-nullable response type on a server endpoint that can return no content?

		// TODO the ContentLength check is copied from Sling, but it's valid for servers to stream responses
		// without a known content length (which this won't handle), which would be a bug. Does it matter?
		return nil, nil
	}

	bodyDecoder := json.NewDecoder(resp.Body)

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		responsePayload := new(TResponse)
		err = bodyDecoder.Decode(responsePayload)
		if err != nil {
			return nil, err
		}
		return responsePayload, nil
	} else {
		errorPayload := new(core.APIError)
		err = bodyDecoder.Decode(errorPayload)
		if err != nil {
			return nil, err
		}
		return nil, errorPayload
	}
	// don't use core.APIErrorChecker, it's overly helpful and gets in the way of error handling.
}

// Generic OctopusDeploy API Update Function.
func ApiUpdate(sling *sling.Sling, inputStruct interface{}, resource interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAPIUpdate, "sling")
	}

	if internal.IsEmpty(path) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAPIUpdate, "path")
	}

	client := sling.New()
	if client == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIUpdate)
	}

	client = client.Put(path)
	if client == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIUpdate)
	}

	request := client.BodyJSON(inputStruct)
	if request == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIUpdate)
	}

	octopusDeployError := new(core.APIError)
	resp, err := request.Receive(resource, &octopusDeployError)

	if apiErrorCheck := core.APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError); apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// Generic OctopusDeploy API Delete Function.
func ApiDelete(sling *sling.Sling, path string) error {
	if sling == nil {
		return internal.CreateInvalidParameterError(constants.OperationAPIDelete, "sling")
	}

	if internal.IsEmpty(path) {
		return internal.CreateInvalidParameterError(constants.OperationAPIDelete, "path")
	}

	client := sling.New()
	if client == nil {
		return internal.CreateClientInitializationError(constants.OperationAPIDelete)
	}

	client = client.Delete(path)
	if client == nil {
		return internal.CreateClientInitializationError(constants.OperationAPIDelete)
	}

	octopusDeployError := new(core.APIError)
	resp, err := client.Receive(nil, &octopusDeployError)

	return core.APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)
}

// ErrItemNotFound is an OctopusDeploy error returned an item cannot be found.
var ErrItemNotFound = errors.New("cannot find the item")

// LoadNextPage checks if the next page should be loaded from the API. Returns
// the new path and a bool if the next page should be checked.
func LoadNextPage(pagedResults resources.PagedResults) (string, bool) {
	if pagedResults.Links.PageNext != "" {
		return pagedResults.Links.PageNext, true
	}

	return "", false
}
