package machines

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"

const (
	gatewayTemplate = "/api/{spaceId}/argocdgateways{/id}{?skip,take,ids,partialName}"
)

//func GetArgoSummariesController(client newclient.Client, spaceId string, grpcClientId string) {
//	res, err := newclient.GetAll[ArgoCDGateway](client, "/api/{spaceId}/argocdgateways", spaceId)
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}

func GetGateway(client newclient.Client, spaceID string, id string) (*ArgoCDGateway, error) {
	res, err := newclient.GetByID[ArgoCDGateway](client, template, spaceID, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TODO(tmm): do we need to move to a different package so we don't run into name conflicts?
// DeleteByID will delete a gateway with the provided id.
func DeleteGatewayByID(client newclient.Client, spaceID string, id string) error {
	return newclient.DeleteByID(client, gatewayTemplate, spaceID, id)
}

func AddGateway(client newclient.Client, gateway ArgoCDGateway) (*ArgoCDGateway, error) {
	res, err := newclient.Add[ArgoCDGateway](client, gatewayTemplate, gateway.SpaceID, gateway)
	if err != nil {
		return nil, err
	}
	return res, nil
}
