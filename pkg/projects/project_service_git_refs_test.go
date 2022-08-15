package projects_test

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	testutil "github.com/OctopusDeploy/go-octopusdeploy/v2/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

	s := testutil.NewMockHttpServer()
	svc := projects.NewProjectService(s.Sling(), "/api/Spaces-1/projects{/id}{?name,skip,ids,clone,take,partialName,clonedFromProjectId}", "", "", "", "")

	t.Run("can get collection of branches", func(t *testing.T) {
		receiver := testutil.GoBegin2(func() ([]*projects.GitReference, error) {
			return svc.GetGitBranches(vcProject)
		})

		// value captured from a real octopus server;
		// note: the real server returns more links, but because links is just a dump map[string]string we only need one thing to test the deserialization.
		s.ExpectRequest(t, "GET", "/api/Spaces-1/projects/Projects-1/git/branches").RespondWithText(`{
  "ItemType": "GitBranch", "TotalResults": 2, "ItemsPerPage": 30, "NumberOfPages": 1, "LastPageNumber": 0,
  "Items": [
    {
      "Name": "develop", "CanonicalName": "refs/heads/develop",
      "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2fheads%2fdevelop/deploymentprocesses" }
    },
    {
      "Name": "main", "CanonicalName": "refs/heads/main",
      "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2fheads%2fmain/deploymentprocesses" }
    }
  ],
  "Links": { "Self": "/api/Spaces-1/projects/Projects-1/git/branches?skip=0&take=30" }
}`)

		branches, err := testutil.ReceivePair(receiver)

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
		receiver := testutil.GoBegin2(func() ([]*projects.GitReference, error) {
			return svc.GetGitTags(vcProject)
		})

		s.ExpectRequest(t, "GET", "/api/Spaces-1/projects/Projects-1/git/tags").RespondWithText(`{
 "ItemType": "GitTag", "TotalResults": 2, "ItemsPerPage": 30, "NumberOfPages": 1, "LastPageNumber": 0,
 "Items": [
   {
	 "Name": "v3", "CanonicalName": "refs/tags/v3",
	 "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv3/deploymentprocesses" }
   },
   {
	 "Name": "v5", "CanonicalName": "refs/tags/v5",
	 "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv5/deploymentprocesses" }
   }
 ],
 "Links": { "Self": "/api/Spaces-1/projects/Projects-1/git/branches?skip=0&take=30" }
}`)

		tags, err := testutil.ReceivePair(receiver)
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
		}, tags)
	})

	t.Run("can't get collection of tags from non version-controlled project", func(t *testing.T) {
		result, err := svc.GetGitTags(dbProject)
		assert.Nil(t, result)
		assert.EqualError(t, err, "cannot get git tags on project proj-db; no Tags link. GetGitTags requires a version-controlled project")
	})

	t.Run("can get collection of tags spanning multiple pages", func(t *testing.T) {
		receiver := testutil.GoBegin2(func() ([]*projects.GitReference, error) {
			return svc.GetGitTags(vcProject)
		})

		// page size of 2; first page
		s.ExpectRequest(t, "GET", "/api/Spaces-1/projects/Projects-1/git/tags").RespondWithText(`{
 "ItemType": "GitTag", "TotalResults": 3, "ItemsPerPage": 2, "NumberOfPages": 2, "LastPageNumber": 1,
 "Items": [
   {
	 "Name": "v3", "CanonicalName": "refs/tags/v3",
	 "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv3/deploymentprocesses" }
   },
   {
	 "Name": "v5", "CanonicalName": "refs/tags/v5",
	 "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv5/deploymentprocesses" }
   }
 ],
 "Links": { 
    "Self": "/api/Spaces-1/projects/Projects-1/git/tags?skip=0&take=2",
    "Page.Next": "/api/Spaces-1/projects/Projects-1/git/tags?skip=2&take=2"
  }
}`)

		// second page
		s.ExpectRequest(t, "GET", "/api/Spaces-1/projects/Projects-1/git/tags?skip=2&take=2").RespondWithText(`{
 "ItemType": "GitTag", "TotalResults": 3, "ItemsPerPage": 1, "NumberOfPages": 2, "LastPageNumber": 1,
 "Items": [
   {
	 "Name": "v7", "CanonicalName": "refs/tags/v7",
	 "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv7/deploymentprocesses" }
   }
 ],
 "Links": { 
    "Self": "/api/Spaces-1/projects/Projects-1/git/tags?skip=0&take=2"
  }
}`)

		tags, err := testutil.ReceivePair(receiver)
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
			{
				Type:          projects.GitRefTypeTag,
				Name:          "v7",
				CanonicalName: "refs/tags/v7",
				Links:         map[string]string{"DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv7/deploymentprocesses"},
			},
		}, tags)
	})

	t.Run("can get single branch", func(t *testing.T) {
		receiver := testutil.GoBegin2(func() (*projects.GitReference, error) {
			return svc.GetGitBranch(vcProject, "develop")
		})

		s.ExpectRequest(t, "GET", "/api/Spaces-1/projects/Projects-1/git/branches/develop").RespondWithText(`{
 "Name": "develop",
 "CanonicalName": "refs/heads/develop",
 "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2fheads%2fdevelop/deploymentprocesses" }
}`)

		branch, err := testutil.ReceivePair(receiver)
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
		receiver := testutil.GoBegin2(func() (*projects.GitReference, error) {
			return svc.GetGitTag(vcProject, "v5")
		})

		s.ExpectRequest(t, "GET", "/api/Spaces-1/projects/Projects-1/git/tags/v5").RespondWithText(`{
 "Name": "v5",
 "CanonicalName": "refs/tags/v5",
 "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv5/deploymentprocesses" }
}`)

		tag, err := testutil.ReceivePair(receiver)
		assert.Nil(t, err)
		assert.Equal(t, &projects.GitReference{
			Type:          projects.GitRefTypeTag,
			Name:          "v5",
			CanonicalName: "refs/tags/v5",
			Links:         map[string]string{"DeploymentProcess": "/api/Spaces-1/projects/Projects-1/refs%2ftags%2fv5/deploymentprocesses"},
		}, tag)
	})

	t.Run("can't get single tag from non version-controlled project", func(t *testing.T) {
		result, err := svc.GetGitTag(dbProject, "v3")
		assert.Nil(t, result)
		assert.EqualError(t, err, "cannot get git tag on project proj-db; no Tags link. GetGitTag requires a version-controlled project")
	})

	t.Run("can get single commit", func(t *testing.T) {
		receiver := testutil.GoBegin2(func() (*projects.GitReference, error) {
			return svc.GetGitCommit(vcProject, "59d550fb")
		})

		s.ExpectRequest(t, "GET", "/api/Spaces-1/projects/Projects-1/git/commits/59d550fb").RespondWithText(`{
 "Name": "59d550fbdf82b83619a72fdbd331cc8fa3cb2f3c",
 "CanonicalName": "59d550fbdf82b83619a72fdbd331cc8fa3cb2f3c",
 "Links": { "DeploymentProcess": "/api/Spaces-1/projects/Projects-1/59d550fbdf82b83619a72fdbd331cc8fa3cb2f3c/deploymentprocesses" }
}`)

		commit, err := testutil.ReceivePair(receiver)
		assert.Nil(t, err)
		assert.Equal(t, &projects.GitReference{
			Type:          projects.GitRefTypeCommit,
			Name:          "59d550fbdf82b83619a72fdbd331cc8fa3cb2f3c",
			CanonicalName: "59d550fbdf82b83619a72fdbd331cc8fa3cb2f3c",
			Links:         map[string]string{"DeploymentProcess": "/api/Spaces-1/projects/Projects-1/59d550fbdf82b83619a72fdbd331cc8fa3cb2f3c/deploymentprocesses"},
		}, commit)
	})

	t.Run("can't get single commit from non version-controlled project", func(t *testing.T) {
		result, err := svc.GetGitCommit(dbProject, "59d550fb")
		assert.Nil(t, result)
		assert.EqualError(t, err, "cannot get git commit on project proj-db; no Commits link. GetGitCommit requires a version-controlled project")
	})
}
