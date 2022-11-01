package projects

// DatabasePersistenceSettings represents database persistence settings associated with a project.
type DatabasePersistenceSettings interface {
	PersistenceSettings
}

type databasePersistenceSettings struct {
	persistenceSettings
}

// NewDatabasePersistenceSettings creates an instance of database persistence settings.
func NewDatabasePersistenceSettings() DatabasePersistenceSettings {
	return &databasePersistenceSettings{
		persistenceSettings: persistenceSettings{settingsType: PersistenceSettingsTypeDatabase},
	}
}

// Type returns the type for this persistence settings.
func (d databasePersistenceSettings) Type() PersistenceSettingsType {
	return d.settingsType
}

var _ DatabasePersistenceSettings = &databasePersistenceSettings{}
