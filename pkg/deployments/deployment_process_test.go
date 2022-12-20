package deployments_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestDeploymentProcessAsJSON(t *testing.T) {
	lastSnapshotID := internal.GetRandomName()
	projectID := internal.GetRandomName()
	spaceID := internal.GetRandomName()

	rand.Seed(time.Now().UnixNano())
	version := rand.Intn(2)

	expectedJSON := fmt.Sprintf(`{
		"LastSnapshotId": "%s",
		"ProjectId": "%s",
		"SpaceId": "%s",
		"Version": %v
	}`,
		lastSnapshotID,
		projectID,
		spaceID,
		version,
	)

	var deploymentProcess deployments.DeploymentProcess
	err := json.Unmarshal([]byte(expectedJSON), &deploymentProcess)
	require.NoError(t, err)
	require.NotNil(t, deploymentProcess)

	// TODO: add checks

	actualJSON, err := json.Marshal(deploymentProcess)
	require.NoError(t, err)
	require.NotNil(t, actualJSON)

	jsonassert.New(t).Assertf(expectedJSON, string(actualJSON))

	stepName := internal.GetRandomName()

	expectedJSON = fmt.Sprintf(`{
		"LastSnapshotId": "%s",
		"ProjectId": "%s",
		"SpaceId": "%s",
		"Steps": [
			{
				"Name": "%s"
			}
		],
		"Version": %v
	}`,
		lastSnapshotID,
		projectID,
		spaceID,
		stepName,
		version,
	)

	err = json.Unmarshal([]byte(expectedJSON), &deploymentProcess)
	require.NoError(t, err)
	require.NotNil(t, deploymentProcess)

	actualJSON, err = json.Marshal(deploymentProcess)
	require.NoError(t, err)
	require.NotNil(t, actualJSON)

	jsonassert.New(t).Assertf(expectedJSON, string(actualJSON))
}
