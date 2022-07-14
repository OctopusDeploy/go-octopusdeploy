package actions

import (
	"encoding/json"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
)

func FromJson(rawMessage *json.RawMessage) (ITriggerAction, error) {
	if rawMessage == nil {
		return nil, internal.CreateInvalidParameterError("FromJson", "rawMessage")
	}

	var action triggerAction
	if err := json.Unmarshal(*rawMessage, &action); err != nil {
		return nil, err
	}

	switch action.Type {
	case AutoDeploy:
		var action *AutoDeployAction
		err := json.Unmarshal(*rawMessage, &action)
		return action, err
	case DeployLatestRelease:
		var action *DeployLatestReleaseAction
		err := json.Unmarshal(*rawMessage, &action)
		return action, err
	case DeployNewRelease:
		var action *DeployNewReleaseAction
		err := json.Unmarshal(*rawMessage, &action)
		return action, err
	case RunRunbook:
		var action *RunRunbookAction
		err := json.Unmarshal(*rawMessage, &action)
		return action, err
	}

	return nil, fmt.Errorf("unable to unmarshal action from JSON")
}
