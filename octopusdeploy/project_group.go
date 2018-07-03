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

	if err != nil {
		return nil, fmt.Errorf("cannot get projectgroup id %s from server. failure from http client %v", projectGroupID, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrItemNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot get projectgroup id %s from server. response from server %s", projectGroupID, resp.Status)
	}

	return &projectGroup, err
}

func (p *ProjectGroupService) GetAll() (*[]ProjectGroup, error) {
	var listOfProjectGroups []ProjectGroup
	path := fmt.Sprintf("projectgroups")

	for {
		var projectGroups ProjectGroups
		var octopusDeployError APIError

		resp, err := p.sling.New().Get(path).Receive(&projectGroups, &octopusDeployError)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if octopusDeployError.Errors != nil {
			return nil, fmt.Errorf("cannot get all projectgroups. response from octopusdeploy %s: ", octopusDeployError.Errors)
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("cannot get all projectgroups. response from server %s", resp.Status)
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
	var octopusDeployError APIError
	resp, err := p.sling.New().Post("projectgroups").BodyJSON(projectGroup).Receive(&created, &octopusDeployError)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if octopusDeployError.Errors != nil {
		return nil, fmt.Errorf("cannot add projectgroup. response from octopus deploy %s: ", octopusDeployError.Errors)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("cannot add projectgroup. response from server %s, req %s", resp.Status, resp.Request.URL)
	}

	return &created, nil
}

func (p *ProjectGroupService) Delete(projectGroupID string) error {
	path := fmt.Sprintf("projectgroups/%s", projectGroupID)
	req, err := p.sling.New().Delete(path).Request()

	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return ErrItemNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("cannot delete projectgroup. response from server %s", resp.Status)
	}

	return nil
}

func (p *ProjectGroupService) Update(projectGroup *ProjectGroup) (*ProjectGroup, error) {
	var updated ProjectGroup
	var octopusDeployError APIError

	path := fmt.Sprintf("projectgroups/%s", projectGroup.ID)
	resp, err := p.sling.New().Put(path).BodyJSON(projectGroup).Receive(&updated, &octopusDeployError)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if octopusDeployError.Errors != nil {
		return nil, fmt.Errorf("cannot update projectgroup. response from octopusdeploy %s: ", octopusDeployError.Errors)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot update projectgroup at url %s. response from server %s", resp.Request.URL, resp.Status)
	}

	return &updated, nil
}
