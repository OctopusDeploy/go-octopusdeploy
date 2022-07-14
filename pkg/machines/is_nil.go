package machines

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *AzureCloudServiceEndpoint:
		return v == nil
	case *AzureServiceFabricEndpoint:
		return v == nil
	case *AzureWebAppEndpoint:
		return v == nil
	case *CloudRegionEndpoint:
		return v == nil
	case *CloudServiceEndpoint:
		return v == nil
	case *DeploymentTarget:
		return v == nil
	case *EndpointResource:
		return v == nil
	case *KubernetesEndpoint:
		return v == nil
	case *ListeningTentacleEndpoint:
		return v == nil
	case *MachineConnectionStatus:
		return v == nil
	case *MachinePolicy:
		return v == nil
	case *PollingTentacleEndpoint:
		return v == nil
	case *SSHEndpoint:
		return v == nil
	case *Worker:
		return v == nil
	default:
		return v == nil
	}
}
