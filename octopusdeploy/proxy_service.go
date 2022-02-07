package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type proxyService struct {
	services.canDeleteService
}

func newProxyService(sling *sling.Sling, uriTemplate string) *proxyService {
	proxyService := &proxyService{}
	proxyService.service = services.newService(ServiceProxyService, sling, uriTemplate)

	return proxyService
}
