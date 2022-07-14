package users

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/authentication"

// UserAuthentication represents enabled authentication providers and whether
// the current user can edit logins for the given user.
type UserAuthentication struct {
	AuthenticationProviders             []authentication.AuthenticationProviderElement
	CanCurrentUserEditIdentitiesForUser bool
	Links                               map[string]string
}
