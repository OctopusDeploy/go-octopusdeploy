package actiontemplates

type ActionTemplateVersionedLogoQuery struct {
	CB       string `uri:"cb,omitempty" url:"cb,omitempty"`
	TypeOrID string `uri:"typeOrId,omitempty" url:"typeOrId,omitempty"`
	Version  string `uri:"version,omitempty" url:"version,omitempty"`
}
