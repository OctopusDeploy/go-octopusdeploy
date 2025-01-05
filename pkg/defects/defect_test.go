package defects

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestDefectAsJSON(t *testing.T) {
	description := internal.GetRandomString(20)

	expectedJSON := fmt.Sprintf(`{
	"Description": "%s"
	}`,
		description,
	)

	var defect Defect
	err := json.Unmarshal([]byte(expectedJSON), &defect)
	require.NoError(t, err)
	require.NotNil(t, defect)

	actualJSON, err := json.Marshal(defect)
	require.NoError(t, err)
	require.NotNil(t, actualJSON)

	jsonassert.New(t).Assertf(expectedJSON, string(actualJSON))
}

func TestDefectWithStatusAsJSON(t *testing.T) {
	description := internal.GetRandomString(20)
	status := DefectStatusUnresolved

	expectedJSON := fmt.Sprintf(`{
		"Description": "%s",
		"Status": "%s"
	}`,
		description,
		status,
	)

	var defect Defect
	err := json.Unmarshal([]byte(expectedJSON), &defect)
	require.NoError(t, err)
	require.NotNil(t, defect)
	require.NotNil(t, defect.Status)

	actualJSON, err := json.Marshal(defect)
	require.NoError(t, err)
	require.NotNil(t, actualJSON)

	jsonassert.New(t).Assertf(expectedJSON, string(actualJSON))

}
