package services

import (
	"github.com/dghubble/sling"
)

type nuGetService struct {
	service
}

func newNuGetService(sling *sling.Sling, uriTemplate string) *nuGetService {
	return &nuGetService{
		service: newService(ServiceNuGetService, sling, uriTemplate),
	}
}
