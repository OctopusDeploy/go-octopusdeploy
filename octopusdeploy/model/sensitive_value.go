package model

type SensitiveValue struct {
	HasValue bool    `json:"HasValue"`
	NewValue *string `json:"NewValue"`
}
