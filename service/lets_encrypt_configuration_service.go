package service

import (
	"github.com/dghubble/sling"
)

type letsEncryptConfigurationService struct {
	service
}

func newLetsEncryptConfigurationService(sling *sling.Sling, uriTemplate string) *letsEncryptConfigurationService {
	return &letsEncryptConfigurationService{
		service: newService(ServiceLetsEncryptConfigurationService, sling, uriTemplate),
	}
}
