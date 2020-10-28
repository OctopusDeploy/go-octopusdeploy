package octopusdeploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpointNew(t *testing.T) {
	testCases := []struct {
		name               string
		communicationStyle string
	}{
		{"Valid", "TentaclePassive"},
		{"Valid", "None"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			endpoint := &endpoint{
				CommunicationStyle: tc.communicationStyle,
			}
			assert.NoError(t, endpoint.Validate())

			endpoint = newEndpoint(tc.communicationStyle)
			assert.NoError(t, endpoint.Validate())
		})
	}
}
