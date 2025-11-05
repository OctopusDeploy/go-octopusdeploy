package e2e

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/retention"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/spaces"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadLifecycleReleaseRetentionDefaultPolicy(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	defaultSpace, _ := spaces.GetDefaultSpace(client)
	query := retention.SpaceDefaultRetentionPolicyQuery{
		RetentionType: retention.LifecycleReleaseRetentionType,
		SpaceID:       defaultSpace.ID,
	}
	res, err := retention.Get(client, query)

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.RetentionType, retention.LifecycleReleaseRetentionType)
}

func TestReadLifecycleTentacleRetentionDefaultPolicy(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	defaultSpace, _ := spaces.GetDefaultSpace(client)
	query := retention.SpaceDefaultRetentionPolicyQuery{
		RetentionType: retention.LifecycleTentacleRetentionType,
		SpaceID:       defaultSpace.ID,
	}
	res, err := retention.Get(client, query)

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.RetentionType, retention.LifecycleTentacleRetentionType)
}

func TestModifyLifecycleReleaseRetentionDefaultPolicy(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)
	defaultSpace, _ := spaces.GetDefaultSpace(client)
	query := retention.SpaceDefaultRetentionPolicyQuery{
		RetentionType: retention.LifecycleReleaseRetentionType,
		SpaceID:       defaultSpace.ID,
	}
	defaultSpaceLifecycleReleasePolicy, err := retention.Get(client, query)
	if err != nil {
		t.Fatal(err)
	}

	policy := retention.NewCountBasedLifecycleReleaseRetentionPolicy(4, retention.RetentionUnitItems, defaultSpace.ID, defaultSpaceLifecycleReleasePolicy.ID)

	res, err := retention.Update(client, policy)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.QuantityToKeep, 4)
	require.Equal(t, res.Strategy, retention.RetentionStrategyCount)
	require.Equal(t, res.Unit, retention.RetentionUnitItems)
}

func TestModifyLifecycleTentacleRetentionDefaultPolicy(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)
	defaultSpace, _ := spaces.GetDefaultSpace(client)
	query := retention.SpaceDefaultRetentionPolicyQuery{
		RetentionType: retention.LifecycleTentacleRetentionType,
		SpaceID:       defaultSpace.ID,
	}
	defaultSpaceLifecycleReleasePolicy, err := retention.Get(client, query)
	if err != nil {
		t.Fatal(err)
	}

	policy := retention.NewCountBasedLifecycleTentacleRetentionPolicy(3, retention.RetentionUnitDays, defaultSpace.ID, defaultSpaceLifecycleReleasePolicy.ID)
	res, err := retention.Update(client, policy)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.QuantityToKeep, 3)
	require.Equal(t, res.Strategy, retention.RetentionStrategyCount)
	require.Equal(t, res.Unit, retention.RetentionUnitDays)
}
