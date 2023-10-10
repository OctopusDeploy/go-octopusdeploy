package libraryvariableset

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

// Add creates a new library variable set.
func Add(client newclient.Client, libraryVariableSet *variables.LibraryVariableSet) (*variables.LibraryVariableSet, error) {
	if libraryVariableSet == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterLibraryVariableSet)
	}

	spaceID, err := internal.GetSpaceID(libraryVariableSet.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.LibraryVariableSets, map[string]any{"spaceId": spaceID})
	if err != nil {
		return nil, err
	}

	return newclient.Post[variables.LibraryVariableSet](client.HttpSession(), expandedUri, libraryVariableSet)
}

// Get returns a collection of library variable sets based on the criteria
// defined by its input query parameter. If an error occurs, an empty
// collection is returned along with the associated error.
func Get(client newclient.Client, spaceID string, libraryVariablesQuery variables.LibraryVariablesQuery) (*resources.Resources[*variables.LibraryVariableSet], error) {
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

	res, err := newclient.Get[resources.Resources[*variables.LibraryVariableSet]](client.HttpSession(), expandedUri)
	if err != nil {
		return &resources.Resources[*variables.LibraryVariableSet]{}, err
	}

	return res, nil
}

// GetByID returns the library variable set that matches the space ID and input ID. If one
// cannot be found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, id string) (*variables.LibraryVariableSet, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.LibraryVariableSets, map[string]any{
		"spaceId": spaceID,
		"id":      id,
	})
	if err != nil {
		return nil, err
	}

	return newclient.Get[variables.LibraryVariableSet](client.HttpSession(), expandedUri)
}

// Update modifies a library variable set based on the one provided as input.
func Update(client newclient.Client, libraryVariableSet *variables.LibraryVariableSet) (*variables.LibraryVariableSet, error) {
	if libraryVariableSet == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterLibraryVariableSet)
	}

	spaceID, err := internal.GetSpaceID(libraryVariableSet.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.LibraryVariableSets, map[string]any{
		"spaceId": spaceID,
		"id":      libraryVariableSet.ID,
	})
	if err != nil {
		return nil, err
	}

	return newclient.Put[variables.LibraryVariableSet](client.HttpSession(), expandedUri, libraryVariableSet)
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
