package e2e

import (
	"archive/zip"
	"bytes"
	"io"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/serverstatus"
	"github.com/stretchr/testify/require"
)

func TestServerStatusServiceGet(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	serverStatus, err := client.ServerStatus.Get()
	require.NoError(t, err)
	require.NotNil(t, serverStatus)
}

func TestServerStatusGetServerStatus(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	serverStatus, err := serverstatus.GetServerStatus(client)
	require.NoError(t, err)
	require.NotNil(t, serverStatus)
	require.NotEmpty(t, serverStatus.Links)
}

func TestServerStatusGetHealthStatus(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	healthStatus, err := serverstatus.GetHealthStatus(client)
	require.NoError(t, err)
	require.NotNil(t, healthStatus)
	require.NotEmpty(t, healthStatus.Description)
}

func TestServerStatusGetTimezones(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	timezones, err := serverstatus.GetTimezones(client)
	require.NoError(t, err)
	require.NotEmpty(t, timezones)

	for _, timezone := range timezones {
		require.NotEmpty(t, timezone.GetID())
		require.NotEmpty(t, timezone.Name)
	}
}

func TestServerStatusGetDocumentCounts(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	documentCounts, err := serverstatus.GetDocumentCounts(client)
	require.NoError(t, err)
	require.NotNil(t, documentCounts)
	require.GreaterOrEqual(t, documentCounts.Global.Spaces, 1)
}

func TestServerStatusGetSystemInfo(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	systemInfo, err := serverstatus.GetSystemInfo(client)
	require.NoError(t, err)
	require.NotNil(t, systemInfo)
	require.NotEmpty(t, systemInfo.Version)
	require.NotEmpty(t, systemInfo.ClrVersion)
	require.NotEmpty(t, systemInfo.OSVersion)
	require.NotEmpty(t, systemInfo.Uptime)
	require.Greater(t, systemInfo.WorkingSetBytes, int64(0))
	require.Greater(t, systemInfo.ThreadCount, 0)
}

func TestServerStatusGetRecentLogs(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	logEntries, err := serverstatus.GetRecentLogs(client, serverstatus.LogsQuery{Take: 5, IncludeDetail: true})
	require.NoError(t, err)
	require.LessOrEqual(t, len(logEntries), 5)

	for _, logEntry := range logEntries {
		require.NotEmpty(t, logEntry.Category)
		require.NotNil(t, logEntry.OccurredAt)
	}
}

func TestServerStatusGetSystemReport(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	report, err := serverstatus.GetSystemReport(client)
	require.NoError(t, err)
	require.NotNil(t, report)
	defer report.Close()

	contents, err := io.ReadAll(report)
	require.NoError(t, err)
	require.NotEmpty(t, contents)

	_, err = zip.NewReader(bytes.NewReader(contents), int64(len(contents)))
	require.NoError(t, err)
}
