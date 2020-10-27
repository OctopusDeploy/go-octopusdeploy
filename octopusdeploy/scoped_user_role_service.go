package octopusdeploy

import "github.com/dghubble/sling"

type scopedUserRoleService struct {
	canDeleteService
}

func newScopedUserRoleService(sling *sling.Sling, uriTemplate string) *scopedUserRoleService {
	scopedUserRoleService := &scopedUserRoleService{}
	scopedUserRoleService.service = newService(ServiceScopedUserRoleService, sling, uriTemplate)

	return scopedUserRoleService
}
