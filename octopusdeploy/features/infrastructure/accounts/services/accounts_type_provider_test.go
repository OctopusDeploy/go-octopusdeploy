package services

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/features/infrastructure/accounts/resources"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleAccountWithoutContainerSerialization(t *testing.T) {
	data := `{"Id":"1","Name":"U/P Account","AccountType":"UsernamePassword","Username":"foo","Password":"bar"}`
	var result *resources.Account
	json.Unmarshal([]byte(data), &result)

	assert.IsType(t, new(resources.Account), result)

	assert.Equal(t, "1", result.GetID())
	assert.Equal(t, "U/P Account", result.Name)
}

func TestSingleAccountSerialization(t *testing.T) {
	data := `{"Id":"1","Name":"U/P Account","AccountType":"UsernamePassword","Username":"foo","Password":"bar"}`
	var result AccountDeserializationContainer
	json.Unmarshal([]byte(data), &result)

	assert.IsType(t, new(resources.UsernamePasswordAccount), result.Account)

	assert.Equal(t, "U/P Account", result.Account.GetName())
	upAccount := result.Account.(*resources.UsernamePasswordAccount)
	assert.IsType(t, new(resources.UsernamePasswordAccount), upAccount)
	assert.Equal(t, "1", upAccount.GetID())
	assert.Equal(t, "foo", upAccount.Username)
}

func TestMultipleAccountsSerialization(t *testing.T) {
	data := `[{"Id":"1","Name":"U/P Account","AccountType":"UsernamePassword","Username":"foo","Password":"bar"},{"Id":"2","Name":"Token Account","AccountType":"Token","Token":"querty"}]`
	var result []AccountDeserializationContainer
	json.Unmarshal([]byte(data), &result)

	assert.IsType(t, new(resources.UsernamePasswordAccount), result[0].Account)
	upAccount := result[0].Account.(*resources.UsernamePasswordAccount)
	assert.IsType(t, new(resources.UsernamePasswordAccount), upAccount)
	assert.Equal(t, "foo", upAccount.Username)

	assert.IsType(t, new(resources.TokenAccount), result[1].Account)
	tokenAccount := result[1].Account.(*resources.TokenAccount)
	assert.NotNil(t, tokenAccount)
	assert.Equal(t, "querty", tokenAccount.Token)
}
