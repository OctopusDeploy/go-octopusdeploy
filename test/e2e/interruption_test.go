package e2e

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"testing"

// 	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
// 	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/interruptions"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// const (
// 	interruptionID    string = "Interruptions-1"
// 	interruptionTitle string = "InterruptionTitle"
// )

// func TestInterruptionsGetAll(t *testing.T) {
// 	client, err := client.GetFakeOctopusClient(t, "/api/interruptions", http.StatusOK, getInterruptionsResponseJSON)
// 	require.NoError(t, err)
// 	require.NotNil(t, client)

// 	interruptions, err := client.Interruptions.GetAll()
// 	require.NoError(t, err)

// 	assert.Equal(t, 1, len(interruptions))
// 	assert.Equal(t, interruptionTitle, (interruptions)[0].Title)
// 	assert.Equal(t, interruptionID, (interruptions)[0].GetID())
// 	assert.Equal(t, true, (interruptions)[0].IsPending)
// 	assert.Equal(t, "/api/interruptions/Interruptions-1", (interruptions)[0].Links["Self"])
// 	assert.Equal(t, "/api/interruptions/Interruptions-1/submit", (interruptions)[0].Links["Submit"])
// 	assert.Equal(t, "/api/interruptions/Interruptions-1/responsible", (interruptions)[0].Links["Responsible"])
// }

// func TestInterruptionsGet(t *testing.T) {
// 	interruptionID := interruptionID
// 	client, err := client.GetFakeOctopusClient(t, "/api/interruptions/"+interruptionID, http.StatusOK, interruptionJSON)
// 	require.NoError(t, err)
// 	require.NotNil(t, client)

// 	interruption, err := client.Interruptions.GetByID(interruptionID)
// 	require.NoError(t, err)
// 	require.NotNil(t, interruption)

// 	assert.Equal(t, interruptionTitle, interruption.Title)
// 	assert.Equal(t, interruptionID, interruption.GetID())
// 	assert.Equal(t, true, interruption.IsPending)
// 	assert.Equal(t, "/api/interruptions/Interruptions-1", interruption.Links["Self"])
// 	assert.Equal(t, "/api/interruptions/Interruptions-1/submit", interruption.Links["Submit"])
// 	assert.Equal(t, "/api/interruptions/Interruptions-1/responsible", interruption.Links["Responsible"])
// }

// func TestInterruptionsTakeResponsibility(t *testing.T) {
// 	interruptionID := interruptionID
// 	client, err := client.GetFakeOctopusClient(t, "/api/interruptions/"+interruptionID+"/responsible", http.StatusOK, interruptionUserJSON)
// 	require.NoError(t, err)
// 	require.NotNil(t, client)

// 	interruption, err := getInterruptionFromJSON(interruptionJSON)
// 	require.NoError(t, err)
// 	require.NotNil(t, interruption)

// 	user, err := client.Interruptions.TakeResponsibility(interruption)
// 	require.NoError(t, err)
// 	require.Equal(t, "user@example.com", user.EmailAddress)
// 	require.Equal(t, "user@example.com", user.Username)
// 	require.Equal(t, "Users-1", user.GetID())
// 	require.Equal(t, "User Name", user.DisplayName)
// }

// func TestInterruptionsGetResponsibilities(t *testing.T) {
// 	interruptionID := "Interruptions-1"
// 	client, err := client.GetFakeOctopusClient(t, "/api/interruptions/"+interruptionID+"/responsible", http.StatusOK, interruptionUserJSON)
// 	require.NoError(t, err)
// 	require.NotNil(t, client)

// 	interruption, err := getInterruptionFromJSON(interruptionJSON)
// 	require.NoError(t, err)
// 	require.NotNil(t, interruption)

