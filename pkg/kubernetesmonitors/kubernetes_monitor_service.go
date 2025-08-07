package kubernetesmonitors

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

// KubernetesMonitorService handles communication with Kubernetes Monitor-related methods of the Octopus API.
type KubernetesMonitorService struct {
	services.CanDeleteService
}

// NewKubernetesMonitorService returns a KubernetesMonitorService with a preconfigured client.
func NewKubernetesMonitorService(sling *sling.Sling, uriTemplate string) *KubernetesMonitorService {
	return &KubernetesMonitorService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceKubernetesMonitorService, sling, uriTemplate),
		},
	}
}

// Register registers a Kubernetes monitor with the Octopus Deploy server.
func (s *KubernetesMonitorService) Register(command *RegisterKubernetesMonitorCommand) (*RegisterKubernetesMonitorResponse, error) {
	if command == nil {
		return nil, internal.CreateInvalidParameterError("register", "command")
	}

	path, err := services.GetPath(s)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiPost(s.GetClient(), command, new(RegisterKubernetesMonitorResponse), path)
	if err != nil {
		return nil, err
	}

	return resp.(*RegisterKubernetesMonitorResponse), nil
}

// GetByID returns the Kubernetes monitor that matches the input ID. If one cannot be found, it returns nil and an error.
func (s *KubernetesMonitorService) GetByID(id string) (*KubernetesMonitor, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(KubernetesMonitor), path)
	if err != nil {
		return nil, err
	}

	return resp.(*KubernetesMonitor), nil
}
