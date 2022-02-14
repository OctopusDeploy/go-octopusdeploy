package octopusdeploy

type ISpaceScopedResource interface {
	getSpaceID() string
}

type SpaceScopedResource struct {
	SpaceID string `json:"SpaceId,omitempty"`
}

func (r SpaceScopedResource) getSpaceID() string {
	return r.SpaceID
}
