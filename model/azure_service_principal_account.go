package model

import (
	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
)

func NewAzureServicePrincipalAccount(name string, subscriptionID uuid.UUID, tenantID uuid.UUID, applicationID uuid.UUID, password SensitiveValue) (*Account, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewAzureServicePrincipalAccount", "name")
	}

	account, err := NewAccount(name, enum.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}

	account.SubscriptionID = &subscriptionID
	account.TenantID = &tenantID
	account.ApplicationID = &applicationID
	account.Password = &password

	return account, nil
}

func validateAzureServicePrincipalAccount(account Account) error {
	validate := validator.New()
	err := validate.Struct(account)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		return err
	}

	validations := []error{
		ValidateRequiredUUID("SubscriptionID", account.SubscriptionID),
		ValidateRequiredUUID("ApplicationID", account.ApplicationID),
		ValidateRequiredUUID("TenantID", account.TenantID),
	}

	return ValidateMultipleProperties(validations)
}
