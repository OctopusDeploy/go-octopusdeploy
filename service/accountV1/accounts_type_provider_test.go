package accountV1

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleAccountWithoutContainerSerialization(t *testing.T) {
	data := `{"Id":"1","Name":"U/P Account","AccountType":"UsernamePassword","Username":"foo","Password":"bar"}`
	var result *Account
	json.Unmarshal([]byte(data), &result)

	assert.IsType(t, new(Account), result)

	assert.Equal(t, "1", result.GetID())
	assert.Equal(t, "U/P Account", result.Name)
}

func TestSingleAccountSerialization(t *testing.T) {
	data := `{"Id":"1","Name":"U/P Account","AccountType":"UsernamePassword","Username":"foo","Password":"bar"}`
	var result AccountDeserializationContainer
	json.Unmarshal([]byte(data), &result)

	assert.IsType(t, new(UsernamePasswordAccount), result.Account)

	assert.Equal(t, "U/P Account", result.Account.GetName())
	upAccount := result.Account.(*UsernamePasswordAccount)
	assert.IsType(t, new(UsernamePasswordAccount), upAccount)
	assert.Equal(t, "1", upAccount.GetID())
	assert.Equal(t, "foo", upAccount.Username)
}

func TestMultipleAccountsSerialization(t *testing.T) {
	data := `[{"Id":"1","Name":"U/P Account","AccountType":"UsernamePassword","Username":"foo","Password":"bar"},{"Id":"2","Name":"Token Account","AccountType":"Token","Token":"querty"}]`
	var result []AccountDeserializationContainer
	json.Unmarshal([]byte(data), &result)

	assert.IsType(t, new(UsernamePasswordAccount), result[0].Account)
	upAccount := result[0].Account.(*UsernamePasswordAccount)
	assert.IsType(t, new(UsernamePasswordAccount), upAccount)
	assert.Equal(t, "foo", upAccount.Username)

	assert.IsType(t, new(TokenAccount), result[1].Account)
	tokenAccount := result[1].Account.(*TokenAccount)
	assert.NotNil(t, tokenAccount)
	assert.Equal(t, "querty", tokenAccount.Token)
}
