package scriptmodules

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"strings"
)

const contentType = "ScriptModule"

// Add creates a new script module.
func Add(client newclient.Client, scriptModule *variables.ScriptModule) (*variables.ScriptModule, error) {
	if scriptModule == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterScriptModule)
	}

	spaceID, err := internal.GetSpaceID(scriptModule.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.LibraryVariableSets, map[string]any{"spaceId": spaceID})
	if err != nil {
		return nil, err
	}

	scriptModuleResponse, err := newclient.Post[variables.ScriptModule](client.HttpSession(), expandedUri, scriptModule)

	// update associated variable set; add script body and syntax
	variablesPath, err := client.URITemplateCache().Expand(uritemplates.Variables, map[string]any{
		"spaceId": spaceID,
		"id":      scriptModuleResponse.VariableSetID,
	})
	if err != nil {
		return nil, err
	}

	variablesResponse, err := newclient.Get[variables.VariableSet](client.HttpSession(), variablesPath)
	if err != nil {
		return nil, err
	}

	variableSet := *variablesResponse
	scriptBodyVariable := variables.NewVariable(fmt.Sprintf("Octopus.Script.Module[%s]", scriptModule.Name))
	scriptBodyVariable.Value = scriptModule.ScriptBody
	variableSet.Variables = append(variableSet.Variables, scriptBodyVariable)

	syntaxVariable := variables.NewVariable(fmt.Sprintf("Octopus.Script.Module.Language[%s]", scriptModule.Name))
	syntaxVariable.Value = scriptModule.Syntax
	variableSet.Variables = append(variableSet.Variables, syntaxVariable)

	_, err = newclient.Put[variables.VariableSet](client.HttpSession(), variablesPath, variableSet)
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
func Get(client newclient.Client, spaceID string, libraryVariablesQuery variables.LibraryVariablesQuery) (*resources.Resources[*variables.ScriptModule], error) {
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	values, _ := uritemplates.Struct2map(libraryVariablesQuery)
	if values == nil {
		values = map[string]any{}
	}
	values["spaceId"] = spaceID

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.LibraryVariableSets, values)
	if err != nil {
		return nil, err
	}

	resp, err := newclient.Get[resources.Resources[*variables.ScriptModule]](client.HttpSession(), expandedUri)
	if err != nil {
		return &resources.Resources[*variables.ScriptModule]{}, err
	}

	return resp, nil
}

// GetByID returns the script module that matches the space ID and input ID. If one
// cannot be found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, id string) (*variables.ScriptModule, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.LibraryVariableSets, map[string]any{
		"spaceId":     spaceID,
		"id":          id,
		"contentType": contentType,
	})

	scriptModuleResponse, err := newclient.Get[variables.ScriptModule](client.HttpSession(), expandedUri)

	// get associated variable set
	variablesPath, err := client.URITemplateCache().Expand(uritemplates.Variables, map[string]any{
		"spaceId": spaceID,
		"id":      scriptModuleResponse.VariableSetID,
	})

	variablesResponse, err := newclient.Get[variables.VariableSet](client.HttpSession(), variablesPath)
	if err != nil {
		return nil, err
	}

	variableSet := *variablesResponse
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

// Update modifies a script module based on the one provided as input.
func Update(client newclient.Client, scriptModule *variables.ScriptModule) (*variables.ScriptModule, error) {
	if scriptModule == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, "scriptModule")
	}

	spaceID, err := internal.GetSpaceID(scriptModule.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.LibraryVariableSets, map[string]any{
		"spaceId": spaceID,
		"id":      scriptModule.ID,
	})
	if err != nil {
		return nil, err
	}

	// update script module
	scriptModuleResponse, err := newclient.Put[variables.ScriptModule](client.HttpSession(), expandedUri, scriptModule)
	if err != nil {
		return nil, err
	}

	// update associated variable set
	variablesPath, err := client.URITemplateCache().Expand(uritemplates.Variables, map[string]any{
		"spaceId": spaceID,
		"id":      scriptModuleResponse.VariableSetID,
	})
	if err != nil {
		return nil, err
	}

	variableSet, err := newclient.Get[variables.VariableSet](client.HttpSession(), variablesPath)
	if err != nil {
		return nil, err
	}

	for _, variable := range variableSet.Variables {
		if strings.HasPrefix(variable.Name, "Octopus.Script.Module[") {
			variable.Value = scriptModule.ScriptBody
		}

		if strings.HasPrefix(variable.Name, "Octopus.Script.Module.Language[") {
			variable.Value = scriptModule.Syntax
		}
	}

	updatedVariableSet, err := newclient.Put[variables.VariableSet](client.HttpSession(), variablesPath, variableSet)

	for _, variable := range updatedVariableSet.Variables {
		if strings.HasPrefix(variable.Name, "Octopus.Script.Module[") {
			scriptModuleResponse.ScriptBody = variable.Value
		}

		if strings.HasPrefix(variable.Name, "Octopus.Script.Module.Language[") {
			scriptModuleResponse.Syntax = variable.Value
		}
	}

	return scriptModuleResponse, nil
}

// DeleteByID deletes the resource that matches the space ID and input ID.
func DeleteByID(client newclient.Client, spaceID string, id string) error {
	if internal.IsEmpty(id) {
		return internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID)
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.LibraryVariableSets, map[string]any{
		"spaceId": spaceID,
		"id":      id,
	})
	if err != nil {
		return err
	}

	return newclient.Delete(client.HttpSession(), expandedUri)
}
