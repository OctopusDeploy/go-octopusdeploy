package octopusdeploy

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type ProjectsService struct {
	sling *sling.Sling
}

func NewProjectService(sling *sling.Sling) *ProjectsService {
	return &ProjectsService{
		sling: sling,
	}
}

type Projects struct {
	Items []Project `json:"Items"`
	PagedResults
}

type Project struct {
	AutoCreateRelease               bool                                `json:"AutoCreateRelease"`
	AutoDeployReleaseOverrides      []AutoDeployReleaseOverrideResource `json:"AutoDeployReleaseOverrides"`
	DefaultGuidedFailureMode        string                              `json:"DefaultGuidedFailureMode,omitempty"`
	DefaultToSkipIfAlreadyInstalled bool                                `json:"DefaultToSkipIfAlreadyInstalled"`
	DeploymentProcessID             string                              `json:"DeploymentProcessId"`
	Description                     string                              `json:"Description"`
	DiscreteChannelRelease          bool                                `json:"DiscreteChannelRelease"`
	ID                              string                              `json:"Id,omitempty"`
	IncludedLibraryVariableSetIds   []string                            `json:"IncludedLibraryVariableSetIds"`
	IsDisabled                      bool                                `json:"IsDisabled"`
	LifecycleID                     string                              `json:"LifecycleId"`
	Name                            string                              `json:"Name"`
	ProjectConnectivityPolicy       ProjectConnectivityPolicy           `json:"ProjectConnectivityPolicy"`
	ProjectGroupID                  string                              `json:"ProjectGroupId"`
	ReleaseCreationStrategy         ReleaseCreationStrategyResource     `json:"ReleaseCreationStrategy"`
	Slug                            string                              `json:"Slug"`
	Templates                       []ActionTemplateParameterResource   `json:"Templates,omitempty"`
	TenantedDeploymentMode          string                              `json:"TenantedDeploymentMode,omitempty"`
	VariableSetID                   string                              `json:"VariableSetId"`
	VersioningStrategy              VersioningStrategyResource          `json:"VersioningStrategy"`
}

func (s *ProjectsService) Get(projectid string) (Project, error) {
	project := new(Project)
	octopusDeployError := new(OctopusDeployError)
	path := fmt.Sprintf("api/projects/%s", projectid)

	resp, err := s.sling.New().Get(path).Receive(project, octopusDeployError)

	if err != nil {
		return *project, fmt.Errorf("cannot get project id %s from server. failure from http client %v", projectid, err)
	}

	if resp.StatusCode != http.StatusOK {
		return *project, fmt.Errorf("cannot get project id %s from server. response from server %s", projectid, resp.Status)
	}

	return *project, err
}

func (s *ProjectsService) GetAll() ([]Project, error) {
	var listOfProjects []Project
	path := fmt.Sprintf("api/projects")

	for {
		projects := new(Projects)
		octopusDeployError := new(OctopusDeployError)

		resp, err := s.sling.New().Get(path).Receive(projects, octopusDeployError)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Response: %s", resp.Status)
		fmt.Printf("Total Results: %d", projects.NumberOfPages)

		for _, project := range projects.Items {
			listOfProjects = append(listOfProjects, project)
		}

		if projects.PagedResults.Links.PageNext != "" {
			fmt.Printf("More pages to go! Next link: %s", projects.PagedResults.Links.PageNext)
			path = projects.PagedResults.Links.PageNext
		} else {
			break
		}
	}

	return listOfProjects, nil // no more pages to go through
}

func (s *ProjectsService) Add(project *Project) (Project, error) {
	var created Project
	path := fmt.Sprintf("api/projects")
	resp, err := s.sling.New().Post(path).BodyJSON(project).ReceiveSuccess(&created)

	defer resp.Body.Close()

	if err != nil {
		return created, err
	}

	if resp.StatusCode != http.StatusAccepted {
		return created, fmt.Errorf("cannot create project. response from server %s", resp.Status)
	}

	return created, nil
}
