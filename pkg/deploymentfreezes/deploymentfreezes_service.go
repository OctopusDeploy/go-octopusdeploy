package deploymentfreezes

import (
	"math"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const template = "/api/deploymentfreezes{/id}{?skip,take,ids,projectIds,tenantIds,environmentIds,includeComplete,status}"

type DeploymentFreezeService struct {
}

func Get(client newclient.Client, deploymentFreezesQuery *DeploymentFreezeQuery) (*DeploymentFreezes, error) {
	path, err := client.URITemplateCache().Expand(template, deploymentFreezesQuery)
	if err != nil {
		return nil, err
	}

	res, err := newclient.Get[DeploymentFreezes](client.HttpSession(), path)
	if err != nil {
		return &DeploymentFreezes{}, err
	}

	return res, nil
}

func GetAll(client newclient.Client) ([]*DeploymentFreeze, error) {
	path, err := client.URITemplateCache().Expand(template, &DeploymentFreezeQuery{Skip: 0, Take: math.MaxInt32})
	if err != nil {
		return nil, err
	}

	res, err := newclient.Get[DeploymentFreezes](client.HttpSession(), path)

	freezes := make([]*DeploymentFreeze, 0)
	for _, freeze := range res.Items {
		freezes = append(freezes, &freeze)
	}

	return freezes, nil
}

func Add(client newclient.Client, deploymentFreeze *DeploymentFreeze) (*DeploymentFreeze, error) {
	if deploymentFreeze == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("deploymentFreeze")
	}

	path, err := client.URITemplateCache().Expand(template, deploymentFreeze)
	if err != nil {
		return nil, err
	}

	res, err := newclient.Post[DeploymentFreeze](client.HttpSession(), path, deploymentFreeze)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Update(client newclient.Client, deploymentFreeze *DeploymentFreeze) (*DeploymentFreeze, error) {
	if deploymentFreeze == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("deploymentFreeze")
	}

	path, err := client.URITemplateCache().Expand(template, deploymentFreeze)
	if err != nil {
		return nil, err
	}

	res, err := newclient.Put[DeploymentFreeze](client.HttpSession(), path, deploymentFreeze)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Delete(client newclient.Client, deploymentFreeze *DeploymentFreeze) error {
	if deploymentFreeze == nil {
		return internal.CreateRequiredParameterIsEmptyOrNilError("deploymentFreeze")
	}

	path, err := client.URITemplateCache().Expand(template, deploymentFreeze)
	if err != nil {
		return err
	}

	err = newclient.Delete(client.HttpSession(), path)
	if err != nil {
		return err
	}
	return nil
}
