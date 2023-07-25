package projectbranches

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
	"net/http"
)

type ProjectBranchesService struct {
	services.Service
}

func NewProjectBranchesService(sling *sling.Sling, uriTemplate string) *ProjectBranchesService {
	return &ProjectBranchesService{
		Service: services.NewService(constants.ServiceProjectService, sling, uriTemplate),
	}
}

func (s *ProjectBranchesService) Add(spaceId, projectID, baseBranch, newBranchName string) (*projects.GitReference, error) {
	err := services.ValidateInternalState(s)
	if err != nil {
		return nil, err
	}
	if internal.IsEmpty(projectID) {
		return nil, internal.CreateInvalidParameterError("CreateBranch", "projectID")
	}
	if internal.IsEmpty(baseBranch) {
		return nil, internal.CreateInvalidParameterError("CreateBranch", "baseBranch")
	}
	if internal.IsEmpty(newBranchName) {
		return nil, internal.CreateInvalidParameterError("CreateBranch", "newBranch")
	}

	request := &projects.CreateBranchRequest{
		BaseGitRef:    baseBranch,
		NewBranchName: newBranchName,
	}

	path, err := getBranchPathV2(spaceId, projectID)
	if err != nil {
		return nil, err
	}

	response := &projects.GitReference{}
	if _, err := services.ApiAddWithResponseStatus(s.GetClient(), request, response, path, http.StatusOK); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *ProjectBranchesService) Get(spaceId, projectId string, branchesQuery ProjectBranchQuery) (*resources.Resources[*projects.GitReference], error) {
	v, _ := query.Values(branchesQuery)
	path, err := getBranchPath(spaceId, projectId)
	if err != nil {
		return nil, err
	}
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*projects.GitReference]), path)
	if err != nil {
		return &resources.Resources[*projects.GitReference]{}, err
	}

	return resp.(*resources.Resources[*projects.GitReference]), nil
}

func getBranchPathV2(spaceId string, projectId string) (string, error) {
	values := map[string]any{
		"spaceId":   spaceId,
		"projectId": projectId,
	}
	template, err := uritemplates.Parse(uritemplates.ProjectBranchesV2)
	if err != nil {
		return "", err
	}
	path, err := template.Expand(values)
	if err != nil {
		return "", err
	}
	return path, nil
}

func getBranchPath(spaceId string, projectId string) (string, error) {
	values := map[string]any{
		"spaceId":   spaceId,
		"projectId": projectId,
	}
	template, err := uritemplates.Parse(uritemplates.ProjectBranches)
	if err != nil {
		return "", err
	}
	path, err := template.Expand(values)
	if err != nil {
		return "", err
	}
	return path, nil
}

type ProjectBranchQuery struct {
	Refresh      string `uri:"refresh,omitempty" url:"refresh,omitempty"`
	SearchByName string `uri:"searchByName,omitempty" url:"searchByName,omitempty"`
	Skip         int    `uri:"skip,omitempty" url:"skip,omitempty"`
	Take         int    `uri:"take,omitempty" url:"take,omitempty"`
}
