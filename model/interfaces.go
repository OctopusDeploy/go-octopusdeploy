package model

import "time"

type IAccount interface {
	GetAccountType() string
	GetDescription() string
	SetDescription(string)

	IHasName
	IResource
}

type IDynamicWorkerPool interface {
	GetWorkerType() string

	IWorkerPool
}

// IEndpoint defines the interface for all endpoints.
type IEndpoint interface {
	GetCommunicationStyle() string
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
	GetFeedType() string
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

// IResource defines the interface for all resources.
type IResource interface {
	GetID() string
	GetLastModifiedBy() string
	GetLastModifiedOn() *time.Time
	GetLinks() map[string]string
	SetID(id string)
	SetLastModifiedBy(name string)
	SetLastModifiedOn(*time.Time)
	Validate() error
}

// IRunsOnAWorker defines the interface for all workers.
type IRunsOnAWorker interface {
	GetDefaultWorkerPoolID() string
	SetDefaultWorkerPoolID(string)
}

type IWorkerPool interface {
	GetWorkerPoolType() string

	IHasName
	IResource
}
