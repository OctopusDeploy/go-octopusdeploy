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
		persistenceSettings: persistenceSettings{Type: PersistenceSettingsTypeDatabase},
	}
}

// GetType returns the type for this persistence settings.
func (d *databasePersistenceSettings) GetType() PersistenceSettingsType {
	return d.Type
}

var _ DatabasePersistenceSettings = &databasePersistenceSettings{}
