package triggers_test

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/filters"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/triggers"
	"github.com/stretchr/testify/assert"
)

func TestTriggerJsonSerialization(t *testing.T) {
	t.Run("serialize including scheduled filter", func(t *testing.T) {
		action := actions.NewAutoDeployAction(true)

		start := time.Date(2022, 0, 14, 9, 0, 0, 0, time.UTC)
		filter := filters.NewDaysPerMonthScheduledTriggerFilter(filters.DayOfMonth, start)
		filter.DayNumberOfMonth = "1"
		filter.TimeZone = "Asia/Kuala_Lumpur"

		project := projects.NewProject("foo", "Lifecycles-123", "ProjectGroup-123")
		project.SetID("Projects-123")

		projectTrigger := triggers.NewProjectTrigger("triggerName", "triggerDescription", false, project, action, filter)

		output := &bytes.Buffer{}
		err := json.NewEncoder(output).Encode(projectTrigger)
		assert.Nil(t, err)

		// note that Go puts a 'Z' on the time
		assert.Equal(t, heredoc.Doc(`
			{"Action":{"ShouldRedeployWhenMachineHasBeenDeployedTo":true,"ActionType":"AutoDeploy"},"Filter":{"MonthlyScheduleType":"DayOfMonth","DayNumberOfMonth":"1","StartTime":"2021-12-14T09:00:00Z","Timezone":"Asia/Kuala_Lumpur","FilterType":"DaysPerMonthSchedule"},"IsDisabled":false,"Name":"triggerName","ProjectId":"Projects-123"}
			`), output.String())
	})

	t.Run("deserialize scheduled filter with standard go timeformat", func(t *testing.T) {
		data := []byte(`{"Action":{"ShouldRedeployWhenMachineHasBeenDeployedTo":true,"ActionType":"AutoDeploy"},"Filter":{"MonthlyScheduleType":"DayOfMonth","DayNumberOfMonth":"1","StartTime":"2021-12-14T09:00:00Z","Timezone":"Asia/Kuala_Lumpur","FilterType":"DaysPerMonthSchedule"},"Name":"triggerName","ProjectId":"Projects-123"}`)
		var output = new(triggers.ProjectTrigger)
		err := json.NewDecoder(bytes.NewReader(data)).Decode(output)
		assert.Nil(t, err)

		action := actions.NewAutoDeployAction(true)
		action.Links = nil

		filter := filters.NewDaysPerMonthScheduledTriggerFilter(filters.DayOfMonth, time.Date(2022, 0, 14, 9, 0, 0, 0, time.UTC))
		filter.DayNumberOfMonth = "1"
		filter.TimeZone = "Asia/Kuala_Lumpur"
		filter.Links = nil

		assert.Equal(t, action, output.Action)
		assert.Equal(t, filter, output.Filter)

		assert.Equal(t, &triggers.ProjectTrigger{
			Action:      action,
			Description: "",
			Filter:      filter,
			IsDisabled:  false,
			Name:        "triggerName",
			ProjectID:   "Projects-123",
		}, output)
	})

	t.Run("deserialize scheduled filter with an unzoned ISO format with milliseconds", func(t *testing.T) {
		// 2022.3 server returns this
		data := []byte(`{"Action":{"ShouldRedeployWhenMachineHasBeenDeployedTo":true,"ActionType":"AutoDeploy"},"Filter":{"MonthlyScheduleType":"DayOfMonth","DayNumberOfMonth":"1","StartTime":"2021-12-14T09:00:00.000","Timezone":"Asia/Kuala_Lumpur","FilterType":"DaysPerMonthSchedule"},"Name":"triggerName","ProjectId":"Projects-123"}`)
		var output = new(triggers.ProjectTrigger)
		err := json.NewDecoder(bytes.NewReader(data)).Decode(output)
		assert.Nil(t, err)

		action := actions.NewAutoDeployAction(true)
		action.Links = nil

		filter := filters.NewDaysPerMonthScheduledTriggerFilter(filters.DayOfMonth, time.Date(2022, 0, 14, 9, 0, 0, 0, time.UTC))
		filter.DayNumberOfMonth = "1"
		filter.TimeZone = "Asia/Kuala_Lumpur"
		filter.Links = nil

		assert.Equal(t, action, output.Action)
		assert.Equal(t, filter, output.Filter)

		assert.Equal(t, &triggers.ProjectTrigger{
			Action:      action,
			Description: "",
			Filter:      filter,
			IsDisabled:  false,
			Name:        "triggerName",
			ProjectID:   "Projects-123",
		}, output)
	})

	t.Run("deserialize OnceDailySchedule filter from a 2022.3 server", func(t *testing.T) {
		// Note the Startime with 3 decimal places of milliseconds, with no timezone info, which Go doesn't normally deal with
		// captured verbatim from real server
		data := []byte(heredoc.Doc(`
		{
		  "Id": "ProjectTriggers-21",
		  "Name": "Daily Scheduled Trigger 9am each day",
		  "ProjectId": "Projects-314",
		  "IsDisabled": false,
		  "Filter": {
			"FilterType": "OnceDailySchedule",
			"StartTime": "2022-09-14T09:00:00.000",
			"DaysOfWeek": ["Tuesday"],
			"Timezone": "Singapore Standard Time",
			"Id": null,
			"LastModifiedOn": null,
			"LastModifiedBy": null,
			"Links": {}
		  },
		  "Action": {
			"ActionType": "DeployLatestRelease",
			"Variables": null,
			"SourceEnvironmentIds": [
			  "Environments-101"
			],
			"DestinationEnvironmentId": "Environments-101",
			"ShouldRedeployWhenReleaseIsCurrent": true,
			"ChannelId": null,
			"TenantIds": [],
			"TenantTags": [],
			"Id": null,
			"LastModifiedOn": null,
			"LastModifiedBy": null,
			"Links": {}
		  },
		  "SpaceId": "Spaces-1",
		  "Description": "",
		  "Links": {
			"Self": "/api/Spaces-1/projects/Projects-314/triggers/ProjectTriggers-21",
			"Project": "/api/Spaces-1/projects/Projects-314"
		  }
		}`))
		var output = new(triggers.ProjectTrigger)
		err := json.NewDecoder(bytes.NewReader(data)).Decode(output)
		assert.Nil(t, err)

		action := actions.NewDeployLatestReleaseAction("Environments-101", true, []string{"Environments-101"}, "")

		filter := filters.NewOnceDailyScheduledTriggerFilter([]filters.Weekday{filters.Tuesday}, time.Date(2022, time.September, 14, 9, 0, 0, 0, time.UTC))
		filter.TimeZone = "Singapore Standard Time"

		assert.Equal(t, action, output.Action)
		assert.Equal(t, filter, output.Filter)

		assert.Equal(t, &triggers.ProjectTrigger{
			Action:      action,
			Description: "",
			Filter:      filter,
			IsDisabled:  false,
			Name:        "Daily Scheduled Trigger 9am each day",
			ProjectID:   "Projects-314",
			SpaceID:     "Spaces-1",
			Resource: resources.Resource{
				ID: "ProjectTriggers-21",
				Links: map[string]string{
					"Project": "/api/Spaces-1/projects/Projects-314",
					"Self":    "/api/Spaces-1/projects/Projects-314/triggers/ProjectTriggers-21",
				},
			},
		}, output)
	})

	t.Run("deserialize DaysPerMonthScheduledTriggerFilter (DateOfMonth) from a 2022.3 server", func(t *testing.T) {
		// captured verbatim from real server
		data := []byte(heredoc.Doc(`
		{
		  "Id": "ProjectTriggers-21",
		  "Name": "Monthly Scheduled Trigger 6th of the month",
		  "ProjectId": "Projects-314",
		  "IsDisabled": false,
		  "Filter": {
			"FilterType": "DaysPerMonthSchedule",
			"StartTime": "2022-09-13T21:00:00.000",
			"MonthlyScheduleType": "DateOfMonth",
			"DateOfMonth": "6",
			"DayNumberOfMonth": null,
			"DayOfWeek": null,
			"Timezone": "Singapore Standard Time",
			"Id": null,
			"LastModifiedOn": null,
			"LastModifiedBy": null,
			"Links": {}
		  },
		  "Action": {
			"ActionType": "DeployLatestRelease",
			"Variables": null,
			"SourceEnvironmentIds": [
			  "Environments-101"
			],
			"DestinationEnvironmentId": "Environments-101",
			"ShouldRedeployWhenReleaseIsCurrent": true,
			"ChannelId": null,
			"TenantIds": [],
			"TenantTags": [],
			"Id": null,
			"LastModifiedOn": null,
			"LastModifiedBy": null,
			"Links": {}
		  },
		  "SpaceId": "Spaces-1",
		  "Description": "",
		  "Links": {
			"Self": "/api/Spaces-1/projects/Projects-314/triggers/ProjectTriggers-21",
			"Project": "/api/Spaces-1/projects/Projects-314"
		  }
		}
		`))
		var output = new(triggers.ProjectTrigger)
		err := json.NewDecoder(bytes.NewReader(data)).Decode(output)
		assert.Nil(t, err)

		action := actions.NewDeployLatestReleaseAction("Environments-101", true, []string{"Environments-101"}, "")

		filter := filters.NewDaysPerMonthScheduledTriggerFilter(filters.DateOfMonth, time.Date(2022, time.September, 13, 21, 0, 0, 0, time.UTC))
		filter.DateOfMonth = "6"
		filter.TimeZone = "Singapore Standard Time"

		assert.Equal(t, action, output.Action)
		assert.Equal(t, filter, output.Filter)

		assert.Equal(t, &triggers.ProjectTrigger{
			Action:      action,
			Description: "",
			Filter:      filter,
			IsDisabled:  false,
			Name:        "Monthly Scheduled Trigger 6th of the month",
			ProjectID:   "Projects-314",
			SpaceID:     "Spaces-1",
			Resource: resources.Resource{
				ID: "ProjectTriggers-21",
				Links: map[string]string{
					"Project": "/api/Spaces-1/projects/Projects-314",
					"Self":    "/api/Spaces-1/projects/Projects-314/triggers/ProjectTriggers-21",
				},
			},
		}, output)
	})

	t.Run("deserialize DaysPerMonthScheduledTriggerFilter (DayNumberOfMonth) from a 2022.3 server", func(t *testing.T) {
		// captured verbatim from real server
		data := []byte(heredoc.Doc(`
		{
		  "Id": "ProjectTriggers-21",
		  "Name": "Monthly Scheduled Trigger Second Wednesday",
		  "ProjectId": "Projects-314",
		  "IsDisabled": false,
		  "Filter": {
			"FilterType": "DaysPerMonthSchedule",
			"StartTime": "2022-09-13T09:00:00.000",
			"MonthlyScheduleType": "DayOfMonth",
			"DateOfMonth": null,
			"DayNumberOfMonth": "2",
			"DayOfWeek": "Wednesday",
			"Timezone": "Singapore Standard Time",
			"Id": null,
			"LastModifiedOn": null,
			"LastModifiedBy": null,
			"Links": {}
		  },
		  "Action": {
			"ActionType": "DeployLatestRelease",
			"Variables": null,
			"SourceEnvironmentIds": [
			  "Environments-101"
			],
			"DestinationEnvironmentId": "Environments-101",
			"ShouldRedeployWhenReleaseIsCurrent": true,
			"ChannelId": null,
			"TenantIds": [],
			"TenantTags": [],
			"Id": null,
			"LastModifiedOn": null,
			"LastModifiedBy": null,
			"Links": {}
		  },
		  "SpaceId": "Spaces-1",
		  "Description": "",
		  "Links": {
			"Self": "/api/Spaces-1/projects/Projects-314/triggers/ProjectTriggers-21",
			"Project": "/api/Spaces-1/projects/Projects-314"
		  }
		}
		`))
		var output = new(triggers.ProjectTrigger)
		err := json.NewDecoder(bytes.NewReader(data)).Decode(output)
		assert.Nil(t, err)

		action := actions.NewDeployLatestReleaseAction("Environments-101", true, []string{"Environments-101"}, "")

		filter := filters.NewDaysPerMonthScheduledTriggerFilter(filters.DayOfMonth, time.Date(2022, time.September, 13, 9, 0, 0, 0, time.UTC))
		filter.DayNumberOfMonth = "2"
		wednesday := filters.Wednesday
		filter.Day = &wednesday
		filter.TimeZone = "Singapore Standard Time"

		assert.Equal(t, action, output.Action)
		assert.Equal(t, filter, output.Filter)

		assert.Equal(t, &triggers.ProjectTrigger{
			Action:      action,
			Description: "",
			Filter:      filter,
			IsDisabled:  false,
			Name:        "Monthly Scheduled Trigger Second Wednesday",
			ProjectID:   "Projects-314",
			SpaceID:     "Spaces-1",
			Resource: resources.Resource{
				ID: "ProjectTriggers-21",
				Links: map[string]string{
					"Project": "/api/Spaces-1/projects/Projects-314",
					"Self":    "/api/Spaces-1/projects/Projects-314/triggers/ProjectTriggers-21",
				},
			},
		}, output)
	})

	t.Run("deserialize CronExpressionSchedule from a 2022.3 server", func(t *testing.T) {
		// captured verbatim from real server
		data := []byte(heredoc.Doc(`
		{
		  "Id": "ProjectTriggers-21",
		  "Name": "Cron Schedule 0 0 06",
		  "ProjectId": "Projects-314",
		  "IsDisabled": false,
		  "Filter": {
			"FilterType": "CronExpressionSchedule",
			"CronExpression": "0 0 6 * * Mon-Fri",
			"Timezone": "Singapore Standard Time",
			"Id": null,
			"LastModifiedOn": null,
			"LastModifiedBy": null,
			"Links": {}
		  },
		  "Action": {
			"ActionType": "DeployLatestRelease",
			"Variables": null,
			"SourceEnvironmentIds": [
			  "Environments-101"
			],
			"DestinationEnvironmentId": "Environments-101",
			"ShouldRedeployWhenReleaseIsCurrent": true,
			"ChannelId": null,
			"TenantIds": [],
			"TenantTags": [],
			"Id": null,
			"LastModifiedOn": null,
			"LastModifiedBy": null,
			"Links": {}
		  },
		  "SpaceId": "Spaces-1",
		  "Description": "",
		  "Links": {
			"Self": "/api/Spaces-1/projects/Projects-314/triggers/ProjectTriggers-21",
			"Project": "/api/Spaces-1/projects/Projects-314"
		  }
		}
		`))
		var output = new(triggers.ProjectTrigger)
		err := json.NewDecoder(bytes.NewReader(data)).Decode(output)
		assert.Nil(t, err)

		action := actions.NewDeployLatestReleaseAction("Environments-101", true, []string{"Environments-101"}, "")

		filter := filters.NewCronScheduledTriggerFilter("0 0 6 * * Mon-Fri", "Singapore Standard Time")
		filter.TimeZone = "Singapore Standard Time"

		assert.Equal(t, action, output.Action)
		assert.Equal(t, filter, output.Filter)

		assert.Equal(t, &triggers.ProjectTrigger{
			Action:      action,
			Description: "",
			Filter:      filter,
			IsDisabled:  false,
			Name:        "Cron Schedule 0 0 06",
			ProjectID:   "Projects-314",
			SpaceID:     "Spaces-1",
			Resource: resources.Resource{
				ID: "ProjectTriggers-21",
				Links: map[string]string{
					"Project": "/api/Spaces-1/projects/Projects-314",
					"Self":    "/api/Spaces-1/projects/Projects-314/triggers/ProjectTriggers-21",
				},
			},
		}, output)
	})
}
