package octopusdeploy

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type ProjectGroupService struct {
	sling *sling.Sling
}

func NewProjectGroupService(sling *sling.Sling) *ProjectGroupService {
	return &ProjectGroupService{
		sling: sling,
	}
}

type ProjectGroups struct {
	Items []ProjectGroup `json:"Items"`
	PagedResults
}

type ProjectGroup struct {

	// description
	Description string `json:"Description,omitempty"`

	// environment ids
	EnvironmentIds []string `json:"EnvironmentIds"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn string `json:"LastModifiedOn,omitempty"`

	// links
	Links Links `json:"Links,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// retention policy Id
	RetentionPolicyID string `json:"RetentionPolicyId,omitempty"`
}

func NewProjectGroup(name string) *ProjectGroup {
	return &ProjectGroup{
		Name:           name,
	}
}

func (p *ProjectGroupService) Get(projectGroupID string) (*ProjectGroup, error) {
	var projectGroup ProjectGroup
	octopusDeployError := new(APIError)
	path := fmt.Sprintf("projectgroups/%s", projectGroupID)

	resp, err := p.sling.New().Get(path).Receive(&projectGroup, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return &projectGroup, err
}

func (p *ProjectGroupService) GetAll() (*[]ProjectGroup, error) {
	var listOfProjectGroups []ProjectGroup
	path := fmt.Sprintf("projectgroups")

	for {
		var projectGroups ProjectGroups
		octopusDeployError := new(APIError)

		resp, err := p.sling.New().Get(path).Receive(&projectGroups, &octopusDeployError)

		apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

		if apiErrorCheck != nil {
			return nil, apiErrorCheck
		}

		for _, projectGroup := range projectGroups.Items {
			listOfProjectGroups = append(listOfProjectGroups, projectGroup)
		}

		if projectGroups.PagedResults.Links.PageNext != "" {
			path = projectGroups.PagedResults.Links.PageNext
		} else {
			break
		}
	}

	return &listOfProjectGroups, nil // no more pages to go through
}

func (p *ProjectGroupService) Add(projectGroup *ProjectGroup) (*ProjectGroup, error) {
	var created ProjectGroup
	octopusDeployError := new(APIError)
	path := "projectgroups"
	resp, err := p.sling.New().Post(path).BodyJSON(projectGroup).Receive(&created, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusCreated, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return &created, nil
}

func (p *ProjectGroupService) Delete(projectGroupID string) error {
	octopusDeployError := new(APIError)
	path := fmt.Sprintf("projectgroups/%s", projectGroupID)
	resp, err := p.sling.New().Delete(path).Receive(nil, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return apiErrorCheck
	}

	return nil
}

func (p *ProjectGroupService) Update(projectGroup *ProjectGroup) (*ProjectGroup, error) {
	var updated ProjectGroup
	octopusDeployError := new(APIError)
	path := fmt.Sprintf("projectgroups/%s", projectGroup.ID)

	resp, err := p.sling.New().Put(path).BodyJSON(projectGroup).Receive(&updated, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return &updated, nil
}
