package resources

type IHasName interface {
	GetName() string
	SetName(string)
}

type IHasSpace interface {
	GetSpaceID() string
	SetSpaceID(string)
}
