package configuration

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *ConfigurationSection:
		return v == nil
	default:
		return v == nil
	}
}
