package approvals

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/stretchr/testify/require"
)

func createClient() newclient.Client {
	return newclient.NewClient(&newclient.HttpSession{})
}

func TestAddApprovalParameterValidation(t *testing.T) {
	client := createClient()
	_, err := AddApproval(client, "Spaces-1", "ServerTaskApprovals-1", nil)
	require.Error(t, err)
}

func TestUpdateApprovalParameterValidation(t *testing.T) {
	client := createClient()
	_, err := UpdateApproval(client, "Spaces-1", "ServerTaskApprovals-1", nil)
	require.Error(t, err)
}
