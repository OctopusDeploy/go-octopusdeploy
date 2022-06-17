package core

type SkipTakeQuery struct {
	Skip int `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int `uri:"take,omitempty" url:"take,omitempty"`
}
