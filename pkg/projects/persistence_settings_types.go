package projects

type PersistenceSettingsType string

const (
	PersistenceSettingsTypeDatabase = PersistenceSettingsType("Database")
	PersistenceSettingsTypeVersionControlled = PersistenceSettingsType("VersionControlled")
)
