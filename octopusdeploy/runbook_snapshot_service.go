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

// Add creates a new runbook snapshot.
func (s runbookSnapshotService) Add(runbookSnapshot *RunbookSnapshot) (*RunbookSnapshot, error) {
	if runbookSnapshot == nil {
		return nil, createInvalidParameterError(OperationAdd, ParameterRunbookSnapshot)
	}

	path, err := getAddPath(s, runbookSnapshot)
	if err != nil {
		return nil, err
	}

	response, err := apiAdd(s.getClient(), runbookSnapshot, new(RunbookSnapshot), path)
	if err != nil {
		return nil, err
	}

	return response.(*RunbookSnapshot), nil
}

// GetByID returns the release that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s runbookSnapshotService) GetByID(id string) (*RunbookSnapshot, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(RunbookSnapshot), path)
	if err != nil {
		return nil, err
	}

	return resp.(*RunbookSnapshot), nil
}
