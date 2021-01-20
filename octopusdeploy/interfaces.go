package octopusdeploy

import "time"

type IAccount interface {
	GetAccountType() AccountType
	GetDescription() string
	GetEnvironmentIDs() []string
	GetTenantedDeploymentMode() TenantedDeploymentMode
	GetTenantIDs() []string
	GetTenantTags() []string
	SetDescription(string)

	IHasName
	IHasSpace
	IResource
}

type IDeploymentTarget interface {
	GetEndpoint() IEndpoint
	GetName() string
	GetHealthStatus() string
	GetIsDisabled() bool
}

type IDynamicWorkerPool interface {
	GetWorkerType() WorkerType

	IWorkerPool
}

// IEndpoint defines the interface for all endpoints.
type IEndpoint interface {
	GetCommunicationStyle() string

	IResource
}

type IEndpointWithAccount interface {
	GetAccountID() string
}

type IEndpointWithFingerprint interface {
	GetFingerprint() string
}

type IEndpointWithHostname interface {
	GetHost() string
}

type IEndpointWithProxy interface {
	GetProxyID() string
	SetProxyID(string)
}

type IFeed interface {
	GetFeedType() FeedType
	GetName() string
	SetName(string)

	IResource
}

type IHasName interface {
	GetName() string
	SetName(string)
}

type IHasSpace interface {
	GetSpaceID() string
	SetSpaceID(string)
}

type IKubernetesAuthentication interface {
	GetAuthenticationType() string
}

// IResource defines the interface for all resources.
type IResource interface {
	GetID() string
	GetModifiedBy() string
	GetModifiedOn() *time.Time
	GetLinks() map[string]string
	Validate() error
}

// IRunsOnAWorker defines the interface for all workers.
type IRunsOnAWorker interface {
	GetDefaultWorkerPoolID() string
	SetDefaultWorkerPoolID(string)
}

type IWorkerPool interface {
	GetWorkerPoolType() WorkerPoolType
	GetIsDefault() bool

	IHasName
	IResource
}
