package lifecycles

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

// TODO: fix test
// func TestLifecycleGet(t *testing.T) {
// 	client, err := octopusdeploy.GetFakeOctopusClient(t, "/api/lifecycles/Lifecycles-41", http.StatusOK, getLifecycleResponseJSON)
// 	assert.NotNil(t, client)
// 	require.NoError(t, err)

// 	lifecycle, err := client.Lifecycles.GetByID("Lifecycles-41")
// 	assert.NotNil(t, lifecycle)
// 	require.NoError(t, err)

// 	assert.Equal(t, "Test", lifecycle.Name)
// 	assert.Equal(t, 2, len(lifecycle.Phases))

// 	phase0 := lifecycle.Phases[0]
// 	assert.Equal(t, "A", phase0.Name)
// 	assert.Equal(t, int32(1), phase0.MinimumEnvironmentsBeforePromotion)
// 	assert.Equal(t, true, phase0.IsOptionalPhase)
// 	assert.Equal(t, 1, len(phase0.AutomaticDeploymentTargets))
// 	assert.Equal(t, "Environments-2", phase0.AutomaticDeploymentTargets[0])
// 	assert.Equal(t, 1, len(phase0.OptionalDeploymentTargets))
// 	assert.Equal(t, "Environments-1", phase0.OptionalDeploymentTargets[0])
// 	assert.Equal(t, octopusdeploy.RetentionUnitDays, phase0.ReleaseRetentionPolicy.Unit)
// 	assert.Equal(t, int32(1), phase0.ReleaseRetentionPolicy.QuantityToKeep)
// 	assert.Equal(t, false, phase0.ReleaseRetentionPolicy.ShouldKeepForever)
// 	assert.Equal(t, octopusdeploy.RetentionUnitItems, phase0.TentacleRetentionPolicy.Unit)
// 	assert.Equal(t, int32(0), phase0.TentacleRetentionPolicy.QuantityToKeep)
// 	assert.Equal(t, true, phase0.TentacleRetentionPolicy.ShouldKeepForever)

// 	assert.Equal(t, (*octopusdeploy.RetentionPeriod)(nil), lifecycle.Phases[1].ReleaseRetentionPolicy)
// 	assert.Equal(t, (*octopusdeploy.RetentionPeriod)(nil), lifecycle.Phases[1].TentacleRetentionPolicy)

// 	assert.Equal(t, octopusdeploy.RetentionUnitDays, lifecycle.ReleaseRetentionPolicy.Unit)
// 	assert.Equal(t, int32(3), lifecycle.ReleaseRetentionPolicy.QuantityToKeep)
// 	assert.Equal(t, false, lifecycle.ReleaseRetentionPolicy.ShouldKeepForever)
// 	assert.Equal(t, octopusdeploy.RetentionUnitItems, lifecycle.TentacleRetentionPolicy.Unit)
// 	assert.Equal(t, int32(2), lifecycle.TentacleRetentionPolicy.QuantityToKeep)
// 	assert.Equal(t, false, lifecycle.TentacleRetentionPolicy.ShouldKeepForever)
// }

// const getLifecycleResponseJSON = `
// {
//   "Id": "Lifecycles-41",
//   "Phases": [
//     {
//       "Id": "61e30a4b-3bdb-4eff-8995-805de61da9ff",
//       "Name": "A",
//       "AutomaticDeploymentTargets": [
// 	    "Environments-2"
//       ],
//       "OptionalDeploymentTargets": [
//         "Environments-1"
//       ],
//       "MinimumEnvironmentsBeforePromotion": 1,
//       "IsOptionalPhase": true,
//       "ReleaseRetentionPolicy": {
//         "Unit": "Days",
//         "QuantityToKeep": 1,
//         "ShouldKeepForever": false
//       },
//       "TentacleRetentionPolicy": {
//         "Unit": "Items",
//         "QuantityToKeep": 0,
//         "ShouldKeepForever": true
//       }
//     },
//     {
//       "Id": "670920c6-1065-4207-8d15-2c5d7947e795",
//       "Name": "B",
//       "AutomaticDeploymentTargets": [],
//       "OptionalDeploymentTargets": [],
//       "MinimumEnvironmentsBeforePromotion": 0,
//       "IsOptionalPhase": false,
//       "ReleaseRetentionPolicy": null,
//       "TentacleRetentionPolicy": null
//     }
//   ],
//   "Name": "Test",
//   "ReleaseRetentionPolicy": {
//     "Unit": "Days",
//     "QuantityToKeep": 3,
//     "ShouldKeepForever": false
//   },
//   "TentacleRetentionPolicy": {
//     "Unit": "Items",
//     "QuantityToKeep": 2,
//     "ShouldKeepForever": false
//   },
//   "Description": "",
//   "Links": {
//     "Self": "/api/lifecycles/Lifecycles-41",
//     "Preview": "/api/lifecycles/Lifecycles-41/preview",
//     "Projects": "/api/lifecycles/Lifecycles-41/projects"
//   }
// }`

