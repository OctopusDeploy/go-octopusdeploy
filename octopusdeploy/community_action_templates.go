package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type CommunityActionTemplateService struct {
	sling   *sling.Sling
	spaceID string
}

func NewCommunityActionTemplateService(sling *sling.Sling, spaceID string) *CommunityActionTemplateService {
	return &CommunityActionTemplateService{
		sling:   sling,
		spaceID: spaceID,
	}
}

type CommunityActionTemplates struct {
	Items []CommunityActionTemplate `json:"Items"`
	PagedResults
}

type CommunityActionTemplate struct {
	ID           string `json:"Id"`
	Name         string `json:"Name,omitempty"`
	Label        string `json:"Label,omitempty"`
	HelpText     string `json:"HelpText,omitempty"`
	DefaultValue string `json:"DefaultValue"`
}

func (p *CommunityActionTemplate) Validate() error {
	validate := validator.New()

	err := validate.Struct(p)

	if err != nil {
		return err
	}

	return nil
}

func NewCommunityActionTemplate(id string) *CommunityActionTemplate {
	return &CommunityActionTemplate{
		ID: id,
	}
}

func (s *CommunityActionTemplateService) Get(communityActionTemplateID string) (*CommunityActionTemplate, error) {
	path := fmt.Sprintf("communityactiontemplates/%s", communityActionTemplateID)
	resp, err := apiGet(s.sling, new(CommunityActionTemplate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*CommunityActionTemplate), nil
}

func (s *CommunityActionTemplateService) GetAll() (*[]CommunityActionTemplate, error) {
	var cat []CommunityActionTemplate

	path := "communityactiontemplates"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(CommunityActionTemplates), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*CommunityActionTemplates)

		for _, item := range r.Items {
			cat = append(cat, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &cat, nil
}

func (s *CommunityActionTemplateService) GetByName(communityActionTemplateName string) (*CommunityActionTemplate, error) {
	var notFound CommunityActionTemplate
	cats, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, cat := range *cats {
		if cat.Name == communityActionTemplateName {
			return &cat, nil
		}
	}

	return &notFound, fmt.Errorf("no communityactiontemplate found with name %s", communityActionTemplateName)
}

func (s *CommunityActionTemplateService) Add(communityActionTemplateID string) (*CommunityActionTemplate, error) {
	create := fmt.Sprintf("communityactiontemplates/%s/installation/%s", communityActionTemplateID, s.spaceID)
	resp, err := apiAdd(s.sling, nil, new(CommunityActionTemplate), create)

	if err != nil {
		return nil, err
	}

	return resp.(*CommunityActionTemplate), nil
}
