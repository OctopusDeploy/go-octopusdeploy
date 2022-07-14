package machines

import "github.com/OctopusDeploy/go-octopusdeploy/internal"

func ToEndpoint(endpointResource *EndpointResource) (IEndpoint, error) {
	if IsNil(endpointResource) {
		return nil, internal.CreateInvalidParameterError("ToEndpoint", "endpointResource")
	}

	if err := endpointResource.Validate(); err != nil {
		return nil, err
	}

	var endpoint IEndpoint

	switch endpointResource.GetCommunicationStyle() {
	case "AzureCloudService":
		azureCloudServiceEndpoint := NewAzureCloudServiceEndpoint()
		azureCloudServiceEndpoint.AccountID = endpointResource.AccountID
		azureCloudServiceEndpoint.CloudServiceName = endpointResource.CloudServiceName
		azureCloudServiceEndpoint.DefaultWorkerPoolID = endpointResource.DefaultWorkerPoolID
		azureCloudServiceEndpoint.Slot = endpointResource.Slot
		azureCloudServiceEndpoint.StorageAccountName = endpointResource.StorageAccountName
		azureCloudServiceEndpoint.SwapIfPossible = endpointResource.SwapIfPossible
		azureCloudServiceEndpoint.UseCurrentInstanceCount = endpointResource.UseCurrentInstanceCount
		endpoint = azureCloudServiceEndpoint
	case "AzureServiceFabricCluster":
		azureServiceFabricEndpoint := NewAzureServiceFabricEndpoint()
		azureServiceFabricEndpoint.AadClientCredentialSecret = endpointResource.AadClientCredentialSecret
		azureServiceFabricEndpoint.AadCredentialType = endpointResource.AadCredentialType
		azureServiceFabricEndpoint.AadUserCredentialPassword = endpointResource.AadUserCredentialPassword
		azureServiceFabricEndpoint.AadUserCredentialUsername = endpointResource.AadUserCredentialUsername
		azureServiceFabricEndpoint.CertificateStoreLocation = endpointResource.CertificateStoreLocation
		azureServiceFabricEndpoint.CertificateStoreName = endpointResource.CertificateStoreName
		azureServiceFabricEndpoint.ClientCertificateVariable = endpointResource.ClientCertificateVariable
		azureServiceFabricEndpoint.ConnectionEndpoint = endpointResource.ConnectionEndpoint
		azureServiceFabricEndpoint.SecurityMode = endpointResource.SecurityMode
		azureServiceFabricEndpoint.ServerCertificateThumbprint = endpointResource.ServerCertificateThumbprint
		endpoint = azureServiceFabricEndpoint
	case "AzureWebApp":
		azureWebAppEndpoint := NewAzureWebAppEndpoint()
		azureWebAppEndpoint.AccountID = endpointResource.AccountID
		azureWebAppEndpoint.ResourceGroupName = endpointResource.ResourceGroupName
		azureWebAppEndpoint.WebAppName = endpointResource.WebAppName
		azureWebAppEndpoint.WebAppSlotName = endpointResource.WebAppSlotName
		endpoint = azureWebAppEndpoint
	case "Kubernetes":
		kubernetesEndpoint := NewKubernetesEndpoint(endpointResource.ClusterURL)
		kubernetesEndpoint.Authentication = endpointResource.Authentication
		kubernetesEndpoint.ClusterCertificate = endpointResource.ClusterCertificate
		kubernetesEndpoint.ClusterCertificatePath = endpointResource.ClusterCertificatePath
		kubernetesEndpoint.Container = endpointResource.Container
		kubernetesEndpoint.DefaultWorkerPoolID = endpointResource.DefaultWorkerPoolID
		kubernetesEndpoint.Namespace = endpointResource.Namespace
		kubernetesEndpoint.ProxyID = endpointResource.ProxyID
		kubernetesEndpoint.RunningInContainer = endpointResource.RunningInContainer
		kubernetesEndpoint.SkipTLSVerification = endpointResource.SkipTLSVerification
		endpoint = kubernetesEndpoint
	case "None":
		cloudRegionEndpoint := NewCloudRegionEndpoint()
		cloudRegionEndpoint.DefaultWorkerPoolID = endpointResource.DefaultWorkerPoolID
		endpoint = cloudRegionEndpoint
	case "OfflineDrop":
		offlinePackageDropEndpoint := NewOfflinePackageDropEndpoint()
		offlinePackageDropEndpoint.ApplicationsDirectory = endpointResource.ApplicationsDirectory
		offlinePackageDropEndpoint.Destination = endpointResource.Destination
		offlinePackageDropEndpoint.SensitiveVariablesEncryptionPassword = endpointResource.SensitiveVariablesEncryptionPassword
		offlinePackageDropEndpoint.WorkingDirectory = endpointResource.WorkingDirectory
		endpoint = offlinePackageDropEndpoint
	case "Ssh":
		sshEndpoint := NewSSHEndpoint(endpointResource.Host, endpointResource.Port, endpointResource.Fingerprint)
		endpoint = sshEndpoint
	case "TentacleActive":
		pollingTentacleEndpoint := NewPollingTentacleEndpoint(endpointResource.URI, endpointResource.Thumbprint)
		endpoint = pollingTentacleEndpoint
	case "TentaclePassive":
		listeningTentacleEndpoint := NewListeningTentacleEndpoint(endpointResource.URI, endpointResource.Thumbprint)
		endpoint = listeningTentacleEndpoint
	}

	endpoint.SetLinks(endpointResource.GetLinks())
	endpoint.SetModifiedBy(endpointResource.GetModifiedBy())
	endpoint.SetModifiedOn(endpointResource.GetModifiedOn())
	endpoint.SetID(endpointResource.GetID())

	return endpoint, nil
}