func TestLifecycleNew(t *testing.T) {
	name := "name"

	lifecycle := Lifecycle{}
	require.Error(t, lifecycle.Validate())

	lifecycle = Lifecycle{
		Name: name,
	}
	require.NoError(t, lifecycle.Validate())

	lifecycleWithNew := NewLifecycle("")
	require.NotNil(t, lifecycleWithNew)
	require.Error(t, lifecycleWithNew.Validate())

	lifecycleWithNew = NewLifecycle(" ")
	require.NotNil(t, lifecycleWithNew)
	require.Error(t, lifecycleWithNew.Validate())

	lifecycleWithNew = NewLifecycle(name)
	require.NotNil(t, lifecycleWithNew)
	require.NoError(t, lifecycleWithNew.Validate())
}

func TestLifecycleToJson(t *testing.T) {
	name := "test-name"

	lifecycle := NewLifecycle(name)
	require.NotNil(t, lifecycle)

	expectedJson := fmt.Sprintf(`{
		"Name": "%s",
		"ReleaseRetentionPolicy": {
			"QuantityToKeep": 30,
			"ShouldKeepForever": false,
			"Unit": "Days"
		},
		"TentacleRetentionPolicy": {
			"QuantityToKeep": 30,
			"ShouldKeepForever": false,
			"Unit": "Days"
		}
	}`, name)

	lifecycleAsJson, err := json.Marshal(lifecycle)
	require.NoError(t, err)
	require.NotNil(t, lifecycleAsJson)

	jsonassert.New(t).Assertf(string(lifecycleAsJson), expectedJson)

	description := "test-description"
	lifecycle.Description = description

	expectedJson = fmt.Sprintf(`{
		"Description": "%s",
		"Name": "%s",
		"ReleaseRetentionPolicy": {
			"QuantityToKeep": 30,
			"ShouldKeepForever": false,
			"Unit": "Days"
		},
		"TentacleRetentionPolicy": {
			"QuantityToKeep": 30,
			"ShouldKeepForever": false,
			"Unit": "Days"
		}
	}`, description, name)

	lifecycleAsJson, err = json.Marshal(lifecycle)
	require.NoError(t, err)
	require.NotNil(t, lifecycleAsJson)

	jsonassert.New(t).Assertf(string(lifecycleAsJson), expectedJson)
}