// 	user, err := client.Interruptions.GetResponsibility(interruption)
// 	require.NoError(t, err)
// 	require.NotNil(t, user)
// 	require.Equal(t, "user@example.com", user.EmailAddress)
// 	require.Equal(t, "user@example.com", user.Username)
// 	require.Equal(t, "Users-1", user.GetID())
// 	require.Equal(t, "User Name", user.DisplayName)
// }

// func TestInterruptionsSubmit(t *testing.T) {
// 	interruptionID := "Interruptions-1"
// 	client, err := client.GetFakeOctopusClient(t, "/api/interruptions/"+interruptionID+"/submit", http.StatusOK, interruptionSubmittedJSON)
// 	require.NoError(t, err)
// 	require.NotNil(t, client)

// 	interruption, err := getInterruptionFromJSON(interruptionJSON)
// 	require.NoError(t, err)
// 	require.NotNil(t, interruption)

// 	submitRequest := interruptions.InterruptionSubmitRequest{
// 		Instructions: "Approve The Deployment",
// 		Notes:        "",
// 		Result:       interruptions.ManualInterventionApprove,
// 	}

// 	i, err := client.Interruptions.Submit(interruption, &submitRequest)
// 	require.NoError(t, err)
// 	require.NotNil(t, i)
// 	require.Equal(t, false, i.IsPending)
// 	require.Equal(t, "Interruptions-1", i.GetID())
// }

// var getInterruptionsResponseJSON = fmt.Sprintf(`
// {
// 	"ItemType": "Interruption",
// 	"TotalResults": 1,
// 	"ItemsPerPage": 30,
// 	"NumberOfPages": 1,
// 	"LastPageNumber": 0,
// 	"Items": [
// 	  %v
// 	],
// 	"Links": {
// 	  "Self": "/api/interruptions?regarding=&pendingOnly=False",
// 	  "Template": "/api/interruptions{?skip,take,regarding,pendingOnly,ids}",
// 	  "Page.All": "/api/interruptions?skip=0&take=2147483647",
// 	  "Page.Current": "/api/interruptions?skip=0&take=30",
// 	  "Page.Last": "/api/interruptions?skip=0&take=30"
// 	}
//   }
// `, interruptionJSON)

// const interruptionJSON = `
// {
// 	"Id": "Interruptions-1",
// 	"Title": "InterruptionTitle",
// 	"Created": "2018-12-31T13:38:39.440+00:00",
// 	"IsPending": true,
// 	"Form": {
// 	  "Values": {
// 		"Instructions": null,
// 		"Notes": null,
// 		"Result": null
// 	  },
// 	  "Elements": [
// 		{
// 		  "Name": "Instructions",
// 		  "Control": {
// 			"Type": "Paragraph",
// 			"Text": "Manual Approval",
// 			"ResolveLinks": false
// 		  },
// 		  "IsValueRequired": false
// 		},
// 		{
// 		  "Name": "Notes",
// 		  "Control": {
// 			"Type": "TextArea",
// 			"Label": "Notes"
// 		  },
// 		  "IsValueRequired": false
// 		},
// 		{
// 		  "Name": "Result",
// 		  "Control": {
// 			"Type": "SubmitButtonGroup",
// 			"Buttons": [
// 			  {
// 				"Text": "Proceed",
// 				"Value": "Proceed",
// 				"RequiresConfirmation": false
// 			  },
// 			  {
// 				"Text": "Abort",
// 				"Value": "Abort",
// 				"RequiresConfirmation": true
// 			  }
// 			]
// 		  },
// 		  "IsValueRequired": false
// 		}
// 	  ]
// 	},
// 	"RelatedDocumentIds": [
// 	  "Deployments-1",
// 	  "ServerTasks-1",
// 	  "Projects-1",
// 	  "Environments-1"
// 	],
// 	"ResponsibleTeamIds": [
// 	  "Teams-1"
// 	],
// 	"ResponsibleUserId": null,
// 	"CanTakeResponsibility": true,
// 	"HasResponsibility": false,
// 	"TaskId": "ServerTasks-1",
// 	"CorrelationId": "ServerTasks-1_CNMPMXUEE6/24921723bcb741409134629931dd6b97/dbfadf8e4aaa4acbb45d45c4c39d0f12",
// 	"IsLinkedToOtherInterruption": false,
// 	"Links": {
// 	  "Self": "/api/interruptions/Interruptions-1",
// 	  "Submit": "/api/interruptions/Interruptions-1/submit",
// 	  "Responsible": "/api/interruptions/Interruptions-1/responsible"
// 	}
//   }
//   `

