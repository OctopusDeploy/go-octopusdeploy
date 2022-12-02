package extensions

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"

type ChangeControlExtensionSettings struct {
	IsChangeControlled bool

	ExtensionSettings
}

// NewChangeControlExtensionSettings creates a settings structure a for change-controlled extension.
func NewChangeControlExtensionSettings(extensionID extensions.ExtensionID, isChangeControlled bool) ChangeControlExtensionSettings {
	return ChangeControlExtensionSettings{
		IsChangeControlled: isChangeControlled,
		ExtensionSettings:  NewExtensionSettings(extensionID),
	}
}
