package octopusdeploy

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

const (
	emptyString      string = ""
	whitespaceString string = " "
)

// IService defines the contract for all services that communicate with the
// Octopus API.
type IService interface {
	getBasePath() string
	getClient() *sling.Sling
	getName() string
	getPath() string
	getURITemplate() *uritemplates.UriTemplate
}

type service struct {
	BasePath    string
	Name        string
	Path        string
	Sling       *sling.Sling
	URITemplate *uritemplates.UriTemplate
	itemType    IResource
}

type canDeleteService struct {
	service
}

func newService(name string, sling *sling.Sling, uriTemplate string, itemType IResource) service {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, _ := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	basePath, _ := template.Expand(make(map[string]interface{}))

	return service{
		itemType:    itemType,
		BasePath:    basePath,
		Name:        name,
		Path:        strings.TrimSpace(uriTemplate),
		Sling:       sling,
		URITemplate: template,
	}
}

func (s service) getBasePath() string {
	return s.BasePath
}

func (s service) getClient() *sling.Sling {
	return s.Sling
}

func (s service) getName() string {
	return s.Name
}

func (s service) getPath() string {
	return s.Path
}

func (s service) getURITemplate() *uritemplates.UriTemplate {
	return s.URITemplate
}

func getAddPath(s IService, r IResource) (string, error) {
	if r == nil || isNil(r) {
		return emptyString, createInvalidParameterError(operationAdd, parameterResource)
	}

	err := r.Validate()
	if err != nil {
		return emptyString, createValidationFailureError(operationAdd, err)
	}

	err = validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	return s.getURITemplate().Expand(values)
}

func getPath(s IService) (string, error) {
	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	return s.getURITemplate().Expand(values)
}

func getAllPath(s IService) (string, error) {
	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	return s.getBasePath() + "/all", nil
}

func getByIDPath(s IService, id string) (string, error) {
	if isEmpty(id) {
		return emptyString, createInvalidParameterError(operationGetByID, parameterID)
	}

	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	values[parameterID] = id

	return s.getURITemplate().Expand(values)
}

func getByIDsPath(s IService, ids []string) (string, error) {
	if len(ids) == 0 {
		return s.getURITemplate().Expand(make(map[string]interface{}))
	}

	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	idValues := emptyString

	for i := 0; i < len(ids); i++ {
		idValues += ids[i]
		if i < len(ids)-1 {
			idValues += ","
		}
	}

	values := make(map[string]interface{})
	values[parameterIDs] = idValues

	return s.getURITemplate().Expand(values)
}

func getByNamePath(s IService, name string) (string, error) {
	if isEmpty(name) {
		return emptyString, createInvalidParameterError(operationGetByName, parameterName)
	}

	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	values[parameterName] = name

	return s.getURITemplate().Expand(values)
}

func getByPartialNamePath(s IService, name string) (string, error) {
	if isEmpty(name) {
		return emptyString, createInvalidParameterError(operationGetByPartialName, parameterName)
	}

	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	values[parameterPartialName] = name

	return s.getURITemplate().Expand(values)
}

func getByAccountTypePath(s IService, accountType string) (string, error) {
	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	values[parameterAccountType] = accountType

	return s.getURITemplate().Expand(values)
}

func (s *service) deleteByID(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError(operationDeleteByID, parameterID)
	}

	err := validateInternalState(s)
	if err != nil {
		return err
	}

	values := make(map[string]interface{})
	values[parameterID] = id

	path, err := s.getURITemplate().Expand(values)
	if err != nil {
		return err
	}

	return apiDelete(s.getClient(), path)
}

func getUpdatePath(s IService, r IResource) (string, error) {
	if isNil(r) {
		return emptyString, createInvalidParameterError(operationUpdate, parameterResource)
	}

	err := r.Validate()
	if err != nil {
		return emptyString, createValidationFailureError(operationUpdate, err)
	}

	err = validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	values[parameterID] = r.GetID()
	return s.getURITemplate().Expand(values)
}

func validateInternalState(s IService) error {
	if s.getClient() == nil {
		return createInvalidClientStateError(s.getName())
	}

	values := make(map[string]interface{})
	path, err := s.getURITemplate().Expand(values)

	if isEmpty(path) {
		return createInvalidPathError(s.getName())
	}

	return err
}

// DeleteByID deletes the resource that matches the input ID.
func (s *canDeleteService) DeleteByID(id string) error {
	err := s.deleteByID(id)
	if err == ErrItemNotFound {
		return createResourceNotFoundError(s.getName(), "ID", id)
	}

	return err
}

var _ IService = &service{}
