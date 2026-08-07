package actiontemplates

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/stretchr/testify/require"
)

func expandCollectionTemplate(t *testing.T, query any) string {
	values, ok := uritemplates.Struct2map(query)
	require.True(t, ok)
	values["spaceId"] = "Spaces-1"

	uri, err := uritemplates.NewUriTemplateCache().Expand(template, values)
	require.NoError(t, err)

	return uri
}

func TestQueryExpandsIntoTheCollectionTemplate(t *testing.T) {
	uri := expandCollectionTemplate(t, Query{
		IDs:         []string{"ActionTemplates-1", "ActionTemplates-2"},
		PartialName: "Hello World",
		Skip:        10,
		Take:        100,
	})

	require.Equal(
		t,
		"/api/Spaces-1/actiontemplates?skip=10&take=100&ids=ActionTemplates-1%2CActionTemplates-2&partialName=Hello%20World",
		uri,
	)
}

// ActionTemplateSearch has no uri tags, so Struct2map emits the Go field names and the template
// drops them. This pins the behaviour that deprecated Get in favour of GetByQuery.
func TestActionTemplateSearchExpandsIntoNothing(t *testing.T) {
	uri := expandCollectionTemplate(t, ActionTemplateSearch{
		ID:   "ActionTemplates-1",
		Name: "Hello World",
	})

	require.Equal(t, "/api/Spaces-1/actiontemplates", uri)
}
