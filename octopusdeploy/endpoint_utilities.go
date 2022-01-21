package octopusdeploy

import (
	"fmt"

	"github.com/jinzhu/copier"
)

func ToEndpoint(endpointResource *EndpointResource) (IEndpoint, error) {
	if isNil(endpointResource) {
		return nil, createInvalidParameterError("ToEndpoint", "endpointResource")
	}

	var endpoint IEndpoint
	var err error
	switch endpointResource.CommunicationStyle {
	case "AzureCloudService":
		endpoint = NewAzureCloudServiceEndpoint()
	case "AzureServiceFabricCluster":
		endpoint = NewAzureServiceFabricEndpoint()
	case "AzureWebApp":
		endpoint = NewAzureWebAppEndpoint()
	case "Kubernetes":
		endpoint = NewKubernetesEndpoint(endpointResource.ClusterURL)
	case "None":
		endpoint = NewCloudRegionEndpoint()
	case "OfflineDrop":
		endpoint = NewOfflinePackageDropEndpoint()
	case "Ssh":
		endpoint = NewSSHEndpoint(endpointResource.Host, endpointResource.Port, endpointResource.Fingerprint)
	case "StepPackage":
		switch endpointResource.DeploymentTargetTypeId {
		case "aws-ecs-target":
			endpoint = NewAmazonECSEndpoint()
		default:
			return nil, fmt.Errorf("StepPackage deployment target with target type id '%s' is not supported", endpointResource.DeploymentTargetTypeId)
		}
	case "TentacleActive":
		endpoint = NewPollingTentacleEndpoint(endpointResource.URI, endpointResource.Thumbprint)
	case "TentaclePassive":
		endpoint = NewListeningTentacleEndpoint(endpointResource.URI, endpointResource.Thumbprint)
	}

	err = copier.Copy(endpoint, endpointResource)
	if err != nil {
		return nil, err
	}

	return endpoint, nil
}

func ToEndpointResource(endpoint IEndpoint) (*EndpointResource, error) {
	if isNil(endpoint) {
		return nil, createInvalidParameterError("ToEndpointResource", "endpoint")
	}

	endpointResource := NewEndpointResource(endpoint.GetCommunicationStyle())

	err := copier.Copy(&endpointResource, endpoint)
	if err != nil {
		return nil, err
	}

	return endpointResource, nil
}

func ToEndpointArray(endpointResources []*EndpointResource) []IEndpoint {
	var items []IEndpoint
	for _, endpointResource := range endpointResources {
		endpoint, err := ToEndpoint(endpointResource)
		if err != nil {
			return nil
		}
		items = append(items, endpoint)
	}
	return items
}
