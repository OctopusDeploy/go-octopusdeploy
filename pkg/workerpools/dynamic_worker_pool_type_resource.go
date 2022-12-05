package workerpools

import "time"

type DynamicWorkerPoolTypes struct {
	ID          string                   `json:"Id"`
	WorkerTypes []*DynamicWorkerPoolType `json:"WorkerTypes"`
}

type DynamicWorkerPoolType struct {
	ID                 string     `json:"Id"`
	Type               string     `json:"Type"`
	Description        string     `json:"Description"`
	DeprecationDateUtc *time.Time `json:"DeprecationDateUtc"`
	EndOfLifeDateUtc   *time.Time `json:"EndOfLifeDateUtc"`
	StartDateUtc       *time.Time `json:"StartDateUtc"`
}
