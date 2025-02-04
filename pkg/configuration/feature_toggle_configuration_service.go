package configuration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

type FeatureToggleConfigurationService struct{}

const template = "/api/configuration/feature-toggles"

func Get(client newclient.Client, request *FeatureToggleConfigurationQuery) (*FeatureToggleConfigurationResponse, error) {
	path, err := client.URITemplateCache().Intern(template)

	if err != nil {
		return nil, err
	}

	test, err := newclient.GetByRequest[FeatureToggleConfigurationResponse](client.HttpSession(), path.String(), request)

	test.ModifiedBy = path.String()

	return test, err
}
