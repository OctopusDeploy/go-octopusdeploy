package extensions

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"

type ExtensionSettings struct {
	ExtensionID extensions.ExtensionID `json:"ExtensionId,omitempty"`
}

// NewExtensionSettings creates a settings structure for an extension.
func NewExtensionSettings(extensionID extensions.ExtensionID) ExtensionSettings {
	return ExtensionSettings{
		ExtensionID: extensionID,
	}
}
