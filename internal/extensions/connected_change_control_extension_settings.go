package extensions

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"

type ConnectedChangeControlExtensionSettings struct {
	ConnectionID string

	ChangeControlExtensionSettings
}

// NewConnectedChangeControlExtensionSettings creates a settings structure a for connected and change-controlled extension.
func NewConnectedChangeControlExtensionSettings(extensionID extensions.ExtensionID, isChangeControlled bool, connectionID string) ConnectedChangeControlExtensionSettings {
	return ConnectedChangeControlExtensionSettings{
		ConnectionID:                   connectionID,
		ChangeControlExtensionSettings: NewChangeControlExtensionSettings(extensionID, isChangeControlled),
	}
}
