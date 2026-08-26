package processtemplates

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type recordedRequest struct {
	uri  string
	body string
}

func newTestClient(t *testing.T, recorded *[]recordedRequest, payload string) newclient.Client {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		*recorded = append(*recorded, recordedRequest{uri: r.URL.RequestURI(), body: string(body)})
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(payload))
	}))
	t.Cleanup(server.Close)

	baseURL, err := url.Parse(server.URL + "/")
	require.NoError(t, err)

	return newclient.NewClient(&newclient.HttpSession{HttpClient: server.Client(), BaseURL: baseURL})
}

func TestListSummaries(t *testing.T) {
	const payload = `{
	  "ProcessTemplateSummaries": [
	    {
	      "Id": "refs/heads/main:my-template",
	      "Name": "My Template",
	      "GitRef": "refs/heads/main",
	      "Slug": "my-template",
	      "Description": "Does a thing",
	      "Icon": { "Id": "rocket", "Color": "#25D384" },
	      "Version": "1.0.0",
	      "PublishedDate": "2026-08-25T01:02:03.000+00:00",
	      "HasError": false
	    },
	    { "Name": "Unpublished", "GitRef": "refs/heads/main", "Slug": "unpublished", "HasError": true }
	  ],
	  "TotalResults": 2,
	  "ItemsPerPage": 30,
	  "TotalNoOfProcessTemplates": 2
	}`

	var recorded []recordedRequest
	client := newTestClient(t, &recorded, payload)

	result, err := ListSummaries(client, SummariesQuery{GitRef: "refs/heads/main", PartialName: "My", Skip: 0, Take: 30})
	require.NoError(t, err)

	assert.Equal(t, "/api/platformhub/refs%2Fheads%2Fmain/processtemplates/summaries?take=30&partialName=My", recorded[0].uri)

	require.Len(t, result.ProcessTemplateSummaries, 2)
	assert.Equal(t, 2, result.TotalResults)
	assert.Equal(t, 2, result.TotalNoOfProcessTemplates)

	first := result.ProcessTemplateSummaries[0]
	assert.Equal(t, "My Template", first.Name)
	assert.Equal(t, "my-template", first.Slug)
	assert.Equal(t, "1.0.0", first.Version)
	require.NotNil(t, first.PublishedDate)
	assert.Equal(t, 2026, first.PublishedDate.Year())
	require.NotNil(t, first.Icon)
	assert.Equal(t, "rocket", first.Icon.ID)

	second := result.ProcessTemplateSummaries[1]
	assert.Nil(t, second.PublishedDate)
	assert.True(t, second.HasError)
}

func TestListSummariesWithEmptyGitRef(t *testing.T) {
	var recorded []recordedRequest
	client := newTestClient(t, &recorded, `{}`)

	result, err := ListSummaries(client, SummariesQuery{})
	require.Error(t, err)
	require.Nil(t, result)
	assert.Empty(t, recorded)
}

func TestList(t *testing.T) {
	const payload = `{
	  "ProcessTemplates": [
	    {
	      "Id": "refs/heads/main:my-template",
	      "Name": "My Template",
	      "GitRef": "refs/heads/main",
	      "Slug": "my-template",
	      "Steps": [{ "Name": "Run a script", "Actions": [] }],
	      "Parameters": []
	    }
	  ],
	  "TotalResults": 1,
	  "ItemsPerPage": 30
	}`

	var recorded []recordedRequest
	client := newTestClient(t, &recorded, payload)

	result, err := List(client, ProcessTemplatesQuery{GitRef: "refs/heads/main", Take: 10})
	require.NoError(t, err)

	assert.Equal(t, "/api/platformhub/refs%2Fheads%2Fmain/processtemplates?take=10", recorded[0].uri)
	require.Len(t, result.ProcessTemplates, 1)
	require.Len(t, result.ProcessTemplates[0].Steps, 1)
	assert.Equal(t, "Run a script", result.ProcessTemplates[0].Steps[0].Name)
}

