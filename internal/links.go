package internal

import "reflect"

func IsLinksEqual(linksA map[string]string, linksB map[string]string) bool {
	return reflect.DeepEqual(linksA, linksB)
}
