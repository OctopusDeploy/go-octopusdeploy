package serverstatus

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/stretchr/testify/require"
)

func TestLogsQueryExpandsIntoTheRecentLogsTemplate(t *testing.T) {
	values, ok := uritemplates.Struct2map(LogsQuery{Skip: 5, Take: 10, IncludeDetail: true})
	require.True(t, ok)

	uri, err := uritemplates.NewUriTemplateCache().Expand(recentLogsTemplate, values)
	require.NoError(t, err)

	require.Equal(t, "/api/serverstatus/logs?skip=5&take=10&includeDetail=true", uri)
}

func TestEmptyLogsQueryExpandsIntoNoParameters(t *testing.T) {
	values, ok := uritemplates.Struct2map(LogsQuery{})
	require.True(t, ok)

	uri, err := uritemplates.NewUriTemplateCache().Expand(recentLogsTemplate, values)
	require.NoError(t, err)

	require.Equal(t, "/api/serverstatus/logs", uri)
}
