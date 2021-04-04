package octopusdeploy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeploymentActionProperties(t *testing.T) {
	properties := map[string]PropertyValue{
		"key.isNotSensitive": NewPropertyValue("isNotSensitive", false),
		"key.isSensitive":    NewPropertyValue("isSensitive", true),
	}

	require.NotNil(t, properties)

	notSensitiveProperty := properties["key.isNotSensitive"]

	require.NotNil(t, notSensitiveProperty)
	require.False(t, notSensitiveProperty.IsSensitive)
	require.Nil(t, notSensitiveProperty.SensitiveValue)
	require.NotNil(t, notSensitiveProperty.Value)
	require.Equal(t, "isNotSensitive", notSensitiveProperty.Value)

	sensitiveProperty := properties["key.isSensitive"]

	require.NotNil(t, sensitiveProperty)
	require.True(t, sensitiveProperty.IsSensitive)
	require.Len(t, sensitiveProperty.Value, 0)
	require.NotNil(t, sensitiveProperty.SensitiveValue)
	require.True(t, sensitiveProperty.SensitiveValue.HasValue)
	require.NotNil(t, sensitiveProperty.SensitiveValue.NewValue)
	require.Equal(t, "isSensitive", sensitiveProperty.SensitiveValue.NewValue)
}
