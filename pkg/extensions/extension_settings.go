package extensions

type ExtensionSettings interface {
	ExtensionID() ExtensionID
	SetExtensionID(ExtensionID)
}

type ChangeControlExtensionSettings interface {
	IsChangeControlled() bool
	SetIsChangeControlled(bool)

	ExtensionSettings
}

type ConnectedChangeControlExtensionSettings interface {
	ConnectionID() string
	SetConnectionID(string)

	ChangeControlExtensionSettings
}
