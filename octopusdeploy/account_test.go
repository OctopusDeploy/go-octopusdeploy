package octopusdeploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAccountValues(t *testing.T) {
	baseAccountValid := &Account{
		Name:        "AccountName",
		Description: "AccountDescription",
	}

	assert.Nil(t, baseAccountValid.Validate())

	baseAccountInvalid := &Account{}

	assert.Error(t, baseAccountInvalid.Validate())

	passwordWithValueValid := SensitiveValue{
		HasValue: true,
		NewValue: "blah",
	}

	passwordWithValueInvalid := SensitiveValue{
		HasValue: true,
	}

	passwordNoValueValid := SensitiveValue{
		HasValue: false,
	}

	azureServicePrincipalAccountValid := &Account{
		AccountType:        AzureServicePrincipal,
		Name:               "AzureServicePrincipalAccount",
		ClientID:           "blah",
		TenantID:           "blah",
		SubscriptionNumber: "blah",
		Password:           passwordWithValueValid,
	}

	assert.Nil(t, azureServicePrincipalAccountValid.Validate())

	azureServicePrincipalAccountPasswordInvalid := &Account{
		AccountType:        AzureServicePrincipal,
		Name:               "AzureServicePrincipalAccount",
		ClientID:           "blah",
		TenantID:           "blah",
		SubscriptionNumber: "blah",
		Password:           passwordWithValueInvalid,
	}

	assert.Error(t, azureServicePrincipalAccountPasswordInvalid.Validate())

	azureServicePrincipalAccountPasswordNoValue := &Account{
		AccountType:        AzureServicePrincipal,
		Name:               "AzureServicePrincipalAccount",
		ClientID:           "blah",
		TenantID:           "blah",
		SubscriptionNumber: "blah",
		Password:           passwordNoValueValid,
	}

	assert.Nil(t, azureServicePrincipalAccountPasswordNoValue.Validate())
}
