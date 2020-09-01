package model

import "time"

type APIKey struct {
	APIKey  string    `json:"ApiKey,omitempty"`
	Created time.Time `json:"Created,omitempty"`
	Purpose string    `json:"Purpose,omitempty"`
	UserID  string    `json:"UserId,omitempty"`
	Resource
}
