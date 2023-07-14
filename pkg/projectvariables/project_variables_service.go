package projectvariables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

type ProjectVariableService struct {
	services.Service
}

func NewProjectVariableService(sling *sling.Sling, uriTemplate string) *ProjectVariableService {
	return &ProjectVariableService{
		Service: services.NewService(constants.ServiceProjectService, sling, uriTemplate),
	}
}

func (s *ProjectVariableService) GetAllByGitRef(spaceId string, projectID string, gitRef string) (*variables.VariableSet, error) {
	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	if internal.IsEmpty(projectID) {
		return nil, internal.CreateInvalidParameterError("GetAllVariablesByGitRef", "projectID")
	}
	if internal.IsEmpty(gitRef) {
		return nil, internal.CreateInvalidParameterError("GetAllVariablesByGitRef", "gitRef")
	}

	path, err := getGitRefPath(spaceId, projectID, gitRef)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(variables.VariableSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*variables.VariableSet), nil
}

// AddSingle adds a single variable to a owner ID. This automates the act of fetching
// the variable set, adding a new item to it, and posting back to Octopus
func (s *ProjectVariableService) AddSingleByGitRef(spaceID string, projectID string, gitRef string, variable *variables.Variable) (*variables.VariableSet, error) {
	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	if internal.IsEmpty(projectID) {
		return nil, internal.CreateInvalidParameterError("AddSingleByGitRef", "projectID")
	}
	if internal.IsEmpty(gitRef) {
		return nil, internal.CreateInvalidParameterError("AddSingleByGitRef", "gitRef")
	}

	variables, err := s.GetAllByGitRef(spaceID, projectID, gitRef)
	if err != nil {
		return nil, err
	}

	variables.Variables = append(variables.Variables, variable)
	return s.UpdateByGitRef(spaceID, projectID, gitRef, variables)
}

// Update takes an entire variable set and posts the entire set back to Octopus Deploy. There are individual
// functions like AddSingle and UpdateSingle that can make this process more of a "typical" CRUD Octopus command.
func (s *ProjectVariableService) UpdateByGitRef(spaceID string, projectID string, gitRef string, variableSet *variables.VariableSet) (*variables.VariableSet, error) {
	err := services.ValidateInternalState(s)
	if err != nil {
		return nil, err
	}

	if internal.IsEmpty(projectID) {
		return nil, internal.CreateInvalidParameterError("UpdateByGitRef", "projectID")
	}
	if internal.IsEmpty(gitRef) {
		return nil, internal.CreateInvalidParameterError("UpdateByGitRef", "gitRef")
	}

	path, err := getGitRefPath(spaceID, projectID, gitRef)
	if err != nil {
		return nil, err
	}
	if _, err := services.ApiUpdate(s.GetClient(), variableSet, new(variables.VariableSet), path); err != nil {
		return nil, err
	}

	// 2021-04-22 (John Bristowe): we need to retrieve the variable set (again)
	// via HTTP GET (below) due to a bug for HTTP POST and HTTP PUT which will
	// provide a null scope value set in their responses

	return s.GetAllByGitRef(spaceID, projectID, gitRef)
}

func getGitRefPath(spaceId string, projectID string, gitRef string) (string, error) {
	values := map[string]any{
		"spaceId":   spaceId,
		"projectId": projectID,
		"gitRef":    gitRef,
	}
	template, err := uritemplates.Parse(uritemplates.ProjectVariablesByGitRef)
	if err != nil {
		return "", err
	}
	path, err := template.Expand(values)
	if err != nil {
		return "", err
	}
	return path, nil
}
