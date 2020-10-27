package octopusdeploy

import "github.com/dghubble/sling"

type runbookSnapshotService struct {
	canDeleteService
}

func newRunbookSnapshotService(sling *sling.Sling, uriTemplate string) *runbookSnapshotService {
	runbookSnapshotService := &runbookSnapshotService{}
	runbookSnapshotService.service = newService(ServiceRunbookSnapshotService, sling, uriTemplate)

	return runbookSnapshotService
}
