package projects_test

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// annoying junk to mock sling. TODO discuss with JB, he's probably done this a better way

type AnyDoer struct {
	impl func(req *http.Request) (*http.Response, error)
}

func NewAnyDoer(impl func(req *http.Request) (*http.Response, error)) *AnyDoer {
	return &AnyDoer{impl: impl}
}

// conformance with sling.Doer
func (f *AnyDoer) Do(req *http.Request) (*http.Response, error) {
	return f.impl(req)
}

// jsonDecoder decodes http response JSON into a JSON-tagged struct value.
type jsonDecoder struct{}

// Decode decodes the Response Body into the value pointed to by v.
// Caller must provide a non-nil v and close the resp.Body.
func (d jsonDecoder) Decode(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}

// Ideally we would have an integration test talking to the octopus server, however setting up VCS projects
// in integration tests is very expensive; unit tests are very nearly as good; the main source of bugs
// is likely to be deserialization errors; we can catch them here fairly easily.

func TestProjectGitReferencesTest(t *testing.T) {
	vcProject := projects.NewProject("proj-vc", "Lifecycles-1", "ProjectGroups-1")
	vcProject.ID = "Projects-1"
	vcProject.Links = map[string]string{
		constants.LinkTags:     "/api/Spaces-1/projects/Projects-1/git/tags{/name}{?skip,take,searchByName,refresh}",
		constants.LinkCommits:  "/api/Spaces-1/projects/Projects-1/git/commits{/hash}{?skip,take,refresh}",
		constants.LinkBranches: "/api/Spaces-1/projects/Projects-1/git/branches{/name}{?skip,take,searchByName,refresh}",
	}

	dbProject := projects.NewProject("proj-db", "Lifecycles-1", "ProjectGroups-1")
	dbProject.ID = "Projects-2"
	// only version controlled projects have the Tags, Commits and Branches links
	dbProject.Links = map[string]string{}

	fakeSling := &sling.Sling{}
	fakeSling.ResponseDecoder(jsonDecoder{})
	svc := projects.NewProjectService(fakeSling, "/api/Spaces-1/projects{/id}{?name,skip,ids,clone,take,partialName,clonedFromProjectId}", "", "", "", "")

	t.Run("can get collection of branches", func(t *testing.T) {
		// value captured from a real octopus server;
		// note: the real server returns more links, but because links is just a dump map[string]string we only need one thing to test the deserialization.
		responseString := `{
  "ItemType": "GitBranch", "TotalResults": 2, "ItemsPerPage": 30, "NumberOfPages": 1, "LastPageNumber": 0,
  "Items": [
    {
      "Name": "develop",
      "CanonicalName": "refs/heads/develop",
      "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2fheads%2fdevelop/deploymentprocesses" }
    },
    {
      "Name": "main",
      "CanonicalName": "refs/heads/main",
      "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2fheads%2fmain/deploymentprocesses" }
    }
  ],
  "Links": { "Self": "/api/Spaces-1/projects/Projects-1/git/branches?skip=0&take=30" }
}`
		fakeSling.Doer(NewAnyDoer(func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, "GET", req.Method)
			assert.Equal(t, "/api/Spaces-1/projects/Projects-1/git/branches", req.URL.String())
			return &http.Response{
				StatusCode:    http.StatusOK,
				Body:          ioutil.NopCloser(strings.NewReader(responseString)),
				ContentLength: int64(len(responseString)),
			}, nil
		}))

		branches, err := svc.GetGitBranches(vcProject)
		assert.Nil(t, err)
		assert.Equal(t, []*projects.GitReference{
			{
				Type:          projects.GitRefTypeBranch,
				Name:          "develop",
				CanonicalName: "refs/heads/develop",
				Links:         map[string]string{"DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2fheads%2fdevelop/deploymentprocesses"},
			},
			{
				Type:          projects.GitRefTypeBranch,
				Name:          "main",
				CanonicalName: "refs/heads/main",
				Links:         map[string]string{"DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2fheads%2fmain/deploymentprocesses"},
			},
		}, branches)
	})

	t.Run("can't get collection of branches from non version-controlled project", func(t *testing.T) {
		result, err := svc.GetGitBranches(dbProject)
		assert.Nil(t, result)
		assert.EqualError(t, err, "cannot get git branches on project proj-db; no Branches link. GetGitBranches requires a version-controlled project")
	})

	t.Run("can get collection of tags", func(t *testing.T) {
		responseString := `{
  "ItemType": "GitTag", "TotalResults": 2, "ItemsPerPage": 30, "NumberOfPages": 1, "LastPageNumber": 0,
  "Items": [
    {
      "Name": "v3",
      "CanonicalName": "refs/tags/v3",
      "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv3/deploymentprocesses" }
    },
    {
      "Name": "v5",
      "CanonicalName": "refs/tags/v5",
      "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv5/deploymentprocesses" }
    }
  ],
  "Links": { "Self": "/api/Spaces-1/projects/Projects-1/git/branches?skip=0&take=30" }
}
`
		fakeSling.Doer(NewAnyDoer(func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, "GET", req.Method)
			assert.Equal(t, "/api/Spaces-1/projects/Projects-1/git/tags", req.URL.String())
			return &http.Response{
				StatusCode:    http.StatusOK,
				Body:          ioutil.NopCloser(strings.NewReader(responseString)),
				ContentLength: int64(len(responseString)),
			}, nil
		}))

		branches, err := svc.GetGitTags(vcProject)
		assert.Nil(t, err)
		assert.Equal(t, []*projects.GitReference{
			{
				Type:          projects.GitRefTypeTag,
				Name:          "v3",
				CanonicalName: "refs/tags/v3",
				Links:         map[string]string{"DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv3/deploymentprocesses"},
			},
			{
				Type:          projects.GitRefTypeTag,
				Name:          "v5",
				CanonicalName: "refs/tags/v5",
				Links:         map[string]string{"DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv5/deploymentprocesses"},
			},
		}, branches)
	})

	t.Run("can't get collection of tags from non version-controlled project", func(t *testing.T) {
		result, err := svc.GetGitTags(dbProject)
		assert.Nil(t, result)
		assert.EqualError(t, err, "cannot get git tags on project proj-db; no Tags link. GetGitTags requires a version-controlled project")
	})

	t.Run("can get single branch", func(t *testing.T) {
		responseString := `{
  "Name": "develop",
  "CanonicalName": "refs/heads/develop",
  "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2fheads%2fdevelop/deploymentprocesses" }
}
`
		fakeSling.Doer(NewAnyDoer(func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, "GET", req.Method)
			assert.Equal(t, "/api/Spaces-1/projects/Projects-1/git/branches/develop", req.URL.String())
			return &http.Response{
				StatusCode:    http.StatusOK,
				Body:          ioutil.NopCloser(strings.NewReader(responseString)),
				ContentLength: int64(len(responseString)),
			}, nil
		}))

		branch, err := svc.GetGitBranch(vcProject, "develop")
		assert.Nil(t, err)
		assert.Equal(t, &projects.GitReference{
			Type:          projects.GitRefTypeBranch,
			Name:          "develop",
			CanonicalName: "refs/heads/develop",
			Links:         map[string]string{"DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2fheads%2fdevelop/deploymentprocesses"},
		}, branch)
	})

	t.Run("can't get single branch from non version-controlled project", func(t *testing.T) {
		result, err := svc.GetGitBranch(dbProject, "main")
		assert.Nil(t, result)
		assert.EqualError(t, err, "cannot get git branch on project proj-db; no Branches link. GetGitBranch requires a version-controlled project")
	})

	t.Run("can get single tag", func(t *testing.T) {
		responseString := `{
  "Name": "v5",
  "CanonicalName": "refs/tags/v5",
  "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv5/deploymentprocesses" }
}
`
		fakeSling.Doer(NewAnyDoer(func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, "GET", req.Method)
			assert.Equal(t, "/api/Spaces-1/projects/Projects-1/git/tags/v5", req.URL.String())
			return &http.Response{
				StatusCode:    http.StatusOK,
				Body:          ioutil.NopCloser(strings.NewReader(responseString)),
				ContentLength: int64(len(responseString)),
			}, nil
		}))

		branch, err := svc.GetGitTag(vcProject, "v5")
		assert.Nil(t, err)
		assert.Equal(t, &projects.GitReference{
			Type:          projects.GitRefTypeTag,
			Name:          "v5",
			CanonicalName: "refs/tags/v5",
			Links:         map[string]string{"DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv5/deploymentprocesses"},
		}, branch)
	})

	t.Run("can't get single tag from non version-controlled project", func(t *testing.T) {
		result, err := svc.GetGitTag(dbProject, "v3")
		assert.Nil(t, result)
		assert.EqualError(t, err, "cannot get git tag on project proj-db; no Tags link. GetGitTag requires a version-controlled project")
	})

	t.Run("can get single commit", func(t *testing.T) {
		responseString := `{
  "Name": "59d550fbdf82b83619a72fdbd331cc8fa3cb2f3c",
  "CanonicalName": "59d550fbdf82b83619a72fdbd331cc8fa3cb2f3c",
  "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/59d550fbdf82b83619a72fdbd331cc8fa3cb2f3c/deploymentprocesses" }
}
`
		fakeSling.Doer(NewAnyDoer(func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, "GET", req.Method)
			assert.Equal(t, "/api/Spaces-1/projects/Projects-1/git/commits/59d550fb", req.URL.String())
			return &http.Response{
				StatusCode:    http.StatusOK,
				Body:          ioutil.NopCloser(strings.NewReader(responseString)),
				ContentLength: int64(len(responseString)),
			}, nil
		}))

		branch, err := svc.GetGitCommit(vcProject, "59d550fb")
		assert.Nil(t, err)
		assert.Equal(t, &projects.GitReference{
			Type:          projects.GitRefTypeCommit,
			Name:          "59d550fbdf82b83619a72fdbd331cc8fa3cb2f3c",
			CanonicalName: "59d550fbdf82b83619a72fdbd331cc8fa3cb2f3c",
			Links:         map[string]string{"DeploymentProcess": "/api/Spaces-1/projects/Projects-1/59d550fbdf82b83619a72fdbd331cc8fa3cb2f3c/deploymentprocesses"},
		}, branch)
	})

	t.Run("can't get single commit from non version-controlled project", func(t *testing.T) {
		result, err := svc.GetGitCommit(dbProject, "59d550fb")
		assert.Nil(t, result)
		assert.EqualError(t, err, "cannot get git commit on project proj-db; no Commits link. GetGitCommit requires a version-controlled project")
	})

}