func ToEndpointResource(endpoint IEndpoint) (*EndpointResource, error) {
	if IsNil(endpoint) {
		return nil, internal.CreateInvalidParameterError("ToEndpointResource", "endpoint")
	}

	// conversion unnecessary if input endpoint is *EndpointResource
	if v, ok := endpoint.(*EndpointResource); ok {
		return v, nil
	}

	endpointResource := NewEndpointResource(endpoint.GetCommunicationStyle())

	switch endpointResource.GetCommunicationStyle() {
	case "AzureCloudService":
		azureCloudServiceEndpoint := endpoint.(*AzureCloudServiceEndpoint)
		endpointResource.AccountID = azureCloudServiceEndpoint.AccountID
		endpointResource.CloudServiceName = azureCloudServiceEndpoint.CloudServiceName
		endpointResource.DefaultWorkerPoolID = azureCloudServiceEndpoint.DefaultWorkerPoolID
		endpointResource.Slot = azureCloudServiceEndpoint.Slot
		endpointResource.StorageAccountName = azureCloudServiceEndpoint.StorageAccountName
		endpointResource.SwapIfPossible = azureCloudServiceEndpoint.SwapIfPossible
		endpointResource.UseCurrentInstanceCount = azureCloudServiceEndpoint.UseCurrentInstanceCount
	case "AzureServiceFabricCluster":
		azureServiceFabricEndpoint := endpoint.(*AzureServiceFabricEndpoint)
		endpointResource.AadClientCredentialSecret = azureServiceFabricEndpoint.AadClientCredentialSecret
		endpointResource.AadCredentialType = azureServiceFabricEndpoint.AadCredentialType
		endpointResource.AadUserCredentialPassword = azureServiceFabricEndpoint.AadUserCredentialPassword
		endpointResource.AadUserCredentialUsername = azureServiceFabricEndpoint.AadUserCredentialUsername
		endpointResource.CertificateStoreLocation = azureServiceFabricEndpoint.CertificateStoreLocation
		endpointResource.CertificateStoreName = azureServiceFabricEndpoint.CertificateStoreName
		endpointResource.ClientCertificateVariable = azureServiceFabricEndpoint.ClientCertificateVariable
		endpointResource.ConnectionEndpoint = azureServiceFabricEndpoint.ConnectionEndpoint
		endpointResource.SecurityMode = azureServiceFabricEndpoint.SecurityMode
		endpointResource.ServerCertificateThumbprint = azureServiceFabricEndpoint.ServerCertificateThumbprint
	case "AzureWebApp":
		azureWebApp := endpoint.(*AzureWebAppEndpoint)
		endpointResource.AccountID = azureWebApp.AccountID
		endpointResource.ResourceGroupName = azureWebApp.ResourceGroupName
		endpointResource.WebAppName = azureWebApp.WebAppName
		endpointResource.WebAppSlotName = azureWebApp.WebAppSlotName
	case "Kubernetes":
		kubernetesEndpoint := endpoint.(*KubernetesEndpoint)
		endpointResource.Authentication = kubernetesEndpoint.Authentication
		endpointResource.ClusterCertificate = kubernetesEndpoint.ClusterCertificate
		endpointResource.ClusterCertificatePath = kubernetesEndpoint.ClusterCertificatePath
		endpointResource.ClusterURL = kubernetesEndpoint.ClusterURL
		endpointResource.Container = kubernetesEndpoint.Container
		endpointResource.DefaultWorkerPoolID = kubernetesEndpoint.DefaultWorkerPoolID
		endpointResource.Namespace = kubernetesEndpoint.Namespace
		endpointResource.ProxyID = kubernetesEndpoint.ProxyID
		endpointResource.RunningInContainer = kubernetesEndpoint.RunningInContainer
		endpointResource.SkipTLSVerification = kubernetesEndpoint.SkipTLSVerification
	case "None":
		cloudRegionEndpoint := endpoint.(*CloudRegionEndpoint)
		endpointResource.DefaultWorkerPoolID = cloudRegionEndpoint.DefaultWorkerPoolID
	case "OfflineDrop":
		offlinePackageDropEndpoint := endpoint.(*OfflinePackageDropEndpoint)
		endpointResource.ApplicationsDirectory = offlinePackageDropEndpoint.ApplicationsDirectory
		endpointResource.Destination = offlinePackageDropEndpoint.Destination
		endpointResource.SensitiveVariablesEncryptionPassword = offlinePackageDropEndpoint.SensitiveVariablesEncryptionPassword
		endpointResource.WorkingDirectory = offlinePackageDropEndpoint.WorkingDirectory
	case "Ssh":
		sshEndpoint := endpoint.(*SSHEndpoint)
		endpointResource.AccountID = sshEndpoint.AccountID
		endpointResource.DotNetCorePlatform = sshEndpoint.DotNetCorePlatform
		endpointResource.Fingerprint = sshEndpoint.Fingerprint
		endpointResource.Host = sshEndpoint.Host
		endpointResource.ProxyID = sshEndpoint.ProxyID
		endpointResource.Port = sshEndpoint.Port
		endpointResource.URI = sshEndpoint.URI
	case "TentacleActive":
		pollingTentacleEndpoint := endpoint.(*PollingTentacleEndpoint)
		endpointResource.CertificateSignatureAlgorithm = pollingTentacleEndpoint.CertificateSignatureAlgorithm
		endpointResource.TentacleVersionDetails = pollingTentacleEndpoint.TentacleVersionDetails
		endpointResource.Thumbprint = pollingTentacleEndpoint.Thumbprint
		endpointResource.URI = pollingTentacleEndpoint.URI
	case "TentaclePassive":
		listeningTentacleEndpoint := endpoint.(*ListeningTentacleEndpoint)
		endpointResource.CertificateSignatureAlgorithm = listeningTentacleEndpoint.CertificateSignatureAlgorithm
		endpointResource.ProxyID = listeningTentacleEndpoint.ProxyID
		endpointResource.TentacleVersionDetails = listeningTentacleEndpoint.TentacleVersionDetails
		endpointResource.Thumbprint = listeningTentacleEndpoint.Thumbprint
		endpointResource.URI = listeningTentacleEndpoint.URI
	}

	endpointResource.SetID(endpoint.GetID())
	endpointResource.SetLinks(endpoint.GetLinks())
	endpointResource.SetModifiedBy(endpoint.GetModifiedBy())
	endpointResource.SetModifiedOn(endpoint.GetModifiedOn())

	return endpointResource, nil
}

func ToEndpointArray(endpointResources []*EndpointResource) []IEndpoint {
	items := []IEndpoint{}
	for _, endpointResource := range endpointResources {
		endpoint, err := ToEndpoint(endpointResource)
		if err != nil {
			return nil
		}
		items = append(items, endpoint)
	}
	return items
}
