package octopusdeploy

// UserAuthentication represents enabled authentication providers and whether
// the current user can edit logins for the given user.
type UserAuthentication struct {
	AuthenticationProviders             []AuthenticationProviderElement
	CanCurrentUserEditIdentitiesForUser bool
	Links                               map[string]string
}
