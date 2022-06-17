package actiontemplates

type ActionTemplateLogoQuery struct {
	CB       string `uri:"cb,omitempty" url:"cb,omitempty"`
	TypeOrID string `uri:"typeOrId,omitempty" url:"typeOrId,omitempty"`
}
