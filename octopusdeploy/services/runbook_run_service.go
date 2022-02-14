package services

import (
	"github.com/dghubble/sling"
)

type runbookRunService struct {
	canDeleteService
}

func newRunbookRunService(sling *sling.Sling, uriTemplate string) *runbookRunService {
	runbookRunService := &runbookRunService{}
	runbookRunService.service = newService(ServiceRunbookRunService, sling, uriTemplate)

	return runbookRunService
}
