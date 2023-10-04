package libraryvariableset

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

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
