package integration

import "strings"

const (
	emptyString      string = ""
	whitespaceString string = " "
)

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
