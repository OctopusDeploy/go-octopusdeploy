package octopusdeploy

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type ProjectTriggerService struct {
	sling *sling.Sling
}

func NewProjectTriggerService(sling *sling.Sling) *ProjectTriggerService {
	return &ProjectTriggerService{
		sling: sling,
	}
}

type ProjectTriggers struct {
	Items []ProjectTrigger `json:"Items"`
	PagedResults
}

type ProjectTrigger struct {
	Action     ProjectTriggerAction `json:"Action"`
	Filter     ProjectTriggerFilter `json:"Filter"`
	ID         string               `json:"Id,omitempty"`
	IsDisabled bool                 `json:"IsDisabled,omitempty"`
	Name       string               `json:"Name"`
	ProjectID  string               `json:"ProjectId,omitempty"`
}

type ProjectTriggerFilter struct {
	DateOfMonth         string   `json:"DateOfMonth"`
	DayNumberOfMonth    string   `json:"DayNumberOfMonth"`
	DayOfWeek           string   `json:"DayOfWeek"`
	EnvironmentIds      []string `json:"EnvironmentIds,omitempty"`
	EventCategories     []string `json:"EventCategories,omitempty"`
	EventGroups         []string `json:"EventGroups,omitempty"`
	FilterType          string   `json:"FilterType"`
	MonthlyScheduleType string   `json:"MonthlyScheduleType"`
	Roles               []string `json:"Roles"`
	StartTime           string   `json:"StartTime"`
	Timezone            string   `json:"Timezone"`
}

type ProjectTriggerAction struct {
	ActionType                                 string `json:"ActionType"`
	DestinationEnvironmentID                   string `json:"DestinationEnvironmentId"`
	ShouldRedeployWhenMachineHasBeenDeployedTo bool   `json:"ShouldRedeployWhenMachineHasBeenDeployedTo"`
	SourceEnvironmentID                        string `json:"SourceEnvironmentId"`
}

func (t *ProjectTrigger) AddEventGroups(eventGroups []string) {
	for _, e := range eventGroups {
		t.Filter.EventGroups = append(t.Filter.EventGroups, e)
	}
}

func (t *ProjectTrigger) AddEventCategories(eventCategories []string) {
	for _, e := range eventCategories {
		t.Filter.EventCategories = append(t.Filter.EventCategories, e)
	}
}

func NewProjectTrigger(name, projectID string, shouldRedeploy bool, roles, eventGroups, eventCategories []string) *ProjectTrigger {
	return &ProjectTrigger{
		Action: ProjectTriggerAction{
			ActionType: "AutoDeploy",
			ShouldRedeployWhenMachineHasBeenDeployedTo: shouldRedeploy,
		},
		Filter: ProjectTriggerFilter{
			EventCategories: eventCategories,
			EventGroups:     eventGroups,
			FilterType:      "MachineFilter",
			Roles:           roles,
		},
		Name:      name,
		ProjectID: projectID,
	}
}

func (s *ProjectTriggerService) Get(projectTriggerID string) (*ProjectTrigger, error) {
	path := fmt.Sprintf("projecttriggers/%s", projectTriggerID)

	resp, err := apiGet(s.sling, new(ProjectTrigger), path)

	if err != nil {
		return nil, err
	}

	return resp.(*ProjectTrigger), nil
}

func (s *ProjectTriggerService) GetByProjectID(projectID string) (*[]ProjectTrigger, error) {
	var triggersByProject []ProjectTrigger

	triggers, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, pt := range *triggers {
		triggersByProject = append(triggersByProject, pt)
	}

	return &triggersByProject, nil
}

func (s *ProjectTriggerService) GetAll() (*[]ProjectTrigger, error) {
	var listOfProjectTriggers []ProjectTrigger
	path := fmt.Sprintf("projecttriggers")

	for {
		var projectTriggers ProjectTriggers
		octopusDeployError := new(APIError)

		resp, err := s.sling.New().Get(path).Receive(&projectTriggers, &octopusDeployError)

		apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

		if apiErrorCheck != nil {
			return nil, apiErrorCheck
		}

		for _, projectTrigger := range projectTriggers.Items {
			listOfProjectTriggers = append(listOfProjectTriggers, projectTrigger)
		}

		if projectTriggers.PagedResults.Links.PageNext != "" {
			path = projectTriggers.PagedResults.Links.PageNext
		} else {
			break
		}
	}

	return &listOfProjectTriggers, nil // no more pages to go through
}

func (s *ProjectTriggerService) Add(projectTrigger *ProjectTrigger) (*ProjectTrigger, error) {
	resp, err := apiAdd(s.sling, projectTrigger, new(ProjectTrigger), "projecttriggers")

	if err != nil {
		return nil, err
	}

	return resp.(*ProjectTrigger), nil
}

func (s *ProjectTriggerService) Delete(projectTriggerID string) error {
	octopusDeployError := new(APIError)
	path := fmt.Sprintf("projecttriggers/%s", projectTriggerID)
	resp, err := s.sling.New().Delete(path).Receive(nil, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return apiErrorCheck
	}

	return nil
}

func (s *ProjectTriggerService) Update(projectTrigger *ProjectTrigger) (*ProjectTrigger, error) {
	var updated ProjectTrigger
	octopusDeployError := new(APIError)
	path := fmt.Sprintf("projecttriggers/%s", projectTrigger.ID)

	resp, err := s.sling.New().Put(path).BodyJSON(projectTrigger).Receive(&updated, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return &updated, nil
}
