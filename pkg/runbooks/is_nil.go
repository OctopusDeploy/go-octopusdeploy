package runbooks

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Runbook:
		return v == nil
	case *RunbookSnapshot:
		return v == nil
	default:
		return v == nil
	}
}
