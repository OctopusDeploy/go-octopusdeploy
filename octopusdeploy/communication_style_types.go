package octopusdeploy

type CommunicationStyle string

const (
	CommunicationStyleAzureCloudService         CommunicationStyle = "AzureCloudService"
	CommunicationStyleAzureServiceFabricCluster CommunicationStyle = "AzureServiceFabricCluster"
	CommunicationStyleFtp                       CommunicationStyle = "Ftp"
	CommunicationStyleKubernetes                CommunicationStyle = "Kubernetes"
	CommunicationStyleNone                      CommunicationStyle = "None"
	CommunicationStyleOfflineDrop               CommunicationStyle = "OfflineDrop"
	CommunicationStyleSSH                       CommunicationStyle = "Ssh"
	CommunicationStyleTentacleActive            CommunicationStyle = "TentacleActive"
	CommunicationStyleTentaclePassive           CommunicationStyle = "TentaclePassive"
)
