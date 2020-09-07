package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
)

const (
	interruptionID    = "Interruptions-1"
	interruptionTitle = "InterruptionTitle"
)

func TestInterruptionsGetAll(t *testing.T) {
	client, err := client.GetFakeOctopusClient(t, "/api/interruptions", http.StatusOK, getInterruptionsResponseJSON)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	interruptions, err := client.Interruptions.GetAll()

	assert.NoError(t, err)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(*interruptions))
	assert.Equal(t, interruptionTitle, (*interruptions)[0].Title)
	assert.Equal(t, interruptionID, (*interruptions)[0].ID)
	assert.Equal(t, true, (*interruptions)[0].IsPending)
	assert.Equal(t, "/api/interruptions/Interruptions-1", (*interruptions)[0].Links["Self"])
	assert.Equal(t, "/api/interruptions/Interruptions-1/submit", (*interruptions)[0].Links["Submit"])
	assert.Equal(t, "/api/interruptions/Interruptions-1/responsible", (*interruptions)[0].Links["Responsible"])
}

func TestInterruptionsGet(t *testing.T) {
	interruptionID := interruptionID
	client, err := client.GetFakeOctopusClient(t, "/api/interruptions/"+interruptionID, http.StatusOK, interruptionJSON)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	interruption, err := client.Interruptions.Get(interruptionID)

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, interruptionTitle, interruption.Title)
	assert.Equal(t, interruptionID, interruption.ID)
	assert.Equal(t, true, interruption.IsPending)
	assert.Equal(t, "/api/interruptions/Interruptions-1", interruption.Links["Self"])
	assert.Equal(t, "/api/interruptions/Interruptions-1/submit", interruption.Links["Submit"])
	assert.Equal(t, "/api/interruptions/Interruptions-1/responsible", interruption.Links["Responsible"])
}

func TestInterruptionsTakeResponsibility(t *testing.T) {
	interruptionID := interruptionID
	client, err := client.GetFakeOctopusClient(t, "/api/interruptions/"+interruptionID+"/responsible", http.StatusOK, interruptionUserJSON)

	assert.NoError(t, err)

	interruption, err := getInterruptonFromJSON(interruptionJSON)

	assert.NoError(t, err)
	assert.Nil(t, err)

	user, err := client.Interruptions.TakeResponsability(interruption)

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, "user@example.com", user.EmailAddress)
	assert.Equal(t, "user@example.com", user.Username)
	assert.Equal(t, "Users-1", user.ID)
	assert.Equal(t, "User Name", user.DisplayName)
}

func TestInterruptionsGetResponsibilities(t *testing.T) {
	interruptionID := "Interruptions-1"
	client, err := client.GetFakeOctopusClient(t, "/api/interruptions/"+interruptionID+"/responsible", http.StatusOK, interruptionUserJSON)

	assert.NoError(t, err)

	interruption, err := getInterruptonFromJSON(interruptionJSON)

	assert.NoError(t, err)
	assert.Nil(t, err)

	user, err := client.Interruptions.GetResponsability(interruption)

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, "user@example.com", user.EmailAddress)
	assert.Equal(t, "user@example.com", user.Username)
	assert.Equal(t, "Users-1", user.ID)
	assert.Equal(t, "User Name", user.DisplayName)
}

func TestInterruptionsSubmit(t *testing.T) {
	interruptionID := "Interruptions-1"
	client, err := client.GetFakeOctopusClient(t, "/api/interruptions/"+interruptionID+"/submit", http.StatusOK, interruptionSubmittedJSON)

	assert.NoError(t, err)

	interruption, err := getInterruptonFromJSON(interruptionJSON)

	assert.NoError(t, err)
	assert.Nil(t, err)

	submitRequest := model.InterruptionSubmitRequest{
		Instructions: "Approve The Deployment",
		Notes:        "",
		Result:       model.ManualInterverventionApprove,
	}

	i, err := client.Interruptions.Submit(interruption, &submitRequest)

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, false, i.IsPending)
	assert.Equal(t, "Interruptions-1", i.ID)
}

var getInterruptionsResponseJSON = fmt.Sprintf(`
{
	"ItemType": "Interruption",
	"TotalResults": 1,
	"ItemsPerPage": 30,
	"NumberOfPages": 1,
	"LastPageNumber": 0,
	"Items": [
	  %v
	],
	"Links": {
	  "Self": "/api/interruptions?regarding=&pendingOnly=False",
	  "Template": "/api/interruptions{?skip,take,regarding,pendingOnly,ids}",
	  "Page.All": "/api/interruptions?skip=0&take=2147483647",
	  "Page.Current": "/api/interruptions?skip=0&take=30",
	  "Page.Last": "/api/interruptions?skip=0&take=30"
	}
  }
`, interruptionJSON)

