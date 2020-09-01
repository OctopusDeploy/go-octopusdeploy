package model

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
)

var accountName = "Account Name"
var accountType = enum.UsernamePassword

func TestEmptyAccount(t *testing.T) {
	account := &Account{}

	if account.Name != "" {
		t.Error("Name should be empty")
	}

	if account.AccountType != enum.None {
		t.Errorf("AccountType should be %s", enum.None)
	}
}

func TestAccountWithName(t *testing.T) {
	account := &Account{Name: accountName}

	if account.Name != accountName {
		t.Errorf("Name should be %s", accountName)
	}

	if account.AccountType != enum.None {
		t.Errorf("AccountType should be %s", enum.None)
	}
}

func TestAccountWithNameAndUsernamePasswordAccountType(t *testing.T) {
	account := &Account{
		AccountType: enum.UsernamePassword,
		Name:        accountName,
	}

	if account.Name != accountName {
		t.Errorf("Name should be %s", accountName)
	}

	if account.AccountType != enum.UsernamePassword {
		t.Errorf("AccountType should be %s", enum.UsernamePassword)
	}
}

func TestNewAccountWithValidParameters(t *testing.T) {
	account, err := NewAccount(accountName, accountType)

	if err != nil {
		t.Errorf("NewAccount() generated an unexpected error: %s", err)
	}

	if account.Name != accountName {
		t.Errorf("Name should be %s", accountName)
	}

	if account.AccountType != accountType {
		t.Errorf("AccountType should be %s", accountType)
	}
}

func TestNewAccountWithEmptyName(t *testing.T) {
	account, err := NewAccount(" ", accountType)

	if account != nil {
		t.Error("NewAccount() returned an account, which was unexpected")
	}

	if err == nil {
		t.Error("NewAccount() was expected to generate an error")
	}
}

func TestNewAccountWithLongEmptyName(t *testing.T) {
	account, err := NewAccount("       ", accountType)

	if account != nil {
		t.Error("NewAccount() returned an account, which was unexpected")
	}

	if err == nil {
		t.Error("NewAccount() was expected to generate an error")
	}
}
