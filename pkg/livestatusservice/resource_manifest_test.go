package livestatusservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResourceManifestRequestTenantMethods(t *testing.T) {
	// Test untenanted request
	untenanteRequest := &GetResourceManifestRequest{
		SpaceID:                                "Spaces-1",
		ProjectID:                              "Projects-1",
		EnvironmentID:                          "Environments-1",
		MachineID:                              "Machines-1",
		TenantID:                               "",
		DesiredOrKubernetesMonitoredResourceID: "Resources-1",
	}

	assert.True(t, !untenanteRequest.IsTenanted())
	assert.False(t, untenanteRequest.IsTenanted())

	// Test tenanted request
	tenantedRequest := &GetResourceManifestRequest{
		SpaceID:                                "Spaces-1",
		ProjectID:                              "Projects-1",
		EnvironmentID:                          "Environments-1",
		MachineID:                              "Machines-1",
		TenantID:                               "Tenants-1",
		DesiredOrKubernetesMonitoredResourceID: "Resources-1",
	}

	assert.False(t, !tenantedRequest.IsTenanted())
	assert.True(t, tenantedRequest.IsTenanted())
}

func TestGetResourceManifestRequestValidation(t *testing.T) {
	// Test valid request
	validRequest := &GetResourceManifestRequest{
		SpaceID:                                "Spaces-1",
		ProjectID:                              "Projects-1",
		EnvironmentID:                          "Environments-1",
		MachineID:                              "Machines-1",
		DesiredOrKubernetesMonitoredResourceID: "Resources-1",
	}

	err := validRequest.Validate()
	assert.NoError(t, err)

	// Test invalid request (missing required fields)
	invalidRequest := &GetResourceManifestRequest{}

	err = invalidRequest.Validate()
	assert.Error(t, err)
}

func TestGetResourceManifestResponseValidation(t *testing.T) {
	// Test valid response
	validResponse := &GetResourceManifestResponse{
		LiveManifest: "valid manifest",
	}

	err := validResponse.Validate()
	assert.NoError(t, err)

	// Test invalid response (missing required LiveManifest)
	invalidResponse := &GetResourceManifestResponse{}

	err = invalidResponse.Validate()
	assert.Error(t, err)
}

func TestLiveResourceDiffValidation(t *testing.T) {
	// Test valid diff
	validDiff := &LiveResourceDiff{
		Left:  "left",
		Right: "right",
		Diff:  "diff",
	}

	err := validDiff.Validate()
	assert.NoError(t, err)

	// Test invalid diff (missing required fields)
	invalidDiff := &LiveResourceDiff{}

	err = invalidDiff.Validate()
	assert.Error(t, err)
}

func TestGetResourceManifestResponseWithDiff(t *testing.T) {
	liveManifest := "live manifest"
	desiredManifest := "desired manifest"
	diff := NewLiveResourceDiff("left", "right", "diff content")

	response := &GetResourceManifestResponse{
		LiveManifest:    liveManifest,
		DesiredManifest: desiredManifest,
		Diff:            diff,
	}

	assert.Equal(t, liveManifest, response.LiveManifest)
	assert.Equal(t, desiredManifest, response.DesiredManifest)
	assert.NotNil(t, response.Diff)
	assert.Equal(t, "left", response.Diff.Left)
	assert.Equal(t, "right", response.Diff.Right)
	assert.Equal(t, "diff content", response.Diff.Diff)

	err := response.Validate()
	assert.NoError(t, err)
}
