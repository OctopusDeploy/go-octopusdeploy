package octopusdeploy

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsAPIKey(t *testing.T) {
	testCases := []struct {
		name    string
		test    string
		isValid bool
	}{
		{"Empty", emptyString, false},
		{"EmptySpace", whitespaceString, false},
		{"StartWithAPI", "API-", false},
		{"Invalid", "API-?OBYCAMCZ7WWBKSTMXT66FCUDPS", false},
		{"Invalid", "API-OBYC👍AMCZ7WWBKSTMXT66FCUDPS", false},
		{"Invalid", "API-***************************", false},
		{"Valid", "API-EOAYCAFCZ7WWBKSTMVT66FCUDPS", true},
		{"Valid", "API-EOBYCAMCZ7WWBKSTMVT66FCUDPR", true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testResult := isAPIKey(tc.test)
			assert.Equal(t, testResult, tc.isValid)
		})
	}
}

func getRandomDuration(mininum time.Duration) time.Duration {
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", rand.Int63n(1000)))
	duration += mininum
	return duration
}

func IsEqualLinks(linksA map[string]string, linksB map[string]string) bool {
	return reflect.DeepEqual(linksA, linksB)
}