// const interruptionUserJSON = `
//   {
// 	"Id": "Users-1",
// 	"Username": "user@example.com",
// 	"DisplayName": "User Name",
// 	"IsActive": true,
// 	"IsService": false,
// 	"EmailAddress": "user@example.com",
// 	"CanPasswordBeEdited": true,
// 	"IsRequestor": true,
// 	"Links": {
// 	  "Self": "/api/users/Users-1",
// 	  "Permissions": "/api/users/Users-1/permissions",
// 	  "ApiKeys": "/api/users/Users-1/apikeys{/id}{?skip,take}",
// 	  "Avatar": "https://www.gravatar.com/avatar/ae0e3d90eeddb248c041469b38cc64fd?d=blank"
// 	}
//   }`

// const interruptionSubmittedJSON = `
//   {
// 	  "Id": "Interruptions-1",
// 	  "Title": "InterruptionTitle",
// 	  "Created": "2018-12-31T13:38:39.440+00:00",
// 	  "IsPending": false,
// 	  "Form": {
// 		"Values": {
// 		  "Instructions": null,
// 		  "Notes": null,
// 		  "Result": null
// 		},
// 		"Elements": [
// 		  {
// 			"Name": "Instructions",
// 			"Control": {
// 			  "Type": "Paragraph",
// 			  "Text": "Manual Approval",
// 			  "ResolveLinks": false
// 			},
// 			"IsValueRequired": false
// 		  },
// 		  {
// 			"Name": "Notes",
// 			"Control": {
// 			  "Type": "TextArea",
// 			  "Label": "Notes"
// 			},
// 			"IsValueRequired": false
// 		  },
// 		  {
// 			"Name": "Result",
// 			"Control": {
// 			  "Type": "SubmitButtonGroup",
// 			  "Buttons": [
// 				{
// 				  "Text": "Proceed",
// 				  "Value": "Proceed",
// 				  "RequiresConfirmation": false
// 				},
// 				{
// 				  "Text": "Abort",
// 				  "Value": "Abort",
// 				  "RequiresConfirmation": true
// 				}
// 			  ]
// 			},
// 			"IsValueRequired": false
// 		  }
// 		]
// 	  },
// 	  "RelatedDocumentIds": [
// 		"Deployments-1",
// 		"ServerTasks-1",
// 		"Projects-1",
// 		"Environments-1"
// 	  ],
// 	  "ResponsibleTeamIds": [
// 		"Teams-1"
// 	  ],
// 	  "ResponsibleUserId": null,
// 	  "CanTakeResponsibility": true,
// 	  "HasResponsibility": false,
// 	  "TaskId": "ServerTasks-1",
// 	  "CorrelationId": "ServerTasks-1_CNMPMXUEE6/24921723bcb741409134629931dd6b97/dbfadf8e4aaa4acbb45d45c4c39d0f12",
// 	  "IsLinkedToOtherInterruption": false,
// 	  "Links": {
// 		"Self": "/api/interruptions/Interruptions-1",
// 		"Submit": "/api/interruptions/Interruptions-1/submit",
// 		"Responsible": "/api/interruptions/Interruptions-1/responsible"
// 	  }
// 	}
// 	`

// func getInterruptionFromJSON(interruptionJSON string) (*interruptions.Interruption, error) {
// 	var interruption interruptions.Interruption
// 	err := json.Unmarshal([]byte(interruptionJSON), &interruption)
// 	return &interruption, err
// }
