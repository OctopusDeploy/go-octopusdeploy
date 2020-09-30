package model

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/stretchr/testify/assert"
)

func TestMachineEndpointNew(t *testing.T) {
	defaultWorkerPoolID := "worker-id"
	proxyID := "proxy-id"
	thumbprint := "thumbprint"
	uri := "https://example.com"

	testCases := []struct {
		name                string
		communicationStyle  enum.CommunicationStyle
		defaultWorkerPoolID string
		proxyID             string
		thumbprint          string
		uri                 string
	}{
		{"Invalid", enum.TentaclePassive, emptyString, emptyString, emptyString, emptyString},
		{"Invalid Communication Style", enum.NoCommunicationStyle, defaultWorkerPoolID, proxyID, thumbprint, uri},
		{"Invalid Worker Pool ID", enum.TentaclePassive, emptyString, emptyString, emptyString, uri},
		{"Invalid Proxy ID", enum.TentaclePassive, defaultWorkerPoolID, emptyString, emptyString, uri},
		{"Invalid Thumbprint", enum.TentaclePassive, defaultWorkerPoolID, proxyID, emptyString, uri},
		{"Invalid URI", enum.TentaclePassive, defaultWorkerPoolID, proxyID, thumbprint, "123"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			endpoint := &MachineEndpoint{
				CommunicationStyle:  tc.communicationStyle,
				DefaultWorkerPoolID: tc.defaultWorkerPoolID,
				ProxyID:             &tc.proxyID,
				Thumbprint:          tc.thumbprint,
				URI:                 tc.uri,
			}
			assert.Error(t, endpoint.Validate())

			endpoint, err := NewMachineEndpoint(
				tc.uri,
				tc.thumbprint,
				tc.communicationStyle,
				tc.proxyID,
				tc.defaultWorkerPoolID,
			)
			assert.NoError(t, err)
			assert.Error(t, endpoint.Validate())
		})
	}
}
