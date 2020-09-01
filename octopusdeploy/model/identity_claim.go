package model

type IdentityClaim struct {
	IsIdentifyingClaim bool   `json:"IsIdentifyingClaim,omitempty"`
	Value              string `json:"Value,omitempty"`
}
