package model

import "time"

type ResourceInterface interface {
	GetID() string
	GetLastModifiedBy() string
	GetLastModifiedOn() *time.Time
	GetLinks() map[string]string
	Validate() error
}

type Resource struct {
	ID             string            `json:"Id,omitempty"`
	LastModifiedBy string            `json:"LastModifiedBy,omitempty"`
	LastModifiedOn *time.Time        `json:"LastModifiedOn,omitempty"`
	Links          map[string]string `json:"Links,omitempty"`
}