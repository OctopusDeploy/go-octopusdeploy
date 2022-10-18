package projects

// DatabasePersistenceSettings represents database persistence settings associated with a project.
type DatabasePersistenceSettings struct {
	persistenceSettings
}

// NewDatabasePersistenceSettings creates an instance of database persistence settings.
func NewDatabasePersistenceSettings() *DatabasePersistenceSettings {
	return &DatabasePersistenceSettings{
		persistenceSettings: persistenceSettings{Type: PersistenceSettingsTypeDatabase},
	}
}

// GetType returns the type for this persistence settings.
func (d *DatabasePersistenceSettings) GetType() PersistenceSettingsType {
	return d.Type
}

var _ IPersistenceSettings = &DatabasePersistenceSettings{}
