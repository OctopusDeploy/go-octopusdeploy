package deploymentfreezes

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/stretchr/testify/require"
	"testing"
)

func createClient() newclient.Client {
	return newclient.NewClient(&newclient.HttpSession{})
}

func TestAddParameterValidation(t *testing.T) {
	client := createClient()
	_, err := Add(client, nil)
	require.Error(t, err)
}
