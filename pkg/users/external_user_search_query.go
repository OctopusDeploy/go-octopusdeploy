package users

type ExternalUserSearchQuery struct {
	PartialName string `uri:"partialName,omitempty" url:"partialName,omitempty"`
}
