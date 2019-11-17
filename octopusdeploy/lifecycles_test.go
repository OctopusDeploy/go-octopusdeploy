package octopusdeploy

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLifecycleGet(t *testing.T) {
	client := getFakeOctopusClient(t, "/api/lifecycles/Lifecycles-41", http.StatusOK, getLifecycleResponseJSON)
	lifecycle, err := client.Lifecycle.Get("Lifecycles-41")

	assert.Nil(t, err)
	assert.Equal(t, "Test", lifecycle.Name)
	assert.Equal(t, 2, len(lifecycle.Phases))

	phase0 := lifecycle.Phases[0]
	assert.Equal(t, "A", phase0.Name)
	assert.Equal(t, int32(1), phase0.MinimumEnvironmentsBeforePromotion)
	assert.Equal(t, true, phase0.IsOptionalPhase)
	assert.Equal(t, 1, len(phase0.AutomaticDeploymentTargets))
	assert.Equal(t, "Environments-2", phase0.AutomaticDeploymentTargets[0])
	assert.Equal(t, 1, len(phase0.OptionalDeploymentTargets))
	assert.Equal(t, "Environments-1", phase0.OptionalDeploymentTargets[0])
	assert.Equal(t, RetentionUnit_Days, phase0.ReleaseRetentionPolicy.Unit)
	assert.Equal(t, int32(1), phase0.ReleaseRetentionPolicy.QuantityToKeep)
	assert.Equal(t, false, phase0.ReleaseRetentionPolicy.ShouldKeepForever)
	assert.Equal(t, RetentionUnit_Items, phase0.TentacleRetentionPolicy.Unit)
	assert.Equal(t, int32(0), phase0.TentacleRetentionPolicy.QuantityToKeep)
	assert.Equal(t, true, phase0.TentacleRetentionPolicy.ShouldKeepForever)

	assert.Equal(t, (*RetentionPeriod)(nil), lifecycle.Phases[1].ReleaseRetentionPolicy)
	assert.Equal(t, (*RetentionPeriod)(nil), lifecycle.Phases[1].TentacleRetentionPolicy)

	assert.Equal(t, RetentionUnit_Days, lifecycle.ReleaseRetentionPolicy.Unit)
	assert.Equal(t, int32(3), lifecycle.ReleaseRetentionPolicy.QuantityToKeep)
	assert.Equal(t, false, lifecycle.ReleaseRetentionPolicy.ShouldKeepForever)
	assert.Equal(t, RetentionUnit_Items, lifecycle.TentacleRetentionPolicy.Unit)
	assert.Equal(t, int32(2), lifecycle.TentacleRetentionPolicy.QuantityToKeep)
	assert.Equal(t, false, lifecycle.TentacleRetentionPolicy.ShouldKeepForever)
}

const getLifecycleResponseJSON = `
{
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

func TestValidateLifecycleValuesJustANamePasses(t *testing.T) {

	lifecycle := &Lifecycle{
		Name: "My Lifecycle",
	}

	assert.Nil(t, ValidateLifecycleValues(lifecycle))
}

func TestValidateLifecycleValuesPhaseWithJustANamePasses(t *testing.T) {

	lifecycle := &Lifecycle{
		Name: "My Lifecycle",
		Phases: []Phase{
			Phase{
				Name: "My Phase",
			},
		},
	}

	assert.Nil(t, ValidateLifecycleValues(lifecycle))
}

func TestValidateLifecycleValuesMissingNameFails(t *testing.T) {

	lifecycle := &Lifecycle{}

	assert.Error(t, ValidateLifecycleValues(lifecycle))
}

func TestValidateLifecycleValuesPhaseWithMissingNameFails(t *testing.T) {

	lifecycle := &Lifecycle{
		Name: "My Lifecycle",
		Phases: []Phase{
			Phase{},
		},
	}

	assert.Error(t, ValidateLifecycleValues(lifecycle))
}
