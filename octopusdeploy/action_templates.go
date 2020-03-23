package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type ActionTemplateService struct {
	sling *sling.Sling
}

func NewActionTemplateService(sling *sling.Sling) *ActionTemplateService {
	return &ActionTemplateService{
		sling: sling,
	}
}

type ActionTemplates struct {
	Items []ActionTemplate `json:"Items"`
	PagedResults
}

type ActionTemplate struct {
	ID                        string                    `json:"Id"`
	Name                      string                    `json:"Name"`
	Description               string                    `json:"Description"`
	ActionType                string                    `json:"ActionType"`
	Version                   int                       `json:"Version"`
	CommunityActionTemplateId string                    `json:"CommunityActionTemplateId"`
	Properties                *ActionTemplateProperties `json:"Properties,omitempty"`
	Parameters                []ActionTemplateParameter `json:"Parameters,omitempty"`
}

type ActionTemplateProperties struct {
	ManualInstructions               string `json:"Octopus.Action.Manual.Instructions,omitempty"`
	ManualResponsibleTeamIds         string `json:"Octopus.Action.Manual.ResponsibleTeamIds,omitempty"`
	ManualBlockConcurrentDeployments bool   `json:"Octopus.Action.Manual.BlockConcurrentDeployments,string,omitempty"`
	RunOnServer                      string `json:"Octopus.Action.RunOnServer,omitempty"`
	ScriptSyntax                     string `json:"Octopus.Action.Script.Syntax,omitempty"`
	ScriptSource                     string `json:"Octopus.Action.Script.ScriptSource,omitempty"`
	ScriptBody                       string `json:"Octopus.Action.Script.ScriptBody,omitempty"`
}

func (p *ActionTemplate) Validate() error {
	validate := validator.New()

	err := validate.Struct(p)

	if err != nil {
		return err
	}

	return nil
}

func NewActionTemplate(name string, description string, actionType string) *ActionTemplate {
	return &ActionTemplate{
		Name:        name,
		Description: description,
		ActionType:  actionType,
	}
}

func (s *ActionTemplateService) Get(actionTemplateID string) (*ActionTemplate, error) {
	path := fmt.Sprintf("actiontemplates/%s", actionTemplateID)
	resp, err := apiGet(s.sling, new(ActionTemplate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*ActionTemplate), nil
}

func (s *ActionTemplateService) GetAll() (*[]ActionTemplate, error) {
	var at []ActionTemplate

	path := "actiontemplates"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(ActionTemplates), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*ActionTemplates)

		for _, item := range r.Items {
			at = append(at, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &at, nil
}

func (s *ActionTemplateService) GetByName(actionTemplateName string) (*ActionTemplate, error) {
	var notFound ActionTemplate
	ats, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, at := range *ats {
		if at.Name == actionTemplateName {
			return &at, nil
		}
	}

	return &notFound, fmt.Errorf("no actiontemplate found with name %s", actionTemplateName)
}

func (s *ActionTemplateService) Add(actionTemplate *ActionTemplate) (*ActionTemplate, error) {
	err := actionTemplate.Validate()
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, actionTemplate, new(ActionTemplate), "actiontemplates")

	if err != nil {
		return nil, err
	}

	return resp.(*ActionTemplate), nil
}

func (s *ActionTemplateService) Delete(actionTemplateID string) error {
	path := fmt.Sprintf("actiontemplates/%s", actionTemplateID)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

func (s *ActionTemplateService) Update(actionTemplate *ActionTemplate) (*ActionTemplate, error) {
	path := fmt.Sprintf("actiontemplates/%s", actionTemplate.ID)
	resp, err := apiUpdate(s.sling, actionTemplate, new(ActionTemplate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*ActionTemplate), nil
}
