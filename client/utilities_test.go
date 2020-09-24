package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	empty = emptyString
	tab   = "\t"
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
			t.Logf("[INFO] API Key: %q", tc.test)

			testResult := isAPIKey(tc.test)

			t.Logf("[INFO] Result is %t, expected is %t", testResult, tc.isValid)

			assert.Equal(t, testResult, tc.isValid)
		})
	}
}

func getRandomName() string {
	fullName := fmt.Sprintf("go-octopusdeploy %s", uuid.New())
	fullName = fullName[0:49] //Some names in Octopus have a max limit of 50 characters (such as Environment Name)
	return fullName
}

func getRandomVarName() string {
	return fmt.Sprintf("go-octo-%v", time.Now().Unix())
}

func generateSensitiveValue() model.SensitiveValue {
	sensitiveValue := model.NewSensitiveValue(getRandomName())
	return sensitiveValue
}

func PrettyJSON(data interface{}) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	encoder.Encode(data)
	return buffer.String()
}
