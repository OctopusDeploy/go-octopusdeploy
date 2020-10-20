package octopusdeploy

import "github.com/dghubble/sling"

type proxyService struct {
	service
}

func newProxyService(sling *sling.Sling, uriTemplate string) *proxyService {
	proxyService := &proxyService{}
	proxyService.service = newService(serviceProxyService, sling, uriTemplate, nil)

	return proxyService
}
