package releases

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSnapshotVariablesByNameMarshall(t *testing.T) {
	command := snapshotVariablesByNameCommand{Variables: []VariableIdentifier{
		{
			Name:    "MyVariable",
			OwnerID: "Projects-1",
		},
	}}
	data, err := json.Marshal(command)
	require.NoError(t, err)
	require.JSONEq(t, `{
		"Variables": [{"Name":"MyVariable","OwnerId":"Projects-1"}]
	}`, string(data))

}