const interruptionJSON = `
{
	"Id": "Interruptions-1",
	"Title": "InterruptionTitle",
	"Created": "2018-12-31T13:38:39.440+00:00",
	"IsPending": true,
	"Form": {
	  "Values": {
		"Instructions": null,
		"Notes": null,
		"Result": null
	  },
	  "Elements": [
		{
		  "Name": "Instructions",
		  "Control": {
			"Type": "Paragraph",
			"Text": "Manual Approval",
			"ResolveLinks": false
		  },
		  "IsValueRequired": false
		},
		{
		  "Name": "Notes",
		  "Control": {
			"Type": "TextArea",
			"Label": "Notes"
		  },
		  "IsValueRequired": false
		},
		{
		  "Name": "Result",
		  "Control": {
			"Type": "SubmitButtonGroup",
			"Buttons": [
			  {
				"Text": "Proceed",
				"Value": "Proceed",
				"RequiresConfirmation": false
			  },
			  {
				"Text": "Abort",
				"Value": "Abort",
				"RequiresConfirmation": true
			  }
			]
		  },
		  "IsValueRequired": false
		}
	  ]
	},
	"RelatedDocumentIds": [
	  "Deployments-1",
	  "ServerTasks-1",
	  "Projects-1",
	  "Environments-1"
	],
	"ResponsibleTeamIds": [
	  "Teams-1"
	],
	"ResponsibleUserId": null,
	"CanTakeResponsibility": true,
	"HasResponsibility": false,
	"TaskId": "ServerTasks-1",
	"CorrelationId": "ServerTasks-1_CNMPMXUEE6/24921723bcb741409134629931dd6b97/dbfadf8e4aaa4acbb45d45c4c39d0f12",
	"IsLinkedToOtherInterruption": false,
	"Links": {
	  "Self": "/api/interruptions/Interruptions-1",
	  "Submit": "/api/interruptions/Interruptions-1/submit",
	  "Responsible": "/api/interruptions/Interruptions-1/responsible"
	}
  }
  `

const interruptionUserJSON = `
  {
	"Id": "Users-1",
	"Username": "user@example.com",
	"DisplayName": "User Name",
	"IsActive": true,
	"IsService": false,
	"EmailAddress": "user@example.com",
	"CanPasswordBeEdited": true,
	"IsRequestor": true,
	"Links": {
	  "Self": "/api/users/Users-1",
	  "Permissions": "/api/users/Users-1/permissions",
	  "ApiKeys": "/api/users/Users-1/apikeys{/id}{?skip,take}",
	  "Avatar": "https://www.gravatar.com/avatar/ae0e3d90eeddb248c041469b38cc64fd?d=blank"
	}
  }`

const interruptionSubmittedJSON = `
  {
	  "Id": "Interruptions-1",
	  "Title": "InterruptionTitle",
	  "Created": "2018-12-31T13:38:39.440+00:00",
	  "IsPending": false,
	  "Form": {
		"Values": {
		  "Instructions": null,
		  "Notes": null,
		  "Result": null
		},
		"Elements": [
		  {
			"Name": "Instructions",
			"Control": {
			  "Type": "Paragraph",
			  "Text": "Manual Approval",
			  "ResolveLinks": false
			},
			"IsValueRequired": false
		  },
		  {
			"Name": "Notes",
			"Control": {
			  "Type": "TextArea",
			  "Label": "Notes"
			},
			"IsValueRequired": false
		  },
		  {
			"Name": "Result",
			"Control": {
			  "Type": "SubmitButtonGroup",
			  "Buttons": [
				{
				  "Text": "Proceed",
				  "Value": "Proceed",
				  "RequiresConfirmation": false
				},
				{
				  "Text": "Abort",
				  "Value": "Abort",
				  "RequiresConfirmation": true
				}
			  ]
			},
			"IsValueRequired": false
		  }
		]
	  },
	  "RelatedDocumentIds": [
		"Deployments-1",
		"ServerTasks-1",
		"Projects-1",
		"Environments-1"
	  ],
	  "ResponsibleTeamIds": [
		"Teams-1"
	  ],
	  "ResponsibleUserId": null,
	  "CanTakeResponsibility": true,
	  "HasResponsibility": false,
	  "TaskId": "ServerTasks-1",
	  "CorrelationId": "ServerTasks-1_CNMPMXUEE6/24921723bcb741409134629931dd6b97/dbfadf8e4aaa4acbb45d45c4c39d0f12",
	  "IsLinkedToOtherInterruption": false,
	  "Links": {
		"Self": "/api/interruptions/Interruptions-1",
		"Submit": "/api/interruptions/Interruptions-1/submit",
		"Responsible": "/api/interruptions/Interruptions-1/responsible"
	  }
	}
	`

func getInterruptonFromJSON(interruptionJSON string) (*model.Interruption, error) {
	var interruption model.Interruption
	err := json.Unmarshal([]byte(interruptionJSON), &interruption)
	return &interruption, err
}
