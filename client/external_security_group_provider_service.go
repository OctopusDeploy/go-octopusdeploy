package client

import "github.com/dghubble/sling"

type externalSecurityGroupProviderService struct {
	service
}

func newExternalSecurityGroupProviderService(sling *sling.Sling, uriTemplate string) *externalSecurityGroupProviderService {
	return &externalSecurityGroupProviderService{
		service: newService(serviceExternalSecurityGroupProviderService, sling, uriTemplate, nil),
	}
}
