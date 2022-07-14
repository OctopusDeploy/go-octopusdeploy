package filters

import (
	"encoding/json"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
)

func FromJson(rawMessage *json.RawMessage) (ITriggerFilter, error) {
	if rawMessage == nil {
		return nil, internal.CreateInvalidParameterError("FromJson", "rawMessage")
	}

	var filter triggerFilter
	if err := json.Unmarshal(*rawMessage, &filter); err != nil {
		return nil, err
	}

	switch filter.Type {
	case ContinuousDailySchedule:
		var filter *ContinuousDailyScheduledTriggerFilter
		err := json.Unmarshal(*rawMessage, &filter)
		return filter, err
	case CronExpressionSchedule:
		var filter *CronScheduledTriggerFilter
		err := json.Unmarshal(*rawMessage, &filter)
		return filter, err
	case DailySchedule:
		var filter *DailyScheduledTriggerFilter
		err := json.Unmarshal(*rawMessage, &filter)
		return filter, err
	case DaysPerMonthSchedule:
		var filter *monthlyScheduledTriggerFilter
		err := json.Unmarshal(*rawMessage, &filter)
		return filter, err
	case DaysPerWeekSchedule:
		// TODO: sort this out
	case MachineFilter:
		var filter *DeploymentTargetFilter
		err := json.Unmarshal(*rawMessage, &filter)
		return filter, err
	case OnceDailySchedule:
		var filter *OnceDailyScheduledTriggerFilter
		err := json.Unmarshal(*rawMessage, &filter)
		return filter, err
	}

	return nil, fmt.Errorf("unable to unmarshal filter from JSON")
}
