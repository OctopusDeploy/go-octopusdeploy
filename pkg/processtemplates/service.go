package processtemplates

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const (
	template          = "/api/platformhub/{gitRef}/processtemplates{/slug}{?skip,take}"
	summariesTemplate = "/api/platformhub/{gitRef}/processtemplates/summaries{?skip,take,partialName}"
)

// ProcessTemplatesQuery represents query parameters for listing process templates.
type ProcessTemplatesQuery struct {
	GitRef string `uri:"gitRef" json:"gitRef"`
	Skip   int    `uri:"skip,omitempty" json:"skip,omitempty"`
	Take   int    `uri:"take,omitempty" json:"take,omitempty"`
}

// ProcessTemplatesQueryResult is a paginated collection of process templates.
type ProcessTemplatesQueryResult struct {
	ProcessTemplates []*ProcessTemplate `json:"ProcessTemplates"`
	TotalResults     int                `json:"TotalResults"`
	ItemsPerPage     int                `json:"ItemsPerPage"`
}

// SummariesQuery represents query parameters for listing process template summaries.
type SummariesQuery struct {
	GitRef      string `uri:"gitRef" json:"gitRef"`
	PartialName string `uri:"partialName,omitempty" json:"partialName,omitempty"`
	Skip        int    `uri:"skip,omitempty" json:"skip,omitempty"`
	Take        int    `uri:"take,omitempty" json:"take,omitempty"`
}

// SummariesQueryResult is a paginated collection of process template summaries.
type SummariesQueryResult struct {
	ProcessTemplateSummaries  []*Summary `json:"ProcessTemplateSummaries"`
	TotalResults              int        `json:"TotalResults"`
	ItemsPerPage              int        `json:"ItemsPerPage"`
	TotalNoOfProcessTemplates int        `json:"TotalNoOfProcessTemplates"`
}

// List returns a paginated collection of process templates, including their steps and
// parameters. Prefer ListSummaries for listings.
func List(client newclient.Client, query ProcessTemplatesQuery) (*ProcessTemplatesQueryResult, error) {
	if internal.IsEmpty(query.GitRef) {
		return nil, internal.CreateInvalidParameterError("List", "GitRef")
	}

	path, err := client.URITemplateCache().Expand(template, query)
	if err != nil {
		return nil, err
	}

	return newclient.Get[ProcessTemplatesQueryResult](client.HttpSession(), path)
}

// ListSummaries returns a paginated collection of process template summaries.
func ListSummaries(client newclient.Client, query SummariesQuery) (*SummariesQueryResult, error) {
	if internal.IsEmpty(query.GitRef) {
		return nil, internal.CreateInvalidParameterError("ListSummaries", "GitRef")
	}

	path, err := client.URITemplateCache().Expand(summariesTemplate, query)
	if err != nil {
		return nil, err
	}

	return newclient.Get[SummariesQueryResult](client.HttpSession(), path)
}

// GetBySlug returns the process template that matches the given slug on the given Git reference.
func GetBySlug(client newclient.Client, gitRef string, slug string) (*ProcessTemplate, error) {
	if internal.IsEmpty(gitRef) {
		return nil, internal.CreateInvalidParameterError("GetBySlug", "gitRef")
	}
	if internal.IsEmpty(slug) {
		return nil, internal.CreateInvalidParameterError("GetBySlug", "slug")
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{"gitRef": gitRef, "slug": slug})
	if err != nil {
		return nil, err
	}

	return newclient.Get[ProcessTemplate](client.HttpSession(), path)
}

// createProcessTemplateCommand creates an empty process template. Steps and parameters
// cannot be set at create time - they are authored afterwards in Git or the portal.
type createProcessTemplateCommand struct {
	GitRef string `json:"GitRef"`
	Name   string `json:"Name"`
	// Description is the process template's description.
	Description string `json:"Description,omitempty"`
	// ChangeDescription becomes the git commit message for the create.
	ChangeDescription string `json:"ChangeDescription,omitempty"`
}

// Add creates a new, empty process template on the given Git reference. changeDescription
// becomes the git commit message.
func Add(client newclient.Client, gitRef string, name string, description string, changeDescription string) (*ProcessTemplate, error) {
	if internal.IsEmpty(gitRef) {
		return nil, internal.CreateInvalidParameterError("Add", "gitRef")
	}
	if internal.IsEmpty(name) {
		return nil, internal.CreateInvalidParameterError("Add", "name")
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{"gitRef": gitRef})
	if err != nil {
		return nil, err
	}

	command := createProcessTemplateCommand{
		GitRef:            gitRef,
		Name:              name,
		Description:       description,
		ChangeDescription: changeDescription,
	}

	return newclient.Post[ProcessTemplate](client.HttpSession(), path, command)
}
