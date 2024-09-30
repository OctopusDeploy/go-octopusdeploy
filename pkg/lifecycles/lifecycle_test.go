package lifecycles

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestLifecycleAsJSON(t *testing.T) {
	lifecycle := NewLifecycle("test-lifecycle-name")
	lifecycle.Description = "test-description"
	lifecycle.ID = "test-lifecycle-id"
	lifecycle.Links["Self"] = "test-self-link"
	lifecycle.Links["Preview"] = "test-preview-link"
	lifecycle.Links["Projects"] = "test-projects-link"
	lifecycle.Phases = append(lifecycle.Phases, NewPhase("test-phase-name-1"))
	lifecycle.Phases[0].AutomaticDeploymentTargets = []string{"test-AutomaticDeploymentTargets-1"}
	lifecycle.Phases[0].ID = "test-phase-id-1"
	lifecycle.Phases[0].IsOptionalPhase = true
	lifecycle.Phases[0].MinimumEnvironmentsBeforePromotion = 123
	lifecycle.Phases[0].OptionalDeploymentTargets = []string{"Environments-1"}
	lifecycle.Phases[0].ReleaseRetentionPolicy = core.NewRetentionPeriod(1, "Days", false)
	lifecycle.Phases[0].TentacleRetentionPolicy = core.NewRetentionPeriod(0, "Items", true)
	lifecycle.Phases = append(lifecycle.Phases, NewPhase("test-phase-name-2"))
	lifecycle.Phases[1].ID = "test-phase-id-2"
	lifecycle.Phases[1].ReleaseRetentionPolicy = nil
	lifecycle.Phases[1].TentacleRetentionPolicy = nil
	lifecycle.Phases[1].IsPriorityPhase = true
	lifecycle.TentacleRetentionPolicy.Unit = "Items"

	lifecycle.ReleaseRetentionPolicy.QuantityToKeep = 3
	lifecycle.TentacleRetentionPolicy.QuantityToKeep = 2

	expectedJSON := `{
		"Description": "test-description",
		"Id": "test-lifecycle-id",
		"Name": "test-lifecycle-name",
		"Phases": [
			{
				"AutomaticDeploymentTargets": ["test-AutomaticDeploymentTargets-1"],
				"Id": "test-phase-id-1",
				"IsOptionalPhase": true,
				"IsPriorityPhase": false,
				"MinimumEnvironmentsBeforePromotion": 123,
				"Name": "test-phase-name-1",
				"OptionalDeploymentTargets": ["Environments-1"],
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
				"Id": "test-phase-id-2",
				"Name": "test-phase-name-2",
				"AutomaticDeploymentTargets": [],
				"OptionalDeploymentTargets": [],
				"MinimumEnvironmentsBeforePromotion": 0,
				"IsOptionalPhase": false,
				"IsPriorityPhase": true,
				"ReleaseRetentionPolicy": null,
				"TentacleRetentionPolicy": null
			}
		],
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
		"Links": {
			"Self": "test-self-link",
			"Preview": "test-preview-link",
			"Projects": "test-projects-link"
		}
	}`

	lifecycleAsJSON, err := json.Marshal(lifecycle)
	require.NoError(t, err)
	require.NotNil(t, lifecycleAsJSON)

	jsonassert.New(t).Assertf(expectedJSON, string(lifecycleAsJSON))
}

func TestLifecycleUnmarshalJSON(t *testing.T) {
	description := internal.GetRandomName()
	name := internal.GetRandomName()
	spaceID := internal.GetRandomName()

	inputJSON := fmt.Sprintf(`{
		"Description": "%s",
		"Name": "%s",
		"SpaceId": "%s"
	}`, description, name, spaceID)

	var lifecycle Lifecycle
	err := json.Unmarshal([]byte(inputJSON), &lifecycle)
	require.NoError(t, err)
	require.NotNil(t, lifecycle)
	require.Equal(t, description, lifecycle.Description)
	require.Equal(t, name, lifecycle.Name)
	require.Nil(t, lifecycle.Phases)
	require.Len(t, lifecycle.Phases, 0)
	require.Nil(t, lifecycle.ReleaseRetentionPolicy)
	require.Equal(t, spaceID, lifecycle.SpaceID)
	require.Nil(t, lifecycle.TentacleRetentionPolicy)
}

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
	name := internal.GetRandomName()
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
      "IsPriorityPhase": false,
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
	require.Equal(t, false, phase0.IsPriorityPhase)
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
		Name:   "My Lifecycle",
		Phases: []*Phase{NewPhase("phase")},
	}

	require.NoError(t, lifecycle.Validate())
}

func TestValidateLifecycleValuesMissingNameFails(t *testing.T) {
	lifecycle := &Lifecycle{}
	require.Error(t, lifecycle.Validate())
}

func TestValidateLifecycleValuesPhaseWithMissingNameFails(t *testing.T) {
	lifecycle := &Lifecycle{
		Name:   "My Lifecycle",
		Phases: []*Phase{NewPhase("")},
	}

	require.Error(t, lifecycle.Validate())
}