func TestLifecycleFromJson(t *testing.T) {
	const lifecycleAsJson = `{
  "Id": "Lifecycles-41",
  "Phases": [
    {
      "Id": "61e30a4b-3bdb-4eff-8995-805de61da9ff",
      "Name": "A",
      "AutomaticDeploymentTargets": [
	    "Environments-2"
      ],
      "OptionalDeploymentTargets": [
        "Environments-1"
      ],
      "MinimumEnvironmentsBeforePromotion": 1,
      "IsOptionalPhase": true,
      "ReleaseRetentionPolicy": {
        "Unit": "Days",
        "QuantityToKeep": 1,
        "ShouldKeepForever": false
      },
      "TentacleRetentionPolicy": {
        "Unit": "Items",
        "QuantityToKeep": 0,
        "ShouldKeepForever": true
      }
    },
    {
      "Id": "670920c6-1065-4207-8d15-2c5d7947e795",
      "Name": "B",
      "AutomaticDeploymentTargets": [],
      "OptionalDeploymentTargets": [],
      "MinimumEnvironmentsBeforePromotion": 0,
      "IsOptionalPhase": false,
      "ReleaseRetentionPolicy": null,
      "TentacleRetentionPolicy": null
    }
  ],
  "Name": "Test",
  "ReleaseRetentionPolicy": {
    "Unit": "Days",
    "QuantityToKeep": 3,
    "ShouldKeepForever": false
  },
  "TentacleRetentionPolicy": {
    "Unit": "Items",
    "QuantityToKeep": 2,
    "ShouldKeepForever": false
  },
  "Description": "",
  "Links": {
    "Self": "/api/lifecycles/Lifecycles-41",
    "Preview": "/api/lifecycles/Lifecycles-41/preview",
    "Projects": "/api/lifecycles/Lifecycles-41/projects"
  }
}`

	var lifecycle *Lifecycle
	err := json.Unmarshal([]byte(lifecycleAsJson), &lifecycle)
	require.NoError(t, err)
	require.NotNil(t, lifecycle)

	phase0 := lifecycle.Phases[0]
	require.Equal(t, "A", phase0.Name)
	require.Equal(t, int32(1), phase0.MinimumEnvironmentsBeforePromotion)
	require.Equal(t, true, phase0.IsOptionalPhase)
	require.Equal(t, 1, len(phase0.AutomaticDeploymentTargets))
	require.Equal(t, "Environments-2", phase0.AutomaticDeploymentTargets[0])
	require.Equal(t, 1, len(phase0.OptionalDeploymentTargets))
	require.Equal(t, "Environments-1", phase0.OptionalDeploymentTargets[0])
	require.Equal(t, RetentionUnitDays, phase0.ReleaseRetentionPolicy.Unit)
	require.Equal(t, int32(1), phase0.ReleaseRetentionPolicy.QuantityToKeep)
	require.Equal(t, false, phase0.ReleaseRetentionPolicy.ShouldKeepForever)
	require.Equal(t, RetentionUnitItems, phase0.TentacleRetentionPolicy.Unit)
	require.Equal(t, int32(0), phase0.TentacleRetentionPolicy.QuantityToKeep)
	require.Equal(t, true, phase0.TentacleRetentionPolicy.ShouldKeepForever)

	require.Equal(t, (*core.RetentionPeriod)(nil), lifecycle.Phases[1].ReleaseRetentionPolicy)
	require.Equal(t, (*core.RetentionPeriod)(nil), lifecycle.Phases[1].TentacleRetentionPolicy)

	require.Equal(t, RetentionUnitDays, lifecycle.ReleaseRetentionPolicy.Unit)
	require.Equal(t, int32(3), lifecycle.ReleaseRetentionPolicy.QuantityToKeep)
	require.Equal(t, false, lifecycle.ReleaseRetentionPolicy.ShouldKeepForever)
	require.Equal(t, RetentionUnitItems, lifecycle.TentacleRetentionPolicy.Unit)
	require.Equal(t, int32(2), lifecycle.TentacleRetentionPolicy.QuantityToKeep)
	require.Equal(t, false, lifecycle.TentacleRetentionPolicy.ShouldKeepForever)
}

func TestValidateLifecycleValuesPhaseWithJustANamePasses(t *testing.T) {
	lifecycle := &Lifecycle{
		Name: "My Lifecycle",
		Phases: []Phase{
			{Name: "My Phase"},
		},
	}

	require.NoError(t, lifecycle.Validate())
}

func TestValidateLifecycleValuesMissingNameFails(t *testing.T) {
	lifecycle := &Lifecycle{}
	require.Error(t, lifecycle.Validate())
}

func TestValidateLifecycleValuesPhaseWithMissingNameFails(t *testing.T) {
	lifecycle := &Lifecycle{
		Name: "My Lifecycle",
		Phases: []Phase{
			{},
		},
	}

	require.Error(t, lifecycle.Validate())
}
