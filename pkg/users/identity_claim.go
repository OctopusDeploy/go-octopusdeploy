package users

type IdentityClaim struct {
	IsIdentifyingClaim bool   `json:"IsIdentifyingClaim"`
	Value              string `json:"Value,omitempty"`
}
