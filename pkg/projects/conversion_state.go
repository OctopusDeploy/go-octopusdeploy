package projects

type ConversionState struct {
	VariablesAreInGit bool
}

func NewConversionState(variablesAreInGit bool) *ConversionState {
	return &ConversionState{
		VariablesAreInGit: variablesAreInGit,
	}
}
