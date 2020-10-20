package octopusdeploy

import "github.com/dghubble/sling"

type scopedUserRoleService struct {
	service
}

func newScopedUserRoleService(sling *sling.Sling, uriTemplate string) *scopedUserRoleService {
	scopedUserRoleService := &scopedUserRoleService{}
	scopedUserRoleService.service = newService(serviceScopedUserRoleService, sling, uriTemplate, nil)

	return scopedUserRoleService
}
