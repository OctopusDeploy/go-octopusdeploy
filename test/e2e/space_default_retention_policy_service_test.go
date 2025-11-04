package e2e

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
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
		retention.LifecycleReleaseRetentionType,
		defaultSpace.ID,
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
		retention.LifecycleTentacleRetentionType,
		defaultSpace.ID,
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
		retention.LifecycleReleaseRetentionType,
		defaultSpace.ID,
	}
	defaultSpaceLifecycleReleasePolicy, err := retention.Get(client, query)
	if err != nil {
		t.Fatal(err)
	}

	policy := retention.LifecycleReleaseRetentionPolicy{
		Unit:           retention.RetentionUnitItems,
		QuantityToKeep: 4,
		Strategy:       retention.RetentionStrategyCount,
		SpaceDefaultRetentionPolicy: retention.SpaceDefaultRetentionPolicy{
			SpaceId:       defaultSpace.ID,
			RetentionType: retention.LifecycleReleaseRetentionType,
			Resource: resources.Resource{
				ID: defaultSpaceLifecycleReleasePolicy.ID,
			},
		},
	}
	res, err := retention.Update(client, &policy)
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
		retention.LifecycleTentacleRetentionType,
		defaultSpace.ID,
	}
	defaultSpaceLifecycleReleasePolicy, err := retention.Get(client, query)
	if err != nil {
		t.Fatal(err)
	}

	policy := retention.LifecycleTentacleRetentionPolicy{
		Unit:           retention.RetentionUnitDays,
		QuantityToKeep: 3,
		Strategy:       retention.RetentionStrategyCount,
		SpaceDefaultRetentionPolicy: retention.SpaceDefaultRetentionPolicy{
			SpaceId:       defaultSpace.ID,
			RetentionType: retention.LifecycleTentacleRetentionType,
			Resource: resources.Resource{
				ID: defaultSpaceLifecycleReleasePolicy.ID,
			},
		},
	}
	res, err := retention.Update(client, &policy)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.QuantityToKeep, 3)
	require.Equal(t, res.Strategy, retention.RetentionStrategyCount)
	require.Equal(t, res.Unit, retention.RetentionUnitDays)
}
