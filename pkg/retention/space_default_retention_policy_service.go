package retention

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type SpaceDefaultRetentionPolicyService struct {
	services.Service
}

const template = "/api/{spaceId}/retentionpolicies{/id}{?RetentionType}"

func Get(client newclient.Client, spaceDefaultRetentionPolicyQuery SpaceDefaultRetentionPolicyQuery) (*SpaceDefaultRetentionPolicy, error) {
	res, err := newclient.GetResourceByQuery[SpaceDefaultRetentionPolicy](client, template, spaceDefaultRetentionPolicyQuery)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Update(client newclient.Client, spaceDefaultRetentionPolicy ISpaceDefaultRetentionPolicy) (*SpaceDefaultRetentionPolicyResource, error) {
	res, err := newclient.Update[SpaceDefaultRetentionPolicyResource](client, template, spaceDefaultRetentionPolicy.GetSpaceID(), spaceDefaultRetentionPolicy.GetID(), spaceDefaultRetentionPolicy)
	if err != nil {
		return nil, err
	}
	return res, err
}
