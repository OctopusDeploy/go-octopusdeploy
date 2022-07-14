package tenants

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Tenant:
		return v == nil
	default:
		return v == nil
	}
}
