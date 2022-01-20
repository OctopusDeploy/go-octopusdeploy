package octopusdeploy

import "time"

// IAccount defines the interface for accounts.
type IAccount interface {
	GetAccountType() AccountType
	GetDescription() string
	GetEnvironmentIDs() []string
	GetTenantedDeploymentMode() TenantedDeploymentMode
	GetTenantIDs() []string
	GetTenantTags() []string
	SetDescription(string)
	SetEnvironmentIDs([]string)
	SetTenantedDeploymentMode(TenantedDeploymentMode)
	SetTenantIDs([]string)
	SetTenantTags([]string)

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

// IEndpoint defines the interface for endpoints.
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

// IGitCredential defines the interface for Git-associated credentials.
type IGitCredential interface {
	GetType() GitCredentialType
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

// IPersistenceSettings defines the interface for persistence settings.
type IPersistenceSettings interface {
	GetType() string
}

// IResource defines the interface for resources.
type IResource interface {
	GetID() string
	GetModifiedBy() string
	GetModifiedOn() *time.Time
	GetLinks() map[string]string
	SetID(string)
	SetLinks(map[string]string)
	SetModifiedBy(string)
	SetModifiedOn(*time.Time)
	Validate() error
}

// IRunsOnAWorker defines the interface for workers.
type IRunsOnAWorker interface {
	GetDefaultWorkerPoolID() string
	SetDefaultWorkerPoolID(string)
}

// ISSHKeyAccount defines the interface for SSH key accounts.
type ISSHKeyAccount interface {
	SetPrivateKeyPassphrase(*SensitiveValue)

	IAccount
}

// ITriggerAction defines the interface for trigger actions.
type ITriggerAction interface {
	GetActionType() ActionType
	SetActionType(actionType ActionType)
}

// ITriggerFilter defines the interface for trigger filters.
type ITriggerFilter interface {
	GetFilterType() FilterType
	SetFilterType(filterType FilterType)
}

// IUsernamePasswordAccount defines the interface for username-password accounts.
type IUsernamePasswordAccount interface {
	GetUsername() string
	SetPassword(*SensitiveValue)
	SetUsername(string)

	IAccount
}

// IWorkerPool defines the interface for worker pools.
type IWorkerPool interface {
	GetWorkerPoolType() WorkerPoolType
	GetIsDefault() bool

	IHasName
	IResource
}
