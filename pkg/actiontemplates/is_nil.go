package actiontemplates

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *ActionTemplate:
		return v == nil
	case *ActionTemplateParameter:
		return v == nil
	default:
		return v == nil
	}
}
