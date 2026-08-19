package releases

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/stretchr/testify/require"
)

func TestSnapshotVariablesByNameMarshall(t *testing.T) {
	command := snapshotVariablesByNameCommand{Variables: []core.VariableIdentifier{
		{
			Name:    internal.GetRandomName(),
			OwnerID: internal.GetRandomName(),
		},
	}}
	data, err := json.Marshal(command)
	require.NoError(t, err)
	jsonString := fmt.Sprintf(`{"Variables":[{"Name":"%s","OwnerId":"%s"}]}`, command.Variables[0].Name, command.Variables[0].OwnerID)

	require.JSONEq(t, jsonString, string(data))

}
