package releases

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/stretchr/testify/assert"
)

func TestReleaseSnapshotVariablesValidation(t *testing.T) {
	_, err := SnapshotVariables(nil, "Spaces-1", "Releases-1")
	assert.Equal(t, internal.CreateInvalidParameterError("SnapshotVariables", "client"), err)

	client := newclient.NewClient(&newclient.HttpSession{})

	_, err = SnapshotVariables(client, "", "Releases-1")
	assert.Equal(t, internal.CreateInvalidParameterError("SnapshotVariables", "spaceID"), err)

	_, err = SnapshotVariables(client, "Spaces-1", "")
	assert.Equal(t, internal.CreateInvalidParameterError("SnapshotVariables", "releaseID"), err)
}
