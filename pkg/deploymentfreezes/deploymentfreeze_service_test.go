package deploymentfreezes

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

func createClient() newclient.Client {
	return newclient.NewClient(&newclient.HttpSession{})
}

func TestAddParameterValidation(t *testing.T) {
	client := createClient()
	_, err := Add(client, nil)
	require.Error(t, err)
}
