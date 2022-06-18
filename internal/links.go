package internal

import "reflect"

func IsEqualLinks(linksA map[string]string, linksB map[string]string) bool {
	return reflect.DeepEqual(linksA, linksB)
}
