package buildinformation

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type BuildInformationService struct {
	bulkPath string

	services.CanDeleteService
}

func NewBuildInformationService(sling *sling.Sling, uriTemplate string, bulkPath string) *BuildInformationService {
	return &BuildInformationService{
		bulkPath: bulkPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceBuildInformationService, sling, uriTemplate),
		},
	}
}

const (
	template     = "/api/{spaceId}/build-information{/id}{?packageId,filter,latest,skip,take,overwriteMode}"
	bulkTemplate = "/api/{spaceId}/build-information/bulk{?ids}"
)

// Add creates a new build information package
func Add(client newclient.Client, command *CreateBuildInformationCommand) (*BuildInformation, error) {
	if IsNil(command) {
		return nil, internal.CreateInvalidParameterError("CreateBuildInformation", "command")
	}

	resp, err := newclient.Add[BuildInformation](client, template, command.SpaceId, command)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get returns a collection of build information based on the criteria defined by its
// input query parameters. If an error occurs, an empty collection is returned along
// with the associated error
func Get(client newclient.Client, spaceId string, buildInformationQuery BuildInformationQuery) (*resources.Resources[*BuildInformation], error) {
	return newclient.GetByQuery[BuildInformation](client, template, spaceId, buildInformationQuery)
}

// GetAll returns all build information. If none can be found or an error occurs, it
// returns an empty collection.
func GetAll(client newclient.Client, spaceId string) ([]*BuildInformation, error) {
	return newclient.GetAll[BuildInformation](client, template, spaceId)
}

// GetById returns the build information that matches the input ID. If one cannot be
// found, it return nil and an error
func GetById(client newclient.Client, spaceId string, id string) (*BuildInformation, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	return newclient.GetByID[BuildInformation](client, template, spaceId, id)
}

// DeleteById deletes the build information based on the ID provided as input.
func DeleteByID(client newclient.Client, spaceId string, id string) error {
	return newclient.DeleteByID(client, template, spaceId, id)
}

// DeleteByIds deletes all build information based on the IDs provided as input.
func DeleteByIDs(client newclient.Client, spaceId string, ids []string) error {
	templateParams := map[string]any{"spaceId": spaceId, "ids": ids}
	expandedUri, err := client.URITemplateCache().Expand(bulkTemplate, templateParams)
	if err != nil {
		return err
	}

	return newclient.Delete(client.HttpSession(), expandedUri)
}
