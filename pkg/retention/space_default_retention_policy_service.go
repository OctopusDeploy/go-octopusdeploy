package retention

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type SpaceDefaultRetentionPolicyService struct {
	services.Service
}

type SpaceDefaultRetentionPolicyQuery struct {
	RetentionType string `url:"type"`
	SpaceID       string `url:"spaceId"`
}

const template = "/api/{spaceId}/retentionpolicies{?type}}"

func Get(client newclient.Client, spaceDefaultRetentionPolicyQuery SpaceDefaultRetentionPolicyQuery) (*SpaceDefaultRetentionPolicy, error) {
	res, err := newclient.GetResourceByQuery[SpaceDefaultRetentionPolicy](client, template, spaceDefaultRetentionPolicyQuery)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Update(client newclient.Client, spaceDefaultRetentionPolicy SpaceDefaultRetentionPolicy) (*SpaceDefaultRetentionPolicy, error) {
	res, err := newclient.Update[SpaceDefaultRetentionPolicy](client, template, spaceDefaultRetentionPolicy.GetSpaceID(), spaceDefaultRetentionPolicy.ID, spaceDefaultRetentionPolicy)
	if err != nil {
		return nil, err
	}
	return res, err
}
