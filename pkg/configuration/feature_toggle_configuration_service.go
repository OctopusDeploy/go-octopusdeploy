package configuration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type FeatureToggleConfigurationService struct{}

const template = "/api/configuration/feature-toggles{?Name}"

func Get(client newclient.Client, query *FeatureToggleConfigurationQuery) (*resources.Resources[*ConfiguredFeatureToggle], error) {
	return newclient.GetByQueryWithoutSpace[ConfiguredFeatureToggle](client, template, query)
}
