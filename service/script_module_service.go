package service

import (
	"fmt"
	"math"
	"strings"

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

	response, err := apiAdd(s.getClient(), scriptModule, new(ScriptModule), path)
	if err != nil {
		return nil, err
	}

	scriptModuleResponse := response.(*ScriptModule)

	// update associated variable set; Add script body and syntax
	variablesPath := scriptModuleResponse.Links[linkVariables]
	variablesResponse, err := apiGet(s.getClient(), new(VariableSet), variablesPath)
	if err != nil {
		return nil, err
	}

	variableSet := variablesResponse.(*VariableSet)
	scriptBodyVariable := NewVariable(fmt.Sprintf("Octopus.Script.Module[%s]", scriptModule.Name))
	scriptBodyVariable.Value = scriptModule.ScriptBody
	variableSet.Variables = append(variableSet.Variables, scriptBodyVariable)

	syntaxVariable := NewVariable(fmt.Sprintf("Octopus.Script.Module.Language[%s]", scriptModule.Name))
	syntaxVariable.Value = scriptModule.Syntax
	variableSet.Variables = append(variableSet.Variables, syntaxVariable)

	_, err = apiUpdate(s.getClient(), variableSet, new(VariableSet), variablesPath)
	if err != nil {
		return nil, err
	}

	scriptModuleResponse.ScriptBody = scriptModule.ScriptBody
	scriptModuleResponse.Syntax = scriptModule.Syntax

	return scriptModuleResponse, nil
}

// Get returns a collection of script modules based on the criteria
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
func (s scriptModuleService) GetAll() (*ScriptModules, error) {
	path, err := s.getURITemplate().Expand(&LibraryVariablesQuery{
		ContentType: "ScriptModule",
		Take:        math.MaxInt32,
	})
	if err != nil {
		return &ScriptModules{}, err
	}

	response, err := apiGet(s.getClient(), new(ScriptModules), path)
	if err != nil {
		return &ScriptModules{}, err
	}

	return response.(*ScriptModules), nil
}

// GetByID returns the script module that matches the input ID. If one
// cannot be found, it returns nil and an error.
func (s scriptModuleService) GetByID(id string) (*ScriptModule, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	response, err := apiGet(s.getClient(), new(ScriptModule), path)
	if err != nil {
		return nil, err
	}

	scriptModuleResponse := response.(*ScriptModule)

	// get associated variable set
	variablesPath := scriptModuleResponse.Links[linkVariables]
	variablesResponse, err := apiGet(s.getClient(), new(VariableSet), variablesPath)
	if err != nil {
		return nil, err
	}

	variableSet := variablesResponse.(*VariableSet)
	for _, variable := range variableSet.Variables {
		if strings.HasPrefix(variable.Name, "Octopus.Script.Module[") {
			scriptModuleResponse.ScriptBody = variable.Value
		}

		if strings.HasPrefix(variable.Name, "Octopus.Script.Module.Language[") {
			scriptModuleResponse.Syntax = variable.Value
		}
	}

	return scriptModuleResponse, nil
}

// GetByPartialName performs a lookup and returns a list of script modules with a matching partial name.
func (s scriptModuleService) GetByPartialName(name string) ([]*ScriptModule, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*ScriptModule{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a script module based on the one provided as input.
func (s scriptModuleService) Update(scriptModule *ScriptModule) (*ScriptModule, error) {
	if scriptModule == nil {
		return nil, createInvalidParameterError(OperationUpdate, "scriptModule")
	}

	path, err := getUpdatePath(s, scriptModule)
	if err != nil {
		return nil, err
	}

	// update script module
	response, err := apiUpdate(s.getClient(), scriptModule, new(ScriptModule), path)
	if err != nil {
		return nil, err
	}

	scriptModuleResponse := response.(*ScriptModule)

	// update associated variable set
	variablesPath := scriptModuleResponse.Links[linkVariables]
	variablesResponse, err := apiGet(s.getClient(), new(VariableSet), variablesPath)
	if err != nil {
		return nil, err
	}

	variableSet := variablesResponse.(*VariableSet)
	for _, variable := range variableSet.Variables {
		if strings.HasPrefix(variable.Name, "Octopus.Script.Module[") {
			variable.Value = scriptModule.ScriptBody
		}

		if strings.HasPrefix(variable.Name, "Octopus.Script.Module.Language[") {
			variable.Value = scriptModule.Syntax
		}
	}

	updatedVariablesResponse, err := apiUpdate(s.getClient(), variableSet, new(VariableSet), variablesPath)
	if err != nil {
		return nil, err
	}

	updatedVriableSet := updatedVariablesResponse.(*VariableSet)
	for _, variable := range updatedVriableSet.Variables {
		if strings.HasPrefix(variable.Name, "Octopus.Script.Module[") {
			scriptModuleResponse.ScriptBody = variable.Value
		}

		if strings.HasPrefix(variable.Name, "Octopus.Script.Module.Language[") {
			scriptModuleResponse.Syntax = variable.Value
		}
	}

	return scriptModuleResponse, nil
}
