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

// TestApprovalUnmarshalIncludesApproverDetails decodes a payload in the shape the
// server actually returns for a vote, including the approver detail fields.
func TestApprovalUnmarshalIncludesApproverDetails(t *testing.T) {
	payload := `{
		"SpaceId": "Spaces-1",
		"Id": "Approvals-1",
		"Name": "Deploy Test Project release 0.0.1 to Staging",
		"ServerTaskApprovalId": "ServerTaskApprovals-1",
		"UserId": "Users-61",
		"ApproverUsername": "michelle.obrien@octopus.com",
		"ApproverDisplayName": "Michelle O'Brien",
		"ApproverEmailAddress": "michelle.obrien@octopus.com",
		"Status": "Approved",
		"Notes": "looks good"
	}`
	var approval Approval
	require.NoError(t, json.Unmarshal([]byte(payload), &approval))

	require.Equal(t, "Approvals-1", approval.Id)
	require.Equal(t, "Users-61", approval.UserId)
	require.Equal(t, "michelle.obrien@octopus.com", approval.ApproverUsername)
	require.Equal(t, "Michelle O'Brien", approval.ApproverDisplayName)
	require.Equal(t, "michelle.obrien@octopus.com", approval.ApproverEmailAddress)
	require.Equal(t, ApprovalStatusApproved, approval.Status)
}

// TestApprovalsListSpansChangeRequest documents that a list of approvals is scoped
// to the change request, so entries can reference sibling server task approvals.
func TestApprovalsListSpansChangeRequest(t *testing.T) {
	payload := `[
		{"Id":"Approvals-1","ServerTaskApprovalId":"ServerTaskApprovals-1","UserId":"Users-1","Status":"Approved","Notes":""},
		{"Id":"Approvals-2","ServerTaskApprovalId":"ServerTaskApprovals-2","UserId":"Users-2","Status":"Rejected","Notes":""}
	]`
	var approvals []*Approval
	require.NoError(t, json.Unmarshal([]byte(payload), &approvals))

	require.Len(t, approvals, 2)
	require.NotEqual(t, approvals[0].ServerTaskApprovalId, approvals[1].ServerTaskApprovalId)
}

// TestServerTaskApprovalUnmarshalIsBare guards the get-by-id decode: the server
// returns the resource directly rather than wrapping it in an envelope.
func TestServerTaskApprovalUnmarshalIsBare(t *testing.T) {
	payload := `{
		"Id": "ServerTaskApprovals-1",
		"ServerTaskId": "ServerTasks-272141",
		"SpaceId": "Spaces-1",
		"ApprovalProviderId": "octopus",
		"Name": "",
		"ChangeRequest": {
			"Id": "446dd494-3f20-45de-8a22-f47159e4b93b",
			"Number": "OCT272141",
			"Description": "Octopus: Deploy \"Test Project\" version 0.0.1 to \"Staging\"",
			"Active": true,
			"Type": null,
			"ChangeRequestApprovalState": "Approved",
			"Self": null
		}
	}`
	var sta ServerTaskApproval
	require.NoError(t, json.Unmarshal([]byte(payload), &sta))

	require.Equal(t, "ServerTaskApprovals-1", sta.Id)
	require.Equal(t, "446dd494-3f20-45de-8a22-f47159e4b93b", sta.ChangeRequest.Id)
	require.Equal(t, ChangeRequestApprovalStateApproved, sta.ChangeRequest.ChangeRequestApprovalState)
}

func TestChangeRequestApprovalStateValues(t *testing.T) {
	cases := map[ChangeRequestApprovalState]string{
		ChangeRequestApprovalStatePreApproval:  "PreApproval",
		ChangeRequestApprovalStateApproved:     "Approved",
		ChangeRequestApprovalStatePostApproval: "PostApproval",
	}
	for state, want := range cases {
		cr := ChangeRequest{ChangeRequestApprovalState: state}
		data, err := json.Marshal(cr)
		require.NoError(t, err)
		require.Contains(t, string(data), `"ChangeRequestApprovalState":"`+want+`"`)
	}
}
