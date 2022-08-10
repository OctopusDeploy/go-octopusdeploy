package projects

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

type GitRefType string

const (
	GitRefTypeBranch = GitRefType("GitBranch")
	GitRefTypeTag    = GitRefType("GitTag")
	GitRefTypeCommit = GitRefType("GitCommit")
)

// GitReference represents the data returned from the Octopus Server
// relating to a git branch or tag in a version controlled project.
// Both branches and tags share the same resource format
type GitReference struct {
	Type          GitRefType        `json:"-"`                       // added by the client library in case you need to disambiguate branches/tags in a single collection
	Name          string            `json:"Name,omitempty"`          // display name of the git item e.g. "main"
	CanonicalName string            `json:"CanonicalName,omitempty"` // underlying git reference e.g. "/refs/heads/main"
	Links         map[string]string `json:"Links,omitempty"`
}

// getGitReference loads a singular item from either
// - api/Spaces-1/projects/Projects-1/git/branches/NAME
// - api/Spaces-1/projects/Projects-1/git/tags/NAME
// - api/Spaces-1/projects/Projects-1/git/commits/HASH
func getGitReference(sling *sling.Sling, linkTemplate string, itemType GitRefType, templateParameters map[string]interface{}) (*GitReference, error) {
	parsedTemplate, err := uritemplates.Parse(linkTemplate)
	if err != nil {
		return nil, err
	}
	linkPath, err := parsedTemplate.Expand(templateParameters)
	if err != nil {
		return nil, err
	}

	result := new(GitReference)
	_, err = api.ApiGet(sling, result, linkPath)
	if err != nil {
		return nil, err
	}
	result.Type = itemType
	return result, nil
}

// getGitBranchesOrTags loads a collection of items from either
// - api/Spaces-1/projects/Projects-1/git/branches
// - api/Spaces-1/projects/Projects-1/git/tags
func getGitBranchesOrTags(sling *sling.Sling, linkTemplate string) ([]*GitReference, error) {
	tagsUrl, err := uritemplates.Parse(linkTemplate)
	if err != nil {
		return nil, err
	}
	linkPath, err := tagsUrl.Expand(make(map[string]interface{}, 0))
	if err != nil {
		return nil, err
	}

	result := make([]*GitReference, 0, 4)
	loadNextPage := true

	for loadNextPage {
		resp, err := api.ApiGet(sling, new(resources.Resources[*GitReference]), linkPath)
		if err != nil {
			return nil, err
		}

		r := resp.(*resources.Resources[*GitReference])
		if r.ItemType != string(GitRefTypeBranch) && r.ItemType != string(GitRefTypeTag) {
			return nil, fmt.Errorf("server returned unsupported git reference type %s", r.ItemType)
		}
		itemType := GitRefType(r.ItemType)

		for _, item := range r.Items {
			item.Type = itemType
			result = append(result, item)
		}
		linkPath, loadNextPage = services.LoadNextPage(r.PagedResults)
	}

	return result, nil
}

func (s *ProjectService) GetGitBranches(project *Project) ([]*GitReference, error) {
	if project == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("project")
	}
	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	return getGitBranchesOrTags(s.GetClient(), project.Links[constants.LinkBranches])
}

func (s *ProjectService) GetGitTags(project *Project) ([]*GitReference, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("GetGitTags", "project")
	}
	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	return getGitBranchesOrTags(s.GetClient(), project.Links[constants.LinkTags])
}

func (s *ProjectService) GetGitBranch(project *Project, name string) (*GitReference, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("GetGitBranch", "project")
	}
	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}
	return getGitReference(s.GetClient(), project.Links[constants.LinkBranches], GitRefTypeBranch, map[string]interface{}{"name": name})
}

func (s *ProjectService) GetGitTag(project *Project, name string) (*GitReference, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("GetGitTag", "project")
	}
	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}
	return getGitReference(s.GetClient(), project.Links[constants.LinkTags], GitRefTypeTag, map[string]interface{}{"name": name})
}

func (s *ProjectService) GetGitCommit(project *Project, hash string) (*GitReference, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("GetGitCommit", "project")
	}
	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}
	return getGitReference(s.GetClient(), project.Links[constants.LinkCommits], GitRefTypeCommit, map[string]interface{}{"hash": hash})
}
