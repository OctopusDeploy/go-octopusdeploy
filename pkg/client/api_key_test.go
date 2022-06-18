package client

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsAPIKey(t *testing.T) {
	testCases := []struct {
		name    string
		test    string
		isValid bool
	}{
		{"Empty", "", false},
		{"EmptySpace", " ", false},
		{"StartWithAPI", "API-", false},
		{"Invalid", "API-?OBYCAMCZ7WWBKSTMXT66FCUDPS", false},
		{"Invalid", "API-OBYC👍AMCZ7WWBKSTMXT66FCUDPS", false},
		{"Invalid", "API-***************************", false},
		{"Valid", "API-EOAYCAFCZ7WWBKSTMVT66FCUDPS", true},
		{"Valid", "API-EOBYCAMCZ7WWBKSTMVT66FCUDPR", true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testResult := IsAPIKey(tc.test)
			assert.Equal(t, testResult, tc.isValid)
		})
	}
}
