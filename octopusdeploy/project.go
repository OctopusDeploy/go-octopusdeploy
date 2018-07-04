package octopusdeploy

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type ProjectService struct {
	sling *sling.Sling
}

func NewProjectService(sling *sling.Sling) *ProjectService {
	return &ProjectService{
		sling: sling,
	}
}

type Projects struct {
	Items []Project `json:"Items"`
	PagedResults
}

type Project struct {
	AutoCreateRelease               bool                        `json:"AutoCreateRelease"`
	AutoDeployReleaseOverrides      []AutoDeployReleaseOverride `json:"AutoDeployReleaseOverrides"`
	DefaultGuidedFailureMode        string                      `json:"DefaultGuidedFailureMode,omitempty"`
	DefaultToSkipIfAlreadyInstalled bool                        `json:"DefaultToSkipIfAlreadyInstalled"`
	DeploymentProcessID             string                      `json:"DeploymentProcessId"`
	Description                     string                      `json:"Description"`
	DiscreteChannelRelease          bool                        `json:"DiscreteChannelRelease"`
	ID                              string                      `json:"Id,omitempty"`
	IncludedLibraryVariableSetIds   []string                    `json:"IncludedLibraryVariableSetIds"`
	IsDisabled                      bool                        `json:"IsDisabled"`
	LifecycleID                     string                      `json:"LifecycleId"`
	Name                            string                      `json:"Name"`
	ProjectConnectivityPolicy       ProjectConnectivityPolicy   `json:"ProjectConnectivityPolicy"`
	ProjectGroupID                  string                      `json:"ProjectGroupId"`
	ReleaseCreationStrategy         ReleaseCreationStrategy     `json:"ReleaseCreationStrategy"`
	Slug                            string                      `json:"Slug"`
	Templates                       []ActionTemplateParameter   `json:"Templates,omitempty"`
	TenantedDeploymentMode          string                      `json:"TenantedDeploymentMode,omitempty"`
	VariableSetID                   string                      `json:"VariableSetId"`
	VersioningStrategy              VersioningStrategy          `json:"VersioningStrategy"`
}

func NewProject(name, lifeCycleID, projectGroupID string) *Project {
	return &Project{
		Name:           name,
		LifecycleID:    lifeCycleID,
		ProjectGroupID: projectGroupID,
		VersioningStrategy: VersioningStrategy{
			Template: "#{Octopus.Version.LastMajor}.#{Octopus.Version.LastMinor}.#{Octopus.Version.NextPatch}",
		},
	}
}

func (s *ProjectService) Get(projectid string) (*Project, error) {
	var project Project
	octopusDeployError := new(APIError)
	path := fmt.Sprintf("projects/%s", projectid)

	resp, err := s.sling.New().Get(path).Receive(&project, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return &project, nil
}

func (s *ProjectService) GetAll() (*[]Project, error) {
	var listOfProjects []Project
	path := fmt.Sprintf("projects")

	for {
		var projects Projects
		octopusDeployError := new(APIError)

		resp, err := s.sling.New().Get(path).Receive(&projects, &octopusDeployError)

		apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

		if apiErrorCheck != nil {
			return nil, apiErrorCheck
		}

		for _, project := range projects.Items {
			listOfProjects = append(listOfProjects, project)
		}

		if projects.PagedResults.Links.PageNext != "" {
			path = projects.PagedResults.Links.PageNext
		} else {
			break
		}
	}

	return &listOfProjects, nil // no more pages to go through
}

func (s *ProjectService) GetByName(projectName string) (*Project, error) {
	var foundProject Project
	projects, err := s.GetAll()

	if err != nil {
		return &foundProject, err
	}

	for _, project := range *projects {
		if project.Name == projectName {
			return &project, nil
		}
	}

	return &foundProject, fmt.Errorf("no project found with project name %s", projectName)
}

func (s *ProjectService) Add(project *Project) (*Project, error) {
	var created Project
	octopusDeployError := new(APIError)
	path := "projects"

	resp, err := s.sling.New().Post(path).BodyJSON(project).Receive(&created, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusCreated, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return &created, nil
}

func (s *ProjectService) Delete(projectid string) error {
	octopusDeployError := new(APIError)
	path := fmt.Sprintf("projects/%s", projectid)
	resp, err := s.sling.New().Delete(path).Receive(nil, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return apiErrorCheck
	}

	return nil
}

func (s *ProjectService) Update(project *Project) (*Project, error) {
	var updated Project
	octopusDeployError := new(APIError)
	path := fmt.Sprintf("projects/%s", project.ID)

	resp, err := s.sling.New().Put(path).BodyJSON(project).Receive(&updated, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return &updated, nil
}


