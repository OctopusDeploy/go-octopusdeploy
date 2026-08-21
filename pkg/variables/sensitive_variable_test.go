package variables

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func sensitiveVariable(name string, value string, id string) *Variable {
	variable := NewVariable(name)
	variable.IsSensitive = true
	variable.Type = "Sensitive"
	variable.Value = value
	variable.ID = id
	return variable
}

func TestValidateSensitiveVariables(t *testing.T) {
	testCases := []struct {
		name        string
		variables   []*Variable
		expectError bool
	}{
		{
			name:        "sensitive with an ID and no value keeps the stored value",
			variables:   []*Variable{sensitiveVariable("secret", "", "abc-123")},
			expectError: false,
		},
		{
			name:        "sensitive with a value and no ID sets it explicitly",
			variables:   []*Variable{sensitiveVariable("secret", "s3cret", "")},
			expectError: false,
		},
		{
			name:        "sensitive with neither would clear the stored value",
			variables:   []*Variable{sensitiveVariable("secret", "", "")},
			expectError: true,
		},
		{
			name:        "non-sensitive with neither is unaffected",
			variables:   []*Variable{NewVariable("plain")},
			expectError: false,
		},
		{
			name:        "reported alongside valid variables",
			variables:   []*Variable{NewVariable("plain"), sensitiveVariable("secret", "", "")},
			expectError: true,
		},
		{
			name:        "nil entries are skipped",
			variables:   []*Variable{nil},
			expectError: false,
		},
		{
			name:        "empty set",
			variables:   nil,
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateSensitiveVariables(VariableSet{Variables: tc.variables})

			if !tc.expectError {
				require.NoError(t, err)
				return
			}

			require.Error(t, err)
			require.ErrorAs(t, err, &errSensitiveVariableWouldBeCleared{})
			require.Contains(t, err.Error(), "secret")
		})
	}
}
