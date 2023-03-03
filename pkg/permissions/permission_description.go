package permissions

type PermissionDescription struct {
	CanApplyAtSpaceLevel  *bool    `json:"CanApplyAtSpaceLevel"`
	CanApplyAtSystemLevel *bool    `json:"CanApplyAtSystemLevel"`
	Description           string   `json:"Description,omitempty"`
	SupportedRestrictions []string `json:"SupportedRestrictions"`
}
