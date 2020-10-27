package octopusdeploy

import "github.com/dghubble/sling"

type proxyService struct {
	canDeleteService
}

func newProxyService(sling *sling.Sling, uriTemplate string) *proxyService {
	proxyService := &proxyService{}
	proxyService.service = newService(ServiceProxyService, sling, uriTemplate)

	return proxyService
}
