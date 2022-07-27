package services

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

var version = "development"
var UserAgentString = GetUserAgentString()

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
		resp, err := ApiGet(s.GetClient(), new(resources.Resources[T]), path)
		if err != nil {
			return resourcesToReturn, err
		}

		responseList := resp.(*resources.Resources[T])
		resourcesToReturn = append(resourcesToReturn, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resourcesToReturn, nil
}

// Generic OctopusDeploy API Get Function.
func ApiGet(sling *sling.Sling, inputStruct interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAPIGet, "sling")
	}

	client := sling.New()
	if client == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIGet)
	}

	client = client.Get(path)
	if client == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIGet)
	}

	client.Set("User-Agent", UserAgentString)

	octopusDeployError := new(core.APIError)
	resp, err := client.Receive(inputStruct, &octopusDeployError)
	// if err != nil {
	// 	return nil, err
	// }

	apiErrorCheck := core.APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)
	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return inputStruct, nil
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

// Returns the User-Agent String "go-octopusdeploy/version (os; arch) go/version"
func GetUserAgentString() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			if dep.Path == "github.com/OctopusDeploy/go-octopusdeploy/v2" {
				if dep.Version != "" {
					version = dep.Version
				}
			}
		}
	}
	return fmt.Sprintf("%s/%s (%s; %s) go/%s", "go-octopusdeploy", version, runtime.GOOS, runtime.GOARCH, runtime.Version())
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

	client.Set("User-Agent", UserAgentString)

	request := client.BodyJSON(inputStruct)
	if request == nil {
		return nil, internal.CreateClientInitializationError(constants.OperationAPIAdd)
	}

	octopusDeployError := new(core.APIError)
	resp, err := request.Receive(resource, &octopusDeployError)

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

	client.Set("User-Agent", UserAgentString)

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

	client.Set("User-Agent", UserAgentString)

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

	client.Set("User-Agent", UserAgentString)

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

	client.Set("User-Agent", UserAgentString)

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
