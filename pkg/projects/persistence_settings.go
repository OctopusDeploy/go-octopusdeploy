package projects

// PersistenceSettings defines the interface for persistence settings.
type PersistenceSettings interface {
	Type() PersistenceSettingsType
}

// persistenceSettings represents persistence settings associated with a project.
type persistenceSettings struct {
	settingsType PersistenceSettingsType `json:"Type"`
}
