package deployments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

type DeploymentProcessService struct {
	services.Service
}

func NewDeploymentProcessService(sling *sling.Sling, uriTemplate string) *DeploymentProcessService {
	return &DeploymentProcessService{
		Service: services.NewService(constants.ServiceDeploymentProcessesService, sling, uriTemplate),
	}
}

// Get returns the deployment process that matches the input project and
// a git reference. If one cannot be found, it returns nil and an error.
//
// Deprecated: use deployments.GetProcessByGitRef
func (s *DeploymentProcessService) Get(project *projects.Project, gitRef string) (*DeploymentProcess, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("Get", "project")
	}

	if project.PersistenceSettings == nil || project.PersistenceSettings.Type() != projects.PersistenceSettingsTypeVersionControlled {
		return s.GetByID(project.DeploymentProcessID)
	}

	gitPersistenceSettings := project.PersistenceSettings.(projects.GitPersistenceSettings)

	if len(gitRef) <= 0 {
		gitRef = gitPersistenceSettings.DefaultBranch()
	}

	template, _ := uritemplates.Parse(project.Links["DeploymentProcess"])
	path, _ := template.Expand(map[string]interface{}{"gitRef": gitRef})

	resp, err := api.ApiGet(s.GetClient(), new(DeploymentProcess), path)
	if err != nil {
		return nil, err
	}

	deploymentProcess := resp.(*DeploymentProcess)
	deploymentProcess.Branch = gitRef

	return deploymentProcess, err
}

// Deprecated: Use deployments.GetDeploymentProcessTemplate
func (s *DeploymentProcessService) GetTemplate(deploymentProcess *DeploymentProcess, channelID string, releaseID string) (*DeploymentProcessTemplate, error) {
	if deploymentProcess == nil {
		return nil, internal.CreateInvalidParameterError("GetTemplate", "deploymentProcess")
	}

	template, _ := uritemplates.Parse(deploymentProcess.Links["Template"])

	values := map[string]interface{}{}

	if len(channelID) > 0 {
		values["channel"] = channelID
	}

	if len(releaseID) > 0 {
		values["releaseId"] = releaseID
	}

	path, _ := template.Expand(values)

	resp, err := api.ApiGet(s.GetClient(), new(DeploymentProcessTemplate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentProcessTemplate), nil
}

// GetAll returns all deployment processes. If none can be found or an error
// occurs, it returns an empty collection.
//
// Deprecated: Use deployments.GetAllDeploymentProcesses
func (s *DeploymentProcessService) GetAll() ([]*DeploymentProcess, error) {
	path, err := services.GetPath(s)
	if err != nil {
		return []*DeploymentProcess{}, err
	}

	return services.GetPagedResponse[DeploymentProcess](s, path)
}

// GetByID returns the deployment process that matches the input ID. If one
// cannot be found, it returns nil and an error.
//
// Deprecated: Use deployments.GetDeploymentProcessByID
func (s *DeploymentProcessService) GetByID(id string) (*DeploymentProcess, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(DeploymentProcess), path)
	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentProcess), nil
}

// Update modifies a deployment process based on the one provided as input.
//
// Deprecated: Use deployments.Update
func (s *DeploymentProcessService) Update(deploymentProcess *DeploymentProcess) (*DeploymentProcess, error) {
	if deploymentProcess == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, "deploymentProcess")
	}

	resp, err := services.ApiUpdate(s.GetClient(), deploymentProcess, new(DeploymentProcess), deploymentProcess.Links["Self"])
	if err != nil {
		return nil, err
	}

	return resp.(*DeploymentProcess), nil
}

// ----- Experimental --------

const (
	deploymentProcessesTemplate = "/api/{spaceId}/deploymentprocesses{/id}{?skip,take,ids}"
)

// GetDeploymentProcessByID fetches a deployment process. This may either be the project level process (template),
// or a process snapshot from a Release, depending on the value of ID
func GetDeploymentProcessByID(client newclient.Client, spaceID string, ID string) (*DeploymentProcess, error) {
	return newclient.GetByID[DeploymentProcess](client, deploymentProcessesTemplate, spaceID, ID)
}

// GetDeploymentProcessByGitRef returns the deployment process that matches the input project and
// a git reference.
func GetDeploymentProcessByGitRef(client newclient.Client, spaceID string, project *projects.Project, gitRef string) (*DeploymentProcess, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("GetByGitRef", "project")
	}

	if project.PersistenceSettings == nil || project.PersistenceSettings.Type() != projects.PersistenceSettingsTypeVersionControlled {
		return GetDeploymentProcessByID(client, spaceID, project.DeploymentProcessID)
	}

	gitPersistenceSettings := project.PersistenceSettings.(projects.GitPersistenceSettings)

	if len(gitRef) <= 0 {
		gitRef = gitPersistenceSettings.DefaultBranch()
	}

	// TODO: remove use of links
	template, _ := uritemplates.Parse(project.Links["DeploymentProcess"])
	path, _ := template.Expand(map[string]interface{}{"gitRef": gitRef})

	deploymentProcess, err := newclient.Get[DeploymentProcess](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	deploymentProcess.Branch = gitRef

	return deploymentProcess, err
}

// UpdateDeploymentProcess modifies a deployment process based on the one provided as input.
func UpdateDeploymentProcess(client newclient.Client, deploymentProcess *DeploymentProcess) (*DeploymentProcess, error) {
	if deploymentProcess == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, "deploymentProcess")
	}

	// TODO: remove use of links
	deploymentProcessTemplate, err := newclient.Update[DeploymentProcess](client, deploymentProcess.Links["Self"], deploymentProcess.SpaceID, deploymentProcess.ID, deploymentProcess)
	if err != nil {
		return nil, err
	}

	return deploymentProcessTemplate, nil
}

func GetDeploymentProcessTemplate(client newclient.Client, deploymentProcess *DeploymentProcess, channelID string, releaseID string) (*DeploymentProcessTemplate, error) {
	if deploymentProcess == nil {
		return nil, internal.CreateInvalidParameterError("GetTemplate", "deploymentProcess")
	}

	template, _ := uritemplates.Parse(deploymentProcess.Links["Template"])

	values := map[string]interface{}{}

	if len(channelID) > 0 {
		values["channel"] = channelID
	}

	if len(releaseID) > 0 {
		values["releaseId"] = releaseID
	}

	path, _ := template.Expand(values)

	deploymentProcessTemplate, err := newclient.Get[DeploymentProcessTemplate](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return deploymentProcessTemplate, nil
}

// GetAllDeploymentProcesses returns all deployment processes. If none can be found or an error
// occurs, it returns an empty collection.
func GetAllDeploymentProcesses(client newclient.Client, spaceID string) ([]*DeploymentProcess, error) {
	return newclient.GetAll[DeploymentProcess](client, deploymentProcessesTemplate, spaceID)
}
