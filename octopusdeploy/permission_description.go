package octopusdeploy

type PermissionDescription struct {
	CanApplyAtSpaceLevel  *bool    `json:"CanApplyAtSpaceLevel,omitempty"`
	CanApplyAtSystemLevel *bool    `json:"CanApplyAtSystemLevel,omitempty"`
	Description           string   `json:"Description,omitempty"`
	SupportedRestrictions []string `json:"SupportedRestrictions"`
}
