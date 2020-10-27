package octopusdeploy

import "github.com/dghubble/sling"

type migrationService struct {
	migrationsImportPath        string
	migrationsPartialExportPath string

	service
}

func newMigrationService(sling *sling.Sling, uriTemplate string, migrationsImportPath string, migrationsPartialExportPath string) *migrationService {
	return &migrationService{
		migrationsImportPath:        migrationsImportPath,
		migrationsPartialExportPath: migrationsPartialExportPath,
		service:                     newService(ServiceMigrationService, sling, uriTemplate),
	}
}
