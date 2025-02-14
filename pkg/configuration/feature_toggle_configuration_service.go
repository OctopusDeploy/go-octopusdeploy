package configuration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

type FeatureToggleConfigurationService struct{}

const template = "/api/configuration/feature-toggles{?Name}"

func Get(client newclient.Client, query *FeatureToggleConfigurationQuery) (*FeatureToggleConfigurationResponse, error) {
	return newclient.GetByQueryWithoutSpace[FeatureToggleConfigurationResponse](client, template, query)
}
