package machines

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type ArgoCDGateway struct {
	SpaceID                 string               `json:"SpaceId,omitempty"`
	Name                    string               `json:"Name"`
	ClientId                string               `json:"ClientId"`
	HealthStatus            string               `json:"HealthStatus,omitempty"`
	HealthLastChecked       string               `json:"HealthLastChecked,omitempty"`
	HealthCheckServerTaskId string               `json:"HealthCheckServerTaskId,omitempty"`
	WebUIUri                string               `json:"WebUIUri,omitempty"`
	InstallDetails          ArgoCDInstallDetails `json:"InstallDetails,omitempty"`
	EnvironmentIDs          []string             `json:"EnvironmentIds,omitempty"`
	TenantIDs               []string             `json:"TenantIds,omitempty"`

	resources.Resource
}

func NewArgoCDGateway(spaceId string, name string, clientId string, healthStatus string, environmentIds []string) *ArgoCDGateway {
	return &ArgoCDGateway{
		SpaceID:        spaceId,
		Name:           name,
		ClientId:       clientId,
		HealthStatus:   healthStatus,
		EnvironmentIDs: environmentIds,
		TenantIDs:      []string{},
		Resource:       *resources.NewResource(),
	}
}

// MarshalJSON returns a deployment target as its JSON encoding.
func (a *ArgoCDGateway) MarshalJSON() ([]byte, error) {
	argoCDGateway := struct {
		SpaceID                 string               `json:"SpaceId,omitempty"`
		Name                    string               `json:"Name"`
		ClientId                string               `json:"ClientId"`
		HealthStatus            string               `json:"HealthStatus,omitempty"`
		HealthLastChecked       string               `json:"HealthLastChecked,omitempty"`
		HealthCheckServerTaskId string               `json:"HealthCheckServerTaskId,omitempty"`
		WebUIUri                string               `json:"WebUIUri,omitempty"`
		InstallDetails          ArgoCDInstallDetails `json:"InstallDetails,omitempty"`
		EnvironmentIDs          []string             `json:"EnvironmentIds,omitempty"`
		TenantIDs               []string             `json:"TenantIds,omitempty"`

		resources.Resource
	}{
		SpaceID:                 a.SpaceID,
		Name:                    a.Name,
		ClientId:                a.ClientId,
		HealthStatus:            a.HealthStatus,
		HealthLastChecked:       a.HealthLastChecked,
		HealthCheckServerTaskId: a.HealthCheckServerTaskId,
		WebUIUri:                a.WebUIUri,
		InstallDetails:          a.InstallDetails,
		EnvironmentIDs:          a.EnvironmentIDs,
		TenantIDs:               a.TenantIDs,
	}

	return json.Marshal(argoCDGateway)
}

// UnmarshalJSON sets this deployment target to its representation in JSON.
func (resource *ArgoCDGateway) UnmarshalJSON(b []byte) error {
	var argoCDGateway map[string]*json.RawMessage
	err := json.Unmarshal(b, &argoCDGateway)
	if err != nil {
		return err
	}

	var install ArgoCDInstallDetails
	err = json.Unmarshal(b, &install)
	if err != nil {
		return err
	}
	resource.InstallDetails = install

	for argoCDGatewayKey, argoCDGatewayValue := range argoCDGateway {
		switch argoCDGatewayKey {
		case "SpaceId":
			err = json.Unmarshal(*argoCDGatewayValue, &resource.SpaceID)
			if err != nil {
				return err
			}
		case "Name":
			err = json.Unmarshal(*argoCDGatewayValue, &resource.Name)
			if err != nil {
				return err
			}
		case "ClientId":
			err = json.Unmarshal(*argoCDGatewayValue, &resource.ClientId)
			if err != nil {
				return err
			}
		case "HealthStatus":
			err = json.Unmarshal(*argoCDGatewayValue, &resource.HealthStatus)
			if err != nil {
				return err
			}
		case "HealthLastChecked":
			err = json.Unmarshal(*argoCDGatewayValue, &resource.HealthLastChecked)
			if err != nil {
				return err
			}
		case "HealthCheckServerTaskId":
			err = json.Unmarshal(*argoCDGatewayValue, &resource.HealthCheckServerTaskId)
			if err != nil {
				return err
			}
		case "WebUIUri":
			err = json.Unmarshal(*argoCDGatewayValue, &resource.WebUIUri)
			if err != nil {
				return err
			}
		case "EnvironmentIds":
			err = json.Unmarshal(*argoCDGatewayValue, &resource.EnvironmentIDs)
			if err != nil {
				return err
			}
		case "TenantIds":
			err = json.Unmarshal(*argoCDGatewayValue, &resource.TenantIDs)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
