package approvals

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServerTaskApprovalUnmarshal(t *testing.T) {
	payload := `{
		"Id": "ServerTaskApprovals-1",
		"ServerTaskId": "ServerTasks-1",
		"SpaceId": "Spaces-1",
		"ApprovalProviderId": "octopus",
		"Name": "Deploy to Prod",
		"ChangeRequest": {
			"Id": "CR-1",
			"Number": "CHG001",
			"Description": "Deploy release 1.0",
			"Active": true,
			"Type": "normal",
			"ChangeRequestApprovalState": "PreApproval",
			"Self": "https://example/cr/1"
		}
	}`
	var sta ServerTaskApproval
	require.NoError(t, json.Unmarshal([]byte(payload), &sta))
	require.Equal(t, "ServerTaskApprovals-1", sta.Id)
	require.Equal(t, "octopus", sta.ApprovalProviderId)
	require.Equal(t, ChangeRequestApprovalStatePreApproval, sta.ChangeRequest.ChangeRequestApprovalState)
	require.Equal(t, "CHG001", sta.ChangeRequest.Number)
}

func TestServerTaskApprovalDetailUnmarshal(t *testing.T) {
	payload := `{
		"Id": "ServerTaskApprovals-1",
		"SpaceId": "Spaces-1",
		"Approvals": [
			{"Id":"Approvals-1","SpaceId":"Spaces-1","Name":"a","ServerTaskApprovalId":"ServerTaskApprovals-1","UserId":"Users-1","Status":"Approved","Notes":"lgtm"}
		],
		"ApprovalState": "Approved",
		"ApprovingUsers": [{"Id":"Users-1","DisplayName":"Alice","DisplayIdAndName":false}],
		"ApprovalsCount": 1,
		"MinimumApproversRequired": 2
	}`
	var detail ServerTaskApprovalDetail
	require.NoError(t, json.Unmarshal([]byte(payload), &detail))
	require.Equal(t, 1, detail.ApprovalsCount)
	require.Equal(t, 2, detail.MinimumApproversRequired)
	require.Len(t, detail.Approvals, 1)
	require.Equal(t, ApprovalStatusApproved, detail.Approvals[0].Status)
	require.Equal(t, "Alice", detail.ApprovingUsers[0].DisplayName)
}

func TestApprovalMarshal(t *testing.T) {
	approval := &Approval{
		ServerTaskApprovalId: "ServerTaskApprovals-1",
		Status:               ApprovalStatusRejected,
		Notes:                "needs change window",
	}
	data, err := json.Marshal(approval)
	require.NoError(t, err)
	expected := `{
		"ServerTaskApprovalId": "ServerTaskApprovals-1",
		"Status": "Rejected",
		"Notes": "needs change window"
	}`
	require.JSONEq(t, expected, string(data))
}
