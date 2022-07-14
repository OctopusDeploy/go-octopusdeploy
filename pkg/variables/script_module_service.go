package variables

import (
	"fmt"
	"math"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type ScriptModuleService struct {
	services.CanDeleteService
}

func NewScriptModuleService(sling *sling.Sling, uriTemplate string) *ScriptModuleService {
	return &ScriptModuleService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceLibraryVariableSetService, sling, uriTemplate),
		},
	}
}

func (s *ScriptModuleService) getPagedResponse(path string) ([]*ScriptModule, error) {
	resources := []*ScriptModule{}
	loadNextPage := true

	for loadNextPage {
		resp, err := services.ApiGet(s.GetClient(), new(ScriptModules), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*ScriptModules)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = services.LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new script module.
func (s *ScriptModuleService) Add(scriptModule *ScriptModule) (*ScriptModule, error) {
	if IsNil(scriptModule) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterScriptModule)
	}

	path, err := services.GetAddPath(s, scriptModule)
	if err != nil {
		return nil, err
	}

	response, err := services.ApiAdd(s.GetClient(), scriptModule, new(ScriptModule), path)
	if err != nil {
		return nil, err
	}

	scriptModuleResponse := response.(*ScriptModule)

	// update associated variable set; add script body and syntax
	variablesPath := scriptModuleResponse.Links[constants.LinkVariables]
	variablesResponse, err := services.ApiGet(s.GetClient(), new(VariableSet), variablesPath)
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

	_, err = services.ApiUpdate(s.GetClient(), variableSet, new(VariableSet), variablesPath)
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
func (s *ScriptModuleService) Get(libraryVariablesQuery LibraryVariablesQuery) (*ScriptModules, error) {
	path, err := s.GetURITemplate().Expand(libraryVariablesQuery)
	if err != nil {
		return &ScriptModules{}, err
	}

	response, err := services.ApiGet(s.GetClient(), new(ScriptModules), path)
	if err != nil {
		return &ScriptModules{}, err
	}

	return response.(*ScriptModules), nil
}

// GetAll returns all script modules. If none can be found or an error
// occurs, it returns an empty collection.
func (s *ScriptModuleService) GetAll() (*ScriptModules, error) {
	path, err := s.GetURITemplate().Expand(&LibraryVariablesQuery{
		ContentType: "ScriptModule",
		Take:        math.MaxInt32,
	})
	if err != nil {
		return &ScriptModules{}, err
	}

	response, err := services.ApiGet(s.GetClient(), new(ScriptModules), path)
	if err != nil {
		return &ScriptModules{}, err
	}

	return response.(*ScriptModules), nil
}

// GetByID returns the script module that matches the input ID. If one
// cannot be found, it returns nil and an error.
func (s *ScriptModuleService) GetByID(id string) (*ScriptModule, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	response, err := services.ApiGet(s.GetClient(), new(ScriptModule), path)
	if err != nil {
		return nil, err
	}

	scriptModuleResponse := response.(*ScriptModule)

	// get associated variable set
	variablesPath := scriptModuleResponse.Links[constants.LinkVariables]
	variablesResponse, err := services.ApiGet(s.GetClient(), new(VariableSet), variablesPath)
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
func (s *ScriptModuleService) GetByPartialName(partialName string) ([]*ScriptModule, error) {
	if internal.IsEmpty(partialName) {
		return []*ScriptModule{}, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName)
	}

	path, err := services.GetByPartialNamePath(s, partialName)
	if err != nil {
		return []*ScriptModule{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a script module based on the one provided as input.
func (s *ScriptModuleService) Update(scriptModule *ScriptModule) (*ScriptModule, error) {
	if scriptModule == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, "scriptModule")
	}

	path, err := services.GetUpdatePath(s, scriptModule)
	if err != nil {
		return nil, err
	}

	// update script module
	response, err := services.ApiUpdate(s.GetClient(), scriptModule, new(ScriptModule), path)
	if err != nil {
		return nil, err
	}

	scriptModuleResponse := response.(*ScriptModule)

	// update associated variable set
	variablesPath := scriptModuleResponse.Links[constants.LinkVariables]
	variablesResponse, err := services.ApiGet(s.GetClient(), new(VariableSet), variablesPath)
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

	updatedVariablesResponse, err := services.ApiUpdate(s.GetClient(), variableSet, new(VariableSet), variablesPath)
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
