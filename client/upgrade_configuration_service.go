package client

import "github.com/dghubble/sling"

type upgradeConfigurationService struct {
	service
}

func newUpgradeConfigurationService(sling *sling.Sling, uriTemplate string) *upgradeConfigurationService {
	return &upgradeConfigurationService{
		service: newService(serviceUpgradeConfigurationService, sling, uriTemplate, nil),
	}
}
