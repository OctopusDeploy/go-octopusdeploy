package integration

import "strings"

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
