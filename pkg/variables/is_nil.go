package variables

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *LibraryVariableSetUsageEntry:
		return v == nil
	case *LibraryVariableSet:
		return v == nil
	case *ScriptModule:
		return v == nil
	default:
		return v == nil
	}
}
