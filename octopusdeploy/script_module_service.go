package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type scriptModuleService struct {
	canDeleteService
}

func newScriptModuleService(sling *sling.Sling, uriTemplate string) *scriptModuleService {
	scriptModuleService := &scriptModuleService{}
	scriptModuleService.service = newService(ServiceLibraryVariableSetService, sling, uriTemplate)

	return scriptModuleService
}

func (s scriptModuleService) getPagedResponse(path string) ([]*ScriptModule, error) {
	resources := []*ScriptModule{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(ScriptModules), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*ScriptModules)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new script module.
func (s scriptModuleService) Add(scriptModule *ScriptModule) (*ScriptModule, error) {
	path, err := getAddPath(s, scriptModule)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), scriptModule, new(ScriptModule), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ScriptModule), nil
}

// Get returns a collection of library variable sets based on the criteria
// defined by its input query parameter. If an error occurs, an empty
// collection is returned along with the associated error.
func (s scriptModuleService) Get(libraryVariablesQuery LibraryVariablesQuery) (*ScriptModules, error) {
	path, err := s.getURITemplate().Expand(libraryVariablesQuery)
	if err != nil {
		return &ScriptModules{}, err
	}

	response, err := apiGet(s.getClient(), new(ScriptModules), path)
	if err != nil {
		return &ScriptModules{}, err
	}

	return response.(*ScriptModules), nil
}

// GetAll returns all script modules. If none can be found or an error
// occurs, it returns an empty collection.
func (s scriptModuleService) GetAll() ([]*ScriptModule, error) {
	items := []*ScriptModule{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the library variable set that matches the input ID. If one
// cannot be found, it returns nil and an error.
func (s scriptModuleService) GetByID(id string) (*ScriptModule, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(ScriptModule), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ScriptModule), nil
}

// GetByPartialName performs a lookup and returns a list of library variable sets with a matching partial name.
func (s scriptModuleService) GetByPartialName(name string) ([]*ScriptModule, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*ScriptModule{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a library variable set based on the one provided as input.
func (s scriptModuleService) Update(scriptModule *ScriptModule) (*ScriptModule, error) {
	if scriptModule == nil {
		return nil, createInvalidParameterError(OperationUpdate, "scriptModule")
	}

	path, err := getUpdatePath(s, scriptModule)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), scriptModule, new(ScriptModule), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ScriptModule), nil
}
