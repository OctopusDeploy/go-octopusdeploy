package newclient

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

// DeleteByID will delete a resource with the provided id.
func DeleteByID(client Client, template string, spaceID string, id string) error {
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return err
	}
	if internal.IsEmpty(id) {
		return internal.CreateRequiredParameterIsEmptyError(constants.ParameterID)
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{
		"spaceId": spaceID,
		"id":      id,
	})
	if err != nil {
		return err
	}

	return Delete(client.HttpSession(), path)
}

// Add creates a new resource.
func Add[TResource any](client Client, template string, spaceID string, resource any) (*TResource, error) {
	if resource == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterAccount)
	}
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{
		"spaceId": spaceID,
	})
	if err != nil {
		return nil, err
	}

	res, err := Post[TResource](client.HttpSession(), path, resource)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update modifies a resource based on the one provided as input.
func Update[TResource any](client Client, template string, spaceID string, ID string, resource any) (*TResource, error) {
	if resource == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterWorkerPool)
	}

	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{
		"spaceId": spaceID,
		"id":      ID,
	})
	if err != nil {
		return nil, err
	}

	res, err := Put[TResource](client.HttpSession(), path, resource)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByQuery returns a collection of resources based on the criteria defined by
// its input query parameter.
func GetByQuery[TResource any](client Client, template string, spaceID string, query any) (*resources.Resources[*TResource], error) {
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}
	values, _ := uritemplates.Struct2map(query)
	if values == nil {
		values = map[string]any{}
	}
	values["spaceId"] = spaceID
	path, err := client.URITemplateCache().Expand(template, values)
	if err != nil {
		return nil, err
	}

	res, err := Get[resources.Resources[*TResource]](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByID returns the resource that matches the input ID.
func GetByID[TResource any](client Client, template string, spaceID string, ID string) (*TResource, error) {
	if ID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyError(constants.ParameterID)
	}
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{
		"spaceId": spaceID,
		"id":      ID,
	})
	if err != nil {
		return nil, err
	}

	res, err := Get[TResource](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return res, nil
}
