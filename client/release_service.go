package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type releaseService struct {
	service
}

func newReleaseService(sling *sling.Sling, uriTemplate string) *releaseService {
	releaseService := &releaseService{}
	releaseService.service = newService(serviceReleaseService, sling, uriTemplate, new(model.Release))

	return releaseService
}
