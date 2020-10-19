package client

import "github.com/dghubble/sling"

type runbookSnapshotService struct {
	service
}

func newRunbookSnapshotService(sling *sling.Sling, uriTemplate string) *runbookSnapshotService {
	runbookSnapshotService := &runbookSnapshotService{}
	runbookSnapshotService.service = newService(serviceRunbookSnapshotService, sling, uriTemplate, nil)

	return runbookSnapshotService
}
