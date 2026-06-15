package gitdependencies

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeploymentActionGitDependencyValidate(t *testing.T) {
	testCases := []struct {
		name        string
		dependency  DeploymentActionGitDependency
		expectError bool
	}{
		{
			name:        "valid - both fields populated",
			dependency:  DeploymentActionGitDependency{DeploymentActionSlug: "deploy-action", GitDependencyName: "my-dependency"},
			expectError: false,
		},
		{
			name:        "valid - empty git dependency name is allowed",
			dependency:  DeploymentActionGitDependency{DeploymentActionSlug: "deploy-action", GitDependencyName: ""},
			expectError: false,
		},
		{
			name:        "valid - whitespace git dependency name is allowed",
			dependency:  DeploymentActionGitDependency{DeploymentActionSlug: "deploy-action", GitDependencyName: "   "},
			expectError: false,
		},
		{
			name:        "invalid - empty deployment action slug",
			dependency:  DeploymentActionGitDependency{DeploymentActionSlug: "", GitDependencyName: "my-dependency"},
			expectError: true,
		},
		{
			name:        "invalid - whitespace-only deployment action slug",
			dependency:  DeploymentActionGitDependency{DeploymentActionSlug: "   ", GitDependencyName: "my-dependency"},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.dependency.Validate()
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// Checks both fields are always present in the payload as these are required properties in Octopus
// and an empty Git dependency name is valid
func TestDeploymentActionGitDependencyMarshalJSON(t *testing.T) {
	data, err := json.Marshal(DeploymentActionGitDependency{})
	require.NoError(t, err)

	var fields map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(data, &fields))

	require.Contains(t, fields, "DeploymentActionSlug")
	require.Contains(t, fields, "GitDependencyName")
}
