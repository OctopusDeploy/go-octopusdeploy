package runbooks

import (
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/stretchr/testify/require"
)

// TestRunbooksQueryExpandsIntoTemplate covers the path Get relies on: RunbooksQuery is
// converted by uritemplates.Struct2map and expanded against the space-wide runbooks
// template. A field present on the struct but absent from the template is silently
// dropped, so assert each filter survives the round trip.
func TestRunbooksQueryExpandsIntoTemplate(t *testing.T) {
	parsedTemplate, err := uritemplates.Parse(template)
	require.NoError(t, err)

	query := RunbooksQuery{
		IDs:         []string{"Runbooks-1", "Runbooks-2"},
		IsClone:     true,
		PartialName: "Deploy Database",
		ProjectIDs:  []string{"Projects-1"},
		Skip:        10,
		Take:        20,
	}

	values, ok := uritemplates.Struct2map(query)
	require.True(t, ok)
	require.NotNil(t, values)
	values["spaceId"] = "Spaces-1"

	expanded, err := parsedTemplate.Expand(values)
	require.NoError(t, err)

	expandedURL, err := url.Parse(expanded)
	require.NoError(t, err)
	require.Equal(t, "/api/Spaces-1/runbooks", expandedURL.Path)

	parameters := expandedURL.Query()
	require.Equal(t, "Runbooks-1,Runbooks-2", parameters.Get("ids"))
	require.Equal(t, "Deploy Database", parameters.Get("partialName"))
	require.Equal(t, "Projects-1", parameters.Get("projectIds"))
	require.Equal(t, "10", parameters.Get("skip"))
	require.Equal(t, "20", parameters.Get("take"))
}

func TestRunbooksQueryOmitsEmptyFilters(t *testing.T) {
	parsedTemplate, err := uritemplates.Parse(template)
	require.NoError(t, err)

	values, ok := uritemplates.Struct2map(RunbooksQuery{})
	require.True(t, ok)
	values["spaceId"] = "Spaces-1"

	expanded, err := parsedTemplate.Expand(values)
	require.NoError(t, err)
	require.Equal(t, "/api/Spaces-1/runbooks", expanded)
}
