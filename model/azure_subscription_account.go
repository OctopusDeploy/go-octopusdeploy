package model

import (
	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
)

func NewAzureSubscriptionAccount(name string, subscriptionID uuid.UUID) (*Account, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewAzureSubscriptionAccount", "name")
	}

	account := &Account{
		Name:        name,
		AccountType: enum.AzureSubscription,
	}
	account.SubscriptionID = &subscriptionID

	return account, nil
}

func validateAzureSubscriptionAccount(account *Account) error {
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
	}

	return ValidateMultipleProperties(validations)
}
