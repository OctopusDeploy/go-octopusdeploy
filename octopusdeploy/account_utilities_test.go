package octopusdeploy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToAccount(t *testing.T) {
	var accountResource *AccountResource
	account, err := ToAccount(accountResource)
	require.Nil(t, account)
	require.Error(t, err)

	accountResource = NewAccountResource("", "")
	account, err = ToAccount(accountResource)
	require.Nil(t, account)
	require.Error(t, err)

	accountResource = NewAccountResource(getRandomName(), "")
	account, err = ToAccount(accountResource)
	require.Nil(t, account)
	require.Error(t, err)

	accountResource = NewAccountResource(getRandomName(), AccountTypeUsernamePassword)
	account, err = ToAccount(accountResource)
	require.NotNil(t, account)
	require.NoError(t, err)
	require.EqualValues(t, account.GetAccountType(), "UsernamePassword")

	account, err = ToAccountResource(accountResource)
	require.NotNil(t, account)
	require.NoError(t, err)
	require.EqualValues(t, account.GetAccountType(), "UsernamePassword")
}