func TestGetBySlug(t *testing.T) {
	const payload = `{
	  "Id": "refs/heads/main:my-template",
	  "Name": "My Template",
	  "GitRef": "refs/heads/main",
	  "Slug": "my-template",
	  "Description": "Does a thing",
	  "Steps": [{ "Name": "Run a script" }],
	  "Parameters": [
	    {
	      "Name": "Environment",
	      "Label": "Environment",
	      "HelpText": "Where to deploy",
	      "IsOptional": false,
	      "DisplaySettings": { "Octopus.ControlType": "SingleLineText" },
	      "Values": [{ "Value": "Production", "Scope": { "Environment": ["Environments-1"] } }]
	    }
	  ]
	}`

	var recorded []recordedRequest
	client := newTestClient(t, &recorded, payload)

	template, err := GetBySlug(client, "refs/heads/main", "my-template")
	require.NoError(t, err)

	assert.Equal(t, "/api/platformhub/refs%2Fheads%2Fmain/processtemplates/my-template", recorded[0].uri)
	assert.Equal(t, "My Template", template.Name)
	require.Len(t, template.Steps, 1)
	require.Len(t, template.Parameters, 1)

	parameter := template.Parameters[0]
	assert.Equal(t, "Environment", parameter.Name)
	assert.Equal(t, "SingleLineText", parameter.DisplaySettings["Octopus.ControlType"])
	require.Len(t, parameter.Values, 1)
	// PropertyValueResource serialises as a bare string when not sensitive.
	assert.Equal(t, "Production", parameter.Values[0].Value.Value)
	assert.False(t, parameter.Values[0].Value.IsSensitive)
	assert.Equal(t, []string{"Environments-1"}, parameter.Values[0].Scope.Environments)
}

func TestGetBySlugWithEmptyArguments(t *testing.T) {
	var recorded []recordedRequest
	client := newTestClient(t, &recorded, `{}`)

	_, err := GetBySlug(client, "", "my-template")
	require.Error(t, err)

	_, err = GetBySlug(client, "refs/heads/main", "")
	require.Error(t, err)

	assert.Empty(t, recorded)
}

func TestAdd(t *testing.T) {
	const payload = `{"Id":"refs/heads/main:my-template","Name":"My Template","GitRef":"refs/heads/main","Slug":"my-template"}`

	var recorded []recordedRequest
	client := newTestClient(t, &recorded, payload)

	created, err := Add(client, "refs/heads/main", "My Template", "Does a thing", "Add My Template")
	require.NoError(t, err)

	require.Len(t, recorded, 1)
	assert.Equal(t, "/api/platformhub/refs%2Fheads%2Fmain/processtemplates", recorded[0].uri)

	var command map[string]any
	require.NoError(t, json.Unmarshal([]byte(recorded[0].body), &command))
	assert.Equal(t, map[string]any{
		"GitRef":            "refs/heads/main",
		"Name":              "My Template",
		"Description":       "Does a thing",
		"ChangeDescription": "Add My Template",
	}, command)

	assert.Equal(t, "my-template", created.Slug)
}

func TestAddOmitsEmptyOptionalFields(t *testing.T) {
	var recorded []recordedRequest
	client := newTestClient(t, &recorded, `{"Name":"My Template","Slug":"my-template"}`)

	_, err := Add(client, "refs/heads/main", "My Template", "", "")
	require.NoError(t, err)

	var command map[string]any
	require.NoError(t, json.Unmarshal([]byte(recorded[0].body), &command))
	assert.Equal(t, map[string]any{"GitRef": "refs/heads/main", "Name": "My Template"}, command)
}

func TestAddWithEmptyArguments(t *testing.T) {
	var recorded []recordedRequest
	client := newTestClient(t, &recorded, `{}`)

	_, err := Add(client, "", "My Template", "", "")
	require.Error(t, err)

	_, err = Add(client, "refs/heads/main", "", "", "")
	require.Error(t, err)

	assert.Empty(t, recorded)
}
