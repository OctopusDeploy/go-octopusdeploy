package releases

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/stretchr/testify/assert"
)

func TestReleaseUpdateSnapshotVariablesValidation(t *testing.T) {
	_, err := UpdateSnapshotVariables(nil, "Spaces-1", "Releases-1")
	assert.Equal(t, internal.CreateInvalidParameterError("UpdateSnapshotVariables", "client"), err)

	client := newclient.NewClient(&newclient.HttpSession{})

	_, err = UpdateSnapshotVariables(client, "", "Releases-1")
	assert.Equal(t, internal.CreateInvalidParameterError("UpdateSnapshotVariables", "spaceID"), err)

	_, err = UpdateSnapshotVariables(client, "Spaces-1", "")
	assert.Equal(t, internal.CreateInvalidParameterError("UpdateSnapshotVariables", "releaseID"), err)
}
